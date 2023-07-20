package validate

import (
	"github.com/l8to/bookies/dto"
)

func ValidateParlayPermission(user dto.User, betType string) bool {
	if (betType == "FH HDP" || betType == "FH OU") && !user.IsFh {
		return false
	} else if (betType == "FH 1X2" || betType == "1X2") && !user.IsFt {
		return false
	} else if (betType == "OE") && !user.IsOe {
		return false
	}
	return true
}

func ValidateParlayMax(user dto.User, parlayCount int32) bool {
	return user.UserProfile.ParlayMax >= parlayCount
}

func ValidateParlayMin(user dto.User, parlayCount int32) bool {
	return user.UserProfile.ParlayMin <= parlayCount
}

func ValidateUserSingleType(user dto.User, parlayCount int32) bool {
	return parlayCount == 1 && user.IsSingle
}

func ValidateMaxPerMatchStake(user dto.User, summaryStake float64, stake float64, parlay []dto.TicketParlay) bool {
	parlayCount := len(parlay)
	return parlayCount == 1 && user.UserProfile.MaxSingle >= (summaryStake+stake)
}

func ValidateMaxPayout(user dto.User, oddsStake float64) bool {
	return user.UserProfile.MaxPayout >= oddsStake
}
