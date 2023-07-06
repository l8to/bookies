package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/l8to/bookies/dto"
)

func TestSportBetDetail_GetValue(t *testing.T) {
	t.Run("valid struct", func(t *testing.T) {
		tests := []struct {
			name     string
			expected interface{}
		}{
			{"CreditBalance", uint(1000)},
			{"Stake", uint(500)},
			{"ParlayCount", uint8(2)},
			{"IsSingleAllowed", false},
			{"MaxSingleLimit", uint(100)},
			{"TotalBetOdds", uint(5)},
			{"MaxPayout", uint(10000)},
			{"MaxStakeLimit", uint(200)},
			{"MinStakeLimit", uint(10)},
			{"IsOverQuotaMatchLimit", false},
			{"MinParlayCount", uint8(1)},
			{"MaxParlayCount", uint8(3)},
			{"ABC", nil},
		}

		sportBetDetail := dto.SportBetDetail{
			CreditBalance:         1000,
			Stake:                 500,
			ParlayCount:           2,
			IsSingleAllowed:       false,
			MaxSingleLimit:        100,
			TotalBetOdds:          5,
			MaxPayout:             10000,
			MaxStakeLimit:         200,
			MinStakeLimit:         10,
			IsOverQuotaMatchLimit: false,
			MinParlayCount:        1,
			MaxParlayCount:        3,
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := GetStructValueByKeyName(sportBetDetail, test.name)
				assert.Equal(t, test.expected, actual)
			})
		}
	})

	t.Run("invalid struct", func(t *testing.T) {
		actual := GetStructValueByKeyName(1, "foo")
		assert.Nil(t, actual)
	})
}

func BenchmarkGetStructValueByKeyName(b *testing.B) {
	sportBet := dto.SportBetDetail{
		TotalBetOdds: 5,
	}
	fieldName := "TotalBetOdds"
	for i := 0; i < b.N; i++ {
		GetStructValueByKeyName(sportBet, fieldName)
	}
}
