package models4arena

//go:generate ffjson $GOFILE

type RivalStat struct {
	Balance    int `json:",omitempty"`
	PlaysCount int `json:",omitempty"`
}
