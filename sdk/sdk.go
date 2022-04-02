package sdk

import (
	"net/http"
	"net/url"

	"github.com/deepsourcelabs/hermes/domain"
	"github.com/deepsourcelabs/hermes/provider/slack"

	slackProvider "github.com/deepsourcelabs/hermes-sdk/sdk/providers/slack"
)

type Client struct {
	Slack *slackProvider.SlackService
}

// NewClient returns a new Client.
func NewClient(baseURL string) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	sc := slack.Client{
		HTTPClient: http.DefaultClient,
	}

	c := &Client{
		Slack: &slackProvider.SlackService{Client: &sc, BaseURL: u},
	}

	return c, nil
}

// GetTemplate returns a template.
func GetTemplate(id string) *domain.Template {
	return &domain.Template{
		ID: id,
	}
}
