package events

type Challenge struct {
	Type   string `json:"type"`   //The type of challenge
	Target string `json:"target"` //The target of the challenge
}

type EventDesc struct {
	Text       string      `json:"text"`
	Challenges []Challenge `json:"challenges"`
}

type Event struct {
	Name        string    `json:"name"`
	Description EventDesc `json:"description"`
}
