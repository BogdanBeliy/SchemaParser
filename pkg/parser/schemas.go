package parser

import (
	"encoding/json"
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
		return nil, err
	}

	defer file.Close()
	defer resp.Body.Close()
	return &schema, nil
}

type (
	url        string
	method     string
	statusCode string
	authType   string
	refName    string
	refField   string
)

type info struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,"`
}
type server struct {
	Url         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

type components struct {
	Schemas         schemas       `json:"schemas,omitempty"`
	Parameters      parameters    `json:"parameters,omitempty"`
	SecuritySchemes secureSchemas `json:"securitySchemes,omitempty"`
}

type parameters map[string]parametersItem
type parametersItem struct {
	Name            string                     `json:"name,omitempty"`
	In              string                     `json:"in,omitempty"`
	Description     string                     `json:"description,omitempty"`
	Required        string                     `json:"required,omitempty"`
	ParameterSchema map[string]interface{}     `json:"schema,omitempty"`
	Security        []map[authType]interface{} `json:"security,omitempty"`
}

type paths map[url]map[method]methodItem
type methodItem struct {
	Summary     string                                `json:"summary,omitempty"`
	Description string                                `json:"description,omitempty"`
	OperationId string                                `json:"operationId,omitempty"`
	Tags        []string                              `json:"tags,omitempty"`
	Parameters  []methodParametersItem                `json:"parameters,omitempty"`
	Responses   map[statusCode]map[string]interface{} `json:"responses,omitempty"`
	RequestBody map[string]interface{}                `json:"requestBody,omitempty"`
	Security    []map[string]interface{}              `json:"security,omitempty"`
}

type methodParametersItem struct {
	parametersItem
}

type schemas map[string]refItem
type refItem struct {
	Type        string                   `json:"type,omitempty"`
	Properties  map[refField]interface{} `json:"properties,omitempty"`
	Required    []string                 `json:"required,omitempty"`
	Description string                   `json:"description,omitempty"`
	Enum        []string                 `json:"enum,omitempty"`
}

type secureSchemas map[string]map[string]string

type Schema struct {
	Openapi       string        `json:"openapi,omitempty"`
	Info          info          `json:"info,omitempty"`
	Servers       []server      `json:"servers,omitempty"`
	Paths         paths         `json:"paths,omitempty"`
	Components    components    `json:"components,omitempty"`
	SecureSchemes secureSchemas `json:"securitySchemes,omitempty"`
}
