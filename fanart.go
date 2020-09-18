package fanart

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// APIClientConfig exposes fields that are mapped to a JSON configuration file.
type APIClientConfig struct {
	API struct {
		Key      string `json:"key"`
		Endpoint string `json:"endpoint"`
	} `json:"api"`

	Account struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"account"`
}

// APIClient wraps the standard library http.Client and maintains any state required
// for interacting with the fanart API.
type APIClient struct {
	http.Client      `json:"-"`
	*APIClientConfig `json:"conf"`
	Endpoint         string `json:"endpoint"`
}

// New creates a new client, initialized with customized standard library http.Client.
func New(config io.Reader, timeoutSeconds int) (*APIClient, error) {
	raw, err := ioutil.ReadAll(config)
	if err != nil {
		return nil, err
	}

	var (
		ac *APIClientConfig = &APIClientConfig{}
	)

	if err := json.Unmarshal(raw, ac); err != nil {
		return nil, err
	}

	return &APIClient{
		APIClientConfig: ac,
		Endpoint:        strings.Trim(ac.API.Endpoint, "/"),
		Client: http.Client{
			Timeout: time.Second * time.Duration(timeoutSeconds),
		},
	}, nil
}

// func (api *APIClient) Movies() {
//
// }
