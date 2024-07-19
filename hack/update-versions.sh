#!/usr/bin/env bash

function latestRelease {
  release=$(gh release list --repo $1 -L 1 | head -n 2 | awk '{ print $1; }')
  echo ${release:1}
}

function updateHaulerManifest {
  envsubst < hack/hauler-manifest-template.yaml > hauler-manifest.yaml
  echo "Updated hauler-manifest.yaml with latest versions."
}

function addChartValues {
  # Use yq to remove env and proxy from each plugin's values.yaml as they're injected from the
  # validator chart's env and proxy config via chart/validator/templates/validator-config.yaml
  values=$(curl -sL https://raw.githubusercontent.com/validator-labs/$1/main/chart/$1/values.yaml | yq eval "del(.env) | del(.proxy)")

  # Fix indentation
  indentedValues=""
  while IFS= read -r line; do
      indentedValues="${indentedValues}    $line"$'\n'
  done <<< "$values"

  # Append the plugin's values to chart/validator/values.yaml
  cat <<EOF >> chart/validator/values.yaml
- chart:
    name: $1
    repository: "https://validator-labs.github.io/$1"
    version: v$2
  values: |-
$indentedValues
EOF
}

function updateValues {
  # Reset values.yaml
  cp -f hack/chart/values-base.yaml chart/validator/values.yaml

  # Add plugin values to values.yaml
  for plugin in "${!versions[@]}"; do
    version=${versions[$plugin]}
    addChartValues validator-plugin-$plugin $version
    truncate -s-1 chart/validator/values.yaml
    echo "Updated values.yaml for validator-plugin-$plugin @ v$version."
  done
}

export AWS_VERSION=$(latestRelease validator-labs/validator-plugin-aws)
export AZURE_VERSION=$(latestRelease validator-labs/validator-plugin-azure)
export KUBESCAPE_VERSION=$(latestRelease validator-labs/validator-plugin-kubescape)
export MAAS_VERSION=$(latestRelease validator-labs/validator-plugin-maas)
export NETWORK_VERSION=$(latestRelease validator-labs/validator-plugin-network)
export OCI_VERSION=$(latestRelease validator-labs/validator-plugin-oci)
export VSPHERE_VERSION=$(latestRelease validator-labs/validator-plugin-vsphere)
export VALIDATORCTL_VERSION=$(latestRelease validator-labs/validatorctl)

declare -A versions
versions["aws"]=$AWS_VERSION
versions["azure"]=$AZURE_VERSION
versions["kubescape"]=$KUBESCAPE_VERSION
versions["maas"]=$MAAS_VERSION
versions["network"]=$NETWORK_VERSION
versions["oci"]=$OCI_VERSION
versions["vsphere"]=$VSPHERE_VERSION

updateHaulerManifest
updateValues