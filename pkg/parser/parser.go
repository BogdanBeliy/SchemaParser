package parser

type Parser interface {
	Set()
	Get()
}

type info struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,"`
}

type paths struct{}
type schemas struct{}
type parameters struct{}
type secureSchemas struct{}
type servers struct{}

type Schema struct {
	Openapi       string        `json:"openapi,omitempty"`
	Info          info          `json:"info,omitempty"`
	Paths         paths         `json:"paths,omitempty"`
	Schemas       schemas       `json:"schemas,omitempty"`
	Parameters    parameters    `json:"parameters,omitempty"`
	SecureSchemes secureSchemas `json:"securitySchemes,omitempty"`
	Servers       servers       `json:"servers,omitempty"`
}

func (rs *Schema) Get(key string) *Schema {
	return rs
}

func (rs *Schema) Set(key string, value interface{}) *Schema {
	return rs
}
