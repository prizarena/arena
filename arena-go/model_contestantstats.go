package arena

// ContestantStats can be used in User entity for storing games statistics
type ContestantStats struct {
	CountOfGames int                            `datastore:",noindex,omitempty"`
	RivalUserIDs CommaSeparatedUniqueValuesList `datastore:",noindex,omitempty"`
	Score        int                            `datastore:",omitempty"`
}

func (cs ContestantStats) GetRivalUserIDs() CommaSeparatedUniqueValuesList {
	return cs.RivalUserIDs
}
