package validate

import (
	"time"

	"github.com/l8to/bookies/dto"
)

func ValidateMatchTime(koTime time.Time, timeNow time.Time, isLive bool) bool {
	if isLive {
		if valid := ValidateMatchKOTimeHT(koTime, timeNow); !valid {
			return false
		}
	}
	if !isLive {
		if valid := ValidateMatchKOTime(koTime, timeNow); !valid {
			return false
		}
	}
	return true
}

func ValidateMatchKOTime(koTime time.Time, timeNow time.Time) bool {
	return koTime.Unix() > timeNow.Unix()
}

func ValidateMatchKOTimeHT(koTime time.Time, timeNow time.Time) bool {
	isBeforeHT := koTime.Add(time.Minute*45).Unix() > timeNow.Unix()
	isAfterHT := koTime.Add(time.Minute*60).Unix() <= timeNow.Unix()
	if isBeforeHT || isAfterHT {
		return false
	}
	return true
}

func ValidateMatchRateActive(user dto.User, matchRateStatus uint, matchActive uint) bool {
	if matchRateStatus != 1 {
		return false
	}
	if matchActive == 0 || matchActive > 3 {
		return false
	}
	return true
}

func ValidateMatchRateAndUserOddsType(user dto.User, rate int32) bool {
	return user.OddsType == rate
}
