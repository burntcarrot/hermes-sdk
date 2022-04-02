package providers

import (
	"context"
	"net/http"
	"net/url"

	"github.com/deepsourcelabs/hermes/domain"
	"github.com/deepsourcelabs/hermes/provider/discord"

	"github.com/deepsourcelabs/hermes-sdk/sdk/utils"
)

type DiscordService struct {
	Client     *discord.Client
	BaseURL    *url.URL
	webhookURI string
}

type DiscordResponse struct {
	Ok bool `json:"ok"`
}

// Setup sets the webhook URI for the Discord service.
func (s *DiscordService) Setup(webhookURI string) {
	s.webhookURI = webhookURI
}

// Send sends a templated message to Discord.
func (s *DiscordService) Send(ctx context.Context, tmp *domain.Template, payload *map[string]interface{}) (DiscordResponse, error) {
	opts := map[string]interface{}{
		"webhook": s.webhookURI,
	}

	body := utils.PrepareReqBody(payload, tmp, "", opts, "discord")

	// build request
	req, err := utils.BuildReq(ctx, http.MethodPost, body, nil, s.BaseURL.String(), "application/json")
	if err != nil {
		return DiscordResponse{}, err
	}

	// perform request
	resp, err := utils.PerformRequest(req, 200)
	if err != nil {
		return DiscordResponse{}, err
	}

	if resp == nil {
		return DiscordResponse{}, err
	}

	message := resp.([]domain.Message)[0]
	discordResp := DiscordResponse{
		Ok: message.Ok,
	}

	return discordResp, nil
}
