package main

import mat4 "humangl/matrice"

var (
	triangle = []float32{
		-0.5,-0.5,-0.5,
		-0.5,-0.5, 0.5,
		-0.5, 0.5, 0.5,
		0.5, 0.5,-0.5, 
		-0.5,-0.5,-0.5,
		-0.5, 0.5,-0.5,
		0.5,-0.5, 0.5,
		-0.5,-0.5,-0.5,
		0.5,-0.5,-0.5,
		0.5, 0.5,-0.5,
		0.5,-0.5,-0.5,
		-0.5,-0.5,-0.5,
		-0.5,-0.5,-0.5,
		-0.5, 0.5, 0.5,
		-0.5, 0.5,-0.5,
		0.5,-0.5, 0.5,
		-0.5,-0.5, 0.5,
		-0.5,-0.5,-0.5,
		-0.5, 0.5, 0.5,
		-0.5,-0.5, 0.5,
		0.5,-0.5, 0.5,
		0.5, 0.5, 0.5,
		0.5,-0.5,-0.5,
		0.5, 0.5,-0.5,
		0.5,-0.5,-0.5,
		0.5, 0.5, 0.5,
		0.5,-0.5, 0.5,
		0.5, 0.5, 0.5,
		0.5, 0.5,-0.5,
		-0.5, 0.5,-0.5,
		0.5, 0.5, 0.5,
		-0.5, 0.5,-0.5,
		-0.5, 0.5, 0.5,
		0.5, 0.5, 0.5,
		-0.5, 0.5, 0.5,
		0.5,-0.5, 0.5,
	}
)

const (
	TORSO = iota
	HEAD  = iota
	LUARM = iota
	RUARM = iota
	LLARM = iota
	RLARM = iota
	LULEG = iota
	RULEG = iota
	LLLEG = iota
	RLLEG = iota
)

type Vec3f32 struct {
	x float32
	y float32
	z float32
}

type BodyConfig struct {
	size Vec3f32
	rotation Vec3f32
}

type Node struct {
	transform mat4.Mat4
	children []Node
	bodyPart int
}

type DrawData struct {
	bodyConfig [10]BodyConfig
	bodyColors [10]Vec3f32
	uniformColor int32
	uniformModel int32
}
