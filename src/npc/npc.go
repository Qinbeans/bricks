package npc

import (
	inv "bricks/inventory"
	qt "bricks/quests"
	ut "bricks/utility"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type NPC struct {
	Name      string        `json:"name"`
	Stats     ut.Stats      `json:"stats"`
	Inventory inv.Inventory `json:"inventory"` //TODO: Add inventory to NPCs
	Dialogue  []ut.Dialogue `json:"dialogue"`
	Quests    []qt.Quest    `json:"quests"`
	Roam      int           `json:"roam"` //The distance the NPC will roam from its spawn point
	TexSkin   string        `json:"skin"`
	Width     int           `json:"width"`
	Height    int           `json:"height"`
	//in-game
	Texture rl.Texture2D `json:"-"`
}

func (n *NPC) Load() {
	n.Texture = rl.LoadTexture(n.TexSkin)
}
