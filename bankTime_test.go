package banktime

import (
	"testing"
	"time"
)

var (
	est *time.Location
)

func init() {
	// US Eastern Time Zone
	est, _ = time.LoadLocation("America/New_York")
}

func TestIsBankingDay(t *testing.T) {
	tests := []struct {
		Date     time.Time
		Expected bool
	}{
		// new years day
		{time.Date(2018, time.January, 1, 1, 0, 0, 0, est), false},
		// Wednesday Canary test
		{time.Date(2018, time.January, 3, 1, 0, 0, 0, est), true},
		// saturday
		{time.Date(2018, time.January, 6, 1, 0, 0, 0, est), false},
		// sunday
		{time.Date(2018, time.January, 7, 1, 0, 0, 0, est), false},
		// Martin Luther King, JR. Day
		{time.Date(2018, time.January, 15, 1, 0, 0, 0, est), false},
		// Presidents' Day
		{time.Date(2018, time.February, 19, 1, 0, 0, 0, est), false},
		// Memorial Day
		{time.Date(2018, time.May, 28, 1, 0, 0, 0, est), false},
		// Independence Day
		{time.Date(2018, time.July, 4, 1, 0, 0, 0, est), false},
		// Labor Day
		{time.Date(2018, time.September, 3, 1, 0, 0, 0, est), false},
		// Columbus Day
		{time.Date(2018, time.October, 8, 1, 0, 0, 0, est), false},
		// Vesterans' Day Observed on the monday
		{time.Date(2018, time.November, 12, 1, 0, 0, 0, est), false},
		// Thanksgiving Day
		{time.Date(2018, time.November, 22, 1, 0, 0, 0, est), false},
		// Christmas Day
		{time.Date(2018, time.December, 25, 1, 0, 0, 0, est), false},
	}
	for _, test := range tests {
		actual := NewBankTime(test.Date).IsBankingDay()
		if actual != test.Expected {
			t.Errorf("Date %s: expected %t, got %t", test.Date, test.Expected, actual)
		}

		actual = New(test.Date, est).IsBankingDay()
		if actual != test.Expected {
			t.Errorf("Date %s: expected %t, got %t", test.Date, test.Expected, actual)
		}
	}
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		Date     time.Time
		Expected bool
	}{
		// saturday
		{time.Date(2018, time.January, 6, 1, 0, 0, 0, est), true},
		// sunday
		{time.Date(2018, time.January, 7, 1, 0, 0, 0, est), true},
		// monday
		{time.Date(2018, time.January, 9, 1, 0, 0, 0, est), false},
	}
	for _, test := range tests {
		actual := NewBankTime(test.Date).IsWeekend()
		if actual != test.Expected {
			t.Errorf("Date %s: expected %t, got %t", test.Date, test.Expected, actual)
		}

		actual = New(test.Date, est).IsWeekend()
		if actual != test.Expected {
			t.Errorf("Date %s: expected %t, got %t", test.Date, test.Expected, actual)
		}
	}
}

func TestAddBankingDay(t *testing.T) {
	tests := []struct {
		Date   time.Time
		Future time.Time
		Days   int
	}{
		// Thursday add two days over a monday holiday abd needs to be following tuesday
		{time.Date(2018, time.January, 11, 1, 0, 0, 0, est), time.Date(2018, time.January, 16, 1, 0, 0, 0, est), 2},
	}
	for _, test := range tests {
		actual := NewBankTime(test.Date).AddBankingDay(test.Days)
		if !actual.Equal(test.Future) {
			t.Errorf("Adding %d days: expected %s, got %s", test.Days, test.Future.Weekday().String(), actual)
		}

		actual = New(test.Date, est).AddBankingDay(test.Days)
		if !actual.Equal(test.Future) {
			t.Errorf("Adding %d days: expected %s, got %s", test.Days, test.Future.Weekday().String(), actual)
		}
	}
}

//Test to ensure that
func TestUTCisEST(t *testing.T) {

	//@TODO create dates that are on a different day earlier or later than a holiday in different time zone but land on the est time date.
}
