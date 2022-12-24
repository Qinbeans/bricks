package players

import (
	inv "bricks/inventory"
	ut "bricks/utility"
	"bufio"
	"encoding/json"
	"errors"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Name      string        `json:"name"`
	Stats     ut.Stats      `json:"stats"`
	Inventory inv.Inventory `json:"inventory"`
	TexSkin   string        `json:"skin"`
	Width     int           `json:"width"`
	Height    int           `json:"height"`
	//in-game
	Current bool         `json:"current"`
	Pos     rl.Vector2   `json:"pos"`
	Skin    rl.Texture2D `json:"-"`
}

func LoadPlayer(world string) (*Player, error) {
	var p *Player
	playerfile := ""
	if _, err := os.Stat(ut.Config); os.IsNotExist(err) {
		//make game directory
		err := os.Mkdir(ut.Config, 0755)
		if err != nil {
			return nil, errors.New("game directory could not be created")
		}
	}
	if _, err := os.Stat(ut.Config + world); os.IsNotExist(err) {
		//make world directory
		err := os.Mkdir(ut.Config+world, 0755)
		if err != nil {
			return nil, errors.New("world directory could not be created")
		}
	}
	if _, err := os.Stat(playerfile); os.IsNotExist(err) {
		//make player file
		playerjson, err := os.Create(playerfile)
		if err != nil {
			return nil, errors.New("player file could not be created")
		}
		defer playerjson.Close()
		writer := bufio.NewWriter(playerjson)
		encoder := json.NewEncoder(writer)
		p = &Player{
			Name:      "Player",
			Stats:     ut.Stats{},
			Inventory: inv.Inventory{},
			TexSkin:   "player.png",
			Width:     32,
			Height:    32,
			Current:   true,
			Pos:       rl.Vector2{X: 0, Y: 0},
		}
		err = encoder.Encode(p)
		if err != nil {
			return nil, errors.New("player file could not be encoded")
		}
		return p, nil
	}
	playerjson, err := os.Open(playerfile)
	if err != nil {
		return nil, errors.New("player file could not be opened")
	}
	defer playerjson.Close()
	reader := bufio.NewReader(playerjson)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(p)
	if err != nil {
		return nil, errors.New("player file could not be decoded")
	}
	return p, nil
}

func NewPlayer(name string, stats ut.Stats, inventory inv.Inventory, path string, width int, height int) (*Player, error) {
	p := &Player{
		Name:      name,
		Stats:     stats,
		Inventory: inventory,
		TexSkin:   path,
		Width:     width,
		Height:    height,
		Current:   true,
		Pos:       rl.Vector2{X: 0, Y: 0},
	}
	return p, nil
}

func (p *Player) SavePlayer(world string) error {
	//open file
	//check os type
	playerfile := ""
	if _, err := os.Stat(ut.Config); os.IsNotExist(err) {
		err := os.Mkdir(ut.Config, 0755)
		if err != nil {
			return errors.New("game directory could not be created")
		}
	}
	if _, err := os.Stat(ut.Config + world); os.IsNotExist(err) {
		err := os.Mkdir(ut.Config+world, 0755)
		if err != nil {
			return errors.New("world directory could not be created")
		}
	}
	if _, err := os.Stat(playerfile); os.IsNotExist(err) {
		_, err := os.Create(playerfile)
		if err != nil {
			return errors.New("player file could not be created")
		}
	}
	playerjson, err := os.OpenFile(playerfile, os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("player file could not be opened")
	}
	defer playerjson.Close()
	writer := bufio.NewWriter(playerjson)
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(p)
	if err != nil {
		return errors.New("player file could not be encoded")
	}
	return nil
}

func (p *Player) Load() {
	p.Skin = rl.LoadTexture(p.TexSkin)
}

func (p *Player) Unload() {
	rl.UnloadTexture(p.Skin)
}
