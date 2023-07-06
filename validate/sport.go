package validate

import "github.com/l8to/bookies/dto"

func SportParlayBet(c dto.SportBetDetail) error {
	var err error

	if !validateCredit(c.CreditBalance, c.Stake) {
		return ErrInsufficientCredit
	}

	if c.ParlayCount == 1 {
		err = validateSingleBet(c)
	} else {
		err = validateMixParlayBet(c)
	}

	if err != nil {
		return err
	}

	win := totalWin(c.TotalBetOdds, c.Stake)
	err = ValidateLimit(c.Stake, c.MaxStakeLimit, c.MinStakeLimit, win, c.MaxPayout)
	if err != nil {
		return err
	}
	return nil
}

func validateSingleBet(c dto.SportBetDetail) error {
	if c.ParlayCount == 1 {
		if !c.IsSingleAllowed {
			return ErrSingleNotAllowed
		}
		if c.MaxSingleLimit < c.Stake {
			return ErrOverSingleStakeLimit
		}
		// find out match quota limit
		if c.IsOverQuotaMatchLimit {
			return ErrOverQuotaProductLimit
		}
	}
	return nil
}

func validateMixParlayBet(c dto.SportBetDetail) error {
	if c.ParlayCount > 1 {
		if c.ParlayCount > c.MinParlayCount {
			return ErrParlayLessThanMinParlay
		}
		if c.MaxSingleLimit < c.Stake {
			return ErrParlayOverMaxParlay
		}
	}
	return nil
}

func totalWin(odds, stake uint) uint {
	return odds * stake
}
