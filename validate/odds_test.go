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
		newOdds  *float64
	}{
		{"Hdp", 2.0, true, utils.PointerOf(2.0)},
		{"HdpAway", 3.14, true, utils.PointerOf(3.14)},
		{"FhHdpHome", 1.8, false, utils.PointerOf(1.5)},
		{"FhHdp", 2.5, false, utils.PointerOf(2.0)},
		{"FhOu", 2.0, false, utils.PointerOf(2.0)},
		{"Hdp", 0.3, false, nil},
		{"InvalidType", 2.0, false, nil},
	}

	for _, tc := range testCases {
		actual, value := ValidateOdds(tc.oddsType, tc.odds, matchRate)
		if actual != tc.expected && value != tc.newOdds {
			t.Errorf("Expected %v for oddsType %s, odds %f, and matchRate %+v, got %v",
				tc.expected, tc.oddsType, tc.odds, matchRate, actual)
		}
	}
}
