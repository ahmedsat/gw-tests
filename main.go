package main

import (
	"embed"

	"github.com/ahmedsat/gw-engine/apps"
	"github.com/ahmedsat/gw-engine/log"
)

//go:embed shaders
var shaders embed.FS

func main() {
	app := &apps.HelloTriangleApplication{
		Shaders: shaders,
	}

	err := app.Run()
	if err != nil {
		log.Error.Panicf("%+v\n", err)
	}
}
