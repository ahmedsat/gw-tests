package main

import (
	"embed"

	"github.com/ahmedsat/go-math-helper/math32/colors"
	"github.com/ahmedsat/gw-engine/apps"
	"github.com/ahmedsat/gw-engine/log"
	vkngmath "github.com/vkngwrapper/math"
)

//go:embed shaders
var shaders embed.FS

var (
	red   = colors.Color{R: 1}
	green = colors.Color{G: 1}
	blue  = colors.Color{B: 1}
	yalow = colors.Color{R: 1, G: 1}
)

var indices = []uint16{0, 1, 2, 2, 3, 0}

var vertices = []apps.Vertex{
	{Position: vkngmath.Vec2[float32]{X: -0.5, Y: -0.5}, Color: vkngmath.Vec3[float32]{X: 1, Y: 0, Z: 0}},
	{Position: vkngmath.Vec2[float32]{X: 0.5, Y: -0.5}, Color: vkngmath.Vec3[float32]{X: 0, Y: 1, Z: 0}},
	{Position: vkngmath.Vec2[float32]{X: 0.5, Y: 0.5}, Color: vkngmath.Vec3[float32]{X: 0, Y: 0, Z: 1}},
	{Position: vkngmath.Vec2[float32]{X: -0.5, Y: 0.5}, Color: vkngmath.Vec3[float32]{X: 1, Y: 1, Z: 1}},
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
