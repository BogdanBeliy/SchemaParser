package main

import (
	"log"

	"github.com/BogdanBeliy/SchemaParser/pkg/parser"
)

func init() {
	err := parser.LoadDocConfig("configs/auto_mode_conf.json")
	if err != nil {
		log.Fatalf("Ошибка чтения конфигурации: %s", err)
	}
}

func main() {
	for _, v := range parser.DocConfigs {
		p := parser.NewConvertor(v)
		p.RunConversion()
	}
}
