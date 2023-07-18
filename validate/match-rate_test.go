package validate

import (
	"testing"
	"time"

	"github.com/l8to/bookies/dto"
)

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
