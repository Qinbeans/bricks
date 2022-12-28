package utility

import (
	"strconv"
	"strings"
)

var (
	FuncMap = map[string]func(){
		"$GORELOAD": GOReload,
	}
)

//Interactions

type GenIntValue struct {
	Current string
	Value   float64
	Min     float64
	Max     float64
	Step    float64
	Values  []string
}

func (n *GenIntValue) Init() {
	for _, v := range n.Values {
		switch v {
		case "$SCREENS":
			n.Values = ScreenSizes.Strings()
		}
	}
	switch n.Current {
	case "$GOSCALE":
		n.Current = strconv.FormatFloat(GameOptions.Scale, 'f', 2, 64)
		n.Value = GameOptions.Scale
	case "$GOFPS":
		n.Current = strconv.Itoa(GameOptions.FPS)
		n.Value = float64(GameOptions.FPS)
	case "$GOWIDTH":
		n.Current = strconv.Itoa(GameOptions.Width)
		n.Value = float64(GameOptions.Width)
	case "$GOHEIGHT":
		n.Current = strconv.Itoa(GameOptions.Height)
		n.Value = float64(GameOptions.Height)
	case "$GOFONTSIZE":
		n.Current = strconv.Itoa(GameOptions.FontSize)
		n.Value = float64(GameOptions.FontSize)
	case "$GOSCREEN":
		tmp, err := ScreenSizes.IndexOf(GameOptions.ScreenString())
		if err != nil {
			n.Current = "0"
			n.Value = 0
		} else {
			n.Current = strconv.Itoa(tmp)
			n.Value = float64(tmp)
		}
	default:
		n.Value, _ = strconv.ParseFloat(n.Current, 64)
	}
}

func (n *GenIntValue) GetFromValues() string {
	if len(n.Values) > 0 {
		return n.Values[int(n.Value)]
	}
	return n.Current
}

type GenGroupValue struct {
	Scene string
	Input string
	Value int
}

func (n *GenGroupValue) Init() {
	if strings.Contains(n.Input, "$CONFIG") {
		path := strings.Split(n.Input, "/")
		for k := 0; k < len(path); k++ {
			if strings.Contains(path[k], "$CONFIG") {
				path[k] = Config
			}
		}
		n.Input = strings.Join(path, "/")
	}
}
