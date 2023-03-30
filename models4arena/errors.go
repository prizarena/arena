package models4arena

import "github.com/pkg/errors"

var (
	ErrRivalUserIsNotBiddingAgainstStranger = errors.New("rival user is not a stranger")
)
