package calendar

import "testing"

var DayofWeekTests = []struct{
	day int
	month int
	year int
	dayOfWeek string
}{
	{
		31,
		1,
		1970,
		"Saturday",
	},
	{
		26,
		4,
		1970,
		"Sunday",
	},
	{
		29,
		3,
		1980,
		"Saturday",
	},
	{
		29,
		2,
		2000,
		"Tuesday",
	},
	{
		19,
		10,
		2017,
		"Thursday",
	},
}

func TestDayOfTheWeek(t *testing.T) {

	for _, dayWeekTest := range DayofWeekTests {
		dayOfWeek := dayWeekTest.dayOfWeek
		day := dayWeekTest.day
		month := dayWeekTest.month
		year := dayWeekTest.year
		calcDayOfWeek, err := GetDayOfTheWeek(day,month,year)

		if err != nil {
			t.Logf("Error in parsing date %s ", err.Error())
		}else if calcDayOfWeek != dayOfWeek{
			t.Errorf("Expected day of week for (%d/%d/%d) is %s but calculated value is %s ", day, month, year, dayOfWeek, calcDayOfWeek)
		}else{
			t.Logf("Correctly calculated day of week for (%d/%d/%d) as %s", day, month, year, dayOfWeek)
		}
	}
}


var monthPrintTest = []struct{
	month int
	year  int
}{
	{3, 1980},
	{10, 2017},
	{3, 2000},
}

func TestPrintMonthlyCalendar(t *testing.T) {

	for _, testCase := range monthPrintTest {
		err := PrintMonthlyCalendar(testCase.month, testCase.year)
		if err != nil {
			t.Errorf("Error in printing Calendar for %d/%d", testCase.month, testCase.year)
		}
	}
}
