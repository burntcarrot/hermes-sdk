package sdk

import (
	"net/http"
	"net/url"

	"github.com/deepsourcelabs/hermes/domain"
	"github.com/deepsourcelabs/hermes/provider/discord"
	"github.com/deepsourcelabs/hermes/provider/slack"

	"github.com/deepsourcelabs/hermes-sdk/sdk/providers"
)

type Client struct {
	Slack   *providers.SlackService
	Discord *providers.DiscordService
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

	dc := discord.Client{
		HTTPClient: http.DefaultClient,
	}

	c := &Client{
		Slack:   &providers.SlackService{Client: &sc, BaseURL: u},
		Discord: &providers.DiscordService{Client: &dc, BaseURL: u},
	}

	return c, nil
}

// GetTemplate returns a template.
func GetTemplate(id string) *domain.Template {
	return &domain.Template{
		ID: id,
	}
}
