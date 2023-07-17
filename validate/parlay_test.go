package validate

import (
	"testing"

	"github.com/l8to/bookies/dto"
)

func TestValidateParlayPermission(t *testing.T) {
	testCases := []struct {
		name           string
		parlay         dto.TicketParlay
		user           dto.User
		expectedResult bool
	}{
		{
			name: "FH HDP and user.IsFh is 0",
			parlay: dto.TicketParlay{
				BetType: ("FH HDP"),
			},
			user: dto.User{
				IsFh: false,
			},
			expectedResult: false,
		},
		{
			name: "FH OU and user.IsFh is 0",
			parlay: dto.TicketParlay{
				BetType: ("FH OU"),
			},
			user: dto.User{
				IsFh: false,
			},
			expectedResult: false,
		},
		{
			name: "FH 1X2 and user.IsFt is 0",
			parlay: dto.TicketParlay{
				BetType: ("FH 1X2"),
			},
			user: dto.User{
				IsFt: false,
			},
			expectedResult: false,
		},
		{
			name: "1X2 and user.IsFt is 0",
			parlay: dto.TicketParlay{
				BetType: ("1X2"),
			},
			user: dto.User{
				IsFt: false,
			},
			expectedResult: false,
		},
		{
			name: "OE and user.IsOe is 0",
			parlay: dto.TicketParlay{
				BetType: ("OE"),
			},
			user: dto.User{
				IsOe: false,
			},
			expectedResult: false,
		},
		{
			name: "Valid permission",
			parlay: dto.TicketParlay{
				BetType: ("FH 1X2"),
			},
			user: dto.User{
				IsFt: true,
			},
			expectedResult: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ValidateParlayPermission(tc.user, tc.parlay)
			if result != tc.expectedResult {
				t.Errorf("Expected %v, but got %v", tc.expectedResult, result)
			}
		})
	}
}

func TestValidateParlayMax(t *testing.T) {
	testCases := []struct {
		user     dto.User
		parlay   []dto.TicketParlay
		expected bool
	}{
		{dto.User{UserProfile: dto.UserProfile{ParlayMax: (int32(3))}}, []dto.TicketParlay{{}, {}}, true},
		{dto.User{UserProfile: dto.UserProfile{ParlayMax: (int32(2))}}, []dto.TicketParlay{{}, {}, {}}, false},
	}
	for i, testCase := range testCases {
		result := ValidateParlayMax(testCase.user, testCase.parlay)
		if result != testCase.expected {
			t.Errorf("Test case %d failed: Expected %v, got %v", i+1, testCase.expected, result)
		}
	}
}

func TestValidateParlayMin(t *testing.T) {
	testCases := []struct {
		user     dto.User
		parlay   []dto.TicketParlay
		expected bool
	}{
		{dto.User{UserProfile: dto.UserProfile{ParlayMin: (int32(1))}}, []dto.TicketParlay{{}, {}}, true},
		{dto.User{UserProfile: dto.UserProfile{ParlayMin: (int32(3))}}, []dto.TicketParlay{{}, {}}, false},
	}
	for i, testCase := range testCases {
		result := ValidateParlayMin(testCase.user, testCase.parlay)
		if result != testCase.expected {
			t.Errorf("Test case %d failed: Expected %v, got %v", i+1, testCase.expected, result)
		}
	}
}

func TestValidateUserSingleType(t *testing.T) {
	testCases := []struct {
		user     dto.User
		parlay   []dto.TicketParlay
		expected bool
	}{
		{dto.User{IsSingle: false}, []dto.TicketParlay{{}}, false},
		{dto.User{IsSingle: true}, []dto.TicketParlay{{}}, true},
		{dto.User{IsSingle: false}, []dto.TicketParlay{{}, {}}, false},
	}
	for i, testCase := range testCases {
		result := ValidateUserSingleType(testCase.user, testCase.parlay)
		if result != testCase.expected {
			t.Errorf("Test case %d failed: Expected %v, got %v", i+1, testCase.expected, result)
		}
	}
}

func TestValidateMaxSingle(t *testing.T) {
	testCases := []struct {
		user     dto.User
		stake    float64
		parlay   []dto.TicketParlay
		expected bool
	}{
		{dto.User{UserProfile: dto.UserProfile{MaxSingle: (float64(100))}}, 50.0, []dto.TicketParlay{{}}, true},
		{dto.User{UserProfile: dto.UserProfile{MaxSingle: (float64(100))}}, 150.0, []dto.TicketParlay{{}}, false},
	}
	for i, testCase := range testCases {
		result := ValidateMaxSingle(testCase.user, testCase.stake, testCase.parlay)
		if result != testCase.expected {
			t.Errorf("Test case %d failed: Expected %v, got %v", i+1, testCase.expected, result)
		}
	}
}

func TestValidateMaxPayout(t *testing.T) {
	testCases := []struct {
		user         dto.User
		summaryStake float64
		stake        float64
		parlay       []dto.TicketParlay
		expected     bool
	}{
		{dto.User{UserProfile: dto.UserProfile{MaxPayout: float64(100)}}, 50.0, 20.0, []dto.TicketParlay{{}}, true},
		{dto.User{UserProfile: dto.UserProfile{MaxPayout: float64(100)}}, 150.0, 20.0, []dto.TicketParlay{{}}, false},
		// Add more test cases here
	}

	for i, testCase := range testCases {
		result := ValidateMaxPayout(testCase.user, testCase.summaryStake, testCase.stake, testCase.parlay)
		if result != testCase.expected {
			t.Errorf("Test case %d failed: Expected %v, got %v", i+1, testCase.expected, result)
		}
	}
}
