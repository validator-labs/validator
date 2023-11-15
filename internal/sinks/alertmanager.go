package sinks

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	// c.hclient.Transport = &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		InsecureSkipVerify: true,
	// 	},
	// }
	s.client = c

	return nil
}

func (s *AlertmanagerSink) Emit(r v1alpha1.ValidationResult) error {
	alerts := []Alert{
		{
			Labels: map[string]string{
				"alertname":       r.Name,
				"plugin":          r.Spec.Plugin,
				"expectedresults": strconv.Itoa(r.Spec.ExpectedResults),
			},
			Annotations: map[string]string{
				"state": string(r.Status.State),
			},
		},
	}

	for i, c := range r.Status.Conditions {
		alerts[0].Annotations[fmt.Sprintf("validationRule%d", i)] = c.ValidationRule
		alerts[0].Annotations[fmt.Sprintf("validationType%d", i)] = c.ValidationType
		alerts[0].Annotations[fmt.Sprintf("message%d", i)] = c.Message
		alerts[0].Annotations[fmt.Sprintf("status%d", i)] = string(c.Status)
		alerts[0].Annotations[fmt.Sprintf("detail%d", i)] = strings.Join(c.Details, "|")
		alerts[0].Annotations[fmt.Sprintf("failure%d", i)] = strings.Join(c.Failures, "|")
		alerts[0].Annotations[fmt.Sprintf("lastValidationTime%d", i)] = c.LastValidationTime.String()
	}

	body, err := json.Marshal(alerts)
	if err != nil {
		s.log.Error(err, "failed to marshal alerts", "alerts", alerts)
		return err
	}
	s.log.V(1).Info("Alertmanager message", "payload", body)

	resp, err := s.client.hclient.Post(s.endpoint, "application/json", bytes.NewReader(body))
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
