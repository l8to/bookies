package validate

type SportBetDetail struct {
	// Balance before make booking
	CreditBalance uint
	// Booking total cost
	Stake                 uint
	ParlayCount           uint8
	IsSingleAllowed       bool
	MaxSingleLimit        uint
	TotalBetOdds          uint
	MaxPayout             uint
	MaxStakeLimit         uint
	MinStakeLimit         uint
	isOverQuotaMatchLimit bool
	MinParlayCount        uint8
	MaxParlayCount        uint8
}

func SportParlayBet(c *SportBetDetail) error {
	var err error

	if validateCredit(c.CreditBalance, c.Stake) == false {
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
	err = validateLimit(c.Stake, c.MaxStakeLimit, c.MinStakeLimit, win, c.MaxPayout)
	if err != nil {
		return err
	}
	return nil
}

func validateSingleBet(c *SportBetDetail) error {
	if c.ParlayCount == 1 {
		if c.IsSingleAllowed == false {
			return ErrSingleNotAllowed
		}
		if c.MaxSingleLimit < c.Stake {
			return ErrOverSingleStakeLimit
		}
		// find out match quota limit
		if c.isOverQuotaMatchLimit == true {
			return ErrOverQuotaProductLimit
		}
	}
	return nil
}

func validateMixParlayBet(c *SportBetDetail) error {
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
