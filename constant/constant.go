package constant

import "time"

const (
	NewPartnerValidDateTime = 7 * 24 * time.Hour
	TimeLayoutTime          = "15:04"
	TimeLayoutDate          = "2006-01-02"
	TimeLayoutDateTime      = "2006-01-02 15:04:05"
	TimeLayoutMonth         = "2006-01"
)

var HtOddsType = []string{
	"ht_hdp_home",
	"ht_hdp_away",
	"ht_ou_over",
	"ht_ou_under",
}

const (
	BetTypeFhHdp = "FH HDP"
	BetTypeFhOu  = "FH OU"
	BetTypeFh1x2 = "FH 1X2"
	BetType1x2   = "1X2"
	BetTypeOe    = "OE"
)
