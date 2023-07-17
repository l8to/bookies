package dto

type (
	User struct {
		ID          int32
		OddsType    int32
		IsSingle    bool
		IsSuspended bool
		IsPrintslip bool
		IsLocked    bool
		IsClosed    bool
		IsFh        bool
		IsFt        bool
		IsOe        bool
		UserProfile UserProfile
	}
	UserProfile struct {
		Balance   float64
		StakeMax  float64
		StakeMin  float64
		MaxSingle float64
		ParlayMax int32
		ParlayMin int32
		MaxPayout float64
	}
)
