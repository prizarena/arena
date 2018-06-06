package arena

import (
	"github.com/strongo/slices"
)

type UserContestantEntity struct {
	BiddingUserIDs slices.CommaSeparatedUniqueValuesList `datastore:",noindex,omitempty"`
	ContestantStats
	RivalStatsEntity
	LastGameID string
}

func (u *UserContestantEntity) UpdateArenaStats(tournamentID, rivalUserID, gameID string, balanceDiff int) {
	rivalKey := NewBattleID(tournamentID, rivalUserID)
	rivalStats := u.GetRivalStats()
	rival := rivalStats[rivalKey]
	rival.GamesCount += 1
	rival.Balance += balanceDiff
	rivalStats[rivalKey] = rival
	u.SetRivalStats(rivalStats)
	u.BiddingUserIDs = u.BiddingUserIDs.Remove(rivalUserID)
	u.LastGameID = gameID
}
