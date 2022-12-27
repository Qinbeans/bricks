package scenes

import (
	ft "bricks/fonts"
	location "bricks/locations"
	ut "bricks/utility"

	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Description struct {
	NewScreen    bool           `json:"NewScreen"`   //if true, the scene will be rendered on a new screen
	Subwindow    bool           `json:"Subwindow"`   //if true, the scene will be rendered in a subwindow
	Player       bool           `json:"Player"`      //if true, the user can move around the scene
	Generate     bool           `json:"Generate"`    //if true, the scene will generate a map
	Locations    []location.Loc `json:"Location"`    //The location of the scene
	Interactions []Interaction  `json:"Interaction"` //The interactions that occur in the scene
	Groups       []Group        `json:"Group"`       //The groups of interactions that occur in the scene
	Spawn        *location.Loc  `json:"-"`           //exclude from json
}

func (d *Description) Load() {
	//load all textures
	for i := 0; i < len(d.Locations); i++ {
		if d.Locations[i].Load() {
			d.Spawn = &d.Locations[i]
		}
	}
}

func (d *Description) Init() {
	//init all textures
	dy := float32(0)
	dx := float32(0)
	var size rl.Vector2
	for i := 0; i < len(d.Groups); i++ {
		d.Groups[i].Init()
	}
	for i := 0; i < len(d.Interactions); i++ {
		index, err := d.Interactions[i].Init()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Interaction Error: %s\n", err)
		}
		if index != -1 {
			d.Interactions[i].Group = &d.Groups[index]
		}
		if d.Interactions[i].Width == -1 || d.Interactions[i].Height == -1 {
			size = rl.MeasureTextEx(*ft.Fonts[0], d.Interactions[i].Name, ut.DEFFONTSIZE, ut.FontSpacing)
		}
		if d.Interactions[i].Width == -1 {
			d.Interactions[i].Width = size.X
		}
		if d.Interactions[i].Height == -1 {
			d.Interactions[i].Height = float32(ut.GameOptions.FontSize)
		}
		if d.Interactions[i].Y == -1 {
			d.Interactions[i].Y = dy
			dy += d.Interactions[i].Height
		} else {
			dy = d.Interactions[i].Y + d.Interactions[i].Height
		}
		if d.Interactions[i].X == -1 {
			d.Interactions[i].X = 0
			dx += d.Interactions[i].Width
		} else {
			dx = d.Interactions[i].X + d.Interactions[i].Width
		}
	}
}

func (d *Description) Draw() {
	//draw all textures
	// for i := 0; i < len(d.Locations); i++ {
	// 	d.Locations[i].Draw()
	// }
	for i := 0; i < len(d.Interactions); i++ {
		d.Interactions[i].Draw()
	}
}

func (d *Description) Update() (*Scene, error) {
	//update all textures
	for i := 0; i < len(d.Interactions); i++ {
		g_intersections := make([]Interaction, 0)
		if d.Interactions[i].Group != nil {
			group := d.Interactions[i].Group
			for j := 0; j < len(group.Interactions); j++ {
				if group.Interactions[j] != i {
					g_intersections = append(g_intersections, d.Interactions[group.Interactions[j]])
				}
			}
		}
		scene, err := d.Interactions[i].Update(g_intersections)
		if err != nil {
			return nil, err
		}
		if scene != nil {
			return scene, nil
		}
	}
	return nil, nil
}
