package parser

type ParserConfig struct {
	LocalSchemaPath  string `json:""`
	RemoteSchemaPath string
	Host             string
	Port             string
}

func GetConfig() {}
