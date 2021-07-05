package main

import (
	mat4 "humangl/matrice"
)

func LeftLeg (body [10]BodyConfig) Node {
	leftLowerLeg := Node {
		bodyPart: LLLEG,
		children: nil,
		transform: mat4.Translate(0, body[LULEG].size.y, 0),
	}

	leftUpperLeg := Node {
		bodyPart: LULEG,
		children: []Node{ leftLowerLeg },
	}
	leftUpperLeg.transform = mat4.Translate(
		-(body[TORSO].size.x * .5 - body[LULEG].size.x * .5), 
		0, 
		body[TORSO].size.z * .5)
	leftUpperLeg.transform = leftUpperLeg.transform.Mult(
		mat4.Rotation(
			body[LULEG].rotation.x, 
			0, 
			0))

	return leftUpperLeg
}

func RightLeg (body [10]BodyConfig) Node {
	rightLowerLeg := Node {
		bodyPart: RLLEG,
		children: nil,
		transform: mat4.Translate(0, body[RULEG].size.y, 0),
	}

	rightUpperLeg := Node {
		bodyPart: RULEG,
		children: []Node{ rightLowerLeg },
	}
	rightUpperLeg.transform = mat4.Translate(
		(body[TORSO].size.x * .5 - body[RULEG].size.x * .5), 
		0, 
		body[TORSO].size.z * .5)
	rightUpperLeg.transform = rightUpperLeg.transform.Mult(
		mat4.Rotation(
			body[RULEG].rotation.x, 
			0, 
			0))

	return rightUpperLeg
}

func leftArm (body [10]BodyConfig) Node {
	
	matLeftLowerArm := mat4.Translate(0, body[LUARM].size.y, 0)
	matLeftLowerArm = matLeftLowerArm.Mult(mat4.Rotation(body[LLARM].rotation.x, 0, 0))
	leftLowerArm := Node{
		transform: matLeftLowerArm,
		bodyPart: LLARM,
		children: nil,
	}

	matLeftUpperArm := mat4.Translate(-(body[TORSO].size.x * .5 + body[LUARM].size.x * .5), body[TORSO].size.y * 0.95, body[TORSO].size.z * .5)
	matLeftUpperArm = matLeftUpperArm.Mult(mat4.Rotation(body[LUARM].rotation.x, body[LUARM].rotation.y, body[LUARM].rotation.z))
	leftUpperArm := Node{
		transform: matLeftUpperArm,
		bodyPart: LUARM,
		children: []Node{leftLowerArm},
	}

	return leftUpperArm
}

func rightArm (body [10]BodyConfig) Node {
	matRightLowerArm := mat4.Translate(0, body[RUARM].size.y, 0)
	matRightLowerArm = matRightLowerArm.Mult(mat4.Rotation(body[RLARM].rotation.x, 0, 0))
	rightLowerArm := Node{
		transform: matRightLowerArm,
		bodyPart: RLARM,
		children: nil,
	}

	matRightUpperArm := mat4.Translate(body[TORSO].size.x * 0.5 + body[RUARM].size.x * .5, body[TORSO].size.y * 0.95, body[TORSO].size.z * .5)
	matRightUpperArm = matRightUpperArm.Mult(mat4.Rotation(body[RUARM].rotation.x, 0, 0))
	rightUpperArm := Node{
		transform: matRightUpperArm,
		bodyPart: RUARM,
		children: []Node{rightLowerArm},
	}

	return rightUpperArm
}

func GenerateHuman(body [10]BodyConfig) *Node {
	leftLeg := LeftLeg(body)
	rightLeg := RightLeg(body)
	rightArm := rightArm(body)
	LeftArm := leftArm(body)

	matHead := mat4.Translate(0, body[TORSO].size.y, 0)
	head := Node {
		transform: matHead,
		bodyPart: HEAD,
		children: nil,
	}

	matTorso := mat4.Rotation(body[TORSO].rotation.x, body[TORSO].rotation.y, body[TORSO].rotation.z)
	torso := Node {
		transform: matTorso,
		bodyPart: TORSO,
		children: []Node{ head, rightArm, LeftArm, leftLeg, rightLeg },
	}

	return &torso
}

func HumanDefaultColor() [10]Vec3f32 {
	var bodyColors [10]Vec3f32

	for x := 0; x < 10; x++ {
		bodyColors[x] = Vec3f32{1.0, 1.0, 1.0}
	}

	bodyColors[TORSO] = Vec3f32{0.4, 0.0, 0.0}
	bodyColors[HEAD] = Vec3f32{0.5, 0.1, 0.0}

	return bodyColors
}

func HumanDefaultConfig() [10]BodyConfig {
	var bodyConfig [10]BodyConfig

	for x := 0; x < 10; x++ {
		bodyConfig[x] = BodyConfig{
			size: Vec3f32{1, 1, 1},
			rotation: Vec3f32{0, 0, 0},
		}
	}

	// Head
	bodyConfig[HEAD].size.x = 2.0
	bodyConfig[HEAD].size.y = 2.0
	bodyConfig[HEAD].size.z = 2.0

	// Torso
	bodyConfig[TORSO].size.x = 4.0
	bodyConfig[TORSO].size.y = 6.0
	bodyConfig[TORSO].size.z = 2.0
	
	// Arms 
	bodyConfig[RUARM].size.y = 4.0
	bodyConfig[RLARM].size.y = 2.0
	bodyConfig[RUARM].rotation.x = mat4.DegToRad(180)

	bodyConfig[LUARM].size.y = 4.0
	bodyConfig[LLARM].size.y = 2.0
	bodyConfig[LUARM].rotation.x = mat4.DegToRad(180)

	// Legs
	bodyConfig[RULEG].rotation.x = mat4.DegToRad(180)
	bodyConfig[RULEG].size.x = 1.0
	bodyConfig[RULEG].size.y = 4.0
	bodyConfig[RLLEG].size.y = 3.0

	bodyConfig[LULEG].rotation.x = mat4.DegToRad(180)
	bodyConfig[LULEG].size.x = 1.0
	bodyConfig[LULEG].size.y = 4.0
	bodyConfig[LLLEG].size.y = 3.0

	return bodyConfig
}