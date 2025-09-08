package parser

type ProxyConf []ProxySettings

type ProxySettings struct {
	Schema          string   `json:"schema"`
	Host            string   `json:"host"`
	Port            string   `json:"port"`
	URL             string   `json:"url"`
	Method          string   `json:"method"`
	RPC             string   `json:"rpc"`
	PathParams      []string `mapstructure:"path_params"`
	BodyParams      []string `mapstructure:"body_params"`
	QueryParams     []string `mapstructure:"query_params"`
	Protected       bool     `json:"protected"`
	TimeoutDuration int      `mapstructure:"timeout_duration,omitempty"`
}
