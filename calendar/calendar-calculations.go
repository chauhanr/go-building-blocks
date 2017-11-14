package calendar

import (
	"errors"
	"fmt"
	"strconv"
)

const BASE_DAY = 1
const BASE_MONTH = 1
const BASE_YEAR = 1970
const BASE_LEAP_YEAR = 1972

var WEEK_DAYS = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday" ,"Saturday" }
var BASE_YEAR_WEEK = []string{"Thursday", "Friday" ,"Saturday","Sunday", "Monday", "Tuesday", "Wednesday"}
var DAYS_IN_MONTH = []int{31,28,31,30,31,30,31,31,30,31,30,31}
var DAYS_IN_MONTH_LEAP = []int{31,29,31,30,31,30,31,31,30,31,30,31}
var MONTHS = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}


func GetDayOfTheWeek(day int, month int, year int) (string, error){
    if year < BASE_YEAR {
    	return "", errors.New("Cannot calculate before 1 Jan 1970")
	}
    isLeapYear := isLeapYear(year)

    err := isDayMonthValid(day, month,isLeapYear)
    if err != nil{
    	return "", err
	}

	offset := calculateDayOffset(day,month,year,isLeapYear)
	return BASE_YEAR_WEEK[offset], nil
}

func calculateDayOffset (day int, month int, year int, isLeapYear bool) int {
	 dYears := year - BASE_YEAR
	 numLeapYears := calculateLeapYearCount(year)
	//fmt.Printf("Number of Leap Years between 1970 and %d is %d\n", year,numLeapYears)
	 totalDays := numLeapYears*366 + (dYears-numLeapYears)*365

	  totalDaysYear := 0;
	  if isLeapYear{
	  	 index := 0
	  	 for ;index < month-1; index++ {
			 totalDaysYear += DAYS_IN_MONTH_LEAP[index]
		 }
	  }else{
		  index := 0
		  for ;index < month-1; index++ {
			  totalDaysYear += DAYS_IN_MONTH[index]
		  }
	  }
	totalDaysYear += day - BASE_DAY
	totalDays += totalDaysYear

	return totalDays%7
}

func calculateLeapYearCount(year int) int{
	yearCounter := BASE_LEAP_YEAR;
	leapYearCounter := 0
	for yearCounter < year{
		leapYearCounter++
		yearCounter += 4
	}
	return leapYearCounter
}



func isLeapYear(year int) bool{
	year2 := year%100
	if year2%4 == 0 {
		return true;
	}
	return false
}

func isDayMonthValid(day int, month int, isLeapYear bool) error{
	if day < 1 || day > 31 {
		return errors.New("Invalid day cannot be less than 1 or greater than 31")
	}
	if month < 1 && month > 12{
		return errors.New("Invalid month cannot be less than 1 or greater than 12")
	}

	if isLeapYear {
		dayInMonth := DAYS_IN_MONTH_LEAP[month-1]
		if day > dayInMonth {
			return errors.New("Invalid day month"+MONTHS[month-1]+" cannot have more than "+string(dayInMonth)+" days")
		}
	}else{
		dayInMonth := DAYS_IN_MONTH[month-1]
		if day > dayInMonth {
			return errors.New("Invalid day month"+MONTHS[month-1]+" cannot have more than "+string(dayInMonth)+" days")
		}
	}
	return nil
}

func PrintMonthlyCalendar(month int, year int) error{
	dayOfWeek, err := GetDayOfTheWeek(1,month,year)
	if err != nil{
		return err
	}

	var monthCalendar = [][]string{{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri" ,"Sat" },}

	for day := 1; day <= DAYS_IN_MONTH[month]; {
		row := []string{}
		for index := 0; index < 7 && day <= DAYS_IN_MONTH[month-1]; index++ {
			if day == 1 {
				if WEEK_DAYS[index] == dayOfWeek {
					row = append(row, strconv.Itoa(day))
					day++
				} else {
					row = append(row, "")
				}
			} else {
				row = append(row, strconv.Itoa(day))
				day++
			}
		}
		monthCalendar = append(monthCalendar, row)
	}
	fmt.Printf("Calnedar for %s, %d\n", MONTHS[month-1],year)
	printCalendar(monthCalendar)
	return nil
}

func printCalendar(month [][]string){

	for _, row := range month{
		for _, ind := range row{
			fmt.Printf("%3s ", ind)
		}
		fmt.Printf("\n")
	}
}