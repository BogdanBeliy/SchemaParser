package parser

import (
	"encoding/json"
	"io"
	"net/http"
)

type RemoteSchema struct {
	Schema
}

func NewRemoteSchema(path string) (*RemoteSchema, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	schema := RemoteSchema{}
	b, _ := io.ReadAll(resp.Body)
	json.Unmarshal(b, &schema)
	defer resp.Body.Close()
	return &schema, nil
}
