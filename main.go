package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(640, 480, "testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	window.SetKeyCallback(keyCallback)
	window.SetFramebufferSizeCallback(frameBufferCallback)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	for !window.ShouldClose() {

		gl.ClearColor(0.0, 1.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		window.SwapBuffers()
		glfw.PollEvents()

	}
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

func frameBufferCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
