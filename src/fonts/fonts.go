package fonts

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	assets = []string{
		"assets/font/FixedsysExcelsiorIIIb_PL3.ttf",
	}
	Fonts []*rl.Font
)

func LoadFonts() {
	defer log.Printf("Loaded %d fonts", len(Fonts))
	for i := 0; i < len(assets); i++ {
		font := rl.LoadFont(assets[i])
		Fonts = append(Fonts, &font)
	}
}

func UnloadFonts() {
	for i := 0; i < len(Fonts); i++ {
		rl.UnloadFont(*Fonts[i])
	}
}
