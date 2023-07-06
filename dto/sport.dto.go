package dto

type SportBetDetail struct {
	CreditBalance         uint
	Stake                 uint
	ParlayCount           uint8
	IsSingleAllowed       bool
	MaxSingleLimit        uint
	TotalBetOdds          uint
	MaxPayout             uint
	MaxStakeLimit         uint
	MinStakeLimit         uint
	IsOverQuotaMatchLimit bool
	MinParlayCount        uint8
	MaxParlayCount        uint8
}
