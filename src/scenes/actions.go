package scenes

import (
	ut "bricks/utility"
	"strconv"
	"strings"
)

// 0 = save to file
// 1 = perform function
// 2 = wait for associated interaction to complete

const (
	SAVE    = 0
	FUNCALL = 1
	WAIT    = 2
)

type Action struct {
	Type  int    `json:"type"`  //The type of action
	Value string `json:"value"` //The value of the action
}

func (a *Action) Init() {
	split := strings.Split(a.Value, ":")
	for j := 0; j < len(split); j++ {
		//look for GameData variables
		if strings.Contains(split[j], "$SCREENS") {
			split[j] = ut.ScreenSizes.String()
		} else if strings.Contains(split[j], "$GOSCREEN") {
			split[j] = ut.GameOptions.ScreenString()
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
	a.Value = strings.Join(split, ":")
}
