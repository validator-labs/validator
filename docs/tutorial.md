# Install Guide

This install guide will help you install the Validator and get started using one the plugins. Validator supports multiple plugins and each plugin may require a set of configuration parameters. This guide will use the AWS plugin and show you how to configure it. Use this guide as a starting point for installing Validator and its plugins.



## Prerequisites

- An AWS account.

- AWS credentials with AdministratorAccess policy attached. You can create a new user with this policy and use the credentials for the installation. See [AWS documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html) for more details.

- [wget](https://www.gnu.org/software/wget/) or similar tool installed on your machine.

- A text editor.

- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) installed on your machine.

- [helm](https://helm.sh/docs/intro/install/) installed on your machine.

- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/) installed on your machine.

> [!NOTE]
> If you already have a Kubernetes cluster, you can skip the kind installation and use your existing cluster. Ensure you have access to the cluster.

## Install Validator

Create a Kubernetes cluster using kind. If you already have a cluster, you can skip this step.


```shell
kind create cluster --name validator
```


Validate that the cluster is available and that you have access to it. Use the following command to get the cluster information.

```
kubectl cluster-info
```

Example output:
```
Kubernetes control plane is running at https://127.0.0.1:49470
CoreDNS is running at https://127.0.0.1:49470/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```

Add the Validator Helm repository and update the local Helm cache.

```shell
helm repo add validator https://spectrocloud-labs.github.io/validator/ && \
helm repo update
```

Before you install Validator, you can configure the installation by editing the `values.yaml` file. The `values.yaml` file contains the default configuration for Validator. You can override the default configuration by editing the `values.yaml` file.



Create a **values.yaml** containing the Validator configuration. Use the command below to download the default **values.yaml** file.

```shell
wget https://raw.githubusercontent.com/spectrocloud-labs/validator/main/chart/validator/values.yaml
```

Use a text editor to edit the **values.yaml** file and change navigate down to the `plugins` section and locate the `AWS` plugin. 

```yaml
plugins:
- chart:
name: validator-plugin-aws
repository: "https://spectrocloud-labs.github.io/validator-plugin-aws"
version: "v0.0.13"
values: |-
# Remainder of the file omitted for brevity
```

The AWS plugin requires credentials to access your AWS account. Uncomment the `auth` section and add your AWS credentials. By default the auth section contains a an empty `{}` block. Replace the empty block with the following configuration. Replace the `accessKeyId` and `secretAccessKey` with your AWS credentials. 

```yaml
auth:
# Leave secret undefined for implicit auth (node instance role, IMDSv2, etc.)
secret:
secretName: aws-creds
accessKeyId: "**********"
secretAccessKey: "**********"
sessionToken: ""
# By default, a secret will be created. Leave the above fields blank and specify 'createSecret: false' to use an existing secret.

# WARNING: the existing secret must match the format used in auth-secret.yaml
createSecret: true
```
You can remove all other plugins from the `plugins` section or leave them as is, but for this guide, only the AWS plugin will be used.


Now that you have configured the AWS plugin, you can install Validator. Use the following command to install Validator and the AWS plugin.

```shell
helm install validator validator/validator --values values.yaml --namespace validator --create-namespace
```

Validate that the Validator pods and the AWS plugin pods are available. Use the following command to get the status of the pods.

```shell
kubectl get pods --namespace validator
```


```shell
NAME                                                       READY   STATUS    RESTARTS   AGE
validator-controller-manager-589664b9c4-7zt7w              2/2     Running   0          73s
validator-plugin-aws-controller-manager-64bf9c5b56-l2j8w   2/2     Running   0          61s
```

Now that you have Validator installed, you can create a validation. A validation is a custom resource that contains the configuration for the validation. Validator will use the configuration to validate the cluster. Validator will create a `ValidationResult` custom resource that contains the result of the validation.

## Create a Validation

The next step is to create a validation configuration. Each plugin contains a set of example validation configurations. You can use the example configurations as a starting point for your validation. For this guide, you will use the [`awsvalidator-spectro-cloud-base`](https://github.com/spectrocloud-labs/validator-plugin-aws/blob/main/config/samples/awsvalidator-iam-role-spectro-cloud-base.yaml) configuration. 

> [!NOTE]
> Check ou the AWS [config samples](https://github.com/spectrocloud-labs/validator-plugin-aws/tree/main/config/samples) directory for more examples.

