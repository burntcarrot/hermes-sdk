package slack

import (
	"context"
	"net/http"
	"net/url"

	"github.com/deepsourcelabs/hermes/domain"
	"github.com/deepsourcelabs/hermes/provider/slack"

	"github.com/deepsourcelabs/hermes-sdk/sdk/utils"
)

type SlackService struct {
	Client  *slack.Client
	BaseURL *url.URL
	token   string
}

type SlackResponse struct {
	Ok bool `json:"ok"`
}

// Setup sets the token for the Slack service.
func (s *SlackService) Setup(token string) {
	s.token = token
}

// Send sends a templated message to the configured channel.
func (s *SlackService) Send(ctx context.Context, tmp *domain.Template, payload *map[string]interface{}, channel string) (SlackResponse, error) {
	opts := map[string]interface{}{
		"channel": channel,
	}

	body := utils.PrepareReqBody(payload, tmp, s.token, opts, "slack")

	// build request
	req, err := utils.BuildReq(ctx, http.MethodPost, body, nil, s.BaseURL.String(), "application/json")
	if err != nil {
		return SlackResponse{}, err
	}

	// perform request
	resp, err := utils.PerformRequest(req, 200)
	if err != nil {
		return SlackResponse{}, err
	}

	if resp == nil {
		return SlackResponse{}, err
	}

	message := resp.([]domain.Message)[0]
	slackResp := SlackResponse{
		Ok: message.Ok,
	}

	return slackResp, nil
}
