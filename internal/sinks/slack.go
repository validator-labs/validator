package sinks

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"github.com/slack-go/slack"

	"github.com/validator-labs/validator/api/v1alpha1"
)

// SlackSink is a sink for sending validation results to Slack.
type SlackSink struct {
	apiToken  string
	channelID string
	client    Client
	log       logr.Logger
}

// Configure configures the SlackSink with the provided configuration.
func (s *SlackSink) Configure(c Client, config map[string][]byte) error {
	apiToken, ok := config["apiToken"]
	if !ok {
		return errors.New("invalid Slack configuration: apiToken required")
	}
	channelID, ok := config["channelId"]
	if !ok {
		return errors.New("invalid Slack configuration: channelId required")
	}
	s.apiToken = string(apiToken)
	s.channelID = string(channelID)
	s.client = c
	return nil
}

// Emit sends a ValidationResult to Slack.
func (s *SlackSink) Emit(r v1alpha1.ValidationResult) error {
	api := slack.New(
		s.apiToken,
		slack.OptionHTTPClient(s.client.hclient),
	)
	_, err := api.AuthTest()
	if err != nil {
		s.log.Error(err, "failed to authenticate with Slack")
		return err
	}

	_, timestamp, err := api.PostMessage(
		s.channelID,
		slack.MsgOptionBlocks(s.buildSlackBlocks(r)...),
		slack.MsgOptionUsername("Validator Bot"),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		s.log.Error(err, "failed to post message", "channelId", s.channelID, "timestamp", timestamp)
		return err
	}
	s.log.V(0).Info("Successfully posted message to channel", "channelId", s.channelID, "timestamp", timestamp)

	return nil
}

func (s *SlackSink) buildSlackBlocks(r v1alpha1.ValidationResult) []slack.Block {

	// Overall State
	state := fmt.Sprintf("*State:* %s :white_check_mark:", r.Status.State)
	if r.Status.State != v1alpha1.ValidationSucceeded {
		state = fmt.Sprintf("*State:* %s :red_circle:", r.Status.State)
	}
	blocks := []slack.Block{
		slack.NewHeaderBlock(&slack.TextBlockObject{
			Type: slack.PlainTextType,
			Text: fmt.Sprintf("%s Validator: %s", r.Spec.Plugin, r.Name),
		}),
		slack.NewSectionBlock(newTextBlockObject(state), nil, nil),
		slack.NewDividerBlock(),
	}

	for i, c := range r.Status.ValidationConditions {

		// Basics
		blocks = append(blocks, []slack.Block{
			slack.NewSectionBlock(newTextBlockObject(fmt.Sprintf("*ValidationResult* %d", i+1)), nil, nil),
			slack.NewSectionBlock(newTextBlockObject(":information_source: Metadata"), nil, nil),
			slack.NewSectionBlock(newTextBlockObject(fmt.Sprintf("*Name:* %s", r.Name)), nil, nil),
			slack.NewSectionBlock(newTextBlockObject(fmt.Sprintf("*Validation Type:* %s", c.ValidationType)), nil, nil),
			slack.NewSectionBlock(newTextBlockObject(fmt.Sprintf("*Validation Rule:* %s", c.ValidationRule)), nil, nil),
			slack.NewSectionBlock(newTextBlockObject(fmt.Sprintf("*Message:* %s", c.Message)), nil, nil),
		}...)

		// Details
		if len(c.Details) > 0 {
			detailsText := newTextBlockObject(":mag_right: Details")
			detailsSection := slack.NewSectionBlock(detailsText, nil, nil)
			blocks = append(blocks, detailsSection)

			for _, d := range c.Details {
				detail := newTextBlockObject(fmt.Sprintf("- %s", d))
				detailSection := slack.NewSectionBlock(detail, nil, nil)
				blocks = append(blocks, detailSection)
			}
		}

		// Failures
		if len(c.Failures) > 0 {
			failuresText := newTextBlockObject(":x: Failures")
			failuresSection := slack.NewSectionBlock(failuresText, nil, nil)
			blocks = append(blocks, failuresSection)

			for i, f := range c.Failures {
				failure := newTextBlockObject(fmt.Sprintf("%d. %s", i+1, f))
				failureSection := slack.NewSectionBlock(failure, nil, nil)
				blocks = append(blocks, failureSection)
			}
		}

		blocks = append(blocks, slack.NewDividerBlock())
	}

	payload, _ := json.Marshal(blocks)
	s.log.V(1).Info("Slack message", "payload", payload)

	return blocks
}

func newTextBlockObject(s string) *slack.TextBlockObject {
	// https://api.slack.com/reference/surfaces/formatting#escaping
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return slack.NewTextBlockObject(slack.MarkdownType, s, false, true)
}
