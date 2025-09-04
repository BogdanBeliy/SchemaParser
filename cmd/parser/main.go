package main

import (
	"log"

	"github.com/BogdanBeliy/SchemaParser/pkg/parser"
)

func main() {
	_, err := parser.NewRemoteSchema("http://localhost:8000/api/schema/?format=json")
	if err != nil {
		log.Fatal(err)
	}
}
