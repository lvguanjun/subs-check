package platform

import (
	"encoding/json"
	"io"
	"net/http"
)

// IPInfoResponse represents the JSON response from ipinfo.io
type IPInfoResponse struct {
	Org string `json:"org"`
}

// CheckIPInfo checks if the node's organization is not "AS13335 Cloudflare, Inc."
func CheckIPInfo(httpClient *http.Client) (bool, error) {
	// Create request
	req, err := http.NewRequest("GET", "https://3.0.3.0/ips", nil)
	if err != nil {
		return false, err
	}

	// Add headers to simulate normal browser access
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "close")

	// Send request
	resp, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Parse JSON response
	var ipInfo IPInfoResponse
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		return false, err
	}

	// Check if organization is Cloudflare
	if ipInfo.Org == "Cloudflare, Inc." {
		return false, nil
	}

	return true, nil
}