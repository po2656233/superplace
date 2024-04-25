package exTime

import (
	"testing"
)

func TestSuperTime_IsNow(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsNow())
}

func TestSuperTime_IsFuture(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsFuture())
}

func TestSuperTime_IsPast(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsPast())
}

func TestSuperTime_IsLeapYear(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsLeapYear())
}

func TestSuperTime_IsLongYear(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsLongYear())
}

func TestSuperTime_IsJanuary(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsJanuary())
}

func TestSuperTime_IsFebruary(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsFebruary())
}

func TestSuperTime_IsMarch(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsMarch())
}

func TestSuperTime_IsApril(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsApril())
}

func TestSuperTime_IsMay(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsMay())
}

func TestSuperTime_IsJune(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsJune())
}

func TestSuperTime_IsJuly(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsJuly())
}

func TestSuperTime_IsAugust(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsAugust())
}

func TestSuperTime_IsSeptember(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsSeptember())
}

func TestSuperTime_IsOctober(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsOctober())
}

func TestSuperTime_IsDecember(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsDecember())
}

func TestSuperTime_IsMonday(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsMonday())
}

func TestSuperTime_IsTuesday(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsTuesday())
}

func TestSuperTime_IsWednesday(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsWednesday())
}

func TestSuperTime_IsThursday(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsThursday())
}

func TestSuperTime_IsFriday(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsFriday())
}

func TestSuperTime_IsSaturday(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsSaturday())
}

func TestSuperTime_IsSunday(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsSunday())
}

func TestSuperTime_IsWeekday(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsWeekday())
}

func TestSuperTime_IsWeekend(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsWeekend())
}

func TestSuperTime_IsYesterday(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsYesterday())
}

func TestSuperTime_IsYesterday1(t *testing.T) {
	now := Now()
	now.SubDay()
	t.Logf("result = %v", now.IsYesterday())
}

func TestSuperTime_IsToday(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsToday())
}

func TestSuperTime_IsTomorrow(t *testing.T) {
	now := Now()
	t.Logf("result = %v", now.IsTomorrow())
}
