package models4arena

import "github.com/pquerna/ffjson/ffjson"

type RivalStats map[BattleID]RivalStat

type RivalStatsEntity struct {
	RivalStats string
}

func (u *RivalStatsEntity) GetRivalStats() (rivalStats RivalStats) {
	rivalStats = make(RivalStats, 1)
	if u.RivalStats == "" {
		return
	}
	if err := ffjson.Unmarshal([]byte(u.RivalStats), &rivalStats); err != nil {
		panic(err)
	}
	return
}

func (u *RivalStatsEntity) SetRivalStats(rivalsStats RivalStats) {
	if len(rivalsStats) == 0 {
		u.RivalStats = ""
		return
	}
	if b, err := ffjson.Marshal(&rivalsStats); err != nil {
		panic(err)
	} else {
		u.RivalStats = string(b)
	}
}
