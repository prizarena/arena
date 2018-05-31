package arena

import (
	"context"
	"github.com/strongo/db"
)

type TournamentDal interface {
	FindStranger(c context.Context, tournamentID, userID string, friends []string) (strangerID string, err error)
}

type User interface {
	db.EntityHolder
	GetRivalUserIDs() CommaSeparatedUniqueValuesList
}
