package main

import (
	"log"

	"github.com/BogdanBeliy/SchemaParser/pkg/parser"
)

func main() {
	_, err := parser.NewRemoteSchema("https://flow-stage.intechs.by/api/schema/?format=json")
	if err != nil {
		log.Fatal(err)
	}
}
