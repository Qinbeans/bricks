package scenes

import (
	ft "bricks/fonts"
	ut "bricks/utility"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//Interactions
// Types:
//  -1 = Exit,
//  0 = Spawn Scene,
//  1 = integer,
//  2 = multiple choice format (value1,value2,value3,...:index),
//  3 = group (group_name) and manages the group,
//  4 = multiple choice from directory and will look for
//      files with specified extension(directory:extension),
//  5 = slider format (min,max,step:value)
//  6 = checkbox 0, 1, or -1, where -1 refers to the default value

const (
	EXIT   = -1
	RENDER = 0
	NUMBER = 1
	CHOICE = 2
	GROUP  = 3
	DIR    = 4
	SLIDER = 5
	TOGGLE = 6
)

type Interaction struct {
	Name     string  `json:"name"`   //The name of the interaction
	X        float32 `json:"x"`      //The x coordinate of the interaction
	Y        float32 `json:"y"`      //The y coordinate of the interaction
	Width    float32 `json:"width"`  //The width of the interaction
	Height   float32 `json:"height"` //The height of the interaction
	Type     int     `json:"type"`   //The type of interaction
	Value    string  `json:"value"`  //The value of the interaction
	Complete bool    `json:"-"`      //exclude from json
	Hover    int     `json:"-"`      //exclude from json
	Group    *Group  `json:"-"`      //Used to check group
}

func (i *Interaction) Draw() {
	switch i.Type {
	case EXIT:
		//exit
		color := rl.Red
		if i.Hover > 0 {
			color = rl.Green
		}
		rl.DrawRectangleRounded(rl.NewRectangle(i.X, i.Y, i.Width+ut.FontSpacing*2, i.Height+ut.FontSpacing*2), 0.2, 0, color)
		rl.DrawTextEx(*ft.Fonts[0], "Exit", rl.NewVector2(i.X+ut.FontSpacing, i.Y+ut.FontSpacing), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
	case RENDER:
		//spawn scene
		color := rl.Red
		if i.Hover > 0 {
			color = rl.Green
		}
		rl.DrawRectangleRounded(rl.NewRectangle(i.X, i.Y, i.Width+ut.FontSpacing*2, i.Height+ut.FontSpacing*2), 0.2, 0, color)
		rl.DrawTextEx(*ft.Fonts[0], i.Name, rl.NewVector2(i.X+ut.FontSpacing, i.Y+ut.FontSpacing), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
	case NUMBER:
		//integer
		color := rl.Red
		if i.Hover > 0 {
			color = rl.Green
		}
		text := i.Name + ": " + i.Value
		rl.DrawRectangleRounded(rl.NewRectangle(i.X, i.Y, i.Width+ut.FontSpacing*2, i.Height+ut.FontSpacing*2), 0.2, 0, color)
		rl.DrawTextEx(*ft.Fonts[0], text, rl.NewVector2(i.X+ut.FontSpacing, i.Y+ut.FontSpacing), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
	case CHOICE:
		split := strings.Split(i.Value, ":")
		choices := strings.Split(split[0], ",")
		choice, err := strconv.Atoi(split[1])
		if err != nil {
			//error
			rl.DrawRectangleRounded(rl.NewRectangle(i.X, i.Y, i.Width, i.Height), 0.2, 0, rl.Red)
			rl.DrawTextEx(*ft.Fonts[0], "Error", rl.NewVector2(i.X+ut.FontSpacing, i.Y+ut.FontSpacing), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
		}
		//draw choice
		rl.DrawRectangleRounded(rl.NewRectangle(i.X+float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Y, i.Width, i.Height), 0.2, 0, rl.Red)
		rl.DrawTextEx(*ft.Fonts[0], choices[choice], rl.NewVector2(i.X+float32(ut.GameOptions.FontSize)+ut.FontSpacing+ut.FontSpacing, i.Y+ut.FontSpacing), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
		//slides left and right
		//left arrow
		if i.Hover < 0 {
			rl.DrawRectangleRounded(rl.NewRectangle(i.X, i.Y, float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Height), 0.2, 0, rl.Green)
		} else {
			rl.DrawRectangleRounded(rl.NewRectangle(i.X, i.Y, float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Height), 0.2, 0, rl.Red)
		}
		rl.DrawTextEx(*ft.Fonts[0], "<", rl.NewVector2(i.X+ut.FontSpacing, i.Y+ut.FontSpacing), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
		//right arrow
		if i.Hover > 0 {
			rl.DrawRectangleRounded(rl.NewRectangle(i.X+i.Width+float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Y, float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Height), 0.2, 0, rl.Green)
		} else {
			rl.DrawRectangleRounded(rl.NewRectangle(i.X+i.Width+float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Y, float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Height), 0.2, 0, rl.Red)
		}
		rl.DrawTextEx(*ft.Fonts[0], ">", rl.NewVector2(i.X+i.Width+float32(ut.GameOptions.FontSize)+ut.FontSpacing+ut.FontSpacing, i.Y+ut.FontSpacing), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)

	case GROUP:
		//group
		color := rl.Red
		if i.Hover > 0 {
			color = rl.Green
		}
		rl.DrawRectangleRounded(rl.NewRectangle(i.X, i.Y, i.Width+ut.FontSpacing*2, i.Height+ut.FontSpacing*2), 0.2, 0, color)
		rl.DrawTextEx(*ft.Fonts[0], i.Name, rl.NewVector2(i.X+ut.FontSpacing, i.Y+ut.FontSpacing), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
	case DIR:
		//directory
	case SLIDER:
		//slider
	case TOGGLE:
		//toggle
	default:
		//error
		rl.DrawRectangleRounded(rl.NewRectangle(float32(i.X), float32(i.Y), float32(i.Width), float32(i.Height)), 0.2, 0, rl.Red)
		rl.DrawText("Error", int32(i.X+5), int32(i.Y+5), 20, rl.White)
	}
}

func (i *Interaction) Update(interactions []Interaction) (*Scene, error) {
	if i.Type != 2 {
		//not a multiple choice
		i.Hover = 1
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(i.X), float32(i.Y), float32(i.Width), float32(i.Height))) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				switch i.Type {
				case EXIT:
					//exit
					return nil, errors.New("exit")
				case RENDER:
					//spawn scene
					//load scene
					scene, err := LoadScene(i.Value)
					if err != nil {
						return nil, err
					}
					return scene, nil
				case NUMBER:
					//integer
					//TODO
				case GROUP:
					//Perform group actions
					for _, action := range i.Group.Actions {
						params := strings.Split(action.Value, ":")
						switch action.Type {
						case SAVE:
							loc := params[0]
							val, err := strconv.Atoi(params[1])
							if err != nil {
								return nil, err
							}
							//save value
							switch val {
							case ut.OPTIONS:
								//save to file as options
								ut.GameOptions.Resolution = ut.ScreenSizes.Strings()[ut.UnformattedInt(interactions[0].Value)]
								ut.GameOptions.FontSize = ut.UnformattedInt(interactions[1].Value)
								ut.GameOptions.FPS = ut.UnformattedInt(interactions[2].Value)
								ut.GameOptions.Scale = ut.UnformattedFloat(interactions[3].Value)
								ut.GameOptions.Fullscreen = interactions[4].Value == "true"

								//save to file
								err := ut.GameOptions.Save()
								if err != nil {
									return nil, err
								}
							default:
								fmt.Fprintf(os.Stderr, "Invalid save location: %s", loc)
							}
						case FUNCALL:
							//call function
							switch params[0] {
							case "$GORELOAD":
								//reload game
								ut.GOReload()
								//prob should reload scene
							default:
							}
						case WAIT:
						default:
						}
					}
				case DIR:
					//directory
				case CHOICE:
					//multiple choice from directory
					//TODO
				case SLIDER:
					//slider
					//TODO
				case TOGGLE:
					//toggle
					//TODO
				default:
					//error
					//TODO
				}
			}
		} else {
			i.Hover = 0
		}
	} else {
		//multiple choice
		split := strings.Split(i.Value, ":")
		choices := strings.Split(split[0], ",")
		choice, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(i.X+i.Width+float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Y, float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Height)) {
			i.Hover = 1
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !i.Complete {
				choice++
				if choice > len(choices)-1 {
					choice = 0
				}
				i.Value = split[0] + ":" + strconv.Itoa(choice)
				i.Complete = true
			} else if rl.IsMouseButtonUp(rl.MouseLeftButton) && i.Complete {
				i.Complete = false
			}
		} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(i.X, i.Y, float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Height)) {
			i.Hover = -1
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !i.Complete {
				choice--
				if choice < 0 {
					choice = len(choices) - 1
				}
				i.Value = split[0] + ":" + strconv.Itoa(choice)
				i.Complete = true
			} else if rl.IsMouseButtonUp(rl.MouseLeftButton) && i.Complete {
				i.Complete = false
			}
		} else {
			i.Hover = 0
		}
	}

	return nil, nil
}

func (i *Interaction) Init() (int, error) {
	//fill values appropriately
	if i.Type == 3 {
		//look for group id
		fmt.Fprintf(os.Stderr, "Interaction value: %s\n", i.Value)
		index, err := strconv.Atoi(i.Value)
		if err != nil {
			return -1, err
		}
		return index, nil
	}
	split := strings.Split(i.Value, ":")

	for j := 0; j < len(split); j++ {
		//look for GameData variables
		if strings.Contains(split[j], "$SCREENS") {
			split[j] = ut.ScreenSizes.String()
		} else if strings.Contains(split[j], "$GOSCREEN") {
			//go through screens and find the one that matches the current screen
			screen := ut.GameOptions.ScreenString()
			index, err := ut.ScreenSizes.IndexOf(screen)
			if err != nil {
				index = 0
			}
			split[j] = strconv.Itoa(index)
		} else if strings.Contains(split[j], "$GOFULLSCREEN") {
			split[j] = strconv.FormatBool(ut.GameOptions.Fullscreen)
		} else if strings.Contains(split[j], "$GOFPS") {
			split[j] = strconv.Itoa(ut.GameOptions.FPS)
		} else if strings.Contains(split[j], "$GOFONT") {
			split[j] = strconv.Itoa(ut.GameOptions.FontSize)
		} else if strings.Contains(split[j], "$GOSCALE") {
			split[j] = strconv.FormatFloat(ut.GameOptions.Scale, 'f', 2, 64)
		} else if strings.Contains(split[j], "$CONFIG") {
			split2 := strings.Split(split[j], "/")
			for k := 0; k < len(split2); k++ {
				if strings.Contains(split2[k], "$CONFIG") {
					split2[k] = ut.Config
				}
			}
			split[j] = strings.Join(split2, "/")
		}
	}
	if len(split) > 1 {
		i.Value = strings.Join(split, ":")
	}
	return -1, nil
}
