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
		newOdds  float64
	}{
		{"Hdp", 2.0, true, 2.0},
		{"HdpAway", 3.14, true, 3.14},
		{"FhHdpHome", 1.8, false, 1.5},
		{"FhHdp", 2.5, false, 2.0},
		{"FhOu", 2.0, false, 2.0},
		{"InvalidType", 2.0, false, 0},
	}

	for _, tc := range testCases {
		actual, value := ValidateOdds(tc.oddsType, tc.odds, matchRate)
		if actual != tc.expected && value != tc.newOdds {
			t.Errorf("Expected %v for oddsType %s, odds %f, and matchRate %+v, got %v",
				tc.expected, tc.oddsType, tc.odds, matchRate, actual)
		}
	}
}
