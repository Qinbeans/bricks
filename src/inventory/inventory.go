package inventory

import (
	item "bricks/items"
)

type Armor struct {
	Head     item.Item `json:"head"`
	Body     item.Item `json:"body"`
	LeftArm  item.Item `json:"leftarm"`
	RightArm item.Item `json:"rightarm"`
	LeftLeg  item.Item `json:"leftleg"`
	RightLeg item.Item `json:"rightleg"`
}

type Equipment struct {
	Armor  Armor     `json:"armor"`
	Weapon item.Item `json:"weapon"`
}

type Inventory struct {
	Items []item.Item `json:"items"`
	Rows  int         `json:"rows"`
	Cols  int         `json:"cols"`
	Bag   item.Item   `json:"bag"`
}
