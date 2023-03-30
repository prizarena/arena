package models4arena

import "strings"

// BattleID has form of 'userId@tournamentId'
type BattleID string

func (k BattleID) IDs() (userID string, tournamentID string) {
	ids := strings.Split(string(k), rivalKeySeparator)
	return ids[0], ids[1]
}

func (k BattleID) RivalID() (rivalID string) {
	rivalID, _ = k.IDs()
	return rivalID
}

func (k BattleID) IsStranger() bool {
	return len(k) > 1 && k[0:1] == RivalKeyStranger
}

const rivalKeySeparator = "@"

func NewBattleID(tournamentID, userID string) BattleID {
	if tournamentID == "" {
		tournamentID = "*"
	}
	return BattleID(userID + rivalKeySeparator + tournamentID)
}

const RivalKeyStranger = "*"

func NewStrangerBattleID(tournamentID string) BattleID {
	return NewBattleID(tournamentID, RivalKeyStranger)
}
