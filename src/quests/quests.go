package quests

import (
	item "bricks/items"
)

type Challenge struct {
	Type   int    `json:"type"`   //The type of challenge
	Target string `json:"target"` //The target of the challenge
	//in-game
	Completed int `json:"completed"` //If the challenge is completed
}

type Quest struct {
	Name        string    `json:"name"`        //The name of the quest
	Description string    `json:"description"` //The description of the quest
	ReqLevel    int       `json:"reqlevel"`    //The required level to start the quest
	Prize       item.Item `json:"prize"`       //The prize of the quest
	//in-game
	Completed bool `json:"-"`
}
