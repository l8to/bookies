package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateCredit(t *testing.T) {
	tests := []struct {
		credit   uint
		stake    uint
		expected bool
	}{
		{credit: 1, stake: 1, expected: true},
		{credit: 1, stake: 2, expected: false},
		{credit: 2, stake: 1, expected: true},
		{credit: 2, stake: 0, expected: false},
		{credit: 0, stake: 1, expected: false},
		{credit: 0, stake: 0, expected: false},
		{credit: 1000, stake: 1000, expected: true},
	}
	for _, test := range tests {
		v := validateCredit(test.credit, test.stake)
		assert.Equal(t, test.expected, v)
	}
}
