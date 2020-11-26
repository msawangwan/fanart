package fanart_test

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"encoding/json"

	"github.com/msawangwan/fanart"
)

const VERBOSE = false

var (
	mockConfig = strings.NewReader(`{
		"api": {
			"key": "SECRET_KEY",
			"endpoint": "http://webservice.fanart.tv/v3"
		},
		"account": {
			"username": "foo",
			"password": "bar"
		}
	}`)
)

func pretty(t *testing.T, o interface{}) {
	raw, err := json.MarshalIndent(o, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	if VERBOSE {
		t.Logf(string(raw))
	}
}

func TestCreateClient(t *testing.T) {
	client, err := fanart.New(mockConfig, 1)
	if err != nil {
		t.Fatal(err)
	}

	pretty(t, client)
}

func TestGetMovieArtResponseRaw(t *testing.T) {
	var (
		secret string
	)

	if v, found := os.LookupEnv("FANART_API_KEY"); found {
		secret = v
	} else {
		t.Skip("no valid api key found, set one with 'FANART_API_KEY'")
	}

	conf := strings.NewReader(fmt.Sprintf(`{
		"api": {
			"key": "%s",
			"endpoint": "http://webservice.fanart.tv/v3"
		}
	}`, secret))

	client, err := fanart.New(conf, 10)
	if err != nil {
		t.Fatal(err)
	}

	var testcases = []struct {
		label string
		req   fanart.MovieRequest
	}{
		{"imdbId/hook", fanart.MovieRequest{MovieID: "tt0102057"}},
		{"imdbId/gone_with_the_wind", fanart.MovieRequest{MovieID: "tt0031381"}},
		{"imdbId/no_such_resource", fanart.MovieRequest{MovieID: "tt3097934"}},
	}

	for _, testcase := range testcases {
		t.Run(testcase.label, func(tt *testing.T) {
			res, err := client.MovieImagesRaw(testcase.req)
			if err != nil {
				if errors.Is(err, fanart.ErrNoSuchResource) {
					return
				}
				t.Error(err)
			}

			o := map[string]interface{}{}

			if err := json.Unmarshal(res, &o); err != nil {
				t.Error(err)
			}

			if VERBOSE {
				pretty(t, o)
			}
		})
	}
}

func TestGetMovieArtResponse(t *testing.T) {
	var (
		secret string
	)

	if v, found := os.LookupEnv("FANART_API_KEY"); found {
		secret = v
	} else {
		t.Skip("no valid api key found, set one with 'FANART_API_KEY'")
	}

	conf := strings.NewReader(fmt.Sprintf(`{
		"api": {
			"key": "%s",
			"endpoint": "http://webservice.fanart.tv/v3"
		}
	}`, secret))

	client, err := fanart.New(conf, 10)
	if err != nil {
		t.Fatal(err)
	}

	var testcases = []struct {
		label string
		req   fanart.MovieRequest
	}{
		{"imdbId/hook", fanart.MovieRequest{MovieID: "tt0102057"}},
		{"imdbId/gone_with_the_wind", fanart.MovieRequest{MovieID: "tt0031381"}},
		{"imdbId/no_such_resource", fanart.MovieRequest{MovieID: "tt3097934"}},
	}

	for _, testcase := range testcases {
		t.Run(testcase.label, func(tt *testing.T) {
			res, err := client.MovieImages(testcase.req)
			if err != nil {
				if errors.Is(err, fanart.ErrNoSuchResource) {
					return
				}
				t.Error(err)
			}

			if VERBOSE {
				pretty(t, res)
			}
		})
	}
}
