package locations

import (
	event "bricks/events"
	item "bricks/items"
	npc "bricks/npc"
	obj "bricks/objects"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LocObj struct {
	X      int     `json:"x"`
	Y      int     `json:"y"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Obj    obj.Obj `json:"obj"`
}

type LocNPC struct {
	X      int     `json:"x"`
	Y      int     `json:"y"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	NPC    npc.NPC `json:"npc"`
}

type LocItem struct {
	X      int       `json:"x"`
	Y      int       `json:"y"`
	Width  int       `json:"width"`
	Height int       `json:"height"`
	Count  int       `json:"count"`
	Item   item.Item `json:"item"`
}

type LocEvent struct {
	X      int         `json:"x"`
	Y      int         `json:"y"`
	Width  int         `json:"width"`
	Height int         `json:"height"`
	Event  event.Event `json:"event"`
}

type Loc struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Spawn       bool       `json:"spawn"` //if true, player will spawn here
	X           int        `json:"x"`
	Y           int        `json:"y"`
	Width       int        `json:"width"`
	Height      int        `json:"height"`
	FloorTex    string     `json:"floor"`
	Objects     []LocObj   `json:"objects"`
	NPCs        []LocNPC   `json:"npcs"`
	Items       []LocItem  `json:"items"`
	Events      []LocEvent `json:"events"`
	//in-game
	Floor rl.Texture2D `json:"-"`
}

func (l *Loc) Load() bool {
	l.Floor = rl.LoadTexture(l.FloorTex)
	for i := 0; i < len(l.Objects); i++ {
		l.Objects[i].Obj.Load()
	}
	for i := 0; i < len(l.NPCs); i++ {
		l.NPCs[i].NPC.Load()
	}
	for i := 0; i < len(l.Items); i++ {
		l.Items[i].Item.Load()
	}
	return l.Spawn
}
