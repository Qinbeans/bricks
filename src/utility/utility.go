package utility

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	OPTIONS = 0
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
	ScreenSizes Resolutions
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
	Resolution string  `json:"resolution"` //The resolution of the window
	Width      int     `json:"-"`          //The width of the window
	Height     int     `json:"-"`          //The height of the window
	FontSize   int     `json:"fontsize"`   //The size of the font
	FPS        int     `json:"fps"`        //The frames per second
	Scale      float64 `json:"scale"`      //The scale of the game
	Fullscreen bool    `json:"fullscreen"`
}

type Resolution struct {
	Width  int
	Height int
}

type Resolutions struct {
	Resolutions []Resolution
}

func (o Option) ScreenString() string {
	return fmt.Sprintf("%dx%d", o.Width, o.Height)
}

func (r Resolution) String() string {
	return fmt.Sprintf("%dx%d", r.Width, r.Height)
}

func UnformattedString(in string) string {
	params := strings.Split(in, ":")
	if len(params) != 2 {
		return ""
	}
	return params[1]
}

func UnformattedInt(in string) int {
	params := strings.Split(in, ":")
	if len(params) != 2 {
		return 0
	}
	i, err := strconv.Atoi(params[1])
	if err != nil {
		return 0
	}
	return i
}

func UnformattedFloat(in string) float64 {
	params := strings.Split(in, ":")
	if len(params) != 2 {
		return 0
	}
	i, err := strconv.ParseFloat(params[1], 64)
	if err != nil {
		return 0
	}
	return i
}

func (r Resolutions) Strings() []string {
	var res []string
	for _, v := range r.Resolutions {
		res = append(res, v.String())
	}
	return res
}

func (r Resolutions) String() string {
	var res string
	for i, v := range r.Resolutions {
		if i == len(r.Resolutions)-1 {
			res += v.String()
		} else {
			res += v.String() + ", "
		}
	}
	return res
}

func (r Resolutions) Contains(s string) bool {
	for _, v := range r.Resolutions {
		if v.String() == s {
			return true
		}
	}
	return false
}

func (r Resolutions) IndexOf(s string) (int, error) {
	for i, v := range r.Resolutions {
		if v.String() == s {
			return i, nil
		}
	}
	return -1, errors.New("resolution not found")
}

// Load options from game directory
func LoadOptions() error {
	var err error
	commonMonitors := Resolutions{
		[]Resolution{
			{640, 480},
			{800, 600},
			{1024, 768},
			{1280, 720},
			{1366, 768},
			{1600, 900},
			{1920, 1080},
			{2560, 1440},
			{3840, 2160},
			{7680, 4320},
		},
	}
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
	if !commonMonitors.Contains(GameOptions.Resolution) {
		GameOptions.Resolution = commonMonitors.Strings()[0]
	}
	resolution := strings.Split(GameOptions.Resolution, "x")
	GameOptions.Width, err = strconv.Atoi(resolution[0])
	if err != nil {
		return errors.New("options file could not be decoded")
	}
	GameOptions.Height, err = strconv.Atoi(resolution[1])
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
	//get current monitor size
	currMonitor := rl.GetCurrentMonitor()
	maxWidth := rl.GetMonitorWidth(currMonitor)
	maxHeight := rl.GetMonitorHeight(currMonitor)
	//check if current monitor size is in common monitor sizes
	for _, v := range commonMonitors.Resolutions {
		if v.Width <= maxWidth && v.Height <= maxHeight {
			//if it is, add it to the list of screen sizes
			ScreenSizes.Resolutions = append(ScreenSizes.Resolutions, v)
		}
	}
	return nil
}

func (o Option) Save() error {
	optionjson, err := os.Create(Config)
	if err != nil {
		return errors.New("options file could not be created")
	}
	defer optionjson.Close()
	encoder := json.NewEncoder(optionjson)
	err = encoder.Encode(&o)
	if err != nil {
		return errors.New("options file could not be encoded")
	}
	return nil
}

func GOReload() {
	//reload options
	err := LoadOptions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s (options)\n", err.Error())
	}
	rl.SetWindowSize(GameOptions.Width, GameOptions.Height)
	rl.SetTargetFPS(int32(GameOptions.FPS))
}
