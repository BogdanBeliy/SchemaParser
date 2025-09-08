package parser

type Convertor struct {
	proxyConf    ProxyConf
	localSchema  LocalSchema
	remoteSchema RemoteSchema
}

func NewConvertor(remote RemoteSchema) {}

func (c *Convertor) RunConvertation() {
	c.paths()
}

func (c *Convertor) paths() {
}
