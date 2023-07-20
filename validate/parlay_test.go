package validate

import (
	"testing"

	"github.com/l8to/bookies/dto"
)

func TestValidateParlayPermission(t *testing.T) {
	testCases := []struct {
		name           string
		betType        string
		user           dto.User
		expectedResult bool
	}{
		{
			name:    "FH HDP and user.IsFh is 0",
			betType: ("FH HDP"),
			user: dto.User{
				IsFh: false,
			},
			expectedResult: false,
		},
		{
			name:    "FH OU and user.IsFh is 0",
			betType: ("FH OU"),
			user: dto.User{
				IsFh: false,
			},
			expectedResult: false,
		},
		{
			name:    "FH 1X2 and user.IsFt is 0",
			betType: ("FH 1X2"),
			user: dto.User{
				IsFt: false,
			},
			expectedResult: false,
		},
		{
			name:    "1X2 and user.IsFt is 0",
			betType: ("1X2"),
			user: dto.User{
				IsFt: false,
			},
			expectedResult: false,
		},
		{
			name:    "OE and user.IsOe is 0",
			betType: ("OE"),
			user: dto.User{
				IsOe: false,
			},
			expectedResult: false,
		},
		{
			name:    "Valid permission",
			betType: ("FH 1X2"),
			user: dto.User{
				IsFt: true,
			},
			expectedResult: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ValidateParlayPermission(tc.user, tc.betType)
			if result != tc.expectedResult {
				t.Errorf("Expected %v, but got %v", tc.expectedResult, result)
			}
		})
	}
}
func TestValidateParlayMax(t *testing.T) {
	user := dto.User{
		UserProfile: dto.UserProfile{ParlayMax: 5},
	}
	parlayCount := int32(3)

	if !ValidateParlayMax(user, parlayCount) {
		t.Error("Expected true, got false")
	}
}

func TestValidateParlayMin(t *testing.T) {
	user := dto.User{
		UserProfile: dto.UserProfile{ParlayMin: 2},
	}
	parlayCount := int32(3)

	if !ValidateParlayMin(user, parlayCount) {
		t.Error("Expected true, got false")
	}
}

func TestValidateUserSingleType(t *testing.T) {
	user := dto.User{
		IsSingle: true,
	}
	parlayCount := int32(1)

	if !ValidateUserSingleType(user, parlayCount) {
		t.Error("Expected true, got false")
	}
}

func TestValidateMaxSingle(t *testing.T) {
	testCases := []struct {
		name           string
		user           dto.User
		parlay         []dto.TicketParlay
		summaryStake   float64
		stake          float64
		expectedResult bool
	}{
		{
			name: "Valid Max Single",
			user: dto.User{
				UserProfile: dto.UserProfile{
					MaxSingle: 100.0,
				},
			},
			parlay: []dto.TicketParlay{
				dto.TicketParlay{},
			},
			summaryStake:   50.0,
			stake:          30.0,
			expectedResult: true,
		},
		{
			name: "Invalid Max Single",
			user: dto.User{
				UserProfile: dto.UserProfile{
					MaxSingle: 50.0,
				},
			},
			parlay: []dto.TicketParlay{
				dto.TicketParlay{},
			},
			summaryStake:   70.0,
			stake:          20.0,
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ValidateMaxPerMatchStake(tc.user, tc.summaryStake, tc.stake, tc.parlay)
			if result != tc.expectedResult {
				t.Errorf("Expected %t, but got %t", tc.expectedResult, result)
			}
		})
	}
}

func TestValidateMaxPayout(t *testing.T) {
	testCases := []struct {
		name           string
		user           dto.User
		oddsStake      float64
		expectedResult bool
	}{
		{
			name: "Valid Max Payout",
			user: dto.User{
				UserProfile: dto.UserProfile{
					MaxPayout: 1000.0,
				},
			},
			oddsStake:      500.0,
			expectedResult: true,
		},
		{
			name: "Invalid Max Payout",
			user: dto.User{
				UserProfile: dto.UserProfile{
					MaxPayout: 200.0,
				},
			},
			oddsStake:      300.0,
			expectedResult: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ValidateMaxPayout(tc.user, tc.oddsStake)
			if result != tc.expectedResult {
				t.Errorf("Expected %t, but got %t", tc.expectedResult, result)
			}
		})
	}
}
