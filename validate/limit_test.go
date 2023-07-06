package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePayoutLimit(t *testing.T) {
	tests := []struct {
		totalWin uint
		limit    uint
		expected bool
	}{
		{totalWin: 1, limit: 1, expected: true},
		{totalWin: 2, limit: 1, expected: false},
		{totalWin: 1, limit: 2, expected: true},
	}
	for _, test := range tests {
		v := payoutLimit(test.totalWin, test.limit)
		assert.Equal(t, test.expected, v)
	}
}

func TestValidateStakeLimit(t *testing.T) {
	tests := []struct {
		stake    uint
		limit    uint
		expected bool
	}{
		{stake: 1, limit: 1, expected: true},
		{stake: 2, limit: 1, expected: false},
		{stake: 1, limit: 2, expected: true},
	}
	for _, test := range tests {
		v := stakeMaxLimit(test.stake, test.limit)
		assert.Equal(t, test.expected, v)
	}
}

func TestValidateLimit(t *testing.T) {
	tests := []struct {
		stake         uint
		maxStakeLimit uint
		minStakeLimit uint
		totalWin      uint
		maxPayout     uint
		expectedError error
	}{
		{stake: 100, maxStakeLimit: 1000, minStakeLimit: 10, totalWin: 500, maxPayout: 10000, expectedError: nil},
		{stake: 5, maxStakeLimit: 100, minStakeLimit: 10, totalWin: 500, maxPayout: 10000, expectedError: ErrLessThanMinStakeLimit},
		{stake: 2000, maxStakeLimit: 1000, minStakeLimit: 10, totalWin: 500, maxPayout: 10000, expectedError: ErrOverMaxStakeLimit},
		{stake: 100, maxStakeLimit: 1000, minStakeLimit: 10, totalWin: 20000, maxPayout: 10000, expectedError: ErrOverMaxPayout},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v,%v,%v,%v,%v", test.stake, test.maxStakeLimit, test.minStakeLimit, test.totalWin, test.maxPayout), func(t *testing.T) {
			err := ValidateLimit(test.stake, test.maxStakeLimit, test.minStakeLimit, test.totalWin, test.maxPayout)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
