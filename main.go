package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	mat4 "humangl/matrice"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func translateBackBodypart(modelMat mat4.Mat4, bodyConfig BodyConfig) mat4.Mat4 {
	instanceMat := modelMat.Mult(mat4.Translate(0, bodyConfig.size.y*0.5, 0))
	instanceMat = instanceMat.Mult(mat4.Scale(bodyConfig.size.x, bodyConfig.size.y, bodyConfig.size.z))
	return instanceMat
}

func iterateChildrens(drawData DrawData, node *Node, matModel mat4.Mat4) {
	if node == nil {
		return
	}
	matInstance := mat4.Identity()
	matModel = matModel.Mult(node.transform)

	if node.bodyPart == TORSO {
		matInstance = matModel.Mult(mat4.Translate(0, drawData.bodyConfig[TORSO].size.y*0.5, drawData.bodyConfig[TORSO].size.z*0.5))
		matInstance = matInstance.Mult(mat4.Scale(drawData.bodyConfig[TORSO].size.x, drawData.bodyConfig[TORSO].size.y, drawData.bodyConfig[TORSO].size.z))
	} else {
		matInstance = translateBackBodypart(matModel, drawData.bodyConfig[node.bodyPart])
	}

	matInstance = mat4.Translate(drawData.bodyTranslation.x, drawData.bodyTranslation.y, drawData.bodyTranslation.z).Mult(matInstance)
	gl.UniformMatrix4fv(drawData.uniformModel, 1, false, &matInstance[0][0])
	if drawData.selectedBodypart == node.bodyPart {
		gl.Uniform3f(drawData.uniformColor, 0.4, 0.4, 0.8)
	} else {
		gl.Uniform3f(drawData.uniformColor, drawData.bodyColors[node.bodyPart].x, drawData.bodyColors[node.bodyPart].y, drawData.bodyColors[node.bodyPart].z)
	}

	gl.DrawArrays(gl.TRIANGLES, 0, 36)
	if node.children != nil {
		for child := range node.children {
			iterateChildrens(drawData, &node.children[child], matModel)
		}
	}
}

func init() {
	runtime.LockOSThread()
}

func main() {
	var vao uint32
	var vbo uint32
	var drawData DrawData

	window := initGlfw()
	program := initOpenGL()
	defer glfw.Terminate()

	drawData.selectedBodypart = -1
	keyTimeout := time.Now()
	frameNumber := 0
	cameraRotation := Vec3f32{0, 0, 0}
	cameraTranslation := Vec3f32{0, 0, -20}
	currentAnimation := Animation{}
	defaultHumanConfig := HumanDefaultConfig()
	drawData.bodyConfig = HumanDefaultConfig()
	drawData.bodyColors = HumanDefaultColor()
	drawData.bodyConfigTmp = SetToZero()
	drawData.bodyTranslation = Vec3f32{0, 0, 0}
	testPwet := 0
	animationType := -1

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(triangle), gl.Ptr(triangle), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.BindVertexArray(vao)
	gl.UseProgram(program)
	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.DEPTH_TEST)

	matProjUniform := gl.GetUniformLocation(program, gl.Str("matProjection\x00"))
	matCameraUniform := gl.GetUniformLocation(program, gl.Str("matCamera\x00"))
	drawData.uniformColor = gl.GetUniformLocation(program, gl.Str("partColor\x00"))
	drawData.uniformModel = gl.GetUniformLocation(program, gl.Str("matModel\x00"))

	matProj := mat4.Perspective(mgl32.DegToRad(60), float32(width)/float32(height), 0.1, 1000)
	gl.UniformMatrix4fv(matProjUniform, 1, false, &matProj[0][0])

	for !window.ShouldClose() {
		start := time.Now()
		if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == 1 {
			os.Exit(0)
		}
		if glfw.GetCurrentContext().GetKey(glfw.KeyLeft) == 1 {
			cameraRotation.y += 0.1
		}
		if glfw.GetCurrentContext().GetKey(glfw.KeyRight) == 1 {
			cameraRotation.y -= 0.1
		}
		if glfw.GetCurrentContext().GetKey(glfw.KeyUp) == 1 {
			cameraRotation.x += 0.065
		}
		if glfw.GetCurrentContext().GetKey(glfw.KeyDown) == 1 {
			cameraRotation.x -= 0.065
		}
		if glfw.GetCurrentContext().GetKey(glfw.KeyA) == 1 {
			cameraTranslation.x -= 0.25
		}
		if glfw.GetCurrentContext().GetKey(glfw.KeyD) == 1 {
			cameraTranslation.x += 0.1
		}

		if glfw.GetCurrentContext().GetKey(glfw.Key1) == 1 {
			currentAnimation = Animation{}
			frameNumber = 0
			drawData.bodyConfig = defaultHumanConfig
			drawData.selectedBodypart = -1
			drawData.bodyTranslation = Vec3f32{0, 0, 0}

		}
		if glfw.GetCurrentContext().GetKey(glfw.Key2) == 1 {
			currentAnimation = createWalkingAnimation()
			animationType = 2
			frameNumber = 0
			drawData.bodyConfig = defaultHumanConfig
			drawData.selectedBodypart = -1
			drawData.bodyTranslation = Vec3f32{0, 0, 0}

		}
		if glfw.GetCurrentContext().GetKey(glfw.Key3) == 1 {
			currentAnimation = createJumpingAnimation()
			animationType = 3
			frameNumber = 0
			drawData.bodyConfig = defaultHumanConfig
			drawData.selectedBodypart = -1
			drawData.bodyTranslation = Vec3f32{0, 0, 0}

		}
		if glfw.GetCurrentContext().GetKey(glfw.Key4) == 1 {
			currentAnimation = createDabAnimation()
			animationType = -1
			frameNumber = 0
			drawData.bodyConfig = defaultHumanConfig
			drawData.selectedBodypart = -1
			drawData.bodyTranslation = Vec3f32{0, 0, 0}

		}
		if glfw.GetCurrentContext().GetKey(glfw.Key5) == 1 {
			currentAnimation = createFuckUAnimation()
			animationType = -1
			frameNumber = 0
			drawData.bodyConfig = defaultHumanConfig
			drawData.selectedBodypart = -1
			drawData.bodyTranslation = Vec3f32{0, 0, 0}

		}

		if glfw.GetCurrentContext().GetKey(glfw.KeyP) == 1 && time.Since(keyTimeout).Milliseconds() > 120 {
			drawData.selectedBodypart += 1
			if drawData.selectedBodypart > 9 {
				drawData.selectedBodypart = -1
			}
			keyTimeout = time.Now()
		}
		if glfw.GetCurrentContext().GetKey(glfw.KeyKPAdd) == 1 && drawData.selectedBodypart >= 0 {
			drawData.bodyConfig[drawData.selectedBodypart].size.x += 0.1
			drawData.bodyConfig[drawData.selectedBodypart].size.y += 0.1
			drawData.bodyConfig[drawData.selectedBodypart].size.z += 0.1
		}
		if glfw.GetCurrentContext().GetKey(glfw.KeyKPSubtract) == 1 && drawData.selectedBodypart >= 0 {
			drawData.bodyConfig[drawData.selectedBodypart].size.x -= 0.1
			drawData.bodyConfig[drawData.selectedBodypart].size.y -= 0.1
			drawData.bodyConfig[drawData.selectedBodypart].size.z -= 0.1
		}
		if drawData.selectedBodypart >= 0 {
			if drawData.bodyConfig[drawData.selectedBodypart].size.x < defaultHumanConfig[drawData.selectedBodypart].size.x*.2 {
				drawData.bodyConfig[drawData.selectedBodypart].size.x = defaultHumanConfig[drawData.selectedBodypart].size.x * .2
			}
			if drawData.bodyConfig[drawData.selectedBodypart].size.y < defaultHumanConfig[drawData.selectedBodypart].size.y*.2 {
				drawData.bodyConfig[drawData.selectedBodypart].size.y = defaultHumanConfig[drawData.selectedBodypart].size.y * .2
			}
			if drawData.bodyConfig[drawData.selectedBodypart].size.z < defaultHumanConfig[drawData.selectedBodypart].size.z*.2 {
				drawData.bodyConfig[drawData.selectedBodypart].size.z = defaultHumanConfig[drawData.selectedBodypart].size.z * .2
			}
		}
		matCamera := mat4.Translate(cameraTranslation.x, -2, cameraTranslation.z).Mult(mat4.Rotation(cameraRotation.x, cameraRotation.y, cameraRotation.z))
		gl.UniformMatrix4fv(matCameraUniform, 1, false, &matCamera[0][0])
		handleDrawHuman(drawData, &frameNumber, currentAnimation, &testPwet, animationType)
		_ = animationType
		glfw.PollEvents()
		window.SwapBuffers()
		if diff := time.Since(start).Milliseconds() - 16; diff > 0 {
			time.Sleep(time.Duration(diff) * time.Millisecond)
		}
	}
}

func handleDrawHuman(drawData DrawData, frame *int, currentAnimation Animation, testPwet *int, animationType int) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	drawData.bodyConfigTmp = SetToZero()
	if *frame > int(currentAnimation.duration) {
		*frame = 0
	}

	for _, keyframe := range currentAnimation.keyframes {
		if *frame >= keyframe.start && *frame >= keyframe.end {
			drawData.bodyConfigTmp[keyframe.BodyPart].rotation.x += mat4.DegToRad(keyframe.rotation.x)
			drawData.bodyConfigTmp[keyframe.BodyPart].rotation.y += mat4.DegToRad(keyframe.rotation.y)
			drawData.bodyConfigTmp[keyframe.BodyPart].rotation.z += mat4.DegToRad(keyframe.rotation.z)
		}
		if *frame >= keyframe.start && *frame < keyframe.end {
			drawData.bodyConfigTmp[keyframe.BodyPart].rotation.x += mat4.DegToRad(keyframe.rotation.x / float32((keyframe.end - keyframe.start)) * float32((*frame - keyframe.start)))
			drawData.bodyConfigTmp[keyframe.BodyPart].rotation.y += mat4.DegToRad(keyframe.rotation.y / float32((keyframe.end - keyframe.start)) * float32((*frame - keyframe.start)))
			drawData.bodyConfigTmp[keyframe.BodyPart].rotation.z += mat4.DegToRad(keyframe.rotation.z / float32((keyframe.end - keyframe.start)) * float32((*frame - keyframe.start)))
		}
	}

	if animationType == 2 {
		drawData.bodyTranslation = Vec3f32{0, 0, float32(*testPwet) * .05} //mat4.Translate(0, 0, float32(*testPwet)*.05)
		(*testPwet)++
	} else if animationType == 3 {
		currentTranslationAnimation := createJumpingTranslation()
		for _, keyframe := range currentTranslationAnimation.keyframes {
			if *frame >= keyframe.start && *frame >= keyframe.end {
				drawData.bodyTranslation.x += mat4.DegToRad(keyframe.translation.x)
				drawData.bodyTranslation.y += mat4.DegToRad(keyframe.translation.y)
				drawData.bodyTranslation.z += mat4.DegToRad(keyframe.translation.z)
			}
			if *frame >= keyframe.start && *frame < keyframe.end {
				drawData.bodyTranslation.x += mat4.DegToRad(keyframe.translation.x / float32((keyframe.end - keyframe.start)) * float32((*frame - keyframe.start)))
				drawData.bodyTranslation.y += mat4.DegToRad(keyframe.translation.y / float32((keyframe.end - keyframe.start)) * float32((*frame - keyframe.start)))
				drawData.bodyTranslation.z += mat4.DegToRad(keyframe.translation.z / float32((keyframe.end - keyframe.start)) * float32((*frame - keyframe.start)))
			}
		}

		// drawData.bodyTranslation = mat4.Translate(0, 0, float32(*testPwet)*.05)
	}

	humanBody := GenerateHuman(drawData.bodyConfig, drawData.bodyConfigTmp)
	iterateChildrens(drawData, humanBody, mat4.Identity())
	(*frame)++
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "HumanGL :[", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
