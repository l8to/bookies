package dto

type MatchRate struct {
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
	Odd       *float64
	Even      *float64
}
