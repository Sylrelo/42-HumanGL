package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Vec3f32 struct {
	x float32
	y float32
	z float32
}

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
		void main() {
			frag_colour = vec4(1, 1, 1, 1.0);
		}
	` + "\x00"
)

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

type BodyConfig struct {
	size Vec3f32
	rotation Vec3f32
}

type Node struct {
	transform mgl32.Mat4
	children []Node
	bodyPart int
}

func generateHumanBody (body [10]BodyConfig) *Node {
	matLeftLowerArm := mgl32.Translate3D(-1, 0, 0)
	matLeftLowerArm = matLeftLowerArm.Mul4(mgl32.Scale3D(1, 1, 1))
	leftLowerArm := Node{
		transform: matLeftLowerArm,
		bodyPart: LLARM,
		children: nil,
	}

	matLeftUpperArm := mgl32.Translate3D(-1, 0, 0)
	matLeftUpperArm = matLeftUpperArm.Mul4(mgl32.Scale3D(1, 1, 1))
	leftUpperArm := Node{
		transform: matLeftUpperArm,
		bodyPart: LUARM,
		children: []Node{leftLowerArm},
	}

	_ = leftUpperArm
	
	matRightLowerArm := mgl32.Translate3D(0, body[RUARM].size.y, 0)
	matRightLowerArm = matRightLowerArm.Mul4(mgl32.HomogRotate3DX(body[RLARM].rotation.x))
	// matRightLowerArm = matRightLowerArm.Mul4(mgl32.Scale3D(1, 1, 1))
	rightLowerArm := Node{
		transform: matRightLowerArm,
		bodyPart: RLARM,
		children: nil,
	}

	matRightUpperArm := mgl32.Translate3D(body[TORSO].size.x + body[RUARM].size.x, body[TORSO].size.y * 0.9, body[TORSO].size.z * .5)
	matRightUpperArm = matRightUpperArm.Mul4(mgl32.HomogRotate3DX(body[RUARM].rotation.x))
	rightUpperArm := Node{
		transform: matRightUpperArm,
		bodyPart: RUARM,
		children: []Node{rightLowerArm},
	}


	matTorso := mgl32.HomogRotate3DX(body[TORSO].rotation.x)
	matTorso = matTorso.Mul4(mgl32.HomogRotate3DY(body[TORSO].rotation.y))
	matTorso = matTorso.Mul4(mgl32.HomogRotate3DZ(body[TORSO].rotation.z))
	// matTorso := mgl32.Translate3D(0, 0, 0)
	// matTorso = matTorso.Mul4(mgl32.Scale3D(1, 1, 1))
	torso := Node {
		transform: matTorso,
		bodyPart: TORSO,
		children: []Node{ rightUpperArm},
	}

	return &torso
}

func iterateChildrens(modelUniform int32, node *Node, spaceBefore int, bodyConfig [10]BodyConfig, stackedMat mgl32.Mat4) {
	modelMat := stackedMat
	instanceMat := mgl32.Ident4()

	if node == nil {
		return
	}
	// fmt.Printf("%*s %d\n", spaceBefore + 16, "Current Part : ", node.bodyPart)
	fmt.Printf("%*s %p\n", spaceBefore + 1, "", node)

	modelMat = modelMat.Mul4(node.transform)

	if node.bodyPart == TORSO {
		instanceMat = modelMat.Mul4(mgl32.Translate3D(bodyConfig[TORSO].size.x * 0.5, bodyConfig[TORSO].size.y * 0.5, bodyConfig[TORSO].size.z * 0.5))
		instanceMat = instanceMat.Mul4(mgl32.Scale3D(bodyConfig[TORSO].size.x, bodyConfig[TORSO].size.y, bodyConfig[TORSO].size.z))
	}
	if node.bodyPart == RUARM {
		instanceMat = modelMat.Mul4(mgl32.Translate3D(0, bodyConfig[RUARM].size.y * 0.5, 0))
		instanceMat = instanceMat.Mul4(mgl32.Scale3D(1, bodyConfig[RUARM].size.y, 1))
	}

	if node.bodyPart == RLARM {
		instanceMat = modelMat.Mul4(mgl32.Translate3D(0, bodyConfig[RLARM].size.y * 0.5, 0))
		instanceMat = instanceMat.Mul4(mgl32.Scale3D(1, bodyConfig[RLARM].size.y, 1))
	}

	// matModel := mgl32.Ident4()
	gl.UniformMatrix4fv(modelUniform, 1, false, &instanceMat[0])
	gl.DrawArrays(gl.TRIANGLES, 0, 36)

	if node.children != nil {
		for child := range node.children {
			iterateChildrens(modelUniform, &node.children[child], spaceBefore + 4, bodyConfig, modelMat)
		}
	}
}

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	// vao := makeVao(triangle)


	var vao uint32
	var vbo uint32
	var glBuffer []float32

	// for x := 0; x < 10; x++ {
	// 	glBuffer = append(glBuffer, triangle...)
	// }


	glBuffer = triangle
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
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)


	var bodyConfig [10]BodyConfig

	for x := 0; x < 10; x++ {
		bodyConfig[x] = BodyConfig{
			size: Vec3f32{1, 1, 1},
			rotation: Vec3f32{0, 0, 0},
		}
	}

	bodyConfig[TORSO].rotation.y = mgl32.DegToRad(-40)
	bodyConfig[TORSO].size.x = 4.0
	bodyConfig[TORSO].size.y = 6.0
	bodyConfig[TORSO].size.z = 2.0

	bodyConfig[RUARM].size.x = 0.5
	bodyConfig[RLARM].size.x = 0.5

	bodyConfig[RUARM].size.y = 3.0
	bodyConfig[RLARM].size.y = 2.0


	bodyConfig[RUARM].rotation.x = mgl32.DegToRad(20)
	bodyConfig[RLARM].rotation.x = mgl32.DegToRad(50)

	// os.Exit(1)
	for !window.ShouldClose() {
		draw(vao, window, program, bodyConfig)
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

func draw(vao uint32, window *glfw.Window, program uint32, bodyConfig [10]BodyConfig) {

	if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == 1 {
		fmt.Println("panic! : tchoin tchoin")
		os.Exit(0)
	}



	matProjUniform := gl.GetUniformLocation(program, gl.Str("matProjection\x00"))
	matCameraUniform := gl.GetUniformLocation(program, gl.Str("matCamera\x00"))
	matModelUniform := gl.GetUniformLocation(program, gl.Str("matModel\x00"))

	matCamera := mgl32.Translate3D(-5, 0, -20)	
	matProj := mgl32.Perspective(mgl32.DegToRad(60), float32(width) / float32(height), 0.1, 100)
	matModel := mgl32.Ident4()
	gl.UniformMatrix4fv(matModelUniform, 1, false, &matModel[0])

	gl.UniformMatrix4fv(matProjUniform, 1, false, &matProj[0])
	gl.UniformMatrix4fv(matCameraUniform, 1, false, &matCamera[0])

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)


	// gl.DrawArrays(gl.TRIANGLES, 0, 36)

	
	humanBody := generateHumanBody(bodyConfig)
	iterateChildrens(matModelUniform, humanBody, 0, bodyConfig, mgl32.Ident4())

	_ = bodyConfig
	// for x := 0; x < 1; x++ {

	// 	matModel = mgl32.Translate3D(float32(x) / 4, 0, 0)
		// gl.UniformMatrix4fv(matModelUniform, 1, false, &matModel[0])
		// gl.DrawArrays(gl.TRIANGLES, 0, 36)
	// }

	glfw.PollEvents()
	window.SwapBuffers()
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

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
