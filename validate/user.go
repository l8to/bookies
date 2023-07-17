package validate

import (
	"github.com/l8to/bookies/dto"
)

func ValidateUserPermission(user dto.User) bool {

	if user.IsSuspended || user.IsLocked || user.IsClosed {
		return false
	}
	return true
}

func ValidateUserCreditBalance(user dto.User, stake float64) bool {
	return user.UserProfile.Balance >= stake
}

func ValidateUserStakeMax(user dto.User, stake float64) bool {
	return user.UserProfile.StakeMax >= stake
}

func ValidateUserStakeMin(user dto.User, stake float64) bool {
	return user.UserProfile.StakeMin <= stake
}
