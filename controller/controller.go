package controller

import (
	"bytes"
	"encoding/json"
	"example/zota_assignment/config"
	"example/zota_assignment/data"
	"fmt"
	"io"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) DepositRequest(request data.ZotaOrderRequest) ([]byte, error) {
	url := "https://" + config.BaseURL + "/api/v1/deposit/request/" + config.EndpointID

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return body, nil
}
