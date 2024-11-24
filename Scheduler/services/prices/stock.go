package prices

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetStockPrice(stockSymbol string) (float64, error) {
	AlphaVantageAPIKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	url := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=5min&apikey=%s",
		stockSymbol,
		AlphaVantageAPIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch stock price: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected API response: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read API response: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	timeSeries, ok := result["Time Series (5min)"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("failed to find time series data")
	}

	for _, v := range timeSeries {
		data, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		closeStr, ok := data["4. close"].(string)
		if !ok {
			continue
		}

		var closePrice float64
		if _, err := fmt.Sscanf(closeStr, "%f", &closePrice); err != nil {
			return 0, fmt.Errorf("failed to parse close price: %v", err)
		}

		return closePrice, nil
	}

	return 0, fmt.Errorf("no stock price data found")
}
