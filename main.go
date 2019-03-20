// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Renders a textured spinning cube using GLFW 3 and OpenGL 4.1 core forward-compatible profile.
package main // import "github.com/artheus/gl-voxel"

import (
	"flag"
	"github.com/artheus/gl-voxel/world"
	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"image"
	"image/draw"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const windowWidth = 1440
const windowHeight = 830

var (
	texturePath  = flag.String("t", "texture.png", "texture file")
	renderRadius = flag.Int("r", 6, "render radius")
)

func loadImage(fname string) ([]uint8, image.Rectangle, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, image.Rectangle{}, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, image.Rectangle{}, err
	}
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)
	return rgba.Pix, img.Bounds(), nil
}

func initGL(w, h int) *glfw.Window {
	err := glfw.Init()
	if err != nil {
		log.Fatal(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, gl.TRUE)

	win, err := glfw.CreateWindow(w, h, "gl-voxel", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	win.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		log.Fatal(err)
	}
	glfw.SwapInterval(1) // enable vsync
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	glhf.Init()
	return win
}

func run() {

	defer func() {
		mainthread.Call(func() {
			glfw.Terminate()
		})
	}()

	var (
		err error
		window *glfw.Window
		blockShader *glhf.Shader
		texture *glhf.Texture
		blockSlices []*glhf.VertexSlice
	)

	img, rect, err := loadImage(*texturePath)
	if err != nil {
		panic(err)
	}

	blockFSFile, err := os.Open("shader/block.fs")
	defer blockFSFile.Close()
	if err != nil {
		panic(err)
	}

	blockFSBytes, err := ioutil.ReadAll(blockFSFile)
	if err != nil {
		panic(err)
	}

	blockFragmentShader := string(blockFSBytes)

	blockVSFile, err := os.Open("shader/block.vs")
	defer blockVSFile.Close()
	if err != nil {
		panic(err)
	}

	blockVSBytes, err := ioutil.ReadAll(blockVSFile)
	if err != nil {
		panic(err)
	}

	blockVertexShader := string(blockVSBytes)

	mainthread.Call(func() {

		window = initGL(windowWidth, windowHeight)
		//win.SetMouseButtonCallback(game.onMouseButtonCallback)
		//win.SetCursorPosCallback(game.onCursorPosCallback)
		//win.SetFramebufferSizeCallback(game.onFrameBufferSizeCallback)
		//win.SetKeyCallback(game.onKeyCallback)
		//game.win = win

		blockShader, err = glhf.NewShader(glhf.AttrFormat{
			glhf.Attr{Name: "aPos", Type: glhf.Vec3},
			glhf.Attr{Name: "aNormal", Type: glhf.Vec3},
			glhf.Attr{Name: "vertTexCoord", Type: glhf.Vec2},
		}, glhf.AttrFormat{
			glhf.Attr{Name: "fragTexCoord", Type: glhf.Vec2},
			glhf.Attr{Name: "Normal", Type: glhf.Vec3},
			glhf.Attr{Name: "FragPos", Type: glhf.Vec3},
		}, blockVertexShader, blockFragmentShader)

		if err != nil {
			panic(err)
		}

		texture = glhf.NewTexture(rect.Dx(), rect.Dy(), false, img)

		gl.UseProgram(blockShader.ID())

		projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/windowHeight, 0.1, 500.0)
		projectionUniform := gl.GetUniformLocation(blockShader.ID(), gl.Str("projection\x00"))
		gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

		camera := mgl32.LookAtV(mgl32.Vec3{150,150,150}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 2, 0})
		cameraUniform := gl.GetUniformLocation(blockShader.ID(), gl.Str("camera\x00"))
		gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

		model := mgl32.Ident4()
		modelUniform := gl.GetUniformLocation(blockShader.ID(), gl.Str("model\x00"))
		gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

		textureUniform := gl.GetUniformLocation(blockShader.ID(), gl.Str("tex\x00"))
		gl.Uniform1i(textureUniform, 0)

		w := world.NewWorld(blockShader)

		log.Printf("%v", w.MeshCache)
	})

	//angle := 0.0
	previousTime := glfw.GetTime()
	fpsPrintTick := 0
	var fps float64

	shouldQuit := false
	for !shouldQuit {
		mainthread.Call(func() {
			if window.ShouldClose() {
				shouldQuit = true
			}
			gl.ClearColor(0.57, 0.71, 0.77, 1)
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			// Update
			timeNow := glfw.GetTime()
			elapsed := timeNow - previousTime
			previousTime = timeNow

			fpsPrintTick += 1
			if fps == 0 {
				fps = elapsed
			} else {
				fps = (fps+elapsed)/2
			}

			if fpsPrintTick == 1000 {
				fpsPrintTick = 0
				log.Printf("FPS: %f", float64(time.Second)/(fps*float64(time.Second)))
			}

			//angle += elapsed

			for _, s := range blockSlices {
				blockShader.Begin()
				texture.Begin()
				s.Begin()
				s.Draw()
				s.End()
				texture.End()
				blockShader.End()
			}

			// Maintenance
			window.SwapBuffers()
			glfw.PollEvents()
		})
	}
}

func main() {
	mainthread.Run(run)
}