package scenes

import (
	mg "bricks/mapgen"
	mt "bricks/materials"
	pl "bricks/players"
	ut "bricks/utility"

	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//If a size or position is -1, try to calculate it
//Value -1 refers to look for default value

// GameData variables
// $SCREENS refers to optional screen sizes
// The default screen size is 800x450
// Results should be in the format of:
//
//	800x450,1024x576,1280x720,1920x1080
//
// $CONFIG refers to the config directory
// $<NAME> refers to a variable in file with the name <NAME>

type Scene struct {
	Name        string      `json:"Name"`        //The name of the scene
	Description Description `json:"Description"` //The description of the scene
	//in-game
	PlayerData   pl.Player  `json:"-"`
	Heartbeat    int        `json:"-"`
	Fps          int        `json:"-"`
	camera       [2]float64 `json:"-"`
	cameraBounds [2]float64 `json:"-"`
	mapData      [][]int8   `json:"-"`
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
	if s.Description.Generate {
		s.mapData = mg.GenMap(len(mt.Materials), ut.MAPSIZE)
	} else {
		//open map file
		filename := "assets/gamedata/maps/pregen.map"
		mapFile, err := os.Open(filename)
		if err != nil {
			return errors.New("map file could not be opened")
		}
		defer mapFile.Close()
		reader := bufio.NewReader(mapFile)
		//read lines of map file
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
			split := strings.Split(line, "")
			conv := make([]int8, len(split)-1)
			for i := 0; i < len(split)-1; i++ {
				tmp, err := strconv.ParseInt(split[i], 10, 8)
				if err != nil {
					return err
				}
				conv[i] = int8(tmp)
			}
			s.mapData = append(s.mapData, conv)
		}
	}
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
			if nx < 0 || ny < 0 || nx >= len(s.mapData) || ny >= len(s.mapData) {
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
