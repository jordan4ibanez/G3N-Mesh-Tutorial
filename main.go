package main

import (
	"fmt"
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"

	chunkBuilder "github.com/jordan4ibanez/G3N-Mesh-Tutorial/engine"
)

func main() {
	fmt.Println("Wow Go is cool!")

	//this creates the application - pointer reference
	var a *app.Application = app.App()

	//this creates a basic scene - pointer reference
	var scene *core.Node = core.NewNode()

	//gui manager now manages scene
	gui.Manager().Set(scene)

	//this creates the camera - pointer reference
	var cam *camera.Camera = camera.New(1)

	// camera is utilizing float32 (could cast 64 to 32 in future?)
	cam.SetPosition(0, 0, 3)

	//set orbital control for the camera
	camera.NewOrbitControl(cam)

	//finally add to scene node - inject pointer reference
	scene.Add(cam)

	//this is the control mechanic of the camera built into G3N for debugging
	//this needs to be changed in the future
	camera.NewOrbitControl(cam)

	//a callback to update the viewport and camera aspect ratio when the window is resized
	//var type is explicit instead of implicit for learning purposes and compilation speedup
	var onResize (func(evname string, ev interface{})) = func(evname string, ev interface{}) {
		//get framebuffer size and update viewport accordingly
		var width, height = a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		//update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}

	//this exports the value into the engine
	a.Subscribe(window.OnWindowSize, onResize)

	//perhaps forces to use part of function so Go will compile it?
	onResize("", nil)

	//clear color (background color) in ARGB float32 0-1 format
	a.Gls().ClearColor(0.5, 0.5, 0.7, 1)

	//this is the function that shows you how to create a mesh from scratch
	chunkBuilder.DebugTest(scene)

	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		a.Gls()
		renderer.Render(scene, cam)
	})

}
