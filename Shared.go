package main

import mat4 "humangl/matrice"

var (
	triangle = []float32{
		-0.5, -0.5, -0.5,
		-0.5, -0.5, 0.5,
		-0.5, 0.5, 0.5,
		0.5, 0.5, -0.5,
		-0.5, -0.5, -0.5,
		-0.5, 0.5, -0.5,
		0.5, -0.5, 0.5,
		-0.5, -0.5, -0.5,
		0.5, -0.5, -0.5,
		0.5, 0.5, -0.5,
		0.5, -0.5, -0.5,
		-0.5, -0.5, -0.5,
		-0.5, -0.5, -0.5,
		-0.5, 0.5, 0.5,
		-0.5, 0.5, -0.5,
		0.5, -0.5, 0.5,
		-0.5, -0.5, 0.5,
		-0.5, -0.5, -0.5,
		-0.5, 0.5, 0.5,
		-0.5, -0.5, 0.5,
		0.5, -0.5, 0.5,
		0.5, 0.5, 0.5,
		0.5, -0.5, -0.5,
		0.5, 0.5, -0.5,
		0.5, -0.5, -0.5,
		0.5, 0.5, 0.5,
		0.5, -0.5, 0.5,
		0.5, 0.5, 0.5,
		0.5, 0.5, -0.5,
		-0.5, 0.5, -0.5,
		0.5, 0.5, 0.5,
		-0.5, 0.5, -0.5,
		-0.5, 0.5, 0.5,
		0.5, 0.5, 0.5,
		-0.5, 0.5, 0.5,
		0.5, -0.5, 0.5,
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

const (
	width  = 1600
	height = 900

	vertexShaderSource = `
		#version 410
		in vec3 vp;
		uniform mat4 matProjection;
		uniform mat4 matCamera;
		uniform mat4 matModel;

		void main() {
			gl_Position = matProjection * matCamera * matModel * vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410
		out vec4 frag_colour;
		uniform vec3 partColor;

		void main() {
			frag_colour = vec4(partColor.xyz, 1.0);
		}
	` + "\x00"
)

type Vec3f32 struct {
	x float32
	y float32
	z float32
}

type BodyConfig struct {
	size     Vec3f32
	rotation Vec3f32
}

type Node struct {
	transform mat4.Mat4
	children  []Node
	bodyPart  int
}

type DrawData struct {
	bodyConfig       [10]BodyConfig
	bodyConfigTmp    [10]BodyConfig
	bodyColors       [10]Vec3f32
	uniformColor     int32
	uniformModel     int32
	selectedBodypart int
}

type AnimationDetail struct {
	rotation Vec3f32
	BodyPart int
	start    int
	end      int
}

type Animation struct {
	duration  float32
	keyframes []AnimationDetail
}
