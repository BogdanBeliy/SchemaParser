package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

type DocConf struct {
	// Конфигурация для проксирвоания и парса схемы
	ServiceName         string `json:"service_name"`
	SchemaUrl           string `json:"schema_url"`
	RSchema             string `json:"r_schema"`
	RUrl                string `json:"r_url"`
	RPort               string `json:"r_port"`
	AddSchemaRaw        bool   `json:"add_raw_schema"`
	FilterPattern       string `json:"filter_pattern,omitempty"`
	VersionRegex        string `json:"version_regex,omitempty"`
	FirstVersion        string `json:"first_version,omitempty"`
	RaiseVersionMissing bool   `json:"raise_version_missing"`
	DelPrefixCount      int    `json:"del_prefix_count,omitempty"`
	Duplicates          bool   `json:"duplicates"`
}

func LoadDocConfig(path string) (*DocConf, error) {
	b, err := os.ReadFile(path)
	var conf DocConf
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, &conf)
	fmt.Println(conf)
}
