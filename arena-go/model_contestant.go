package arena

import (
	"github.com/strongo/db"
	"strings"
	"time"
)

var ContestantKind = "ArenaContestant"

type ContestantEntity struct {
	TimeJoined   time.Time
	TournamentID string
	UserID       string
	Stranger     time.Time `datastore:",omitempty"`
	ContestantStats
}

type Contestant struct {
	db.StringID
	*ContestantEntity
}

type ContestantID string

const contestantIdSeparator = "@"

func (id ContestantID) UserID() string {
	s := string(id)
	if i := strings.Index(s, contestantIdSeparator); i > 0 {
		return s[:i]
	}
	return s
}

func NewContestantID(tournamentID, userID string) string {
	if tournamentID == "" {
		tournamentID = "*"
	}
	return userID + contestantIdSeparator + tournamentID
}

var _ db.EntityHolder = (*Tournament)(nil)

func (Contestant) Kind() string {
	return ContestantKind
}

func (Contestant) NewEntity() interface{} {
	return new(ContestantEntity)
}

func (t Contestant) Entity() interface{} {
	return t.ContestantEntity
}

func (t *Contestant) SetEntity(v interface{}) {
	if v == nil {
		t.ContestantEntity = nil
	} else {
		t.ContestantEntity = v.(*ContestantEntity)
	}
}
