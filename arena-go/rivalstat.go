package arena

//go:generate ffjson $GOFILE

type RivalStat struct {
	Balance    int `json:",omitempty"`
	GamesCount int `json:",omitempty"`
}
