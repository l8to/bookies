package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/l8to/bookies/dto"
)

type UserTestCase struct {
	User        dto.User
	Stake       float64
	Expectation bool
}

func TestValidateUserPermission(t *testing.T) {
	testCases := []UserTestCase{
		{
			User:        dto.User{IsSuspended: true},
			Expectation: false,
		},
		{
			User:        dto.User{IsLocked: true},
			Expectation: false,
		},
		{
			User:        dto.User{IsClosed: true},
			Expectation: false,
		},
		{
			User:        dto.User{},
			Expectation: true,
		},
	}
	for _, tc := range testCases {
		result := ValidateUserPermission(tc.User)
		assert.Equal(t, tc.Expectation, result)
	}
}

func TestValidateUserCreditBalance(t *testing.T) {
	testCases := []UserTestCase{
		{
			User:        dto.User{UserProfile: dto.UserProfile{Balance: 100}},
			Stake:       50,
			Expectation: true,
		},
		{
			User:        dto.User{UserProfile: dto.UserProfile{Balance: 100}},
			Stake:       150,
			Expectation: false,
		},
	}
	for _, tc := range testCases {
		result := ValidateUserCreditBalance(tc.User, tc.Stake)
		assert.Equal(t, tc.Expectation, result)
	}
}

func TestValidateUserStakeMax(t *testing.T) {
	testCases := []UserTestCase{
		{
			User:        dto.User{UserProfile: dto.UserProfile{StakeMax: 100}},
			Stake:       50,
			Expectation: true,
		},
		{
			User:        dto.User{UserProfile: dto.UserProfile{StakeMax: 100}},
			Stake:       150,
			Expectation: false,
		},
	}
	for _, tc := range testCases {
		result := ValidateUserStakeMax(tc.User, tc.Stake)
		assert.Equal(t, tc.Expectation, result)
	}
}

func TestValidateUserStakeMin(t *testing.T) {
	testCases := []UserTestCase{
		{
			User:        dto.User{UserProfile: dto.UserProfile{StakeMin: 50}},
			Stake:       100,
			Expectation: true,
		},
		{
			User:        dto.User{UserProfile: dto.UserProfile{StakeMin: 50}},
			Stake:       25,
			Expectation: false,
		},
	}
	for _, tc := range testCases {
		result := ValidateUserStakeMin(tc.User, tc.Stake)
		assert.Equal(t, tc.Expectation, result)
	}
}
