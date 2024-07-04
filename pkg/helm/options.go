package helm

// Options holds all the options for installing/pulling/upgrading a chart.
type Options struct {
	Chart string
	Path  string

	Repo            string
	Version         string
	Values          string
	SetValues       map[string]string
	SetStringValues map[string]string

	Username string
	Password string

	Atomic          bool
	Force           bool
	CreateNamespace bool

	Untar    bool
	UntarDir string

	CaFile                string
	InsecureSkipTLSVerify bool

	ExtraArgs []string
}

// ConfigureRepo adds the repo flag to the command.
func (o Options) ConfigureRepo(args []string) []string {
	args = append(args, "--repo", o.Repo)
	return args
}

// ConfigureVersion adds the version flag to the command.
func (o Options) ConfigureVersion(args []string) []string {
	if o.Version != "" {
		args = append(args, "--version", o.Version)
	}
	return args
}

// ConfigureArchive adds the archive flags to the command.
func (o Options) ConfigureArchive(args []string) []string {
	if o.Untar {
		args = append(args, "--untar")
	}
	if o.UntarDir != "" {
		args = append(args, "--untardir", o.UntarDir)
	}
	return args
}

// ConfigureAuth adds basic auth flags to the command.
func (o Options) ConfigureAuth(args []string) []string {
	if o.Username != "" {
		args = append(args, "--username", o.Username)
	}
	if o.Password != "" {
		args = append(args, "--password", o.Password)
	}
	return args
}

// ConfigureTLS adds TLS flags to the command.
func (o Options) ConfigureTLS(args []string) []string {
	if o.CaFile != "" {
		args = append(args, "--ca-file", o.CaFile)
	}
	if o.InsecureSkipTLSVerify {
		args = append(args, "--insecure-skip-tls-verify")
	}
	return args
}
