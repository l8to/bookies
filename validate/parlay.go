package validate

import (
	"github.com/l8to/bookies/dto"
)

func ValidateParlayPermission(user dto.User, parlay dto.TicketParlay) bool {
	betType := parlay.BetType
	if (betType == "FH HDP" || betType == "FH OU") && !user.IsFh {
		return false
	} else if (betType == "FH 1X2" || betType == "1X2") && !user.IsFt {
		return false
	} else if (betType == "OE") && !user.IsOe {
		return false
	}
	return true
}

func ValidateParlayMax(user dto.User, parlay []dto.TicketParlay) bool {
	parlayCount := len(parlay)
	return user.UserProfile.ParlayMax >= int32(parlayCount)
}

func ValidateParlayMin(user dto.User, parlay []dto.TicketParlay) bool {
	parlayCount := len(parlay)
	return user.UserProfile.ParlayMin <= int32(parlayCount)
}

func ValidateUserSingleType(user dto.User, parlay []dto.TicketParlay) bool {
	parlayCount := len(parlay)
	return parlayCount == 1 && user.IsSingle
}

func ValidateMaxSingle(user dto.User, stake float64, parlay []dto.TicketParlay) bool {
	parlayCount := len(parlay)
	return parlayCount == 1 && user.UserProfile.MaxSingle >= stake
}

func ValidateMaxPayout(user dto.User, summaryStake float64, stake float64, parlay []dto.TicketParlay) bool {
	parlayCount := len(parlay)
	return parlayCount == 1 && user.UserProfile.MaxPayout >= (summaryStake+stake)
}
