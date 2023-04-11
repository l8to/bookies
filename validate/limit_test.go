package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
