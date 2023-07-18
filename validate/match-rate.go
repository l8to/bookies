package validate

import (
	"time"

	"github.com/l8to/bookies/dto"
)

func ValidateMatchKOTime(koTime time.Time, isLive bool) bool {
	unixTimeNow := time.Now().Unix()
	if isLive && koTime.Unix() <= unixTimeNow {
		return false
	}
	return true
}

func ValidateMatchKOTimeHT(koTime time.Time, isLive bool) bool {
	unixTimeNow := time.Now().Unix()
	isBeforeHT := koTime.Add(time.Minute*45).Unix() > int64(unixTimeNow)
	isAfterHT := koTime.Add(time.Minute*60).Unix() < int64(unixTimeNow)
	if isLive && (isBeforeHT || isAfterHT) {
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
