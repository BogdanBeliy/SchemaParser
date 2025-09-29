package parser

import (
	"fmt"
	"strings"
)

type LocalSchema struct {
	Schema
}

func NewLocalSchema(openapi string, i info, servers []server, sSchemes secureSchemas, pathKey url, methodData methodItem, m method, schemasData schemas, confs DocConf) *LocalSchema {
	l := LocalSchema{
		Schema{
			Openapi:       openapi,
			Info:          i,
			Servers:       servers,
			SecureSchemes: sSchemes,
		},
	}
	l.makeUrl(pathKey)
	return &l
}

func (l *LocalSchema) makePaths() {
	fmt.Println("Make paths")
}
func (l *LocalSchema) makeSchemas() {
	fmt.Println("Make schemas")
}

func (l *LocalSchema) makeUrl(pathKey url) {
	path := strings.ReplaceAll(string(pathKey), "}", "")
	path = strings.ReplaceAll(path, "{", ":")
	p := make(paths)
	_ = p
}
