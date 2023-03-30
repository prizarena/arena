package models4arena

import (
	"github.com/strongo/slice"
	"strings"
)

// We need this file as players can play without participating in a tournament
// But we still want to collect stats

// UserEntity extends application specific `User` entity with arena properties
type UserEntity struct {
	BiddingUserIDs string `datastore:",noindex,omitempty"`
	ContestantStats
	RivalStatsEntity
	LastPlayIDs string `datastore:",noindex,omitempty"`
}

func (u *UserEntity) UpdateArenaStats(tournamentID, rivalUserID, playID string, balanceDiff int) (updated bool) {
	if slice.Index(strings.Split(u.LastPlayIDs, ","), playID) >= 0 {
		return
	}
	rivalKey := NewBattleID(tournamentID, rivalUserID)
	rivalStats := u.GetRivalStats()
	rival := rivalStats[rivalKey]
	rival.PlaysCount += 1
	rival.Balance += balanceDiff
	rivalStats[rivalKey] = rival
	u.SetRivalStats(rivalStats)
	//u.BiddingUserIDs = u.BiddingUserIDs.Remove(rivalUserID)
	//u.LastPlayIDs.Add(playID, 100)
	updated = true
	return
}
