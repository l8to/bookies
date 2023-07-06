package validate

// ValidateLimit checks if the stake, total win, and payout limits are within acceptable ranges.
// It returns an error if any of the limits exceed their respective thresholds.
func ValidateLimit(stake, maxStakeLimit, minStakeLimit, totalWin, maxPayout uint) error {
	if !payoutLimit(totalWin, maxPayout) {
		return ErrOverMaxPayout
	}
	if !stakeMaxLimit(stake, maxStakeLimit) {
		return ErrOverMaxStakeLimit
	}
	if !stakeMinLimit(stake, minStakeLimit) {
		return ErrLessThanMinStakeLimit
	}
	return nil
}

func payoutLimit(totalWin, maxPayout uint) bool {
	return totalWin <= maxPayout
}

func stakeMaxLimit(stake, maxStakeLimit uint) bool {
	return stake <= maxStakeLimit
}

func stakeMinLimit(stake, minStakeLimit uint) bool {
	return stake >= minStakeLimit
}
