package validate

import (
	"github.com/dollarsignteam/go-utils"
	"github.com/iancoleman/strcase"

	"github.com/l8to/bookies/dto"
	"github.com/l8to/bookies/helper"
)

func ValidateOdds(oddsType string, odds float64, matchRate dto.MatchRate) (bool, *float64) {
	oddsRate := helper.GetStructValueByKeyName(matchRate, strcase.ToCamel(oddsType))
	floatValue, ok := oddsRate.(*float64)
	if !ok || floatValue == nil || odds < 1 {
		return false, nil
	}
	if odds != utils.ValueOf(floatValue) {
		return false, floatValue
	}
	return true, floatValue
}
