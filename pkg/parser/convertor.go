package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Convertor struct {
	autoDocConf  DocConf
	proxyConfigs map[string]ProxySettings
	localSchema  *LocalSchema
	remoteSchema *RemoteSchema
}

func NewConvertor(conf DocConf) *Convertor {
	remote, err := NewRemoteSchema(conf.SchemaUrl)
	if err != nil {
		log.Fatal(err)
	}
	c := Convertor{
		proxyConfigs: ProxyConfigs,
		remoteSchema: remote,
		autoDocConf:  conf,
	}
	return &c
}

func (c *Convertor) RunConversion() {
	c.convertToProxyConf()
	fileName, err := c.makeFileName("rest_proxy_conf")
	if err != nil {
		log.Fatalf("Error of making name: %s", err)
	}
	c.write("configs", c.proxyConfigs, fileName)
	c.convertToLocalSchema()
}

func (c *Convertor) convertToProxyConf() {
	log.Println("START of schema conversion")
	for k, v := range c.remoteSchema.Paths {
		for m, value := range v {
			NewProxySettings(c.autoDocConf.RSchema, c.autoDocConf.RUrl, c.autoDocConf.RPort, c.remoteSchema.Info.Title, k, m, value, c.remoteSchema.Components.Schemas, c.autoDocConf)
		}
	}
	log.Println("SUCCESS schema conversion")
}

func (c *Convertor) convertToLocalSchema() {
	for k, v := range c.remoteSchema.Paths {
		for m, value := range v {
			NewLocalSchema(c.remoteSchema.Openapi, c.remoteSchema.Info, c.remoteSchema.Servers, c.remoteSchema.SecureSchemes, k, value, m, c.remoteSchema.Components.Schemas, c.autoDocConf)
		}
	}
}

func (c *Convertor) makeFileName(prefix string) (string, error) {
	fileName := c.remoteSchema.Info.Title
	regex, err := regexp.Compile("[^\\w\\s]")
	if err != nil {
		return "", err
	}
	fileName = regex.ReplaceAllString(fileName, " ")
	fileName = strings.TrimSpace(fileName)
	fileName = strings.ReplaceAll(fileName, " ", "_")
	fileName = strings.ToLower(fileName)
	return fmt.Sprintf("%s_%s", prefix, fileName), nil
}

func (c *Convertor) write(path string, data any, fileName string) {
	b, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		log.Fatalf("Marshal proxy configs: %s", err)
	}
	err = os.WriteFile(fmt.Sprintf("%s/%s.json", path, fileName), b, 0666)
	if err != nil {
		log.Fatalf("Write file: %s", err)
	}
}
