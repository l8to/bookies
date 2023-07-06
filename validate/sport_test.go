package validate

import (
	"testing"

	"github.com/l8to/bookies/dto"
)

func TestValidateSingleBet(t *testing.T) {
	errOver := SportParlayBet(dto.SportBetDetail{
		Stake:           1000,
		CreditBalance:   1000,
		ParlayCount:     1,
		IsSingleAllowed: true,
		MaxSingleLimit:  100,
	})
	if errOver != ErrOverSingleStakeLimit {
		t.Errorf("should return ErrOverSingleStakeLimit, but got %v", errOver)
	}

	errAllowed := validateSingleBet(dto.SportBetDetail{
		ParlayCount:     1,
		IsSingleAllowed: false,
	})
	if errAllowed != ErrSingleNotAllowed {
		t.Errorf("should return ErrSingleNotAllowed, but got %v", errAllowed)
	}
}

func TestStakeOverCreditBalance(t *testing.T) {
	err := SportParlayBet(dto.SportBetDetail{
		CreditBalance: 10,
		Stake:         11,
	})
	if err != ErrInsufficientCredit {
		t.Errorf("should return ErrInsufficientCredit, but got %v", err)
	}
}

func TestOverSingleBetLimit(t *testing.T) {
	err := SportParlayBet(dto.SportBetDetail{
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
