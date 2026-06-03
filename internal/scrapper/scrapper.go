package scrapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

const (
	requestTimeout = 15 * time.Second
	mlAPIBaseURL   = "https://api.mercadolibre.com/items/"
)

var productIDPattern = regexp.MustCompile(`MLB-?(\d+)`)

type Scrapper struct {
	client *http.Client
}

func NewScrapper() *Scrapper {
	return &Scrapper{
		client: &http.Client{Timeout: requestTimeout},
	}
}

func (s *Scrapper) PriceColector(url string) (float64, error) {
	productID, err := extractProductID(url)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Get(mlAPIBaseURL + productID)
	if err != nil {
		return 0, fmt.Errorf("api request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var payload struct {
		Price  float64 `json:"price"`
		Status string  `json:"status"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return 0, fmt.Errorf("decode json: %w", err)
	}

	if payload.Price <= 0 {
		return 0, fmt.Errorf("invalid price returned by api: %v (status=%s)", payload.Price, payload.Status)
	}

	return payload.Price, nil
}

func extractProductID(url string) (string, error) {
	match := productIDPattern.FindStringSubmatch(url)
	if len(match) < 2 {
		return "", errors.New("could not extract Mercado Livre product id from url")
	}
	return "MLB" + match[1], nil
}
