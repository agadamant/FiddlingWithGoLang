/*
Package that implements on-duty calculation logic for the specified year.
Date attributes and values are stored as contants.

Output sample:
2020-04-02 02:17:00 - 2020-04-09 02:17:00
2020-05-07 02:17:00 - 2020-05-14 02:17:00
2020-06-11 02:17:00 - 2020-06-18 02:17:00
2020-07-16 02:17:00 - 2020-07-23 02:17:00
2020-08-20 02:17:00 - 2020-08-27 02:17:00
2020-09-24 02:17:00 - 2020-10-01 02:17:00
2020-10-29 02:17:00 - 2020-11-05 02:17:00
2020-12-03 02:17:00 - 2020-12-10 02:17:00

*/
package main

import (
	"fmt"
	"time"
)

// Location - tz database time zone
const Location = "Europe/Vilnius"

// DateFormat - date format to be used for output
const DateFormat = "2006-01-02 15:04:05"

// NumberOfPeople - number of people in the on-duty cycle
const NumberOfPeople = 5

// OnDutyLengthDays - length of one on-duty shift in days
const OnDutyLengthDays = 7

// InitYear - initial year
const InitYear = 2020

// InitMonth - initial month
const InitMonth = time.February

// InitDay - initial day
const InitDay = 26

// InitHour - initial hour
const InitHour = 17

func main() {
	printOnCallDutyUntilTheEndOfTheYear()
}

// Returns Date struct based on value of Init* constants
func getInitDate() time.Time {
	location, _ := time.LoadLocation(Location)
	initDate :=
		time.Date(InitYear, InitMonth, InitDay, InitDay, InitHour, 0, 0, location)
	return initDate
}

// Calculated and prints on-duty date ranges for year value taken from InitYear,
// starting from date speciefied in Init* constants
func printOnCallDutyUntilTheEndOfTheYear() {
	iterLength := NumberOfPeople * OnDutyLengthDays
	currentDate := time.Now()

	for iterDate := getInitDate(); iterDate.Year() == currentDate.Year(); iterDate = iterDate.AddDate(0, 0, iterLength) {
		if iterDate.After(currentDate) {
			fmt.Println(
				iterDate.Format(DateFormat),
				"-",
				iterDate.AddDate(0, 0, OnDutyLengthDays).Format(DateFormat))
		}
	}
}
