# Install Guide

This guide will guide you through the process of installing the Validator controller on your cluster. Validator supports multiple plugins and each plugin may require a set of configuration parameters. This guide will use the AWS plugin and show you how to configure it. Use this guide as a starting point for installing Validator and its plugins.



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