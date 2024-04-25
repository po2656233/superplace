package exTime

import (
	"testing"
)

func TestSuperTime_DiffInYears(t *testing.T) {
	ct1 := CreateFromDate(2012, 12, 1)
	ct2 := CreateFromDate(2022, 2, 1)

	years := ct1.DiffInYears(ct2)
	t.Logf("result = %v", years)
}

func TestSuperTime_DiffInYearsWithAbs(t *testing.T) {
	ct1 := CreateFromDate(2012, 12, 1)
	ct2 := CreateFromDate(2022, 2, 1)

	years := ct1.DiffInYearsWithAbs(ct2)

	t.Logf("result = %v", years)
}

func TestSuperTime_DiffInMonths(t *testing.T) {
	ct1 := CreateFromDate(2021, 12, 15)
	ct2 := CreateFromDate(2022, 1, 1)

	years := ct1.DiffInMonths(ct2)

	t.Logf("result = %v", years)
}

func TestSuperTime_DiffInMonthsWithAbs(t *testing.T) {
	ct1 := CreateFromDate(2021, 12, 15)
	ct2 := CreateFromDate(2022, 1, 1)

	years := ct1.DiffInMonthsWithAbs(ct2)

	t.Logf("result = %v", years)
}

func TestSuperTime_DiffInWeeks(t *testing.T) {
	ct1 := CreateFromDate(2021, 12, 15)
	ct2 := CreateFromDate(2022, 1, 1)

	years := ct1.DiffInWeeks(ct2)

	t.Logf("result = %v", years)
}

func TestSuperTime_DiffInWeeksWithAbs(t *testing.T) {
	ct1 := CreateFromDate(2021, 12, 15)
	ct2 := CreateFromDate(2022, 1, 1)

	years := ct1.DiffInWeeksWithAbs(ct2)

	t.Logf("result = %v", years)
}

func TestSuperTime_DiffInDays(t *testing.T) {
	ct1 := CreateFromDate(2021, 12, 15)
	ct2 := CreateFromDate(2022, 1, 1)

	years := ct1.DiffInDays(ct2)

	t.Logf("result = %v", years)
}

func TestSuperTime_DiffInDaysWithAbs(t *testing.T) {
	ct1 := CreateFromDate(2021, 12, 15)
	ct2 := CreateFromDate(2022, 1, 1)

	years := ct1.DiffInDaysWithAbs(ct2)

	t.Logf("result = %v", years)
}

func TestSuperTime_DiffInHours(t *testing.T) {
	ct1 := CreateFromDate(2021, 12, 15)
	ct2 := CreateFromDate(2022, 1, 1)

	years := ct1.DiffInHours(ct2)

	t.Logf("result = %v", years)
}

func TestSuperTime_DiffInHoursWithAbs(t *testing.T) {
	ct1 := CreateFromDate(2021, 12, 15)
	ct2 := CreateFromDate(2022, 1, 1)

	years := ct1.DiffInHoursWithAbs(ct2)

	t.Logf("result = %v", years)
}

func TestSuperTime_DiffInSeconds(t *testing.T) {
	ct1 := CreateFromDate(2021, 12, 15)
	ct2 := CreateFromDate(2022, 1, 1)

	years := ct1.DiffInSeconds(ct2)

	t.Logf("result = %v", years)
}

func TestSuperTime_DiffInSecondsWithAbs(t *testing.T) {
	ct1 := CreateFromDate(2021, 12, 15)
	ct2 := CreateFromDate(2022, 1, 1)

	years := ct1.DiffInSecondsWithAbs(ct2)

	t.Logf("result = %v", years)
}
