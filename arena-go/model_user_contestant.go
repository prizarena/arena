package arena

import (
	"github.com/strongo/slices"
)

// We need this file as players can play without participating in a tournament
// But we still want to collect stats

// UserEntity extends application specific `User` entity with arena properties
type UserEntity struct {
	BiddingUserIDs slices.CommaSeparatedUniqueValuesList `datastore:",noindex,omitempty"`
	ContestantStats
	RivalStatsEntity
	LastPlayIDs    slices.CommaSeparatedUniqueValuesList
}

func (u *UserEntity) UpdateArenaStats(tournamentID, rivalUserID, playID string, balanceDiff int) (updated bool){
	if u.LastPlayIDs.Contains(playID) {
		return
	}
	rivalKey := NewBattleID(tournamentID, rivalUserID)
	rivalStats := u.GetRivalStats()
	rival := rivalStats[rivalKey]
	rival.PlaysCount += 1
	rival.Balance += balanceDiff
	rivalStats[rivalKey] = rival
	u.SetRivalStats(rivalStats)
	u.BiddingUserIDs = u.BiddingUserIDs.Remove(rivalUserID)
	u.LastPlayIDs.Add(playID, 100)
	updated = true
	return
}
