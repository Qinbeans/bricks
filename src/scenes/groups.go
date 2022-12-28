package scenes

type Group struct {
	Name         string   `json:"Name"`         //The name of the group
	Interactions []int    `json:"Interactions"` //The indexes of the interactions in the group
	Actions      []Action `json:"Actions"`      //The actions to be performed when the group is complete
	Complete     int      `json:"-"`            //exclude from json
}

func (g *Group) Init() {
	for i := 0; i < len(g.Actions); i++ {
		g.Actions[i].Init()
	}
}
