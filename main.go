package main

import (
	"fmt"
	"log"
	"math"
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

	rightLowerLeg := Node {
		bodyPart: RLLEG,
		children: nil,
	}
	rightLowerLeg.transform = mgl32.Translate3D(0, body[RULEG].size.y, 0)

	rightUpperLeg := Node {
		bodyPart: RULEG,
		children: []Node{ rightLowerLeg },
	}
	rightUpperLeg.transform = mgl32.Translate3D((body[TORSO].size.x * .5 - body[RULEG].size.x * .5), 0, body[TORSO].size.z * .5)
	rightUpperLeg.transform = rightUpperLeg.transform.Mul4(mgl32.HomogRotate3DX(body[RULEG].rotation.x))


	leftLowerLeg := Node {
		bodyPart: LLLEG,
		children: nil,
	}
	leftLowerLeg.transform = mgl32.Translate3D(0, body[LULEG].size.y, 0)

	leftUpperLeg := Node {
		bodyPart: LULEG,
		children: []Node{leftLowerLeg},
	}
	leftUpperLeg.transform = mgl32.Translate3D(-(body[TORSO].size.x * .5 - body[LULEG].size.x * .5), 0, body[TORSO].size.z * .5)
	leftUpperLeg.transform = leftUpperLeg.transform.Mul4(mgl32.HomogRotate3DX(body[LULEG].rotation.x))

	matLeftLowerArm := mgl32.Translate3D(0, body[LUARM].size.y, 0)
	matLeftLowerArm = matLeftLowerArm.Mul4(mgl32.HomogRotate3DX(body[LLARM].rotation.x))
	leftLowerArm := Node{
		transform: matLeftLowerArm,
		bodyPart: LLARM,
		children: nil,
	}

	matLeftUpperArm := mgl32.Translate3D(-(body[TORSO].size.x * .5 + body[LUARM].size.x * .5), body[TORSO].size.y * 0.9, body[TORSO].size.z * .5)
	matLeftUpperArm = matLeftUpperArm.Mul4(mgl32.HomogRotate3DX(body[LUARM].rotation.x))
	matLeftUpperArm = matLeftUpperArm.Mul4(mgl32.HomogRotate3DX(body[LUARM].rotation.z))
	matLeftUpperArm = matLeftUpperArm.Mul4(mgl32.HomogRotate3DY(body[LUARM].rotation.y))
	leftUpperArm := Node{
		transform: matLeftUpperArm,
		bodyPart: LUARM,
		children: []Node{leftLowerArm},
	}

	
	matRightLowerArm := mgl32.Translate3D(0, body[RUARM].size.y, 0)
	matRightLowerArm = matRightLowerArm.Mul4(mgl32.HomogRotate3DX(body[RLARM].rotation.x))
	rightLowerArm := Node{
		transform: matRightLowerArm,
		bodyPart: RLARM,
		children: nil,
	}

	matRightUpperArm := mgl32.Translate3D(body[TORSO].size.x * 0.5 + body[RUARM].size.x * .5, body[TORSO].size.y * 0.9, body[TORSO].size.z * .5)
	matRightUpperArm = matRightUpperArm.Mul4(mgl32.HomogRotate3DX(body[RUARM].rotation.x))
	rightUpperArm := Node{
		transform: matRightUpperArm,
		bodyPart: RUARM,
		children: []Node{rightLowerArm},
	}

	matHead := mgl32.Translate3D(0, body[TORSO].size.y, 0)
	head := Node {
		transform: matHead,
		bodyPart: HEAD,
		children: nil,
	}

	matTorso := mgl32.HomogRotate3DX(body[TORSO].rotation.x)
	matTorso = matTorso.Mul4(mgl32.HomogRotate3DY(body[TORSO].rotation.y))
	matTorso = matTorso.Mul4(mgl32.HomogRotate3DZ(body[TORSO].rotation.z))
	torso := Node {
		transform: matTorso,
		bodyPart: TORSO,
		children: []Node{ head, rightUpperArm, leftUpperArm, leftUpperLeg, rightUpperLeg},
	}

	return &torso
}

func translateBackBodypart(modelMat mgl32.Mat4, bodyConfig BodyConfig) mgl32.Mat4 {
	instanceMat := modelMat.Mul4(mgl32.Translate3D(0, bodyConfig.size.y * 0.5, 0))
	instanceMat = instanceMat.Mul4(mgl32.Scale3D(bodyConfig.size.x, bodyConfig.size.y, bodyConfig.size.z))
	return instanceMat
}

func iterateChildrens(modelUniform int32, node *Node, bodyConfig [10]BodyConfig, matModel mgl32.Mat4) {
	if node == nil {
		return
	}
	matInstance := mgl32.Ident4()
	matModel = matModel.Mul4(node.transform)

	if node.bodyPart == TORSO {
		matInstance = matModel.Mul4(mgl32.Translate3D(0, bodyConfig[TORSO].size.y * 0.5, bodyConfig[TORSO].size.z * 0.5))
		matInstance = matInstance.Mul4(mgl32.Scale3D(bodyConfig[TORSO].size.x, bodyConfig[TORSO].size.y, bodyConfig[TORSO].size.z))
	} else {
		matInstance = translateBackBodypart(matModel, bodyConfig[node.bodyPart])
	}

	gl.UniformMatrix4fv(modelUniform, 1, false, &matInstance[0])
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
	if node.children != nil {
		for child := range node.children {
			iterateChildrens(modelUniform, &node.children[child], bodyConfig, matModel)
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

	bodyConfig[HEAD].size.x = 2.0
	bodyConfig[HEAD].size.y = 2.0
	bodyConfig[HEAD].size.z = 2.0

	bodyConfig[TORSO].rotation.y = mgl32.DegToRad(0)
	bodyConfig[TORSO].size.x = 5.0
	bodyConfig[TORSO].size.y = 6.0
	bodyConfig[TORSO].size.z = 2.0

	bodyConfig[RUARM].size.x = 2.0
	bodyConfig[LUARM].size.x = 2.0
	bodyConfig[RLARM].size.x = 1.0
	bodyConfig[LLARM].size.x = 1.0

	bodyConfig[RUARM].size.y = 4.0
	bodyConfig[LUARM].size.y = 4.0

	bodyConfig[RLARM].size.y = 2.0
	bodyConfig[LLARM].size.y = 2.0

	bodyConfig[RUARM].rotation.x = mgl32.DegToRad(20)
	bodyConfig[RLARM].rotation.x = mgl32.DegToRad(50)

	bodyConfig[LUARM].rotation.x = mgl32.DegToRad(180)
	bodyConfig[LUARM].rotation.y = mgl32.DegToRad(0)
	bodyConfig[LLARM].rotation.x = mgl32.DegToRad(0)


	bodyConfig[LULEG].rotation.x = mgl32.DegToRad(180)
	bodyConfig[LULEG].size.y = 3.0
	bodyConfig[LULEG].size.x = 2.0
	bodyConfig[LLLEG].size.y = 1.5

	bodyConfig[RULEG].rotation.x = mgl32.DegToRad(180)
	bodyConfig[RULEG].size.y = 3.0
	bodyConfig[RULEG].size.x = 2.0
	bodyConfig[RLLEG].size.y = 2.5
	// os.Exit(1)


	matProjUniform := gl.GetUniformLocation(program, gl.Str("matProjection\x00"))
	matCameraUniform := gl.GetUniformLocation(program, gl.Str("matCamera\x00"))

	matCamera := mgl32.Translate3D(0, -2, -20)	
	matProj := mgl32.Perspective(mgl32.DegToRad(60), float32(width) / float32(height), 0.1, 100)

	a := 0.0
	b := 0.0
	c := 0.0
	for !window.ShouldClose() {
		gl.UniformMatrix4fv(matProjUniform, 1, false, &matProj[0])
		gl.UniformMatrix4fv(matCameraUniform, 1, false, &matCamera[0])

		bodyConfig[TORSO].rotation.y = float32(math.Cos(a))
		bodyConfig[TORSO].size.x = 2 + float32(math.Abs(float64(float32(math.Cos(a) * 8))))
		
		bodyConfig[RUARM].rotation.x =  float32(math.Cos(b))
		bodyConfig[RLARM].rotation.x =  float32(math.Cos(c))
		bodyConfig[LUARM].rotation.x =  -float32(math.Cos(b))
		bodyConfig[LLARM].rotation.x =  -float32(math.Cos(c))
	

		draw(vao, window, program, bodyConfig)

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

func draw(vao uint32, window *glfw.Window, program uint32, bodyConfig [10]BodyConfig) {

	if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == 1 {
		fmt.Println("panic! : tchoin tchoin")
		os.Exit(0)
	}


	matModelUniform := gl.GetUniformLocation(program, gl.Str("matModel\x00"))

	matModel := mgl32.Ident4()
	gl.UniformMatrix4fv(matModelUniform, 1, false, &matModel[0])


	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)


	// gl.DrawArrays(gl.TRIANGLES, 0, 36)

	
	humanBody := generateHumanBody(bodyConfig)
	iterateChildrens(matModelUniform, humanBody, bodyConfig, mgl32.Ident4())

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
