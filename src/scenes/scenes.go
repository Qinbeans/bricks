package scenes

import (
	ft "bricks/fonts"
	location "bricks/locations"
	mg "bricks/mapgen"
	mt "bricks/materials"
	pl "bricks/players"
	ut "bricks/utility"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//Types: -1 = Exit, 0 = Spawn Scene, 1 = integer, 2 = multiple choice integer, 3 = complete others then spawn scene, 4 = multiple choice from directory

type Interaction struct {
	Name     string  `json:"name"`   //The name of the interaction
	X        float32 `json:"x"`      //The x coordinate of the interaction
	Y        float32 `json:"y"`      //The y coordinate of the interaction
	Width    float32 `json:"width"`  //The width of the interaction
	Height   float32 `json:"height"` //The height of the interaction
	Type     int     `json:"type"`   //The type of interaction
	Value    string  `json:"value"`  //The value of the interaction
	Complete bool    `json:"-"`      //exclude from json
	Hover    bool    `json:"-"`      //exclude from json
}

type Description struct {
	NewScreen    bool           `json:"NewScreen"`   //if true, the scene will be rendered on a new screen
	Subwindow    bool           `json:"Subwindow"`   //if true, the scene will be rendered in a subwindow
	Player       bool           `json:"Player"`      //if true, the user can move around the scene
	Locations    []location.Loc `json:"Location"`    //The location of the scene
	Interactions []Interaction  `json:"Interaction"` //The interactions that occur in the scene
	Spawn        *location.Loc  `json:"-"`           //exclude from json
}

type Scene struct {
	Name        string      `json:"name"`        //The name of the scene
	Description Description `json:"description"` //The description of the scene
	//in-game
	PlayerData   pl.Player  `json:"-"`
	Heartbeat    int        `json:"-"`
	Fps          int        `json:"-"`
	camera       [2]float64 `json:"-"`
	cameraBounds [2]float64 `json:"-"`
	mapData      [][]int8   `json:"-"`
}

func (i *Interaction) Draw() {
	switch i.Type {
	case -1:
		//exit
		color := rl.Red
		if i.Hover {
			color = rl.Green
		}
		rl.DrawRectangleRounded(rl.NewRectangle(float32(i.X), float32(i.Y), float32(i.Width+ut.FontSpacing*2), float32(i.Height+ut.FontSpacing*2)), 0.2, 0, color)
		rl.DrawTextEx(*ft.Fonts[0], "Exit", rl.NewVector2(float32(i.X+ut.FontSpacing), float32(i.Y+ut.FontSpacing)), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
	case 0:
		//spawn scene
		color := rl.Red
		if i.Hover {
			color = rl.Green
		}
		rl.DrawRectangleRounded(rl.NewRectangle(float32(i.X), float32(i.Y), float32(i.Width+ut.FontSpacing*2), float32(i.Height+ut.FontSpacing*2)), 0.2, 0, color)
		rl.DrawTextEx(*ft.Fonts[0], i.Name, rl.NewVector2(float32(i.X+ut.FontSpacing), float32(i.Y+ut.FontSpacing)), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
	case 1:
		//integer
		color := rl.Red
		if i.Hover {
			color = rl.Green
		}
		text := i.Name + ": " + i.Value
		rl.DrawRectangleRounded(rl.NewRectangle(float32(i.X), float32(i.Y), float32(i.Width+ut.FontSpacing*2), float32(i.Height+ut.FontSpacing*2)), 0.2, 0, color)
		rl.DrawTextEx(*ft.Fonts[0], text, rl.NewVector2(float32(i.X+ut.FontSpacing), float32(i.Y+ut.FontSpacing)), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
	case 2:
		//multiple choice integer
		split := strings.Split(i.Value, ":")
		size, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}
		choice, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}
		for j := 1; j <= size; j++ {
			//draw all options
			if j == choice {
				rl.DrawRectangleRounded(rl.NewRectangle(float32(i.X), float32(i.Y), float32(i.Width), float32(i.Height)), 0.2, 0, rl.Red)
			} else {
				rl.DrawRectangleRounded(rl.NewRectangle(float32(i.X), float32(i.Y), float32(i.Width), float32(i.Height)), 0.2, 0, rl.White)
			}
			rl.DrawTextEx(*ft.Fonts[0], strconv.Itoa(j), rl.NewVector2(float32(i.X+5), float32(i.Y+5)), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
		}
	case 3:
		//complete others then spawn scene
	case 4:
		//multiple choice from directory
	default:
		//error
		rl.DrawRectangleRounded(rl.NewRectangle(float32(i.X), float32(i.Y), float32(i.Width), float32(i.Height)), 0.2, 0, rl.Red)
		rl.DrawText("Error", int32(i.X+5), int32(i.Y+5), 20, rl.White)
	}
}

func (i *Interaction) Update() (*Scene, error) {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(i.X), float32(i.Y), float32(i.Width), float32(i.Height))) {
		i.Hover = true
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			switch i.Type {
			case -1:
				//exit
				return nil, errors.New("exit")
			case 0:
				//spawn scene
				//load scene
				scene, err := LoadScene(i.Value)
				if err != nil {
					return nil, err
				}
				return scene, nil
			case 1:
				//integer
				//TODO
			case 2:
				//multiple choice integer
				split := strings.Split(i.Value, ":")
				size, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v", err)
					return nil, err
				}
				choice, err := strconv.Atoi(split[1])
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v", err)
					return nil, err
				}
				if choice < size {
					choice++
				} else {
					choice = 1
				}
				i.Value = split[0] + ":" + strconv.Itoa(choice)
			case 3:
				//complete others then spawn scene
				//TODO
			case 4:
				//multiple choice from directory
				//TODO
			default:
				//error
				//TODO
			}
		}
	} else {
		i.Hover = false
	}
	return nil, nil
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
	for i := 0; i < len(d.Interactions); i++ {
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
	fmt.Fprintf(os.Stderr, "%f, %f\n", dx, dy)
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
		scene, err := d.Interactions[i].Update()
		if err != nil {
			return nil, err
		}
		if scene != nil {
			return scene, nil
		}
	}
	return nil, nil
}

// load necessary assets
func LoadScene(scenename string) (*Scene, error) {
	defer log.Printf("Loaded scene: %s", scenename)
	//scenes are stored in assets/gamedata/scenes
	filename := "assets/gamedata/scenes/" + scenename
	sceneFile, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("scene file could not be opened")
	}
	defer sceneFile.Close()
	reader := bufio.NewReader(sceneFile)
	decoder := json.NewDecoder(reader)
	scene := &Scene{}
	err = decoder.Decode(scene)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return nil, errors.New("scene file could not be decoded")
	}
	scene.Description.Load()
	return scene, nil
}

// init necessary variables
func (s *Scene) Init() error {
	defer log.Printf("Initialized scene: %s", s.Name)
	//generate map
	s.mapData = mg.GenMap(len(mt.Materials), ut.MAPSIZE)
	//pair for location
	s.camera = [2]float64{0, 0}
	s.cameraBounds = [2]float64{0.1, 0.1}
	s.Description.Init()
	return nil
}

func (s *Scene) Draw() error {
	//draw all objects, items, etc
	//draw map from camera position to camera position + screen size
	for x := 0; x < ut.GameOptions.Width; x += int(ut.Step) {
		for y := 0; y < ut.GameOptions.Height; y += int(ut.Step) {
			nx := (x / int(ut.Step)) + int(s.camera[0]*ut.Step)
			ny := (y / int(ut.Step)) + int(s.camera[1]*ut.Step)
			ax := float32(x)
			ay := float32(y)
			if nx < 0 || ny < 0 || nx >= ut.MAPSIZE || ny >= ut.MAPSIZE {
				continue
			}
			mt.Materials[s.mapData[nx][ny]].Draw(
				rl.Rectangle{X: 0, Y: 0, Width: float32(ut.Step), Height: float32(ut.Step)},
				rl.Rectangle{X: ax, Y: ay, Width: float32(ut.Step), Height: float32(ut.Step)},
				rl.White,
			)
		}
	}
	//draw interactions
	s.Description.Draw()
	return nil
}

func (s *Scene) Update() (*Scene, error) {
	//update all objects, items, etc
	if s.Description.Player {
		if rl.IsKeyDown(rl.KeyW) && s.camera[1] > 0 {
			if ut.Heartbeat%ut.DELAY == 0 {
				s.camera[1] -= s.cameraBounds[1]
			}
		}
		if rl.IsKeyDown(rl.KeyS) && s.camera[1] < float64(ut.MAPSIZE-ut.Modheight) {
			if ut.Heartbeat%ut.DELAY == 0 {
				s.camera[1] += s.cameraBounds[1]
			}
		}
		if rl.IsKeyDown(rl.KeyA) && s.camera[0] > 0 {
			if ut.Heartbeat%ut.DELAY == 0 {
				s.camera[0] -= s.cameraBounds[0]
			}
		}
		if rl.IsKeyDown(rl.KeyD) && s.camera[0] < float64(ut.MAPSIZE-ut.Modwidth) {
			if ut.Heartbeat%ut.DELAY == 0 {
				s.camera[0] += s.cameraBounds[0]
			}
		}
		if ut.Heartbeat == math.MaxInt {
			ut.Heartbeat = 0
		} else {
			ut.Heartbeat++
		}
	}
	//check for interactions
	for i := 0; i < len(s.Description.Interactions); i++ {
		scene, err := s.Description.Update()
		if err != nil {
			return nil, err
		}
		if scene != nil {
			return scene, nil
		}
	}
	return nil, nil
}

func (s *Scene) Unload() error {
	//unload all textures
	return nil
}
