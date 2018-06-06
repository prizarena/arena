package arena

import (
	"github.com/strongo/db"
	"github.com/strongo/decimal"
	"strconv"
	"time"
	"strings"
	"github.com/strongo/slices"
)

type TournamentEntity struct {
	CreatorGameUserID     string
	CreatorUserID         string
	GameID                string
	Status                string
	Name                  string                          `datastore:",noindex"`
	Note                  string                          `datastore:",noindex"`
	Created               time.Time
	Starts                time.Time
	Ends                  time.Time
	MinGamesToScore       int                             `datastore:",noindex,omitempty" json:",omitempty"`
	IsDiscoverable        bool                            `datastore:",omitempty" json:",omitempty"`
	SponsorshipIsActive   bool                            `datastore:",omitempty" json:",omitempty"`
	SponsorName           string                          `datastore:",noindex,omitempty" json:",omitempty"`
	SponsorUrl            string                          `datastore:",noindex,omitempty" json:",omitempty"`
	SponsorText           string                          `datastore:",noindex,omitempty" json:",omitempty"`
	SponsorPrizeMedium    string                          `datastore:",noindex,omitempty" json:",omitempty"` // e.g. 'amazon_gift_card', 'apple_gift_card', etc.
	SponsorPrizeCurrency  string                          `datastore:",noindex,omitempty" json:",omitempty"`
	SponsorPrizeValue     decimal.Decimal64p2             `datastore:",noindex,omitempty" json:",omitempty"`
	CountOfContestants    int                             `datastore:",noindex,omitempty" json:",omitempty"`
	CountOfPlaysCompleted int                             `datastore:",noindex,omitempty" json:",omitempty"`
	LastPlayIDs           slices.CommaSeparatedValuesList `datastore:",noindex,omitempty" json:",omitempty"`
	// SponsorHttpReferers  string              `datastore:",noindex,omitempty"`
}

var TournamentKind = "ArenaTournament"

const TournamentStarID = "*"

type Tournament struct {
	db.StringID
	*TournamentEntity
}

var _ db.EntityHolder = (*Tournament)(nil)

func (Tournament) Kind() string {
	return TournamentKind
}

func (Tournament) NewEntity() interface{} {
	return new(TournamentEntity)
}

func (t Tournament) Entity() interface{} {
	return t.TournamentEntity
}

func (t *Tournament) SetEntity(v interface{}) {
	if v == nil {
		t.TournamentEntity = nil
	} else {
		t.TournamentEntity = v.(*TournamentEntity)
	}
}

func IsMonthlyTournamentID(tournamentID string) bool {
	if len(tournamentID) <= 7 || strings.Count(tournamentID, ":") != 1 {
		return false
	}
	if tournamentID = strings.Split(tournamentID, ":")[1]; len(tournamentID) == 6 {
		if v, err := strconv.ParseInt(tournamentID, 10, 32); err == nil && v > 201801 {
			return true
		}
	}
	return false
}
