package main

import (
	"embed"

	"github.com/ahmedsat/gw-engine/apps"
	"github.com/ahmedsat/gw-engine/log"
	vkngmath "github.com/vkngwrapper/math"
)

//go:embed shaders
var shaders embed.FS

var (
	red   = vkngmath.Vec3[float32]{X: 1, Y: 0, Z: 0}
	green = vkngmath.Vec3[float32]{X: 0, Y: 1, Z: 0}
	blue  = vkngmath.Vec3[float32]{X: 0, Y: 0, Z: 1}
	yalow = vkngmath.Vec3[float32]{X: 1, Y: 1, Z: 0}
)

var indices = []uint16{0, 1, 2, 2, 3, 0}

var vertices = []apps.Vertex{
	{Position: vkngmath.Vec2[float32]{X: -0.5, Y: -0.5}, Color: red},
	{Position: vkngmath.Vec2[float32]{X: 0.5, Y: -0.5}, Color: green},
	{Position: vkngmath.Vec2[float32]{X: 0.5, Y: 0.5}, Color: blue},
	{Position: vkngmath.Vec2[float32]{X: -0.5, Y: 0.5}, Color: yalow},
}

func main() {
	app := &apps.HelloTriangleApplication{
		Shaders:  shaders,
		Vertices: vertices,
		Indices:  indices,
	}

	err := app.Run()
	if err != nil {
		log.Error.Panicf("%+v\n", err)
	}
}
