package parser

import (
	"fmt"
	"strings"
)

var ProxyConfigs map[string]ProxySettings = make(map[string]ProxySettings)

type ProxySettings struct {
	ServiceName     string   `json:"service_name"`
	Schema          string   `json:"schema"`
	Host            string   `json:"host"`
	Port            string   `json:"port"`
	URL             string   `json:"url"`
	Method          string   `json:"method"`
	RPC             string   `json:"rpc"`
	PathParams      []string `json:"path_params"`
	BodyParams      []string `json:"body_params"`
	QueryParams     []string `json:"query_params"`
	Protected       bool     `json:"protected"`
	TimeoutDuration int      `json:"timeout_duration,omitempty"`
}

func NewProxySettings(schema, host, port, serviceName string, path url, m method, methodData methodItem, schemasData schemas) {
	p := ProxySettings{
		Schema: schema,
		Host:   host,
		Port:   port,
		Method: string(m),
	}
	p.makeUrl(path)
	p.makeRpcName(path)
	p.makeQueryAndPathParams(methodData)
	p.makeBodyParams(methodData, schemasData)
	p.checkSecurity(methodData)
	p.checkDuplicate(ProxyConfigs)
	ProxyConfigs[p.RPC] = p
}

func (p *ProxySettings) checkDuplicate(memData map[string]ProxySettings) {
	_, ok := memData[p.RPC]
	if !ok {
		return
	}
	p.RPC = fmt.Sprintf("%s%s", p.RPC, "_Duplicate")
}

func (p *ProxySettings) makeUrl(pathKey url) {
	path := strings.ReplaceAll(string(pathKey), "}", "")
	path = strings.ReplaceAll(path, "{", ":")
	p.URL = path
}

func (p *ProxySettings) makeRpcName(pathKey url) {
	path := strings.ReplaceAll(string(pathKey), "}", "")
	path = strings.ReplaceAll(path, "{", ":")
	path = strings.ReplaceAll(path, ":id", "by-id")
	path = strings.ReplaceAll(path, "-", "/")
	rpcSplit := strings.Split(path, "/")
	rpcTitle := make([]string, len(rpcSplit))
	for _, v := range rpcSplit {
		rpcTitle = append(rpcTitle, strings.Title(v))
	}
	p.RPC = fmt.Sprintf("%s%s", strings.Title(p.Method), strings.Join(rpcTitle, ""))
}

func (p *ProxySettings) makeQueryAndPathParams(methodData methodItem) {
	queryParams := make([]string, 0)
	pathParams := make([]string, 0)
	for _, v := range methodData.Parameters {
		if v.In == "query" {
			queryParams = append(queryParams, v.Name)
		} else if v.In == "path" {
			pathParams = append(pathParams, v.Name)
		}
	}
	p.QueryParams = queryParams
	p.PathParams = pathParams
}

func (p *ProxySettings) makeBodyParams(methodData methodItem, schemasData schemas) {
	if methodData.RequestBody == nil {
		return
	}
	// TODO нужно переписать на структуры, слищко много ифов
	bodyParams := make([]string, 0)
	if schema, ok := methodData.RequestBody["content"]; ok {
		if appJson, ok := schema.(map[string]interface{})["application/json"]; ok {
			if schemaPath, ok := appJson.(map[string]interface{})["schema"]; ok {
				if ref, ok := schemaPath.(map[string]interface{})["$ref"]; ok {
					refSplit := strings.Split(fmt.Sprintf("%v", ref), "/")
					ref = refSplit[len(refSplit)-1]
					s, ok := schemasData[refSplit[len(refSplit)-1]]
					if ok {
						for k, _ := range s.Properties {
							bodyParams = append(bodyParams, string(k))
						}
						p.BodyParams = bodyParams
					}
				}
			}
		}
	}
	p.BodyParams = bodyParams
}

func (p *ProxySettings) checkSecurity(methodData methodItem) {
	if len(methodData.Security) == 0 {
		p.Protected = false
		return
	}
	p.Protected = true

}
