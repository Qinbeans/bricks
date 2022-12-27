package main

import (
	ft "bricks/fonts"
	mt "bricks/materials"
	sc "bricks/scenes"
	ut "bricks/utility"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	curr *sc.Scene
	err  error
)

func main() {
	defer func() {
		mt.UnloadMaterials()
		ft.UnloadFonts()
		rl.CloseWindow()
	}()
	rl.InitWindow(int32(ut.DEFWIDTH), int32(ut.DEFHEIGHT), ut.TITLE)

	if err = ut.LoadOptions(); err != nil {
		panic(err)
	}

	rl.SetTargetFPS(60)
	//load assets
	mt.LoadMaterials()
	ft.LoadFonts()

	//load loading screen
	loadTex := rl.LoadTexture("assets/mat/loading.png")
	rl.BeginDrawing()
	rl.DrawTexturePro(loadTex, rl.NewRectangle(0, 0, float32(loadTex.Width), float32(loadTex.Height)), rl.NewRectangle(0, 0, float32(rl.GetScreenHeight()), float32(rl.GetScreenHeight())), rl.NewVector2(0, 0), 0, rl.Black)
	rl.EndDrawing()

	curr, err = sc.LoadScene("home.json")
	if err != nil {
		panic(err)
	}
	if err = curr.Init(); err != nil {
		panic(err)
	}
	rl.SetWindowSize(int(ut.GameOptions.Width), int(ut.GameOptions.Height))
	ut.IsFullscreen()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if err := curr.Draw(); err != nil {
			panic(err)
		}
		scene, err := curr.Update()
		if err != nil {
			if err.Error() == "exit" {
				break
			}
			panic(err)
		}
		if scene != nil {
			curr.Unload()
			if err = curr.Init(); err != nil {
				panic(err)
			}
			curr = scene
			curr.Init()
		}
		rl.EndDrawing()
	}
	if err = curr.Unload(); err != nil {
		panic(err)
	}

	rl.CloseWindow()
}
