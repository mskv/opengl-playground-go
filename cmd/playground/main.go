package main

import (
	"fmt"
	"runtime"
	"time"

	"example.com/playground/playground/pkg/core"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	windowWidth := 640
	windowHeight := 480

	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Printf("OpenGL version: %#v\n", version)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	system := new(core.System)
	if err := system.Init(windowWidth, windowHeight); err != nil {
		panic(err)
	}

	renderNo := 0
	programStartNs := time.Now().UnixNano()
	for !window.ShouldClose() {
		nowNs := time.Now().UnixNano()
		elapsedMs := float32(nowNs-programStartNs) / 1000000

		if renderNo%50 == 0 {
			fmt.Printf("Render number %#v time %#v\n", renderNo, elapsedMs)
		}
		renderNo++

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		system.Run(elapsedMs)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
