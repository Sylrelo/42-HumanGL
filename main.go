package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

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

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	vao := makeVao(triangle)
	gl.BindVertexArray(vao)
	gl.UseProgram(program)

	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	for !window.ShouldClose() {
		draw(vao, window, program)
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

func draw(vao uint32, window *glfw.Window, program uint32) {
	matProjUniform := gl.GetUniformLocation(program, gl.Str("matProjection\x00"))
	matCameraUniform := gl.GetUniformLocation(program, gl.Str("matCamera\x00"))
	matModelUniform := gl.GetUniformLocation(program, gl.Str("matModel\x00"))

	matCamera := mgl32.Translate3D(0, 0, -5)
	matProj := mgl32.Perspective(mgl32.DegToRad(60), float32(width) / float32(height), 0.1, 100)
	matModel := mgl32.Ident4()

	gl.UniformMatrix4fv(matProjUniform, 1, false, &matProj[0])
	gl.UniformMatrix4fv(matCameraUniform, 1, false, &matCamera[0])
	gl.UniformMatrix4fv(matModelUniform, 1, false, &matModel[0])

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

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

func makeVao(points []float32) uint32 {
	var vbo uint32
	var vao uint32

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
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
