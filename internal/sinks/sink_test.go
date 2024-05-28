package sinks

import (
	"reflect"
	"testing"

	"github.com/go-logr/logr"

	"github.com/validator-labs/validator/pkg/types"
)

func TestNewSink(t *testing.T) {
	cs := []struct {
		name     string
		sinkType types.SinkType
		expected Sink
	}{
		{
			name:     "Pass (slack)",
			sinkType: types.SinkTypeSlack,
			expected: &SlackSink{},
		},
		{
			name:     "Pass (alertmanager)",
			sinkType: types.SinkTypeAlertmanager,
			expected: &AlertmanagerSink{},
		},
		{
			name:     "Pass (default)",
			sinkType: "foo",
			expected: &SlackSink{},
		},
	}
	for _, c := range cs {
		t.Log(c.name)
		sink := NewSink(c.sinkType, logr.Logger{})
		if !reflect.DeepEqual(sink, c.expected) {
			t.Errorf("expected (%+v), got (%+v)", c.expected, sink)
		}
	}
}
