package utility

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	RENDERSIZE  = 32
	DEFSCALE    = float64(0.01)
	DEFHEIGHT   = 450
	DEFWIDTH    = 800
	MAPSIZE     = 10000
	TITLE       = "Bricks"
	DEFFPS      = 60
	DEFFONTSIZE = 32
	FONTMOD     = 43
	DELAY       = 8
	DEFFS       = false
)

var (
	GameOptions Option
	Step        float64
	Modwidth    int
	Modheight   int
	Heartbeat   int
	BPS         int
	Config      string
	FontSpacing float32
)

type Stats struct {
	Level        int `json:"level"` //The level of the item/character
	Strength     int `json:"str"`
	Dexterity    int `json:"dex"`
	Agility      int `json:"agi"`
	Intelligence int `json:"int"`
	Charisma     int `json:"cha"`
	Luck         int `json:"lck"`
	Vitality     int `json:"vit"`
	Defense      int `json:"def"`
	Mana         int `json:"man"`
	Health       int `json:"hp"`
	Stamina      int `json:"stam"`
	Mp           int `json:"mp"`
}

type Dialogue struct {
	Text       string   `json:"text"`
	Selections []string `json:"selections"`
}

type Option struct {
	Width      int     `json:"width"`    //The width of the window
	Height     int     `json:"height"`   //The height of the window
	FontSize   int     `json:"fontsize"` //The size of the font
	FPS        int     `json:"fps"`      //The frames per second
	Scale      float64 `json:"scale"`    //The scale of the game
	Fullscreen bool    `json:"fullscreen"`
}

// Load options from game directory
func LoadOptions() error {
	var err error

	Config, err = os.UserConfigDir()
	if err != nil {
		return errors.New("could not get user config directory")
	}
	Config += "/bricks/"

	err = os.MkdirAll(Config, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\nConfig: %s\n", err.Error(), Config)
		return errors.New("options directory could not be created")
	}
	Config += "options.json"
	optionjson, err := os.Open(Config)
	if err != nil {
		//create file
		GameOptions = Option{Width: DEFWIDTH, Height: DEFHEIGHT, FPS: DEFFPS, Scale: DEFSCALE, FontSize: DEFFONTSIZE, Fullscreen: DEFFS}
		optionjson, err := os.Create(Config)
		if err != nil {
			return errors.New("options file could not be created")
		}
		encoder := json.NewEncoder(optionjson)
		err = encoder.Encode(&GameOptions)
		if err != nil {
			return errors.New("options file could not be encoded")
		}
	}
	defer optionjson.Close()
	decoder := json.NewDecoder(optionjson)
	err = decoder.Decode(&GameOptions)
	if err != nil {
		return errors.New("options file could not be decoded")
	}
	if GameOptions.Width == 0 {
		GameOptions.Width = DEFWIDTH
	}
	if GameOptions.Height == 0 {
		GameOptions.Height = DEFHEIGHT
	}
	if GameOptions.FPS == 0 {
		GameOptions.FPS = DEFFPS
	}
	if GameOptions.Scale == 0 {
		GameOptions.Scale = DEFSCALE
	}
	if GameOptions.FontSize == 0 {
		GameOptions.FontSize = DEFFONTSIZE
	}
	Step = float64(GameOptions.Width) * float64(GameOptions.Scale)
	Modwidth = GameOptions.Width / int(Step)
	Modheight = GameOptions.Height / int(Step)
	BPS = GameOptions.FPS / DELAY
	FontSpacing = float32(GameOptions.FontSize) / FONTMOD
	return nil
}
