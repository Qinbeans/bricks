package items

import (
	ut "bricks/utility"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Item struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"` //Weapon, Armor, Consumable, etc.
	Stats    ut.Stats `json:"stats"`
	TexPath  string   `json:"texture"`
	Width    int      `json:"width"`
	Height   int      `json:"height"`
	Count    int      `json:"count"`
	MaxStack int      `json:"maxstack"`
	//in-game
	Texture rl.Texture2D `json:"-"`
}

func (i *Item) Load() {
	i.Texture = rl.LoadTexture(i.TexPath)
}
