package main

import (
	"embed"

	"github.com/noemacias/beacon/cmd"
)

//go:embed templates
var templates embed.FS

func main() {

	cmd.Execute(templates)

}
