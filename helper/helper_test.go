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

func TestCalculateWlCommission(t *testing.T) {
	testCases := []struct {
		name           string
		input          CalculateWlCommissionInput
		expectedResult []float64
	}{
		{
			name: "test 1",
			input: CalculateWlCommissionInput{
				Stake: 20,

				CommMember: 6,
				CommAgent:  13,
				CommMaster: 13,
				CommSenior: 13,
				CommSuper:  13,
				CommWeb:    13,

				ShAgent:  0.3,
				ShMaster: 0,
				ShSenior: 0.6,
				ShSuper:  0,
				ShWeb:    0.1,
			},
			expectedResult: []float64{1.2, 0.62, 0, -1.56, 0, -0.26},
		},
		{
			name: "test 2",
			input: CalculateWlCommissionInput{
				Stake: 50,

				CommMember: 0,
				CommAgent:  12,
				CommMaster: 12,
				CommSenior: 20,
				CommSuper:  20,
				CommWeb:    20,

				ShAgent:  0.75,
				ShMaster: 0,
				ShSenior: 0.13,
				ShSuper:  0.02,
				ShWeb:    0.1,
			},
			expectedResult: []float64{0, 1.5, 0, -0.3, -0.2, -1},
		},
	}

	for _, testCase := range testCases {
		wlCommMember, wlCommAgent, WlCommMaster, WlCommSenior, WlCommSuper, WlCommWeb := CalculateWlCommission(testCase.input)
		commissions := []float64{wlCommMember, wlCommAgent, WlCommMaster, WlCommSenior, WlCommSuper, WlCommWeb}
		for i, commission := range commissions {
			if commission != testCase.expectedResult[i] {
				t.Errorf("%s commission%d: expected %.2f, got %.2f", testCase.name, i+1, testCase.expectedResult[i], commission)
			}
		}
	}
}
