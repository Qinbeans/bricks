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
	params := strings.Split(a.Value, ":")
	for j := 0; j < len(params); j++ {
		if strings.Contains(params[j], "$CONFIG") {
			path := strings.Split(params[j], "/")
			for k := 0; k < len(path); k++ {
				if strings.Contains(path[k], "$CONFIG") {
					path[k] = ut.Config
				}
			}
			params[j] = strings.Join(path, "/")
		} else {
			switch params[j] {
			case "$SCREENS":
				params[j] = ut.ScreenSizes.String()
			case "$GOSCREEN":
				//go through screens and find the one that matches the current screen
				screen := ut.GameOptions.ScreenString()
				index, err := ut.ScreenSizes.IndexOf(screen)
				if err != nil {
					index = 0
				}
				params[j] = strconv.Itoa(index)
			case "$GOFULLSCREEN":
				params[j] = strconv.FormatBool(ut.GameOptions.Fullscreen)
			case "$GOFPS":
				params[j] = strconv.Itoa(ut.GameOptions.FPS)
			case "$GOFONT":
				params[j] = strconv.Itoa(ut.GameOptions.FontSize)
			case "$GOSCALE":
				params[j] = strconv.FormatFloat(ut.GameOptions.Scale, 'f', 2, 64)
			}
		}
	}
	a.Value = strings.Join(params, ":")
}
