package validate

import (
	"testing"
	"time"

	"github.com/l8to/bookies/dto"
)

func TestValidateMatchTime(t *testing.T) {
	koTime := time.Date(2022, time.January, 1, 9, 0, 0, 0, time.UTC)
	isLive := false

	testCases := []struct {
		koTime   time.Time
		isLive   bool
		oddsType string
		expected bool
	}{
		{koTime, isLive, "ht_hdp_home", true},                           // oddsType is "ht_hdp_home", should return true
		{koTime, isLive, "ht_hdp_away", true},                           // oddsType is "ht_hdp_away", should return true
		{koTime, isLive, "ht_ou_over", true},                            // oddsType is "ht_ou_over", should return true
		{koTime, isLive, "ht_ou_under", true},                           // oddsType is "ht_ou_under", should return true
		{koTime, isLive, "ft_hdp_home", true},                           // oddsType is not an HT odds type, should return true
		{koTime, isLive, "ft_hdp_away", true},                           // oddsType is not an HT odds type, should return true
		{koTime, isLive, "ft_ou_over", true},                            // oddsType is not an HT odds type, should return true
		{koTime, isLive, "ft_ou_under", true},                           // oddsType is not an HT odds type, should return true
		{koTime, true, "ht_hdp_home", false},                            // isLive is true, should return false
		{time.Now().Add(-time.Minute * 50), true, "ht_ou_under", false}, // oddsType is "ht_ou_under", should return true
		{time.Now().Add(time.Hour), true, "ft_hdp_home", false},         // oddsType is not an HT odds type, should return true
	}

	for _, tc := range testCases {
		actual := ValidateMatchTime(tc.koTime, tc.isLive, tc.oddsType)
		if actual != tc.expected {
			t.Errorf("Expected %v for koTime %v, isLive %v, and oddsType %s, got %v",
				tc.expected, tc.koTime, tc.isLive, tc.oddsType, actual)
		}
	}
}

func TestValidateMatchKOTime(t *testing.T) {
	testCases := []struct {
		name     string
		koTime   time.Time
		isLive   bool
		expected bool
	}{
		{
			name:     "WithLiveMatch",
			koTime:   time.Now().Add(time.Hour),
			isLive:   true,
			expected: true,
		},
		{
			name:     "WithExpiredLiveMatch",
			koTime:   time.Now().Add(-time.Hour),
			isLive:   true,
			expected: false,
		},
		{
			name:     "WithNonLiveMatch",
			koTime:   time.Now().Add(time.Hour),
			isLive:   false,
			expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := ValidateMatchKOTime(testCase.koTime, testCase.isLive)
			if result != testCase.expected {
				t.Errorf("Expected %v but got %v", testCase.expected, result)
			}
		})
	}
}

func TestValidateMatchKOTimeHT(t *testing.T) {
	testCases := []struct {
		name     string
		koTime   time.Time
		isLive   bool
		expected bool
	}{
		{
			name:     "Before Half Time and Is Live",
			koTime:   time.Now().Add(-time.Minute * 30),
			isLive:   true,
			expected: false,
		},
		{
			name:     "After Half Time and Is Live",
			koTime:   time.Now().Add(time.Minute * 70),
			isLive:   true,
			expected: false,
		},
		{
			name:     "Between Half Time and Full Time and Is Live",
			koTime:   time.Now().Add(-time.Minute * 50),
			isLive:   true,
			expected: true,
		},
		{
			name:     "Before Half Time and Not Live",
			koTime:   time.Now().Add(-time.Minute * 30),
			isLive:   false,
			expected: true,
		},
		{
			name:     "After Half Time and Not Live",
			koTime:   time.Now().Add(time.Minute * 70),
			isLive:   false,
			expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := ValidateMatchKOTimeHT(testCase.koTime, testCase.isLive)
			if result != testCase.expected {
				t.Errorf("Expected %v but got %v", testCase.expected, result)
			}
		})
	}
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
