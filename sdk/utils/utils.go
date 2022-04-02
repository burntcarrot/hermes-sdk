package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/deepsourcelabs/hermes/domain"
	"github.com/deepsourcelabs/hermes/service"
)

// PrepareReqBody returns the prepared request body.
func PrepareReqBody(payload *map[string]interface{}, tmp *domain.Template, token string, opts map[string]interface{}, providerType string) *service.SendMessageRequest {
	return &service.SendMessageRequest{
		Payload: payload,
		Recipients: []struct {
			Notifier *domain.Notifier `json:"notifier"`
			Template *domain.Template `json:"template"`
		}{
			{
				Notifier: &domain.Notifier{
					Config: &domain.NotifierConfiguration{
						Secret: &domain.NotifierSecret{
							Token: token,
						},
						Opts: opts,
					},
					Type: domain.ProviderType(providerType),
				},
				Template: tmp,
			},
		},
	}
}

// BuildReq builds and returns the request.
func BuildReq(ctx context.Context, method string, r interface{}, conf interface{}, baseURL string, contentType string) (*http.Request, error) {
	// prepare request body
	var reqBody []byte
	var err error
	if r != nil {
		reqBody, err = json.MarshalIndent(r, "", "	")
		if err != nil {
			return nil, err
		}
	}

	// create request
	req, err := http.NewRequestWithContext(ctx, method, baseURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	return req, nil
}

// PerformRequest executes the request and validates its status code.
func PerformRequest(req *http.Request, expectedStatus int) (interface{}, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != expectedStatus {
		return nil, errors.New("status code doesn't match")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	message := []domain.Message{}
	err = json.Unmarshal(body, &message)
	if err != nil {
		return nil, err
	}

	return message, nil
}
