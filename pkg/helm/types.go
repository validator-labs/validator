package helm

import corev1 "k8s.io/api/core/v1"

// Release describes a deployment of a chart, together with the chart
// and the variables used to deploy that chart.
type Release struct {
	// Name is the name of the release
	Name string `json:"name,omitempty"`
	// Info provides information about a release
	Info *Info `json:"info,omitempty"`
	// Chart is the chart that was released.
	Chart *Chart `json:"chart,omitempty"`
	// Config is the set of extra Values added to the chart.
	// These values override the default values inside of the chart.
	Config map[string]interface{} `json:"config,omitempty"`
	// Version is an int which represents the version of the release.
	Version int `json:"version,omitempty"`
	// Namespace is the kubernetes namespace of the release.
	Namespace string `json:"namespace,omitempty"`

	Secret *corev1.Secret `json:"-"`
}

// Info describes release information.
type Info struct {
	// FirstDeployed is when the release was first deployed.
	// +optional
	FirstDeployed Time `json:"first_deployed,omitempty"`
	// LastDeployed is when the release was last deployed.
	// +optional
	LastDeployed Time `json:"last_deployed,omitempty"`
	// Deleted tracks when this object was deleted.
	// +optional
	Deleted Time `json:"deleted,omitempty"`
	// Description is human-friendly "log entry" about this release.
	// +optional
	Description string `json:"description,omitempty"`
	// Status is the current state of the release
	// +optional
	Status string `json:"status,omitempty"`
	// Contains the rendered templates/NOTES.txt if available
	// +optional
	Notes string `json:"notes,omitempty"`
}

// Chart holds the chart metadata
type Chart struct {
	Metadata *Metadata `json:"metadata,omitempty"`
}

// UpgradeOptions holds all the options for upgrading / installing a chart
type UpgradeOptions struct {
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

	InsecureSkipTlsVerify bool

	ExtraArgs []string
}

type Maintainer struct {
	// Name is a user name or organization name
	// +optional
	Name string `json:"name,omitempty"`
	// Email is an optional email address to contact the named maintainer
	// +optional
	Email string `json:"email,omitempty"`
	// URL is an optional URL to an address for the named maintainer
	// +optional
	URL string `json:"url,omitempty"`
}

type Metadata struct {
	// The name of the chart
	// +optional
	Name string `json:"name,omitempty"`
	// The URL to a relevant project page, git repo, or contact person
	// +optional
	Home string `json:"home,omitempty"`
	// Source is the URL to the source code of this chart
	// +optional
	Sources []string `json:"sources,omitempty"`
	// A SemVer 2 conformant version string of the chart
	// +optional
	Version string `json:"version,omitempty"`
	// A one-sentence description of the chart
	// +optional
	Description string `json:"description,omitempty"`
	// A list of string keywords
	// +optional
	Keywords []string `json:"keywords,omitempty"`
	// A list of name and URL/email address combinations for the maintainer(s)
	// +optional
	Maintainers []*Maintainer `json:"maintainers,omitempty"`
	// The URL to an icon file.
	// +optional
	Icon string `json:"icon,omitempty"`
	// The API Version of this chart.
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`
	// The condition to check to enable chart
	// +optional
	Condition string `json:"condition,omitempty"`
	// The tags to check to enable chart
	// +optional
	Tags string `json:"tags,omitempty"`
	// The version of the application enclosed inside of this chart.
	// +optional
	AppVersion string `json:"appVersion,omitempty"`
	// Whether or not this chart is deprecated
	// +optional
	Deprecated bool `json:"deprecated,omitempty"`
	// Annotations are additional mappings uninterpreted by Helm,
	// made available for inspection by other applications.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// KubeVersion is a SemVer constraint specifying the version of Kubernetes required.
	// +optional
	KubeVersion string `json:"kubeVersion,omitempty"`
	// Specifies the chart type: application or library
	// +optional
	Type string `json:"type,omitempty"`
	// Urls where to find the chart contents
	// +optional
	Urls []string `json:"urls,omitempty"`
}
