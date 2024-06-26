package sinks

import (
	"net/http"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/validator-labs/validator/api/v1alpha1"
	"github.com/validator-labs/validator/pkg/types"
)

var SinkEmissionFailed = errors.New("sink emission failed")

type Sink interface {
	Configure(c Client, config map[string][]byte) error
	Emit(result v1alpha1.ValidationResult) error
}

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

type Client struct {
	hclient *http.Client
}

func NewClient(timeout time.Duration) *Client {
	client := &http.Client{
		Timeout: timeout,
	}
	return &Client{
		hclient: client,
	}
}
