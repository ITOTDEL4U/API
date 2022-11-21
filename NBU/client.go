package NBU

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("tomeout can't be zero")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil

}

func (c Client) GetAssets() ([]AssetData, error) {

	today := time.Now().Format("20060102")

	resp, err := c.client.Get("https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?date=" + today + "&json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := make([]AssetData, 10)

	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
