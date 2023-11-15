package sinks

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

var sinkClient = NewClient(1 * time.Second)

func TestAlertmanagerConfigure(t *testing.T) {
	cs := []struct {
		name     string
		sink     AlertmanagerSink
		config   map[string][]byte
		expected error
	}{
		{
			name: "Pass",
			sink: AlertmanagerSink{},
			config: map[string][]byte{
				"endpoint": []byte("http://fake.alertmanager.com:9093/api/v2/alerts"),
				"caCert":   []byte("_fake_ca_cert"),
			},
			expected: nil,
		},
		{
			name: "Fail (invalid endpoint)",
			sink: AlertmanagerSink{},
			config: map[string][]byte{
				"endpoint": []byte("_not_an_endpoint_"),
			},
			expected: InvalidEndpoint,
		},
		{
			name: "Fail (invalid insecureSkipVerify)",
			sink: AlertmanagerSink{},
			config: map[string][]byte{
				"endpoint":           []byte("https://fake.com"),
				"insecureSkipVerify": []byte("_not_a_bool_"),
			},
			expected: errors.New(`invalid Alertmanager config: failed to parse insecureSkipVerify: strconv.ParseBool: parsing "_not_a_bool_": invalid syntax`),
		},
	}
	for _, c := range cs {
		t.Log(c.name)
		err := c.sink.Configure(*sinkClient, c.config)
		if err != nil && !reflect.DeepEqual(err.Error(), c.expected.Error()) {
			t.Errorf("expected (%v), got (%v)", c.expected, err)
		}
	}
}
