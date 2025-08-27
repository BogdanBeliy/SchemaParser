package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

	file, _ := os.Create("schema.json")
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(schema); err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()
	defer resp.Body.Close()
	return &schema, nil
}
