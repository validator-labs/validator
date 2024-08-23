// Package plugins provides the interface that must be implemented by all plugins.
package plugins

// PluginSpec is the interface that must be implemented by all plugins.
type PluginSpec interface {
	PluginCode() string
	ResultCount() int
}
