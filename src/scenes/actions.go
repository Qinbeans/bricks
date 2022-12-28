package scenes

import (
	ut "bricks/utility"
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
	Type  int              `json:"type"`  //The type of action
	Value ut.GenGroupValue `json:"value"` //The value of the action
}

func (a *Action) Init() {
	a.Value.Init()
}
