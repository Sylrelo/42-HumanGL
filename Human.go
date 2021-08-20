package main

import (
	mat4 "humangl/matrice"
)

func LeftLeg(body [10]BodyConfig, bodyTmp [10]BodyConfig) Node {
	leftLowerLeg := Node{
		bodyPart:  LLLEG,
		children:  nil,
		transform: mat4.Translate(0, body[LULEG].size.y, 0),
	}
	leftLowerLeg.transform = leftLowerLeg.transform.Mult(mat4.Rotation(
		body[LLLEG].rotation.x+bodyTmp[LLLEG].rotation.x,
		body[LLLEG].rotation.y+bodyTmp[LLLEG].rotation.y,
		body[LLLEG].rotation.z+bodyTmp[LLLEG].rotation.z))

	leftUpperLeg := Node{
		bodyPart: LULEG,
		children: []Node{leftLowerLeg},
	}
	leftUpperLeg.transform = mat4.Translate(
		-(body[TORSO].size.x*.5 - body[LULEG].size.x*.5),
		0,
		body[TORSO].size.z*.5)
	leftUpperLeg.transform = leftUpperLeg.transform.Mult(
		mat4.Rotation(
			body[LULEG].rotation.x+bodyTmp[LULEG].rotation.x,
			0,
			0))

	return leftUpperLeg
}

func RightLeg(body [10]BodyConfig, bodyTmp [10]BodyConfig) Node {
	rightLowerLeg := Node{
		bodyPart:  RLLEG,
		children:  nil,
		transform: mat4.Translate(0, body[RULEG].size.y, 0),
	}
	rightLowerLeg.transform = rightLowerLeg.transform.Mult(mat4.Rotation(
		body[RLLEG].rotation.x+bodyTmp[RLLEG].rotation.x,
		body[RLLEG].rotation.y+bodyTmp[RLLEG].rotation.y,
		body[RLLEG].rotation.z+bodyTmp[RLLEG].rotation.z))

	rightUpperLeg := Node{
		bodyPart: RULEG,
		children: []Node{rightLowerLeg},
	}
	rightUpperLeg.transform = mat4.Translate(
		(body[TORSO].size.x*.5 - body[RULEG].size.x*.5),
		0,
		body[TORSO].size.z*.5)
	rightUpperLeg.transform = rightUpperLeg.transform.Mult(
		mat4.Rotation(
			body[RULEG].rotation.x+bodyTmp[RULEG].rotation.x,
			0,
			0))

	return rightUpperLeg
}

func leftArm(body [10]BodyConfig, bodyTmp [10]BodyConfig) Node {

	matLeftLowerArm := mat4.Translate(0, body[LUARM].size.y, 0)
	matLeftLowerArm = matLeftLowerArm.Mult(mat4.Rotation(
		body[LLARM].rotation.x+bodyTmp[LLARM].rotation.x,
		body[LLARM].rotation.y+bodyTmp[LLARM].rotation.y,
		body[LLARM].rotation.z+bodyTmp[LLARM].rotation.z,
	))
	leftLowerArm := Node{
		transform: matLeftLowerArm,
		bodyPart:  LLARM,
		children:  nil,
	}

	matLeftUpperArm := mat4.Translate(-(body[TORSO].size.x*.5 + body[LUARM].size.x*.5), body[TORSO].size.y*0.95, body[TORSO].size.z*.5)
	matLeftUpperArm = matLeftUpperArm.Mult(mat4.Rotation(
		body[LUARM].rotation.x+bodyTmp[LUARM].rotation.x,
		body[LUARM].rotation.y+bodyTmp[LUARM].rotation.y,
		body[LUARM].rotation.z+bodyTmp[LUARM].rotation.z,
	))
	leftUpperArm := Node{
		transform: matLeftUpperArm,
		bodyPart:  LUARM,
		children:  []Node{leftLowerArm},
	}

	return leftUpperArm
}

func rightArm(body [10]BodyConfig, bodyTmp [10]BodyConfig) Node {
	matRightLowerArm := mat4.Translate(0, body[RUARM].size.y, 0)
	matRightLowerArm = matRightLowerArm.Mult(mat4.Rotation(
		body[RLARM].rotation.x+bodyTmp[RLARM].rotation.x,
		body[RLARM].rotation.y+bodyTmp[RLARM].rotation.y,
		body[RLARM].rotation.z+bodyTmp[RLARM].rotation.z,
	))
	rightLowerArm := Node{
		transform: matRightLowerArm,
		bodyPart:  RLARM,
		children:  nil,
	}

	matRightUpperArm := mat4.Translate(body[TORSO].size.x*0.5+body[RUARM].size.x*.5, body[TORSO].size.y*0.95, body[TORSO].size.z*.5)
	matRightUpperArm = matRightUpperArm.Mult(mat4.Rotation(
		body[RUARM].rotation.x+bodyTmp[RUARM].rotation.x,
		body[RUARM].rotation.y+bodyTmp[RUARM].rotation.y,
		body[RUARM].rotation.z+bodyTmp[RUARM].rotation.z,
	))
	rightUpperArm := Node{
		transform: matRightUpperArm,
		bodyPart:  RUARM,
		children:  []Node{rightLowerArm},
	}

	return rightUpperArm
}

func GenerateHuman(body [10]BodyConfig, bodyTmp [10]BodyConfig) *Node {
	leftLeg := LeftLeg(body, bodyTmp)
	rightLeg := RightLeg(body, bodyTmp)
	rightArm := rightArm(body, bodyTmp)
	LeftArm := leftArm(body, bodyTmp)

	matHead := mat4.Translate(0, body[TORSO].size.y, body[TORSO].size.z*.5)
	matHead = matHead.Mult(mat4.Rotation(
		body[HEAD].rotation.x+bodyTmp[HEAD].rotation.x,
		body[HEAD].rotation.y+bodyTmp[HEAD].rotation.y,
		body[HEAD].rotation.z+bodyTmp[HEAD].rotation.z,
	))
	head := Node{
		transform: matHead,
		bodyPart:  HEAD,
		children:  nil,
	}

	matTorso := mat4.Rotation(
		body[TORSO].rotation.x+bodyTmp[TORSO].rotation.x,
		body[TORSO].rotation.y+bodyTmp[TORSO].rotation.y,
		body[TORSO].rotation.z+bodyTmp[TORSO].rotation.z,
	)
	// matTorso = matTorso.Mult(mat4.Translate(0, 0, 1))

	torso := Node{
		transform: matTorso,
		bodyPart:  TORSO,
		children:  []Node{head, rightArm, LeftArm, leftLeg, rightLeg},
	}

	return &torso
}

func HumanDefaultColor() [10]Vec3f32 {
	var bodyColors [10]Vec3f32

	for x := 0; x < 10; x++ {
		bodyColors[x] = Vec3f32{1.0, 1.0, 1.0}
	}

	bodyColors[TORSO] = Vec3f32{0.75, 0.75, 0.75}
	bodyColors[HEAD] = Vec3f32{0.5, 0.5, 0.5}

	bodyColors[RUARM] = Vec3f32{0.6, 0.6, 0.6}
	bodyColors[LUARM] = Vec3f32{0.625, 0.625, 0.625}

	bodyColors[LLARM] = Vec3f32{0.7, 0.7, 0.7}
	bodyColors[RLARM] = Vec3f32{0.725, 0.725, 0.725}

	bodyColors[RULEG] = Vec3f32{0.4, 0.4, 0.4}
	bodyColors[LULEG] = Vec3f32{0.425, 0.425, 0.425}
	bodyColors[LLLEG] = Vec3f32{0.8, 0.8, 0.8}
	bodyColors[RLLEG] = Vec3f32{0.825, 0.825, 0.825}

	return bodyColors
}

func SetToZero() [10]BodyConfig {
	var bodyConfig [10]BodyConfig
	for x := 0; x < 10; x++ {
		bodyConfig[x] = BodyConfig{
			size:     Vec3f32{0, 0, 0},
			rotation: Vec3f32{0, 0, 0},
		}
	}
	return bodyConfig
}

func HumanDefaultConfig() [10]BodyConfig {
	var bodyConfig [10]BodyConfig

	for x := 0; x < 10; x++ {
		bodyConfig[x] = BodyConfig{
			size:     Vec3f32{1, 1, 1},
			rotation: Vec3f32{0, 0, 0},
		}
	}

	// Head
	bodyConfig[HEAD].size.x = 2.0
	bodyConfig[HEAD].size.y = 2.0
	bodyConfig[HEAD].size.z = 2.0

	bodyConfig[TORSO].rotation.y = mat4.DegToRad(0)
	bodyConfig[TORSO].rotation.x = mat4.DegToRad(0)

	// Arms

	bodyConfig[RUARM].size = Vec3f32{1, 4, 1}
	bodyConfig[RLARM].size = Vec3f32{0.8, 4.5, 0.8}

	bodyConfig[LUARM].size = Vec3f32{1, 4, 1}
	bodyConfig[LLARM].size = Vec3f32{0.8, 4.5, 0.8}

	bodyConfig[RUARM].size.y = 4.0
	bodyConfig[RLARM].size.y = 2.0
	bodyConfig[RUARM].rotation.x = mat4.DegToRad(180)
	bodyConfig[LUARM].rotation.x = mat4.DegToRad(180)

	// Legs
	bodyConfig[RULEG].rotation.x = mat4.DegToRad(180)
	// bodyConfig[RLLEG].rotation.x = mat4.DegToRad(-30)

	bodyConfig[RULEG].size.x = 1.0
	bodyConfig[RULEG].size.y = 4.0
	bodyConfig[RLLEG].size.y = 3.0

	bodyConfig[LULEG].rotation.x = mat4.DegToRad(180)
	bodyConfig[LULEG].size.x = 1.0
	bodyConfig[LULEG].size.y = 4.0
	bodyConfig[LLLEG].size.y = 3.0

	bodyConfig[TORSO].size = Vec3f32{4, 5, 2.5}
	bodyConfig[RUARM].size = Vec3f32{1, 3, 1}
	bodyConfig[RLARM].size = Vec3f32{0.8, 4, 0.8}
	bodyConfig[LUARM].size = Vec3f32{1, 3, 1}
	bodyConfig[LLARM].size = Vec3f32{0.8, 4, 0.8}
	bodyConfig[LULEG].size = Vec3f32{1.8, 4, 1.8}
	bodyConfig[LLLEG].size = Vec3f32{1.5, 4, 1.5}
	bodyConfig[RULEG].size = Vec3f32{1.8, 4, 1.8}
	bodyConfig[RLLEG].size = Vec3f32{1.5, 4, 1.5}
	bodyConfig[HEAD].size = Vec3f32{1.5, 2, 1.5}
	return bodyConfig
}
