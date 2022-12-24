package objects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Obj struct {
	Name    string `json:"name"`
	TexPath string `json:"texture"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	//in-game
	Texture rl.Texture2D `json:"-"`
}

func (o *Obj) Load() {
	o.Texture = rl.LoadTexture(o.TexPath)
}
