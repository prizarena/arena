package models4arena

import "github.com/strongo/slices"

// ContestantStats can be used in User entity for storing games statistics
type ContestantStats struct {
	CountOfPlaysCompleted int                                   `datastore:",noindex,omitempty"`
	CountOfWins           int                                   `datastore:",noindex,omitempty"`
	CountOfDraws          int                                   `datastore:",noindex,omitempty"`
	CountOfLoses          int                                   `datastore:",noindex,omitempty"`
	Score                 int                                   `datastore:",omitempty"`
	RivalGameUserIDs      slices.CommaSeparatedUniqueValuesList `datastore:",noindex,omitempty"`
	LastPlayIDs           slices.CommaSeparatedValuesList       `datastore:",noindex,omitempty"`
}

func (cs ContestantStats) GetRivalUserIDs() slices.CommaSeparatedUniqueValuesList {
	return cs.RivalGameUserIDs
}
