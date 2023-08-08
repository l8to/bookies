package validate

import (
	"testing"

	"github.com/dollarsignteam/go-utils"
	"github.com/stretchr/testify/assert"

	"github.com/l8to/bookies/constant"
	"github.com/l8to/bookies/dto"
)

func TestValidateMatchTime(t *testing.T) {
	t.Run("Invalid HT", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-20 18:00:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-20 18:30:00")
		isValid := ValidateMatchTime(koTime, timeNow, true)
		assert.False(t, isValid)
	})

	t.Run("Invalid FT", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-21 18:00:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-21 18:10:00")
		isValid := ValidateMatchTime(koTime, timeNow, false)
		assert.False(t, isValid)
	})

	t.Run("Valid HT", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-21 18:00:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-21 18:50:00")
		isValid := ValidateMatchTime(koTime, timeNow, true)
		assert.True(t, isValid)
	})

	t.Run("Valid FT", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-21 18:10:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-21 18:00:00")
		isValid := ValidateMatchTime(koTime, timeNow, false)
		assert.True(t, isValid)
	})
}

func TestValidateMatchKOTime(t *testing.T) {
	t.Run("koTime < timeNow", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-21 18:00:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-21 18:10:00")
		expected := false

		result := ValidateMatchKOTime(koTime, timeNow)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("koTime = timeNow", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-21 18:00:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-21 18:00:00")
		expected := false

		result := ValidateMatchKOTime(koTime, timeNow)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("koTime > timeNow", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-21 18:00:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-21 17:00:00")
		expected := true

		result := ValidateMatchKOTime(koTime, timeNow)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}

func TestValidateMatchKOTimeHT(t *testing.T) {
	t.Run("timeNow = a day before koTime", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-20 18:00:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-19 17:30:00")
		expected := false

		result := ValidateMatchKOTimeHT(koTime, timeNow)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("timeNow = a day after koTime", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-20 18:00:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-23 09:30:00")
		expected := false

		result := ValidateMatchKOTimeHT(koTime, timeNow)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("timeNow <  koTime", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-20 18:00:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-20 17:30:00")
		expected := false

		result := ValidateMatchKOTimeHT(koTime, timeNow)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("timeNow < 45 minutes after koTime", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-20 18:00:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-20 18:30:00")
		expected := false

		result := ValidateMatchKOTimeHT(koTime, timeNow)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("timeNow > 70 minutes after koTime", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-20 18:00:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-20 19:10:00")
		expected := false

		result := ValidateMatchKOTimeHT(koTime, timeNow)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("timeNow > 45 minutes after koTime & timeNow < 60 minutes after koTime", func(t *testing.T) {
		koTime, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-20 18:00:00")
		timeNow, _ := utils.Time.ParseInBangkokLocation(constant.TimeLayoutDateTime, "2023-07-20 18:50:00")
		expected := true

		result := ValidateMatchKOTimeHT(koTime, timeNow)

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}

func TestValidateMatchRateActive(t *testing.T) {
	testCases := []struct {
		user            dto.User
		matchRateStatus uint
		matchActive     uint
		expected        bool
	}{
		{dto.User{}, 1, 0, false},
		{dto.User{}, 2, 2, false},
		{dto.User{}, 1, 4, false},
		{dto.User{}, 1, 1, true},
		{dto.User{}, 1, 2, true},
		{dto.User{}, 1, 3, true},
	}

	for _, tc := range testCases {
		actual := ValidateMatchRateActive(tc.user, tc.matchRateStatus, tc.matchActive)
		if actual != tc.expected {
			t.Errorf("Expected %v for user %+v, matchRateStatus %d, and matchActive %d, got %v",
				tc.expected, tc.user, tc.matchRateStatus, tc.matchActive, actual)
		}
	}
}

func TestValidateMatchRateAndUserOddsType(t *testing.T) {
	testCases := []struct {
		user     dto.User
		rate     int32
		expected bool
	}{
		{dto.User{OddsType: 2}, 2, true},
		{dto.User{OddsType: 2}, 1, false},
		{dto.User{OddsType: 3}, 3, true},
		{dto.User{OddsType: 3}, 5, false},
	}

	for _, tc := range testCases {
		actual := ValidateMatchRateAndUserOddsType(tc.user, tc.rate)
		if actual != tc.expected {
			t.Errorf("Expected %v for user %+v and rate %d, got %v", tc.expected, tc.user, tc.rate, actual)
		}
	}
}
