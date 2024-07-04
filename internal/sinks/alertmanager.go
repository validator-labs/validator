package sinks

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/validator-labs/validator/api/v1alpha1"
)

// AlertmanagerSink is a sink for sending validation results to Alertmanager.
type AlertmanagerSink struct {
	client Client
	log    logr.Logger

	endpoint string
	username string
	password string
}

// Alert is an Alertmanager alert.
type Alert struct {
	// Annotations are arbitrary key-value pairs.
	Annotations map[string]string `json:"annotations"`

	// Labels are key-value pairs that can be used to group and filter alerts.
	Labels map[string]string `json:"labels"`
}

var (
	// ErrInvalidEndpoint is returned when an Alertmanager endpoint is invalid.
	ErrInvalidEndpoint = errors.New("invalid Alertmanager config: endpoint scheme and host are required")

	// ErrEndpointRequired is returned when the Alertmanager endpoint is not provided.
	ErrEndpointRequired = errors.New("invalid Alertmanager config: endpoint required")
)

// Configure configures the AlertmanagerSink with the provided configuration.
func (s *AlertmanagerSink) Configure(c Client, config map[string][]byte) error {
	// endpoint
	endpoint, ok := config["endpoint"]
	if !ok {
		return ErrEndpointRequired
	}
	u, err := url.Parse(string(endpoint))
	if err != nil {
		return errors.Wrap(err, "invalid Alertmanager config: failed to parse endpoint")
	}
	if u.Scheme == "" || u.Host == "" {
		return ErrInvalidEndpoint
	}
	if u.Path != "" {
		s.log.V(1).Info("stripping path from Alertmanager endpoint", "path", u.Path)
		u.Path = ""
	}
	s.endpoint = fmt.Sprintf("%s/api/v2/alerts", u.String())

	// basic auth
	s.username = string(config["username"])
	s.password = string(config["password"])

	// tls
	var caCertPool *x509.CertPool
	var insecureSkipVerify bool

	insecure, ok := config["insecureSkipVerify"]
	if ok {
		insecureSkipVerify, err = strconv.ParseBool(string(insecure))
		if err != nil {
			return errors.Wrap(err, "invalid Alertmanager config: failed to parse insecureSkipVerify")
		}
	}
	caCert, ok := config["caCert"]
	if ok {
		caCertPool, err = x509.SystemCertPool()
		if err != nil {
			return errors.Wrap(err, "invalid Alertmanager config: failed to get system cert pool")
		}
		caCertPool.AppendCertsFromPEM(caCert)
	}
	c.hclient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecureSkipVerify, // #nosec G402
			MinVersion:         tls.VersionTLS12,
			RootCAs:            caCertPool,
		},
	}
	s.client = c

	return nil
}

// Emit sends a ValidationResult to Alertmanager.
func (s *AlertmanagerSink) Emit(r v1alpha1.ValidationResult) error {
	alerts := make([]Alert, 0, len(r.Status.ValidationConditions))

	for i, c := range r.Status.ValidationConditions {
		alerts = append(alerts, Alert{
			Labels: map[string]string{
				"alertname":         r.Name,
				"plugin":            r.Spec.Plugin,
				"validation_result": strconv.Itoa(i + 1),
				"expected_results":  strconv.Itoa(r.Spec.ExpectedResults),
			},
			Annotations: map[string]string{
				"state":                string(r.Status.State),
				"validation_rule":      c.ValidationRule,
				"validation_type":      c.ValidationType,
				"message":              c.Message,
				"status":               string(c.Status),
				"detail":               strings.Join(c.Details, "|"),
				"failure":              strings.Join(c.Failures, "|"),
				"last_validation_time": c.LastValidationTime.String(),
			},
		})
	}

	body, err := json.Marshal(alerts)
	if err != nil {
		s.log.Error(err, "failed to marshal alerts", "alerts", alerts)
		return err
	}
	s.log.V(1).Info("Alertmanager message", "payload", body)

	req, err := http.NewRequest(http.MethodPost, s.endpoint, bytes.NewReader(body))
	if err != nil {
		s.log.Error(err, "failed to create HTTP POST request", "endpoint", s.endpoint)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	if s.username != "" && s.password != "" {
		req.Header.Add(basicAuthHeader(s.username, s.password))
	}

	resp, err := s.client.hclient.Do(req)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		s.log.Error(err, "failed to post alert", "endpoint", s.endpoint)
		return err
	}
	if resp.StatusCode != 200 {
		s.log.V(0).Info("failed to post alert", "endpoint", s.endpoint, "status", resp.Status, "code", resp.StatusCode)
		return ErrSinkEmissionFailed
	}

	s.log.V(0).Info("Successfully posted alert to Alertmanager", "endpoint", s.endpoint, "status", resp.Status, "code", resp.StatusCode)
	return nil
}

func basicAuthHeader(username, password string) (string, string) {
	auth := base64.StdEncoding.EncodeToString(
		bytes.Join([][]byte{[]byte(username), []byte(password)}, []byte(":")),
	)
	return "Authorization", fmt.Sprintf("Basic %s", auth)
}
