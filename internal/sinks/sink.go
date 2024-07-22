// Package sinks contains sinks for emitting ValidationResults.
package sinks

import (
	"net/http"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/validator-labs/validator/api/v1alpha1"
	"github.com/validator-labs/validator/pkg/types"
)

// ErrSinkEmissionFailed is returned when emitting a validation result to a sink fails.
var ErrSinkEmissionFailed = errors.New("sink emission failed")

// Sink is an interface for sending validation results to a sink.
type Sink interface {
	Configure(c Client, config map[string][]byte) error
	Emit(result v1alpha1.ValidationResult) error
}

// NewSink returns a new Sink based on the provided SinkType.
func NewSink(sinkType types.SinkType, log logr.Logger) Sink {
	switch sinkType {
	case types.SinkTypeAlertmanager:
		return &AlertmanagerSink{log: log}
	case types.SinkTypeSlack:
		return &SlackSink{log: log}
	default:
		return &SlackSink{log: log}
	}
}

// Client is an HTTP client for a Sink.
type Client struct {
	hclient *http.Client
}

// NewClient returns a new Client with the provided timeout.
func NewClient(timeout time.Duration) *Client {
	client := http.DefaultClient // inherit http.ProxyFromEnvironment
	client.Timeout = timeout
	return &Client{
		hclient: client,
	}
}
