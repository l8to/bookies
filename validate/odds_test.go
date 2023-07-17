package validate

import (
	"testing"

	"github.com/dollarsignteam/go-utils"
	"github.com/l8to/bookies/dto"
)

func TestValidateOdds(t *testing.T) {
	parlay := dto.TicketParlay{
		OddsType: "HdpHome",
		Odds:     2.5,
	}
	matchRate := dto.MatchRate{
		HdpHome: utils.PointerOf(2.5),
	}
	result := ValidateOdds(parlay, matchRate)
	if !result {
		t.Errorf("Expected true but got false")
	}
}

func TestValidateOdds_InvalidOddsType(t *testing.T) {
	parlay := dto.TicketParlay{
		OddsType: "invalidOddsType",
		Odds:     2.5,
	}
	matchRate := dto.MatchRate{
		HdpHome: utils.PointerOf(2.5),
	}
	result := ValidateOdds(parlay, matchRate)
	if result {
		t.Errorf("Expected false but got true")
	}
}

func TestValidateOdds_MismatchedOdds(t *testing.T) {
	parlay := dto.TicketParlay{
		OddsType: "HdpHome",
		Odds:     2.5,
	}
	matchRate := dto.MatchRate{
		HdpHome: utils.PointerOf(3.0),
	}
	result := ValidateOdds(parlay, matchRate)
	if result {
		t.Errorf("Expected false but got true")
	}
}
