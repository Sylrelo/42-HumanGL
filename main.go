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

const (
	width  = 1280
	height = 720

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

	gl.UniformMatrix4fv(drawData.uniformModel, 1, false, &matInstance[0][0])
	gl.Uniform3f(drawData.uniformColor, drawData.bodyColors[node.bodyPart].x, drawData.bodyColors[node.bodyPart].y, drawData.bodyColors[node.bodyPart].z)

	gl.DrawArrays(gl.TRIANGLES, 0, 36)
	if node.children != nil {
		for child := range node.children {
			iterateChildrens(drawData, &node.children[child], matModel)
		}
	}
}

func main() {
	runtime.LockOSThread()
	var vao uint32
	var vbo uint32

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	glBuffer := triangle
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(glBuffer), gl.Ptr(glBuffer), gl.STATIC_DRAW)
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

	cameraRotation := Vec3f32{0, 0, 0}
	cameraTranslation := Vec3f32{0, 0, 0}

	// matCamera := mgl32.Translate3D(0, -2, -20)
	matProj := mat4.Perspective(mgl32.DegToRad(60), float32(width)/float32(height), 0.1, 1000)
	gl.UniformMatrix4fv(matProjUniform, 1, false, &matProj[0][0])

	var drawData DrawData

	drawData.bodyColors = HumanDefaultColor()

	drawData.bodyConfig = HumanDefaultConfig()
	drawData.bodyConfigTmp = SetToZero()

	drawData.uniformColor = gl.GetUniformLocation(program, gl.Str("partColor\x00"))
	drawData.uniformModel = gl.GetUniformLocation(program, gl.Str("matModel\x00"))

	a := 0.0
	b := 0.0
	c := 0.0

	currentAnimation := createJumpingAnimation()

	for id, keyframe := range currentAnimation.keyframes {
		fmt.Printf("%3d [%3d  ] [%4.1d - %4.1d ] (%4.2f, %4.2f, %4.2f)\n",
			id,
			keyframe.BodyPart,
			keyframe.start,
			keyframe.end,
			keyframe.rotation.x,
			keyframe.rotation.y,
			keyframe.rotation.z,
		)
	}

	frameNumber := 0
	//drawData.bodyConfig[LULEG].rotation = Vec3f32{-15, 0, 0}
	//drawData.bodyConfig[LLLEG].rotation = Vec3f32{-45, 0, 0}

	for !window.ShouldClose() {
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

		matCamera := mat4.Translate(cameraTranslation.x, -2, -20).Mult(mat4.Rotation(cameraRotation.x, cameraRotation.y, cameraRotation.z))
		gl.UniformMatrix4fv(matCameraUniform, 1, false, &matCamera[0][0])

		// bodyConfig[TORSO].rotation.y = float32(math.Cos(a))
		// bodyConfig[TORSO].size.x = 2 + float32(math.Abs(float64(float32(math.Cos(a) * 8))))

		// bodyConfig[RUARM].rotation.x =  float32(math.Cos(b))
		// bodyConfig[RLARM].rotation.x =  float32(math.Cos(c))
		// bodyConfig[LUARM].rotation.x =  -float32(math.Cos(b))
		// bodyConfig[LLARM].rotation.x =  -float32(math.Cos(c))

		draw(vao, window, program, drawData, &frameNumber)

		a += 0.005
		b += 0.005
		c += 0.005
	}
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

func draw(vao uint32, window *glfw.Window, program uint32, drawData DrawData, frame *int) {
	start := time.Now()
	// fmt.Printf("Frame: %d\n", *frame)

	if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == 1 {
		os.Exit(0)
	}

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	currentAnimation := createJumpingAnimation()

	if *frame > int(currentAnimation.duration) {
		*frame = 0
	}
	drawData.bodyConfigTmp = SetToZero()
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

	humanBody := GenerateHuman(drawData.bodyConfig, drawData.bodyConfigTmp)

	iterateChildrens(drawData, humanBody, mat4.Identity())

	glfw.PollEvents()
	window.SwapBuffers()
	if diff := time.Since(start).Milliseconds() - 16; diff > 0 {
		time.Sleep(time.Duration(diff) * time.Millisecond)
	}
	(*frame)++
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
