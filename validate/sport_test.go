package validate

import (
	"testing"
)

func TestValidateSingleBet(t *testing.T) {
	errOver := SportParlayBet(&SportBetDetail{
		Stake:           1000,
		CreditBalance:   1000,
		ParlayCount:     1,
		IsSingleAllowed: true,
		MaxSingleLimit:  100,
	})
	if errOver != ErrOverSingleStakeLimit {
		t.Errorf("should return ErrOverSingleStakeLimit, but got %v", errOver)
	}

	errAllowed := validateSingleBet(&SportBetDetail{
		ParlayCount:     1,
		IsSingleAllowed: false,
	})
	if errAllowed != ErrSingleNotAllowed {
		t.Errorf("should return ErrSingleNotAllowed, but got %v", errAllowed)
	}
}

func TestStakeOverCreditBalance(t *testing.T) {
	err := SportParlayBet(&SportBetDetail{
		CreditBalance: 10,
		Stake:         11,
	})
	if err != ErrInsufficientCredit {
		t.Errorf("should return ErrInsufficientCredit, but got %v", err)
	}
}

func TestOverSingleBetLimit(t *testing.T) {
	err := SportParlayBet(&SportBetDetail{
		CreditBalance:   1000,
		Stake:           1000,
		ParlayCount:     1,
		IsSingleAllowed: true,
		MaxSingleLimit:  100,
	})
	if err != ErrOverSingleStakeLimit {
		t.Errorf("should return ErrOverSingleStakeLimit, but got %v", err)
	}
}
