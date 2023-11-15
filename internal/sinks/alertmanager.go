package sinks

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"github.com/spectrocloud-labs/validator/api/v1alpha1"
)

type AlertmanagerSink struct {
	client Client
	log    logr.Logger

	// TODO: add auth [ Basic | mTLS ]
	endpoint string
	username string
	password string
}

type Alert struct {
	Annotations map[string]string `json:"annotations"`
	Labels      map[string]string `json:"labels"`
}

func (s *AlertmanagerSink) Configure(c Client, vc v1alpha1.ValidatorConfig, config map[string][]byte) error {
	endpoint, ok := config["endpoint"]
	if !ok {
		return errors.New("invalid Alertmanager configuration: endpoint required")
	}
	u, err := url.Parse(string(endpoint))
	if err != nil {
		return errors.Wrap(err, "invalid Alertmanager endpoint")
	}
	s.endpoint = fmt.Sprintf("%s/api/v2/alerts", u.String())

	s.username = string(config["username"])
	s.password = string(config["password"])

	// c.hclient.Transport = &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		InsecureSkipVerify: true,
	// 	},
	// }
	s.client = c

	return nil
}

func (s *AlertmanagerSink) Emit(r v1alpha1.ValidationResult) error {
	alerts := make([]Alert, 0, len(r.Status.Conditions))

	for i, c := range r.Status.Conditions {
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
		s.log.Error(err, "failed to post alert", "endpoint", s.endpoint, "status", resp.Status, "code", resp.StatusCode)
		return err
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
