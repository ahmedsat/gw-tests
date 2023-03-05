package main

import (
	"embed"

	"github.com/ahmedsat/gw-engine/apps"
	"github.com/ahmedsat/gw-engine/log"
)

//go:embed shaders
var shaders embed.FS

var (
	red   = apps.Color{R: 1}
	green = apps.Color{G: 1}
	blue  = apps.Color{B: 1}
)

var vertices = []apps.Vertex{
	{Vec2: apps.Vec2{X: 0.0, Y: -0.8}, Color: red},
	{Vec2: apps.Vec2{X: 0.8, Y: 0.8}, Color: green},
	{Vec2: apps.Vec2{X: -0.8, Y: 0.8}, Color: blue},

	{Vec2: apps.Vec2{X: -0.5, Y: -0.5}, Color: red},
	{Vec2: apps.Vec2{X: 0.5, Y: -0.5}, Color: green},
	{Vec2: apps.Vec2{X: 0.5, Y: 0.5}, Color: blue},

	{Vec2: apps.Vec2{X: 0.5, Y: 0.5}, Color: blue},
	{Vec2: apps.Vec2{X: -0.5, Y: 0.5}, Color: green},
	{Vec2: apps.Vec2{X: -0.5, Y: -0.5}, Color: red},
}

func main() {
	app := &apps.HelloTriangleApplication{
		Shaders:  shaders,
		Vertices: vertices,
	}

	err := app.Run()
	if err != nil {
		log.Error.Panicf("%+v\n", err)
	}
}
