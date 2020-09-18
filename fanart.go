package fanart

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// MovieRequest contains the parameters to an API request.
type MovieRequest struct {
	MovieID string `json:"MovieID,omitempty"`
}

func (mr MovieRequest) String() string {
	return mr.MovieID
}

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

// MovieImages is like MovieImagesRaw except that it returns a serialized struct.
func (client *APIClient) MovieImages(payload MovieRequest) (*MovieResponse, error) {
	data, err := client.MovieImagesRaw(payload)
	if err != nil {
		return nil, err
	}

	mr := &MovieResponse{}

	if err := json.Unmarshal(data, mr); err != nil {
		return nil, err
	}

	return mr, nil
}

// MovieImagesRaw returns the raw byte payload of a call to the /movies endpoint.
func (client *APIClient) MovieImagesRaw(payload MovieRequest) ([]byte, error) {
	resource := fmt.Sprintf("%s/movies/%s", client.Endpoint, payload)

	if payload.MovieID == "" {
		return nil, fmt.Errorf("missing required query parameter: need imdb id or movie id")
	}

	req, err := http.NewRequest(http.MethodGet, resource, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("api-key", client.API.Key)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	return data, nil
}

// TVImages is like TVImagesRaw except that it returns a serialized struct.
func (client *APIClient) TVImages(payload MovieRequest) ([]byte, error) {
	return nil, nil
}

// TVImagesRaw returns the raw byte payload of a call to the /tv endpoint.
func (client *APIClient) TVImagesRaw(payload MovieRequest) ([]byte, error) {
	return nil, nil
}

// MusicImages is like TVImagesRaw except that it returns a serialized struct.
func (client *APIClient) MusicImages(payload MovieRequest) ([]byte, error) {
	return nil, nil
}

// MusicImagesRaw returns the raw byte payload of a call to the /tv endpoint.
func (client *APIClient) MusicImagesRaw(payload MovieRequest) ([]byte, error) {
	return nil, nil
}
