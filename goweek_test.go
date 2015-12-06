package goweek

import (
	"reflect"
	"testing"
	"time"
)

var expectedDays = []time.Time{
	time.Date(2015, 11, 8, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 9, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 10, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 11, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 12, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 13, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 14, 0, 0, 0, 0, time.UTC),
}

var expectedDaysWithMondayAsFirstDayOfTheWeek = []time.Time {
	time.Date(2015, 11, 9, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 10, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 11, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 12, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 13, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 14, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 15, 0, 0, 0, 0, time.UTC),
}

var expectedDaysForNextWeek = []time.Time {
	time.Date(2015, 11, 15, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 16, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 17, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 18, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 19, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 20, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 21, 0, 0, 0, 0, time.UTC),
}

// assuming that in case of 2015 year last week of 2015 in first week of 2016
var expectedDaysForNextWeekWithYearSwitch = []time.Time{
	time.Date(2016, 1,  3,  0, 0, 0, 0, time.UTC),
	time.Date(2016, 1,  4,  0, 0, 0, 0, time.UTC),
	time.Date(2016, 1,  5,  0, 0, 0, 0, time.UTC),
	time.Date(2016, 1,  6,  0, 0, 0, 0, time.UTC),
	time.Date(2016, 1,  7,  0, 0, 0, 0, time.UTC),
	time.Date(2016, 1,  8,  0, 0, 0, 0, time.UTC),
	time.Date(2016, 1,  9,  0, 0, 0, 0, time.UTC),
}

var expectedDaysForPreviousWeek = []time.Time{
	time.Date(2015, 11, 1, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 2, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 3, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 4, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 5, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 6, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 7, 0, 0, 0, 0, time.UTC),
}

// assuming that in case of 2015 year first week of 2015 in last week of 2014
var expectedDaysForPreviousWeekWithYearSwitch = []time.Time{
	time.Date(2014, 12, 21, 0, 0, 0, 0, time.UTC),
	time.Date(2014, 12, 22, 0, 0, 0, 0, time.UTC),
	time.Date(2014, 12, 23, 0, 0, 0, 0, time.UTC),
	time.Date(2014, 12, 24, 0, 0, 0, 0, time.UTC),
	time.Date(2014, 12, 25, 0, 0, 0, 0, time.UTC),
	time.Date(2014, 12, 26, 0, 0, 0, 0, time.UTC),
	time.Date(2014, 12, 27, 0, 0, 0, 0, time.UTC),
}

func Test_NormalUsage(t *testing.T) {
	var week, _ = NewWeek(2015, 46)
	if len(week.Days) != 7 {
		t.Error("Unexpected number of Week.Days, \n expected %v, \n given %v", 7, len(week.Days))
	}
	if week.Year != 2015 {
		t.Error("Unexpected Week.Year, \n expected %v, \n given %v", 2015, week.Year)
	}
	if week.Number != 46 {
		t.Error("Unexpected Week.Number, \n expected %v, \n given %v", 46, week.Number)
	}
	if week.FirstDay != 0 {
		t.Error("Unexpected Week.FirstDay, \n expected %v, \n given %v", 0, week.FirstDay)
	}
	if !reflect.DeepEqual(expectedDays, week.Days) {
		t.Errorf("Unexpected Week.Days, \n expected %v, \n given %v", expectedDays, week.Days)
	}
}

func Test_MondayAsFirstWeekDay(t *testing.T) {
	var week, _ = NewWeek(2015, 46, 1)
	if week.FirstDay != 1 {
		t.Error("Unexpected Week.FirstDay, \n expected %v, \n given %v", 1, week.FirstDay)
	}
	if !reflect.DeepEqual(expectedDaysWithMondayAsFirstDayOfTheWeek, week.Days) {
		t.Errorf("Unexpected Week.Days, \n expected %v, \n given %v", expectedDaysWithMondayAsFirstDayOfTheWeek, week.Days)
	}
}

func Test_ErrorThrowing(t *testing.T) {
	var _, errorA = NewWeek()
	if errorA.Error() != "NewWeek(): too few arguments, specify year and number of week" {
		t.Error("Error expected when passing too few arguments (no args given)")
	}

	var _, errorB = NewWeek(2015)
	if errorB.Error() != "NewWeek(): too few arguments, specify year and number of week" {
		t.Error("Error expected when passing too few arguments (only year given)")
	}

	var _, errorC = NewWeek(2015, 54)
	if errorC.Error() != "NewWeek(): number of week can't be less than 1 or greater than 53" {
		t.Error("Error expected when passing incorrect week number")
	}

	var _, errorD = NewWeek(-1, 53)
	if errorD.Error() != "NewWeek(): year can't be less than zero" {
		t.Error("Error expected when passing incorrect year number")
	}
}

func Test_NextWeek (t *testing.T) {
	var week, _ = NewWeek(2015, 46)
	var nextWeek, _ = week.Next()
	if !reflect.DeepEqual(expectedDaysForNextWeek, nextWeek.Days) {
		t.Errorf("Unexpected Week.Next(), \n expected %v, \n given %v", expectedDaysForNextWeek, nextWeek.Days)
	}

	var weekA, _ = NewWeek(2015, 53)
	var nextWeekA, errA = weekA.Next()
	if !reflect.DeepEqual(expectedDaysForNextWeekWithYearSwitch, nextWeekA.Days) {
		t.Errorf("Unexpected Week.Next() with year switch, \n expected %v, \n given %v", expectedDaysForNextWeekWithYearSwitch, nextWeekA.Days)
	}

	if errA != nil {
		t.Error(errA.Error())
	}
}

func Test_PreviousWeek (t *testing.T) {
	var week, _ = NewWeek(2015, 46)
	var previousWeek, _ = week.Previous()
	if !reflect.DeepEqual(expectedDaysForPreviousWeek, previousWeek.Days) {
		t.Errorf("Unexpected Week.Previous(), \n expected %v, \n given %v", expectedDaysForPreviousWeek, previousWeek.Days)
	}

	var weekA, _ = NewWeek(2015, 1)
	var previousWeekA, errA = weekA.Previous()
	if !reflect.DeepEqual(expectedDaysForPreviousWeekWithYearSwitch, previousWeekA.Days) {
		t.Errorf("Unexpected Week.Previous() with year switch, \n expected %v, \n given %v", expectedDaysForPreviousWeekWithYearSwitch, previousWeekA.Days)
	}

	if errA != nil {
		t.Error(errA.Error())
	}
}

func Test_YearsVerge (t *testing.T) {
	var weekA, _ = NewWeek(2015, 53)
	var nextWeek, _ = weekA.Next()
	if nextWeek.Number != 2 {
		t.Errorf("Unexpected Week.Next() at the verge of years, \n expected %v, \n given %v", 2, nextWeek.Number)
	}

	var weekB, _ = NewWeek(2015, 1)
	var previousWeek, _ = weekB.Previous()
	if previousWeek.Number != 52 {
		t.Errorf("Unexpected Week.Previous() at the verge of years, \n expected %v, \n given %v", 52, previousWeek.Number)
	}
}

func Test_YearsVergeWithoutNextAndPreviousWeekFixing (t *testing.T) {
	var weekA, _ = NewWeek(2016, 53)
	var nextWeek, _ = weekA.Next()
	if nextWeek.Number != 1 {
		t.Errorf("Unexpected Week.Next() at the verge of years (without fixing), \n expected %v, \n given %v", 1, nextWeek.Number)
	}

	var weekB, _ = NewWeek(2017, 1)
	var previousWeek, _ = weekB.Previous()
	if previousWeek.Number != 53 {
		t.Errorf("Unexpected Week.Previous() at the verge of years (without fixing), \n expected %v, \n given %v", 53, previousWeek.Number)
	}
}