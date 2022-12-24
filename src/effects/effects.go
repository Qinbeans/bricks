package effects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const FRAMEDELAY int32 = 5

type Effect struct {
	Img     *rl.Image
	Tex     rl.Texture2D
	frames  int32
	counter int32
	curr    int32
}

var (
	assets = []string{
		"assets/effect/flame.gif",
	}
	Effects []*Effect
)

func LoadEffects() {
	for i := 0; i < len(assets); i++ {
		effect := &Effect{}
		effect.Img = rl.LoadImageAnim(assets[i], &effect.frames)
		effect.Tex = rl.LoadTextureFromImage(effect.Img)
		effect.counter = 0
		effect.curr = 0
		Effects = append(Effects, effect)
	}
}

func UpdateEffects() {
	for i := 0; i < len(Effects); i++ {
		if Effects[i].counter >= Effects[i].frames {
			Effects[i].counter = 0
			Effects[i].curr++
			if Effects[i].curr >= Effects[i].frames {
				Effects[i].curr = 0
			}
		}
		Effects[i].counter++
	}
}

func UnloadEffects() {
	for i := 0; i < len(Effects); i++ {
		rl.UnloadImage(Effects[i].Img)
		rl.UnloadTexture(Effects[i].Tex)
	}
}
