package fanart_test

import (
	"testing"

	"encoding/json"
)

func pretty(t *testing.T, o interface{}) {
	raw, err := json.MarshalIndent(o, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf(string(raw))
}

func TestCreateClient(t *testing.T) {
}

func TestQuery(t *testing.T) {
}

func TestSearch(t *testing.T) {
}
