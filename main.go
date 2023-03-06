package main

import (
	"embed"

	"github.com/ahmedsat/go-math-helper/math32/colors"
	"github.com/ahmedsat/go-math-helper/math32/vectors"
	"github.com/ahmedsat/gw-engine/apps"
	"github.com/ahmedsat/gw-engine/log"
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
	// {Vec3: vectors.Vec3{X: 0.0, Y: -0.8}, Color: red},
	// {Vec3: vectors.Vec3{X: 0.8, Y: 0.8}, Color: green},
	// {Vec3: vectors.Vec3{X: -0.8, Y: 0.8}, Color: blue},
	{Vec3: vectors.Vec3{X: -0.5, Y: -0.5}, Color: red},
	{Vec3: vectors.Vec3{X: 0.5, Y: -0.5}, Color: green},
	{Vec3: vectors.Vec3{X: 0.5, Y: 0.5}, Color: blue},
	// {Vec3: vectors.Vec3{X: 0.5, Y: 0.5}, Color: blue},
	{Vec3: vectors.Vec3{X: -0.5, Y: 0.5}, Color: yalow},
	// {Vec3: vectors.Vec3{X: -0.5, Y: -0.5}, Color: red},
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
