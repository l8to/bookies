package validate

import (
	"github.com/dollarsignteam/go-utils"
	"github.com/iancoleman/strcase"

	"github.com/l8to/bookies/dto"
	"github.com/l8to/bookies/helper"
)

func ValidateOdds(parlay dto.TicketParlay, matchRate dto.MatchRate) bool {
	oddsRate := helper.GetStructValueByKeyName(matchRate, strcase.ToCamel(parlay.OddsType))
	floatValue, ok := oddsRate.(*float64)
	if !ok {
		return false
	}
	if parlay.Odds != utils.ValueOf(floatValue) {
		return false
	}
	return true
}
