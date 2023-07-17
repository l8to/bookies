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
	user := dto.User{}
	matchRate := dto.MatchRate{}
	testCases := []struct {
		desc     string
		status   int32
		active   int32
		expected bool
	}{
		{
			desc:     "match rate status is not 1",
			status:   0,
			active:   2,
			expected: false,
		},
		{
			desc:     "match rate match active is 0",
			status:   1,
			active:   0,
			expected: false,
		},
		{
			desc:     "match rate match active is greater than 3",
			status:   1,
			active:   4,
			expected: false,
		},
		{
			desc:     "valid match rate",
			status:   1,
			active:   2,
			expected: true,
		},
	}
	for _, tc := range testCases {
		matchRate.Status = tc.status
		matchRate.Match.Active = tc.active
		if result := ValidateMatchRateActive(user, matchRate); result != tc.expected {
			t.Errorf("Test case failed for %s: Expected %v, got %v", tc.desc, tc.expected, result)
		}
	}
}

func TestValidateMatchRateAndUserOddsType(t *testing.T) {
	testCases := []struct {
		user      dto.User
		matchRate dto.MatchRate
		expected  bool
	}{
		{
			dto.User{
				OddsType: 20,
			},
			dto.MatchRate{
				Rate: int32(20),
			},
			true},
		{
			dto.User{
				OddsType: 20,
			},
			dto.MatchRate{
				Rate: int32(10),
			},
			false},
	}
	for i, testCase := range testCases {
		result := ValidateMatchRateAndUserOddsType(testCase.user, testCase.matchRate)
		if result != testCase.expected {
			t.Errorf("Test case %d failed: Expected %v, got %v", i+1, testCase.expected, result)
		}
	}
}
