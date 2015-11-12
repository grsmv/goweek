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

var expectedDaysWithMondayAsFirstDayOfTheWeek = []time.Time{
	time.Date(2015, 11, 9, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 10, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 11, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 12, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 13, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 14, 0, 0, 0, 0, time.UTC),
	time.Date(2015, 11, 15, 0, 0, 0, 0, time.UTC),
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
