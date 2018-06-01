package canvas

import (
	"github.com/strongo/db"
	"time"
	"github.com/strongo-games/arena/arena-go"
)

const BoardKind = "B"

type BoardEntity struct {
	Created   time.Time
	Round     int                            `datastore:",noindex"`
	UserIDs   []string
	UserTimes []time.Time                    `datastore:",noindex"`
	UserMoves arena.CommaSeparatedValuesList `datastore:",noindex"`
}

type Board struct {
	db.StringID
	*BoardEntity
}

var _ db.EntityHolder = (*Board)(nil)

func (Board) Kind() string {
	return BoardKind
}

func (canvas *Board) SetEntity(v interface{}) {
	canvas.BoardEntity = v.(*BoardEntity)
}

func (canvas Board) Entity() interface{} {
	return canvas.BoardEntity
}

func (canvas Board) NewEntity() interface{} {
	return &BoardEntity{}
}
