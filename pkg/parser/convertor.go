package parser

import "fmt"

type Convertor struct {
	proxyConf    ProxyConf
	localSchema  LocalSchema
	remoteSchema RemoteSchema
}

func NewConvertor(path string) {
	_, err := LoadDocConfig("configs/auto_mode_conf.json")
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Convertor) RunConvertation() {
	c.paths()
}

func (c *Convertor) paths() {
	paths := c.remoteSchema.Paths
	fmt.Println(paths)
}
