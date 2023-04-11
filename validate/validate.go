package validate

import (
	"errors"
)

// common error
var (
	ErrInsufficientCredit    = errors.New("insufficient Credit")
	ErrOverMaxPayout         = errors.New("max payout limit reached")
	ErrOverMaxStakeLimit     = errors.New("max stake limit reached")
	ErrLessThanMinStakeLimit = errors.New("stake less than minimum limit")
	ErrOverQuotaProductLimit = errors.New("max match stake reached")
)

// sport error
var (
	ErrSingleNotAllowed        = errors.New("single is not allowed")
	ErrOverSingleStakeLimit    = errors.New("single stake limit reached")
	ErrParlayLessThanMinParlay = errors.New("bet parlay less than minimum parlay allowed")
	ErrParlayOverMaxParlay     = errors.New("bet parlay over than max parlay allowed")
)
