package sinks

import (
	"reflect"
	"testing"

	"github.com/go-logr/logr"
)

func TestNewSink(t *testing.T) {
	cs := []struct {
		name     string
		sinkType string
		expected Sink
	}{
		{
			name:     "Pass (slack)",
			sinkType: "slack",
			expected: &SlackSink{},
		},
		{
			name:     "Pass (alertmanager)",
			sinkType: "alertmanager",
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
