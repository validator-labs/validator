package sinks

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"

	"github.com/validator-labs/validator/api/v1alpha1"
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
			name:     "Fail (no endpoint)",
			sink:     AlertmanagerSink{},
			config:   map[string][]byte{},
			expected: EndpointRequired,
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

func TestAlertManagerEmit(t *testing.T) {
	cs := []struct {
		name     string
		sink     AlertmanagerSink
		res      v1alpha1.ValidationResult
		server   *httptest.Server
		expected error
	}{
		{
			name: "Pass",
			sink: AlertmanagerSink{},
			res: v1alpha1.ValidationResult{
				Status: v1alpha1.ValidationResultStatus{
					ValidationConditions: []v1alpha1.ValidationCondition{
						{
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "ok")
			})),
			expected: nil,
		},
		{
			name: "Fail",
			sink: AlertmanagerSink{},
			res:  v1alpha1.ValidationResult{},
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "invalid auth", http.StatusUnauthorized)
			})),
			expected: SinkEmissionFailed,
		},
	}
	for _, c := range cs {
		t.Log(c.name)
		defer c.server.Close()
		_ = c.sink.Configure(*sinkClient, map[string][]byte{
			"endpoint": []byte(c.server.URL),
		})
		err := c.sink.Emit(c.res)
		if err != nil && !reflect.DeepEqual(err.Error(), c.expected.Error()) {
			t.Errorf("expected (%v), got (%v)", c.expected, err)
		}
	}
}

func TestBasicAuthHeader(t *testing.T) {
	cs := []struct {
		name     string
		username string
		password string
		expected string
	}{
		{
			name:     "Pass",
			username: "bob",
			password: "frogs",
			expected: "Basic Ym9iOmZyb2dz",
		},
	}
	for _, c := range cs {
		t.Log(c.name)
		_, v := basicAuthHeader(c.username, c.password)
		if !reflect.DeepEqual(c.expected, v) {
			t.Errorf("expected (%s), got (%s)", c.expected, v)
		}
	}
}
