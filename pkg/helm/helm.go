package helm

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	klog "k8s.io/klog/v2"
)

var (
	CommandPath   = "./helm"
	preserveFiles = false // whether to preserve kubeconfig and Helm values files
)

func init() {
	if os.Getenv("HELM_PRESERVE_FILES") == "true" {
		preserveFiles = true
	}
}

// HelmClient is an interface for interacting with Helm
type HelmClient interface {
	Delete(name, namespace string) error
	Pull(options Options) error
	Upgrade(name, namespace string, options Options) error
}

type helmClient struct {
	config   *clientcmdapi.Config
	helmPath string

	stderr io.Writer
	stdout io.Writer
}

// NewHelmClient creates a new Helm client from the given config
func NewHelmClient(config *clientcmdapi.Config) *helmClient {
	return &helmClient{
		config:   config,
		helmPath: CommandPath,
	}
}

func (c *helmClient) Delete(name, namespace string) error {
	kubeConfig, err := writeKubeConfig(c.config)
	if err != nil {
		return err
	}
	if !preserveFiles {
		defer func() {
			if err := os.Remove(kubeConfig); err != nil {
				klog.Errorf("failed to remove temp file %s: %v", kubeConfig, err)
			}
		}()
	}

	args := []string{"delete", name, "--namespace", namespace, "--kubeconfig", kubeConfig}
	return c.exec(args)
}

func (c *helmClient) Pull(options Options) error {
	if options.Repo == "" {
		return fmt.Errorf("chart repo cannot be null")
	}
	args := []string{"pull", options.Repo}
	args = options.ConfigureVersion(args)
	args = options.ConfigureArchive(args)
	args = options.ConfigureAuth(args)
	args = options.ConfigureTLS(args)
	return c.exec(args)
}

func (c *helmClient) Upgrade(name, namespace string, options Options) error {
	options.ExtraArgs = append(options.ExtraArgs, "--install")
	return c.run(name, namespace, options, "upgrade", options.ExtraArgs)
}

func (c *helmClient) run(name, namespace string, options Options, command string, extraArgs []string) error {
	kubeConfig, err := writeKubeConfig(c.config)
	if err != nil {
		return err
	}
	if !preserveFiles {
		defer func() {
			if err := os.Remove(kubeConfig); err != nil {
				klog.Errorf("failed to remove temp file %s: %v", kubeConfig, err)
			}
		}()
	}

	args := []string{command, name}
	if options.Path != "" {
		args = append(args, options.Path)
	} else if options.Chart != "" {
		args = append(args, options.Chart)
		if options.Repo == "" {
			return fmt.Errorf("chart repo cannot be null")
		}
		args = options.ConfigureRepo(args)
		args = options.ConfigureVersion(args)
		args = options.ConfigureAuth(args)
		args = options.ConfigureTLS(args)
	}

	args = append(args, "--kubeconfig", kubeConfig, "--namespace", namespace)
	args = append(args, extraArgs...)
	if options.CreateNamespace {
		args = append(args, "--create-namespace")
	}

	// Values
	if options.Values != "" {
		// Create temp file
		tempFile, err := os.CreateTemp("", "")
		if err != nil {
			return errors.Wrap(err, "create temp file")
		}

		// Write to temp file
		_, err = tempFile.Write([]byte(options.Values))
		if err != nil {
			if removeErr := os.Remove(tempFile.Name()); removeErr != nil {
				klog.Errorf("failed to remove temp file %s: %v", tempFile.Name(), err)
			}
			return errors.Wrap(err, "write temp file")
		}

		// Close temp file
		if err := tempFile.Close(); err != nil {
			return errors.Wrap(err, "close temp file")
		}
		if !preserveFiles {
			defer func() {
				if err := os.Remove(tempFile.Name()); err != nil {
					klog.Errorf("failed to remove temp file %s: %v", tempFile.Name(), err)
				}
			}()
		}

		// Wait quickly so helm will find the file
		time.Sleep(time.Millisecond)

		args = append(args, "--values", tempFile.Name())
	}

	// Set values
	if options.SetValues != nil && len(options.SetValues) > 0 {
		args = append(args, "--set")

		setString := ""
		for key, value := range options.SetValues {
			if setString != "" {
				setString += ","
			}

			setString += key + "=" + value
		}

		args = append(args, setString)
	}

	// Set string values
	if options.SetStringValues != nil && len(options.SetStringValues) > 0 {
		args = append(args, "--set-string")

		setString := ""
		for key, value := range options.SetStringValues {
			if setString != "" {
				setString += ","
			}
			setString += key + "=" + value
		}
		args = append(args, setString)
	}

	if options.Force {
		args = append(args, "--force")
	}
	if options.Atomic {
		args = append(args, "--atomic")
	}

	return c.exec(args)
}

func (c *helmClient) exec(args []string) error {
	if len(args) == 0 {
		return nil
	}

	sb := strings.Builder{}
	mask := false
	for _, a := range args {
		if mask {
			sb.WriteString("***** ")
			mask = false
			continue
		}
		if a == "--password" {
			mask = true
		}
		sb.WriteString(a)
		sb.WriteString(" ")
	}
	sanitizedArgs := sb.String()

	fmt.Println("helm " + sanitizedArgs)
	cmd := exec.Command(c.helmPath, args...) // #nosec G204
	if c.stdout != nil {
		cmd.Stdout = c.stdout
		cmd.Stderr = c.stderr
		return cmd.Run()
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "release: not found") {
			return nil
		}
		klog.Errorf("Error executing command: helm %s", sanitizedArgs)
		klog.Errorf("Output: %s, Error: %v", string(output), err)
		return fmt.Errorf("error executing helm %s: %s", args[0], string(output))
	}

	return nil
}

// writeKubeConfig writes the kubeconfig to a file and returns the filename
func writeKubeConfig(configRaw *clientcmdapi.Config) (string, error) {
	data, err := clientcmd.Write(*configRaw)
	if err != nil {
		return "", err
	}

	// Create temp file
	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		return "", errors.Wrap(err, "create temp file")
	}

	// Write to temp file
	_, err = tempFile.Write(data)
	if err != nil {
		if removeErr := os.Remove(tempFile.Name()); removeErr != nil {
			klog.Errorf("failed to remove temp file %s: %v", tempFile.Name(), err)
		}
		return "", errors.Wrap(err, "write temp file")
	}

	// Close temp file
	if err := tempFile.Close(); err != nil {
		return "", errors.Wrap(err, "close temp file")
	}

	// Okay sometimes the file is written so quickly that helm somehow
	// cannot read it immediately which causes errors
	// so we wait here till the file is ready
	now := time.Now()
	for time.Since(now) < time.Minute {
		_, err = os.Stat(tempFile.Name())
		if err != nil {
			if os.IsNotExist(err) {
				time.Sleep(time.Millisecond * 50)
				continue
			}
			if removeErr := os.Remove(tempFile.Name()); removeErr != nil {
				klog.Errorf("failed to remove temp file %s: %v", tempFile.Name(), err)
			}
			return "", err
		}
		break
	}

	return tempFile.Name(), nil
}
