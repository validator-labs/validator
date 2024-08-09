[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/validator-labs/validator/issues)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Test](https://github.com/validator-labs/validator/actions/workflows/test.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/validator-labs/validator)](https://goreportcard.com/report/github.com/validator-labs/validator)
[![codecov](https://codecov.io/gh/validator-labs/validator/graph/badge.svg?token=GVZ4LZ5SOY)](https://codecov.io/gh/validator-labs/validator)
[![Go Reference](https://pkg.go.dev/badge/github.com/validator-labs/validator.svg)](https://pkg.go.dev/github.com/validator-labs/validator)

validator (AKA Validation Controller) monitors ValidationResults created by one or more validator plugins and uploads them to a sink of your choosing, e.g., Slack or Alertmanager.

<img width="1364" alt="image" src=./img/install_use_flow_diagram.png>

## Description
The validator repository is fairly minimal - all the heavy lifting is done by the validator plugins. Installation of validator and one or more plugins is accomplished via Helm.

Plugins:
- [AWS](https://github.com/validator-labs/validator-plugin-aws)
- [Azure](https://github.com/validator-labs/validator-plugin-azure)
- [Network](https://github.com/validator-labs/validator-plugin-network)
- [OCI](https://github.com/validator-labs/validator-plugin-oci)
- [vSphere](https://github.com/validator-labs/validator-plugin-vsphere)
- [Kubescape](https://github.com/validator-labs/validator-plugin-kubescape)

## Installation

### Connected
For connected installations, two options are supported: the validator CLI, `validatorctl`, and Helm. Using validatorctl is **recommended**, as it provides a text-based user interface (TUI) for configuring validator.

#### Validator CLI
1. Download the latest release of validatorctl from https://github.com/validator-labs/validatorctl/releases
2. Execute validatorctl
   ```bash
   validatorctl install
   ```

#### Helm
Install Validator by pulling the latest Helm chart and installing it in your cluster. Use the following commands to install the latest version of the chart.

```bash
helm repo add validator https://validator-labs.github.io/validator
helm repo update
helm install validator validator/validator -n validator --create-namespace
```

Check out the [Helm install guide](./docs/install.md) for a step-by-step guide for installing and using Validator.

### Air-gapped
For air-gapped installations, the recommended approach is to use [Hauler](https://github.com/rancherfederal/hauler). Hauls containing all validator artifacts (container images, Helm charts, and the validator CLI) are generated for multiple platforms (linux/amd64 and linux/arm64) during each validator release.

Prerequisites:
* A Linux-based air-gapped workstation with:
  * A container runtime, e.g., [containerd](https://containerd.io/docs/getting-started/), [Docker Engine](https://docs.docker.com/engine/), etc.
  * If using a container runtime other than Docker Engine, [podman](https://podman.io/docs/installation) must be installed and [Docker CLI emulation](https://podman-desktop.io/docs/migrating-from-docker/emulating-docker-cli-with-podman) configured
  * [kind](https://github.com/kubernetes-sigs/kind/releases) installed and on your PATH
  * [hauler](https://github.com/hauler-dev/hauler/releases) installed and on your PATH

Once the prerequisites are met, the following steps document the air-gapped installation procedure:

1. Download the Hauler Store (then somehow get it across the air-gap)
   ```bash
   # Download the Haul for your chosen release and platform, e.g.:
   curl -L https://github.com/validator-labs/validator/releases/download/v0.0.46/validator-haul-linux-amd64.tar.zst -o validator-haul-linux-amd64.tar.zst
   ```
2. Load the Hauler Store (on the air-gapped workstation)
   ```bash
   # Load the air-gapped content to your local hauler store.
   hauler store load validator-haul-linux-amd64.tar.zst
   ```
3. Extract validatorctl from the Hauler Store
   ```bash
   # Extract the validator CLI binary, validatorctl, from the hauler store.
   # It's always tagged as "latest" within the store, despite being versioned.
   # This is a hauler defect. The version can be verified via `validatorctl version`.
   hauler store extract -s store hauler/validatorctl:latest
   chmod +x validatorctl && mv validatorctl /usr/local/bin
   ```
4. Serve the Hauler Store
   ```bash
   # Serve the content as a registry from the hauler store.
   # (Defaults to <FQDN or IP>:5000).
   nohup hauler store serve registry | tee -a hauler.log &
   
   # Optionally tail the hauler registry logs
   tail -f hauler.log
   ```
5. Execute validatorctl
   ```bash
   validatorctl install
   ```

## Sinks
Validator can be configured to emit updates to various event sinks whenever a `ValidationResult` is created or updated. See configuration details below for each supported sink.

### Alertmanager
Integrate with the Alertmanager API to emit alerts to all [supported Alertmanager receivers](https://prometheus.io/docs/alerting/latest/configuration/#receiver-integration-settings), including generic webhooks. The only required configuration is an Alertmanager endpoint. HTTP basic authentication and TLS are also supported. See [values.yaml](https://github.com/validator-labs/validator/blob/main/chart/validator/values.yaml) for configuration details.

#### Sample Output
![Screen Shot 2023-11-15 at 10 42 20 AM](https://github.com/validator-labs/validator/assets/1795270/ce958b8e-96d7-4f5e-8efc-80e2fc2b0b4d)

#### Setup
1. Install Alertmanager in your cluster (if it isn't installed already)
2. Configure Alertmanager alert content. Alerts can be formatted/customized via the following labels and annotations:

   Labels
   - alertname
   - plugin
   - validation_result
   - expected_results

   Annotations
   - state
   - validation_rule
   - validation_type
   - message
   - status
   - detail
      - pipe-delimited array of detail messages, see sample config for parsing example
   - failure (also pipe-delimited)
   - last_validation_time

   Example Alertmanager ConfigMap used to produce the sample output above:
   ```yaml
   apiVersion: v1
   data:
   alertmanager.yml: |
      global:
         slack_api_url: https://slack.com/api/chat.postMessage
      receivers:
      - name: default-receiver
         slack_configs:
         - channel: <channel-id>
         text: |-
            {{ range .Alerts.Firing -}}
            *Validation Result: {{ .Labels.validation_result }}/{{ .Labels.expected_results }}*

            {{ range $k, $v := .Annotations }}
            {{- if $v }}*{{ $k | title }}*:
            {{- if match "\\|" $v }}
            - {{ reReplaceAll "\\|" "\n- " $v -}}
            {{- else }}
            {{- printf " %s" $v -}}
            {{- end }}
            {{- end }}
            {{ end }}

            {{ end }}
         title: "{{ (index .Alerts 0).Labels.plugin }}: {{ (index .Alerts 0).Labels.alertname }}\n"
         http_config:
            authorization:
               credentials: xoxb--<bot>-<token>
         send_resolved: false
      route:
         group_interval: 10s
         group_wait: 10s
         receiver: default-receiver
         repeat_interval: 1h
      templates:
      - /etc/alertmanager/*.tmpl
   kind: ConfigMap
   metadata:
   name: alertmanager
   namespace: alertmanager
   ```

2. Install validator and/or upgrade your validator Helm release, configuring `values.sink` accordingly.

### Slack

#### Sample Output
<img width="704" alt="Screen Shot 2023-11-10 at 4 30 12 PM" src="https://github.com/validator-labs/validator/assets/1795270/c011143a-4d4b-4299-b88b-699188f4bda2">
<img width="700" alt="Screen Shot 2023-11-10 at 4 18 22 PM" src="https://github.com/validator-labs/validator/assets/1795270/9f2c4ab7-34d6-496a-9f60-68655a7ee3d6">

#### Setup
1. Go to https://api.slack.com/apps and click **Create New App**, then select **From scratch**. Pick an App Name and Slack Workspace, then click **Create App**.

   <img src="https://github.com/validator-labs/validator/assets/1795270/58cbb5a0-12a4-4a83-a0dd-20ae87a8105d" width="500">

2. Go to `OAuth & Permissions` and copy the `Bot User OAuth Token` under the `OAuth Tokens for Your Workspace` section. Save it somewhere for later. Scroll down to `Scopes` and click **Add an OAuth Scope**. Enable the `chat:write` scope for your bot.

   <img src="https://github.com/validator-labs/validator/assets/1795270/7b4d80be-5799-497a-9a4b-480793b26d59" width="500">

3. Find and/or create a channel in Slack and note its Channel ID (at the very bottom of the modal when you view channel details). Add the bot you just created to the channel via `View channel details > Integrations > Apps > Add apps`.

   <img src="https://github.com/validator-labs/validator/assets/1795270/a78c852c-7aeb-41a4-aa76-6afbe9b2ec81" width="500">

4. Install validator and/or upgrade your validator Helm release, configuring `values.sink` accordingly.

## Development
You’ll need a Kubernetes cluster to run against. You can use [kind](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/validator:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/validator:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing
All contributions are welcome! Feel free to reach out on the [Spectro Cloud community Slack](https://spectrocloudcommunity.slack.com/join/shared_invite/zt-g8gfzrhf-cKavsGD_myOh30K24pImLA#/shared-invite/email).

Make sure `pre-commit` is [installed](https://pre-commit.com#install).

Install the `pre-commit` scripts:

```console
pre-commit install --hook-type commit-msg
pre-commit install --hook-type pre-commit
```

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)
