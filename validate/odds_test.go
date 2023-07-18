package validate

import (
	"testing"

	"github.com/dollarsignteam/go-utils"

	"github.com/l8to/bookies/dto"
)

func TestValidateOdds(t *testing.T) {
	matchRate := dto.MatchRate{
		Hdp:       utils.PointerOf(2.0),
		HdpAway:   utils.PointerOf(3.14),
		FhHdpHome: utils.PointerOf(1.5),
	}

	testCases := []struct {
		oddsType string
		odds     float64
		expected bool
	}{
		{"Hdp", 2.0, true},
		{"HdpAway", 3.14, true},
		{"FhHdpHome", 1.8, false},
		{"FhHdp", 2.5, false},
		{"FhOu", 2.0, false},
		{"InvalidType", 2.0, false},
	}

	for _, tc := range testCases {
		actual := ValidateOdds(tc.oddsType, tc.odds, matchRate)
		if actual != tc.expected {
			t.Errorf("Expected %v for oddsType %s, odds %f, and matchRate %+v, got %v",
				tc.expected, tc.oddsType, tc.odds, matchRate, actual)
		}
	}
}
