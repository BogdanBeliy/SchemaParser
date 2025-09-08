package parser

import (
	"fmt"
	"log"
)

type Convertor struct {
	autoDocConf  DocConf
	proxyConf    *ProxyConfigs
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

func (c *Convertor) RunConvertation() {
	c.paths()
}

func (c *Convertor) paths() paths {
	path := c.remoteSchema.Paths
	for k := range path {
		// TODO продолжить с разбором путей
		fmt.Println(k)
	}
	return nil
}

// func (c *Convertor) servers() []server {
// 	servers := c.remoteSchema.Servers
// 	_ = servers
// 	return nil
// }

// type SchemaInterface interface {
// 	GetInfo() info
// 	GetServers() []server
// 	GetPaths() paths
// 	GetComponents() components
// 	GetSecureSchemes() secureSchemas
// }
