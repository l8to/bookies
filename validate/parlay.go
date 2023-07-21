package validate

import (
	"github.com/l8to/bookies/constant"
	"github.com/l8to/bookies/dto"
)

func ValidateParlayPermission(user dto.User, betType string) bool {
	if (betType == constant.BetTypeFhHdp || betType == constant.BetTypeFhOu) && !user.IsFh {
		return false
	} else if (betType == constant.BetTypeFh1x2 || betType == constant.BetType1x2) && !user.IsFt {
		return false
	} else if (betType == constant.BetTypeOe) && !user.IsOe {
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

func ValidateMaxPerMatchStake(user dto.User, summaryStake float64, stake float64, parlayCount int32) bool {
	return parlayCount == 1 && user.UserProfile.MaxSingle >= (summaryStake+stake)
}

func ValidateMaxPayout(user dto.User, oddsStake float64) bool {
	return user.UserProfile.MaxPayout >= oddsStake
}
