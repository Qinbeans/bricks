package scenes

import (
	ft "bricks/fonts"
	ut "bricks/utility"
	"errors"
	"fmt"
	"os"
	"strconv"

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
	Name     string         `json:"Name"`   //The name of the interaction
	X        float32        `json:"X"`      //The x coordinate of the interaction
	Y        float32        `json:"Y"`      //The y coordinate of the interaction
	Width    float32        `json:"Width"`  //The width of the interaction
	Height   float32        `json:"Height"` //The height of the interaction
	Type     int            `json:"Type"`   //The type of interaction
	Value    ut.GenIntValue `json:"Value"`  //The value of the interaction
	Modified bool           `json:"-"`      //exclude from json
	Hover    int            `json:"-"`      //exclude from json
	Group    *Group         `json:"-"`      //Used to check group
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
		number := i.Value
		color := rl.Red
		if i.Hover > 0 {
			color = rl.Green
		}
		text := i.Name + ": " + number.Current
		rl.DrawRectangleRounded(rl.NewRectangle(i.X, i.Y, i.Width+ut.FontSpacing*2, i.Height+ut.FontSpacing*2), 0.2, 0, color)
		rl.DrawTextEx(*ft.Fonts[0], text, rl.NewVector2(i.X+ut.FontSpacing, i.Y+ut.FontSpacing), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
	case CHOICE:
		choice := i.Value
		//draw choice
		rl.DrawRectangleRounded(rl.NewRectangle(i.X+float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Y, i.Width, i.Height), 0.2, 0, rl.Red)
		rl.DrawTextEx(*ft.Fonts[0], choice.GetFromValues(), rl.NewVector2(i.X+float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Y+ut.FontSpacing), float32(ut.GameOptions.FontSize), ut.FontSpacing, rl.White)
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
	switch i.Type {
	case CHOICE:
		choice := i.Value
		tmp, err := strconv.Atoi(choice.Current)
		choice.Value = float64(tmp)
		if err != nil {
			return nil, err
		}
		//multiple choice
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(i.X+i.Width+float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Y, float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Height)) {
			i.Hover = 1
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !i.Modified {
				choice.Value++
				if int(choice.Value) > len(choice.Values)-1 {
					choice.Value = 0
				}
				choice.Current = strconv.FormatFloat(choice.Value, 'f', 0, 64)
				//hot beauty
				i.Value = choice
				i.Modified = true
			} else if rl.IsMouseButtonUp(rl.MouseLeftButton) && i.Modified {
				i.Modified = false
			}
		} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(i.X, i.Y, float32(ut.GameOptions.FontSize)+ut.FontSpacing, i.Height)) {
			i.Hover = -1
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !i.Modified {
				choice.Value--
				if choice.Value < 0 {
					choice.Value = float64(len(choice.Values) - 1)
				}
				choice.Current = strconv.FormatFloat(choice.Value, 'f', 0, 64)
				i.Value = choice
				i.Modified = true
			} else if rl.IsMouseButtonUp(rl.MouseLeftButton) && i.Modified {
				i.Modified = false
			}
		} else {
			i.Hover = 0
		}
		size := rl.MeasureTextEx(*ft.Fonts[0], choice.GetFromValues(), float32(ut.GameOptions.FontSize), ut.FontSpacing)
		i.Width = size.X
	default:
		//not a multiple choice
		i.Hover = 1
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(i.X), float32(i.Y), float32(i.Width), float32(i.Height))) {
			if i.Type == SLIDER {
				//render width is fontSize * 4
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					//follow mouse
					// start := i.X + i.Width
					// end := i.X + i.Width + float32(ut.GameOptions.FontSize)*4
					//
				} else if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
					//save value
					start := i.X + i.Width
					end := i.X + i.Width + float32(ut.GameOptions.FontSize)*4
					pos := rl.GetMousePosition().X
					value := (pos - start) / (end - start)
					i.Value.Current = strconv.FormatFloat(float64(value), 'f', 2, 64)
				}
			} else {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !i.Modified {
					i.Modified = true
					switch i.Type {
					case EXIT:
						//exit
						return nil, errors.New("exit")
					case RENDER:
						//spawn scene
						//load scene
						fmt.Fprintf(os.Stderr, "Render: %v\n", i.Value)
						render := i.Value
						scene, err := LoadScene(render.Current)
						if err != nil {
							return nil, err
						}
						return scene, nil
					case NUMBER:
						//integer
						//TODO
						//modifiable number
					case GROUP:
						//Perform group actions
						for _, action := range i.Group.Actions {
							switch action.Type {
							case SAVE:
								save := action.Value
								loc := save.Input
								val := save.Value
								//save value
								switch val {
								case ut.OPTIONS:
									//save to file as options
									ut.GameOptions.Resolution = interactions[0].Value.GetFromValues()
									ut.GameOptions.FontSize = int(interactions[1].Value.Value)
									ut.GameOptions.FPS = int(interactions[2].Value.Value)
									ut.GameOptions.Scale = interactions[3].Value.Value
									ut.GameOptions.Fullscreen = int(interactions[4].Value.Value) == 1

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
								function := action.Value
								ut.FuncMap[function.Input]()
							case WAIT:
							default:
							}
						}
					case DIR:
						//directory
					case TOGGLE:
						//toggle
						//TODO
					default:
						//error
						//TODO
					}
				} else if rl.IsMouseButtonReleased(rl.MouseLeftButton) && i.Modified {
					i.Modified = false
				}
			}
		} else {
			i.Hover = 0
		}
	}
	return nil, nil
}

func (i *Interaction) Init() (int, error) {
	//fill values appropriately
	i.Value.Init()
	if i.Type == GROUP {
		//look for group id
		group := i.Value
		fmt.Fprintf(os.Stderr, "Interaction value: %d\n", int(group.Value))
		index := int(group.Value)
		return index, nil
	}
	return -1, nil
}
