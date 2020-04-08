package alphavantage

import (
	"fmt"
	"os"
	"time"

	"github.com/SimpleApplicationsOrg/stock/client"
)

const (
	envURL      = "ALPHA_VANTAGE_URL"
	envKeyName  = "ALPHA_VANTAGE_KEY_NAME"
	envKeyValue = "ALPHA_VANTAGE_KEY_VALUE"
)

// AVClient is the client for the Alpha Vantage API
type AVClient struct {
	client *client.APIClient
}

// NewAVClient gives a specific Alpha Vantage API client
func NewAVClient(cfg map[string]string) (*AVClient, error) {
	config, err := configuration()
	if err != nil {
		return nil, err
	}
	for k,v:=range cfg{
		config.AddKey(k,v)
	}
	client := client.NewClient(config)
	return &AVClient{client}, nil
}

func configuration() (*client.Configuration, error) {
	url := os.Getenv(envURL)
	keyName := os.Getenv(envKeyName)
	keyValue := os.Getenv(envKeyValue)

	if url == "" {
		return nil, fmt.Errorf("missing %s", envURL)
	}

	if keyName == "" {
		return nil, fmt.Errorf("missing %s", envKeyName)
	}

	if keyValue == "" {
		return nil, fmt.Errorf("missing %s", envKeyValue)
	}

	config := client.NewConfiguration(url)
	config.AddKey(keyName, keyValue)
	config.WithTimeout(20 * time.Second)

	return config, nil
}
