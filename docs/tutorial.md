# Install Guide

This install guide will help you install the Validator and get started using one the plugins. Validator supports multiple plugins and each plugin may require a set of configuration parameters. This guide uses the AWS plugin and shows you how to configure it. Use this guide as a starting point for installing Validator and its plugins.


## Prerequisites

- An AWS account.

- AWS credentials with [AdministratorAccess](https://docs.aws.amazon.com/aws-managed-policy/latest/reference/AdministratorAccess.html) policy attached. You can create a new user with this policy and use the credentials with the AWS plguin. Refer to the [AWS IAM User](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html) documentation for more details.

> [!IMPORTANT]
> You can use a different policy, but ensure that the policy has the required permissions to check the resources required by the plugin. Refer to the AWS plugin documentation for more details.

- [wget](https://www.gnu.org/software/wget/) or a similar tool installed on your machine.

- A text editor.

- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) installed on your machine.

- [helm](https://helm.sh/docs/intro/install/) installed on your machine.

- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/) installed on your machine.

> [!NOTE]
> If you already have a Kubernetes cluster, you can skip the kind installation and use your existing cluster. Ensure you have access to the cluster.

## Overview

The following diagram displays at a high-level the steps you will complete. You will install Validator in a Kubernetes cluster, configure a plugin, deploy a validation plugin's custom resource, and review the results.

![An illustration of sequence of steps](./img/install_use_flow_diagram.png)


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

The `values.yaml` file contains the default configuration for Validator. Before you install Validator, you can configure the installation by editing the `values.yaml` file and overwriting the default values. 



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

Now that you have Validator installed, you can create a validation. A validation is a custom resource that contains the configuration for the validation. Validator will create a `ValidationResult` custom resource that contains the result of the validation.

## Create a Validation

The next step is to create a validation configuration. Each plugin contains a set of example validation configurations. You can use the example configurations as a starting point for your validation. For this guide, you will use the [`awsvalidator-spectro-cloud-base`](https://github.com/spectrocloud-labs/validator-plugin-aws/blob/main/config/samples/awsvalidator-iam-role-spectro-cloud-base.yaml) configuration. 

> [!NOTE]
> Check ou the AWS [config samples](https://github.com/spectrocloud-labs/validator-plugin-aws/tree/main/config/samples) directory for more examples.


Create a file named `validation.yaml` and copy the contents of the [`awsvalidator-spectro-cloud-base`](https://github.com/spectrocloud-labs/validator-plugin-aws/blob/main/config/samples/awsvalidator-iam-role-spectro-cloud-base.yaml) Use the following command to download the file and save it as `validation.yaml`.

```shell
wget https://raw.githubusercontent.com/spectrocloud-labs/validator-plugin-aws/main/config/samples/awsvalidator-iam-role-spectro-cloud-base.yaml --output-document validation.yaml
```

Review the contents of the `validation.yaml` file. The file contains the configuration for the validation. The `spec` section contains the configuration for the validation.  Change the `spec` section to match your requirements. For example, you can change the `defaultRegion` to match your prefered AWS region. 


```yaml
apiVersion: validation.spectrocloud.labs/v1alpha1
kind: AwsValidator
metadata:
  name: awsvalidator-spectro-cloud-base
spec:
  auth: {}
  defaultRegion: us-west-1
  iamRoleRules:
  - iamPolicies:
    - name: Controllers Policy
      statements:
      - actions:
      # Remainder of the file omitted for brevity
```


Now that you have the validation configuration, you can create the validation. Use the following command to create the validation.

```shell
kubectl apply --values validation.yaml
```

You can verify the Custom Resource (CR) for the AWS plugin was created by using the following command. 

```shell
kubectl get crd
```

```shell
NAME                                             CREATED AT
awsvalidators.validation.spectrocloud.labs       2023-11-14T23:47:03Z
validationresults.validation.spectrocloud.labs   2023-11-14T23:46:54Z
validatorconfigs.validation.spectrocloud.labs    2023-11-14T23:46:54Z
```

Review the result of the validation by using the following command.

```shell
kubectl describe validationresults
```

```shell
Name:         validator-plugin-aws-awsvalidator-spectro-cloud-base
Namespace:    default
Labels:       <none>
Annotations:  validator/validation-result-hash: 2IjYPs4C+8fu6MnuBq09lg==
API Version:  validation.spectrocloud.labs/v1alpha1
Kind:         ValidationResult
Metadata:
  Creation Timestamp:  2023-11-15T17:22:33Z
  Generation:          1
  Resource Version:    141250
  UID:                 18d37dd9-9e4a-4397-b857-03116cf62975
Spec:
  Expected Results:  1
  Plugin:            AWS
Status:
  Conditions:
    Failures:
      v1alpha1.IamRoleRule SpectroCloudRole missing action(s): [s3:PutObject s3:DeleteObject s3:PutBucketOwnershipControls s3:PutBucketPolicy s3:PutBucketPublicAccessBlock s3:PutObjectAcl] for resource arn:*:s3:::* from policy Controllers Policy
    Last Validation Time:  2023-11-15T17:26:37Z
    Message:               One or more required IAM permissions was not found, or a condition was not met
    Status:                False
    Validation Rule:       validation-SpectroCloudRole
    Validation Type:       aws-iam-role-policy
  Sink State:              N/A
  State:                   Failed
Events:                    <none>
```

If the validation fails, you can review the `Failures` section of the `Conditions` section to determine the cause of the failure. In this example, the IAM role is missing the required permissions. You can update the IAM role to add the missing permissions and rerun the validation.


The Validator will continouslly re-issue a validation cand update the `ValidationResult` CR with the result of the validation. You can use the following command to get the status of the validation. Once you addressed the cause of the failure, the validation will pass.


If you encountered an error and fixed the error, after about 30 seconds check the validation results again. In this example, the IAM policy was updated to add the missing permissions. 


```shell
kubectl describe validationresults
```

```shell
Name:         validator-plugin-aws-awsvalidator-spectro-cloud-base
Namespace:    default
Labels:       <none>
Annotations:  validator/validation-result-hash: 2A5jQj0W4SBN8IKGC0zRbQ==
API Version:  validation.spectrocloud.labs/v1alpha1
Kind:         ValidationResult
Metadata:
  Creation Timestamp:  2023-11-15T17:43:11Z
  Generation:          1
  Resource Version:    1937
  UID:                 876b29bd-8d96-47f4-8d0c-c4d6eeae039c
Spec:
  Expected Results:  1
  Plugin:            AWS
Status:
  Conditions:
    Last Validation Time:  2023-11-15T17:53:21Z
    Message:               All required aws-iam-role-policy permissions were found
    Status:                True
    Validation Rule:       validation-SpectroCloudRole
    Validation Type:       aws-iam-role-policy
  Sink State:              N/A
  State:                   Succeeded
Events:                    <none>
```


The validation results are hashed and result events are only updated if the result has changed. In this example, the validation check was successful and the `ValidationResult` CR was updated with the result of the validation.



## Cleanup

To delete the Validator installation, use the following command. This command deletes the Validator and all deployed resources including the `ValidationResult` CRs.

```shell
helm uninstall validator --namespace validator
```

If you deployed a kind cluster, use the following command to delete the cluster.

```shell
kind delete clusters validator
```


## Next Steps

In this guide you learned how to install Validator and configure one of the plugins. You also learned how to create a validation and review the results. Use the knowledge you gained in this guide to configure the other plugins and create your own validations.

We encourage you to explore the other plugins and create your own validations as you gain more experience with Validator.

## Additional Resources


Below are links to the documentation for the other plugins.

- [AWS](https://github.com/spectrocloud-labs/validator-plugin-aws)
- [Azure](https://github.com/spectrocloud-labs/validator-plugin-azure)
- [Network](https://github.com/spectrocloud-labs/validator-plugin-network)
- [OCI](https://github.com/spectrocloud-labs/validator-plugin-oci)
- [vSphere](https://github.com/spectrocloud-labs/validator-plugin-vsphere)