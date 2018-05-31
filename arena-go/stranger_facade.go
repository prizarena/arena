package arena

import (
	"context"
	"github.com/pkg/errors"
	"github.com/strongo/db"
	"github.com/strongo/log"
	"strconv"
	"time"
)

var DB db.Database
var TournamentDAL TournamentDal

var ErrRivalUserIsNotBiddingAgainstStranger = errors.New("rival user is not a stranger")

func VerifyUserAndTorunamentIDs(userID string, tournamentID *string) (err error) {
	if *tournamentID == "" || IsMonthlyTournamentID(*tournamentID) {
		tID := TournamentStarID
		tournamentID = &tID
	}
	if userID == "" {
		err = errors.New("Parameter userID is empty string")
		return
	}
	return
}

func MakeMoveAgainstStranger(
	c context.Context, now time.Time,
	tournamentID string,
	user User,
	onRivalFound func(rivalUserID string) error,
	onStranger func(contestant *Contestant) error,
) (err error) {
	var rivalUserIDs []string

	contestant := new(Contestant)

	userID := user.StrID()
	if userID == "" {
		userID = strconv.FormatInt(user.IntID(), 10)
	}

	contestant.ID = NewContestantID(tournamentID, userID)

	if err = DB.Get(c, contestant); err != nil {
		if db.IsNotFound(err) {
			if err = DB.Get(c, user); err != nil {
				return
			}
			rivalUserIDs = user.GetRivalUserIDs().Strings()
		} else {
			return
		}
	} else {
		rivalUserIDs = contestant.RivalUserIDs.Strings()
	}

	for {
		var rivalUserID string
		if rivalUserID, err = TournamentDAL.FindStranger(c, tournamentID, userID, rivalUserIDs); err != nil {
			err = errors.WithMessage(err, "failed to find stranger")
			return
		}
		log.Debugf(c, "strangerFacade.PlaceBidAgainstStranger() => rivalUserID: %v", rivalUserID)

		switch rivalUserID {
		case userID:
			err = errors.WithMessage(err, "FindStranger returned rivalUserID equal to current userID")
			return
		case "": // no strangers with existing open bids found
			err = onStranger(contestant)
			return
		default: // Link 2 strangers

			if err = onRivalFound(rivalUserID); errors.Cause(err) == ErrRivalUserIsNotBiddingAgainstStranger {
				err = nil
				continue
			}
			return
		}
	}
	return
}

func RegisterStranger(c context.Context,
	now time.Time,
	tournamentID, userID string,
	contestant *Contestant,
	updateUser func(tc context.Context, strangerRivalKey BattleID) (user db.EntityHolder, err error),
) (err error) {
	strangerRivalKey := NewStrangerBattleID(tournamentID)

	return DB.RunInTransaction(c, func(tc context.Context) (err error) {
		var user db.EntityHolder
		if user, err = updateUser(tc, strangerRivalKey); err != nil {
			err = errors.WithMessage(err, "arena.RegisterStranger() failed to update user")
			return
		}

		if contestant.ContestantEntity == nil {
			contestant.ContestantEntity = &ContestantEntity{
				TimeJoined:   now,
				TournamentID: tournamentID,
				UserID:       userID,
				Stranger:     now,
			}
		} else {
			contestant.Stranger = now
		}
		if err = DB.UpdateMulti(tc, []db.EntityHolder{contestant, user}); err != nil {
			return errors.WithMessage(err, "failed to update user & contestant entities")
		}
		return
	}, db.CrossGroupTransaction)

	return
}
