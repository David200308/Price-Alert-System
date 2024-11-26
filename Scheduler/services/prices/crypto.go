package prices

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type CryptoPriceResponse struct {
	Symbol    string `json:"symbol"`
	Price     string `json:"price"`
	Timestamp int64  `json:"timestamp"`
}

func GetCryptoPrice(cryptoPair string) (float64, error) {
	apiURL := fmt.Sprintf("https://api.api-ninjas.com/v1/cryptoprice?symbol=%s", cryptoPair)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Api-Key", os.Getenv("NINJAS_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API returned error: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	var response CryptoPriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("failed to decode response: %w", err)
	}

	price, err := strconv.ParseFloat(response.Price, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse price: %w", err)
	}

	return price, nil
}
