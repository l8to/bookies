package validate

func validateLimit(stake, maxStakeLimit, minStakeLimit, totalWin, maxPayout uint) error {
	if payoutLimit(totalWin, maxPayout) == false {
		return ErrOverMaxPayout
	}

	if stakeMaxLimit(stake, maxStakeLimit) == false {
		return ErrOverMaxStakeLimit
	}

	if stakeMinLimit(stake, minStakeLimit) == false {
		return ErrLessThanMinStakeLimit
	}
	return nil
}

func payoutLimit(totalWin, maxPayout uint) bool {
	if totalWin > maxPayout {
		return false
	}
	return true
}

func stakeMaxLimit(stake, maxStakeLimit uint) bool {
	if stake > maxStakeLimit {
		return false
	}
	return true
}
func stakeMinLimit(stake, minStakeLimit uint) bool {
	if stake < minStakeLimit {
		return false
	}
	return true
}
