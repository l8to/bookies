package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		v := ValidateCredit(test.credit, test.stake)
		assert.Equal(t, test.expected, v)
	}
}
