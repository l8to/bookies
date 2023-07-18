package validate

import (
	"github.com/dollarsignteam/go-utils"
	"github.com/iancoleman/strcase"

	"github.com/l8to/bookies/dto"
	"github.com/l8to/bookies/helper"
)

func ValidateOdds(oddsType string, odds float64, matchRate dto.MatchRate) bool {
	oddsRate := helper.GetStructValueByKeyName(matchRate, strcase.ToCamel(oddsType))
	floatValue, ok := oddsRate.(*float64)
	if !ok {
		return false
	}
	if odds != utils.ValueOf(floatValue) {
		return false
	}
	return true
}
