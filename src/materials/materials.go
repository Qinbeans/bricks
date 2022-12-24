package materials

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	assets = []string{
		"assets/mat/green_grass.jpg",
		"assets/mat/dirt.jpg",
		"assets/mat/gravel.jpg",
		"assets/mat/sand.jpg",
		"assets/mat/water.jpg",
		"assets/mat/rock.jpg",
		"assets/mat/dry_grass.jpg",
	}
	Materials []*Material
)

type Material struct {
	Tex rl.Texture2D
}

func LoadMaterials() {
	defer log.Printf("Loaded %d materials\n", len(assets))
	for i := 0; i < len(assets); i++ {
		mat := &Material{}
		mat.Tex = rl.LoadTexture(assets[i])
		Materials = append(Materials, mat)
	}
}

func (m *Material) Draw(src rl.Rectangle, dest rl.Rectangle, tint rl.Color) {
	rl.DrawTexturePro(m.Tex, src, dest, rl.Vector2{X: 0, Y: 0}, 0, tint)
}

func UnloadMaterials() {
	for i := 0; i < len(Materials); i++ {
		rl.UnloadTexture(Materials[i].Tex)
	}
}
