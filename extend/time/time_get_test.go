package exTime

import (
	"testing"
)

func TestSuperTime_DaysInYear(t *testing.T) {
	t.Logf("result = %v", Now().DaysInYear())
}

func TestSuperTime_DaysInMonth(t *testing.T) {
	t.Logf("result = %v", Now().DaysInMonth())
}

func TestSuperTime_MonthOfYear(t *testing.T) {
	t.Logf("result = %v", Now().MonthOfYear())
}

func TestSuperTime_DayOfYear(t *testing.T) {
	t.Logf("result = %v", Now().DayOfYear())
}

func TestSuperTime_DayOfMonth(t *testing.T) {
	t.Logf("result = %v", Now().DayOfMonth())
}

func TestSuperTime_DayOfWeek(t *testing.T) {
	t.Logf("result = %v", Now().DayOfWeek())
}

func TestSuperTime_WeekOfYear(t *testing.T) {
	t.Logf("result = %v", Now().WeekOfYear())
}

func TestSuperTime_WeekOfMonth(t *testing.T) {
	t.Logf("result = %v", Now().WeekOfMonth())
}

func TestSuperTime_Year(t *testing.T) {
	t.Logf("result = %v", Now().Year())
}

func TestSuperTime_Quarter(t *testing.T) {
	t.Logf("result = %v", Now().Quarter())
}

func TestSuperTime_Month(t *testing.T) {
	t.Logf("result = %v", Now().Month())
}

func TestSuperTime_Week(t *testing.T) {
	t.Logf("result = %v", Now().Week())
}

func TestSuperTime_Day(t *testing.T) {
	t.Logf("result = %v", Now().Day())
}

func TestSuperTime_Hour(t *testing.T) {
	t.Logf("result = %v", Now().Hour())
}

func TestSuperTime_Minute(t *testing.T) {
	t.Logf("result = %v", Now().Minute())
}

func TestSuperTime_Second(t *testing.T) {
	t.Logf("result = %v", Now().Second())
}

func TestSuperTime_Millisecond(t *testing.T) {
	t.Logf("result = %v", Now().Millisecond())
}

func TestSuperTime_Microsecond(t *testing.T) {
	t.Logf("result = %v", Now().Microsecond())
}

func TestSuperTime_Nanosecond(t *testing.T) {
	t.Logf("result = %v", Now().Nanosecond())
}

func TestSuperTime_Timezone(t *testing.T) {
	t.Logf("result = %v", Now().Timezone())
}
