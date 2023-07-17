package dto

type Match struct {
	ID     int64
	Active int32
}

type MatchRate struct {
	ID        uint64
	MatchID   int64
	Rate      int32
	Status    int32
	IsLive    bool
	Hdp       *float64
	HdpHome   *float64
	HdpAway   *float64
	Ou        *float64
	OuOver    *float64
	OuUnder   *float64
	FhHdp     *float64
	FhHdpHome *float64
	FhHdpAway *float64
	FhOu      *float64
	FhOuOver  *float64
	FhOuUnder *float64
	Match     Match
}
