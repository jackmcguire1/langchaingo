package cloudflareclient

import (
	"net/http"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient         httpClient
	accountID          string
	token              string
	baseURL            string
	modelName          string
	embeddingModelName string
	endpointURL        string
	bearerToken        string
}

func NewClient(client httpClient, accountID, baseURL, token, modelName, embeddingModelName string) *Client {
	return &Client{
		httpClient:         client,
		accountID:          accountID,
		baseURL:            baseURL,
		token:              token,
		modelName:          modelName,
		embeddingModelName: embeddingModelName,
		endpointURL:        "%s/%s/ai/run/%s",
		bearerToken:        "Bearer " + token,
	}
}

func (c *Client) SetModel(model string) {
	c.modelName = model
}
