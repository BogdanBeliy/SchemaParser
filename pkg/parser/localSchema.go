package parser

type LocalSchema struct {
	Schema
}

func NewLocalSchema(path string) *LocalSchema {
	schema := LocalSchema{}
	return &schema
}
