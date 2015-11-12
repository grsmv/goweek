package goweek

import (
	"time"
	"errors"
)

type Week struct {
	Days     []time.Time
	Year     int
	Number   int
	FirstDay int
}

func NewWeek(params ...int) (*Week, error) {
	if len(params) < 2 {
		return &Week{}, errors.New("NewWeek(): too few arguments, specify year and number of week")
	} else if params[0] < 0 {
		return &Week{}, errors.New("NewWeek(): year can't be less than zero")
	} else if params[1] < 1 || params[1] > 53 {
		return &Week{}, errors.New("NewWeek(): number of week can't be less than 1 or greater than 53")
	} else {
		var (
			week                = initWeek(params...)
			approximateDay      = (week.Number-1) * 7  // converting from human-readable to machine notation
			approximateFirstDay = 0
			commonNumberOfDays  = 0
			monthNumber         = 0
		)

		for index, numberOfDaysInMonth := range numberOfDays(week.Year) {
			if approximateDay >= commonNumberOfDays && approximateDay <= commonNumberOfDays+numberOfDaysInMonth {
				monthNumber = index
				break
			} else {
				commonNumberOfDays += numberOfDaysInMonth
			}
		}

		approximateFirstDay = approximateDay - commonNumberOfDays

		// workaround for calculation of day number in first week of the year
		if approximateFirstDay == 0 {
			approximateFirstDay = 1
		}

		// workaround for calculation of day number in last week of the year
		if approximateFirstDay == -1 {
			monthNumber = 11
			approximateFirstDay = numberOfDays(week.Year)[monthNumber] - 2
		}

		// composing week listing
		var estimatedDate = time.Date(week.Year, time.Month(monthNumber+1), approximateFirstDay, 0, 0, 0, 0, time.UTC)
		var estimatedFirstDayOfTheWeek = estimatedDate.AddDate(0, 0, -1*int(estimatedDate.Weekday()))

		for i := week.FirstDay; i <= week.FirstDay+6; i++ {
			week.Days = append(week.Days, estimatedFirstDayOfTheWeek.AddDate(0, 0, i))
		}

		return &week, nil
	}
}

func (week *Week) Next() Week {
	return *week
}

func (week *Week) Previous() Week {
	return *week
}

func initWeek(params ...int) Week {
	var week = Week{
		Year:   params[0],
		Number: params[1],
	}

	if len(params) < 3 {
		week.FirstDay = 0
	} else {
		week.FirstDay = params[2]
	}
	return week
}

func isLeapYear(year int) bool {
	if year%4 > 0 {
		return false
	} else if year%100 > 0 {
		return true
	} else if year%400 > 0 {
		return false
	} else {
		return true
	}
}

func numberOfDays(year int) (numbers []int) {
	numbers = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isLeapYear(year) {
		numbers[1] = 29
	}
	return
}