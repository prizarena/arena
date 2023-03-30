package models4arena

// ContestantStats can be used in User entity for storing games statistics
type ContestantStats struct {
	CountOfPlaysCompleted int    `datastore:",noindex,omitempty"`
	CountOfWins           int    `datastore:",noindex,omitempty"`
	CountOfDraws          int    `datastore:",noindex,omitempty"`
	CountOfLoses          int    `datastore:",noindex,omitempty"`
	Score                 int    `datastore:",omitempty"`
	RivalGameUserIDs      string `datastore:",noindex,omitempty"`
	LastPlayIDs           string `datastore:",noindex,omitempty"`
}

func (cs ContestantStats) GetRivalUserIDs() string {
	return cs.RivalGameUserIDs
}
