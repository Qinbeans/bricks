package scenes

type Group struct {
	Name         string   `json:"name"`         //The name of the group
	Interactions []int    `json:"interactions"` //The indexes of the interactions in the group
	Actions      []Action `json:"actions"`      //The actions to be performed when the group is complete
	Complete     bool     `json:"-"`            //exclude from json
}

func (g *Group) Init() {
	for i := 0; i < len(g.Actions); i++ {
		g.Actions[i].Init()
	}
}
