package parser

import (
	"encoding/json"
	"log"
	"os"
)

type Convertor struct {
	autoDocConf  DocConf
	localSchema  *LocalSchema
	remoteSchema *RemoteSchema
}

func NewConvertor(conf DocConf) *Convertor {
	remote, err := NewRemoteSchema(conf.SchemaUrl)
	if err != nil {
		log.Fatal(err)
	}
	c := Convertor{
		remoteSchema: remote,
		autoDocConf:  conf,
	}
	return &c
}

func (c *Convertor) RunConversion() {
	log.Println("start of schema conversion")
	c.convertToProxyConf()
	log.Println("end of schema conversion: SUCCESS")

	b, err := json.MarshalIndent(ProxyConfigs, "", "   ")
	if err != nil {
		log.Fatalf("Marshal proxy configs: %s", err)
	}
	err = os.WriteFile("rpc_conf.json", b, 0666)
	if err != nil {
		log.Fatalf("Write file: %s", err)
	}
}

func (c *Convertor) convertToProxyConf() paths {
	path := c.remoteSchema.Paths
	for k, v := range path {
		for m, value := range v {
			//c.createProxyConf(k, m, value, c.remoteSchema.Info.Title)
			NewProxySettings(c.autoDocConf.RSchema, c.autoDocConf.RUrl, c.autoDocConf.RPort, c.remoteSchema.Info.Title, k, m, value, c.remoteSchema.Components.Schemas)
		}
	}
	return nil
}

func (c *Convertor) createProxyConf(pathKey url, m method, methodData methodItem, serviceName string) {
	NewProxySettings(c.autoDocConf.RSchema, c.autoDocConf.RUrl, c.autoDocConf.RPort, serviceName, pathKey, m, methodData, c.remoteSchema.Components.Schemas)
}
