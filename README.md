![Build](https://github.com/spectrocloud-labs/validator/actions/workflows/build_container.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/spectrocloud-labs/validator)](https://goreportcard.com/report/github.com/spectrocloud-labs/validator)
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/spectrocloud-labs/validator/issues)
[![Go Reference](https://pkg.go.dev/badge/github.com/spectrocloud-labs/validator.svg)](https://pkg.go.dev/github.com/spectrocloud-labs/validator)

# validator
validator (AKA Validation Controller) monitors ValidationResults created by one or more validator plugins and uploads them to a sink of your choosing, e.g., Slack or Spectro Cloud Palette.
<img width="1364" alt="image" src="https://github.com/spectrocloud-labs/validator/assets/1795270/e68dfdf5-25bf-4336-ad87-f783c4825c7e">

## Description
The validator repository is fairly minimal - all the heavy lifting is done by the validator plugins. Installation of validator and one or more plugins is accomplished via Helm.

Plugins:
- [AWS](https://github.com/spectrocloud-labs/validator-plugin-aws)
- [Azure](https://github.com/spectrocloud-labs/validator-plugin-azure)
- [Network](https://github.com/spectrocloud-labs/validator-plugin-network)
- [OCI](https://github.com/spectrocloud-labs/validator-plugin-oci)
- [vSphere](https://github.com/spectrocloud-labs/validator-plugin-vsphere)

## Installation
```bash
helm repo add validator https://spectrocloud-labs.github.io/validator/
helm repo update
helm install validator validator/validator -n validator --create-namespace
```

## Sinks
Validator can be configured to emit updates to various event sinks whenever a `ValidationResult` is created or updated. See configuration details below for each supported sink.

### Slack
1. Go to https://api.slack.com/apps and click **Create New App**, then select **From scratch**. Pick an App Name and Slack Workspace, then click **Create App**.
2. Go to `OAuth & Permissions` and copy the `Bot User OAuth Token` under the `OAuth Tokens for Your Workspace` section. Save it somewhere for later. Scroll down to `Scopes` and click **Add an OAuth Scope**. Enable the `chat:write` scope for your bot.
3. Find and/or create a channel in Slack and note its Channel ID (at the very bottom of the model when you view channel details). Add the bot you just created to the channel via `View channel details > Integrations > Apps > Add apps`.
4. Install validator and/or upgrade your validator Helm release, configuring `values.sink` accordingly.

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
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

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

