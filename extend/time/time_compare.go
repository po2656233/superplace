package exTime

import "time"

// IsNow 是否是当前时间
func (c SuperTime) IsNow() bool {
	return c.ToSecond() == Now().ToSecond()
}

// IsFuture 是否是未来时间
func (c SuperTime) IsFuture() bool {
	return c.ToSecond() > Now().ToSecond()
}

// IsPast 是否是过去时间
func (c SuperTime) IsPast() bool {
	return c.ToSecond() < Now().ToSecond()
}

// IsLeapYear 是否是闰年
func (c SuperTime) IsLeapYear() bool {
	year := c.Year()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// IsLongYear 是否是长年
func (c SuperTime) IsLongYear() bool {
	_, w := time.Date(c.Year(), time.December, 31, 0, 0, 0, 0, c.Location()).ISOWeek()
	return w == WeeksPerLongYear
}

// IsJanuary 是否是一月
func (c SuperTime) IsJanuary() bool {
	return c.Time.Month() == time.January
}

// IsFebruary 是否是二月
func (c SuperTime) IsFebruary() bool {
	return c.Time.Month() == time.February
}

// IsMarch 是否是三月
func (c SuperTime) IsMarch() bool {
	return c.Time.Month() == time.March
}

// IsApril 是否是四月
func (c SuperTime) IsApril() bool {
	return c.Time.Month() == time.April
}

// IsMay 是否是五月
func (c SuperTime) IsMay() bool {
	return c.Time.Month() == time.May
}

// IsJune 是否是六月
func (c SuperTime) IsJune() bool {
	return c.Time.Month() == time.June
}

// IsJuly 是否是七月
func (c SuperTime) IsJuly() bool {
	return c.Time.Month() == time.July
}

// IsAugust 是否是八月
func (c SuperTime) IsAugust() bool {
	return c.Time.Month() == time.August
}

// IsSeptember 是否是九月
func (c SuperTime) IsSeptember() bool {
	return c.Time.Month() == time.September
}

// IsOctober 是否是十月
func (c SuperTime) IsOctober() bool {
	return c.Time.Month() == time.October
}

// IsNovember 是否是十一月
func (c SuperTime) IsNovember() bool {
	return c.Time.Month() == time.November
}

// IsDecember 是否是十二月
func (c SuperTime) IsDecember() bool {
	return c.Time.Month() == time.December
}

// IsMonday 是否是周一
func (c SuperTime) IsMonday() bool {
	return c.Time.Weekday() == time.Monday
}

// IsTuesday 是否是周二
func (c SuperTime) IsTuesday() bool {
	return c.Time.Weekday() == time.Tuesday
}

// IsWednesday 是否是周三
func (c SuperTime) IsWednesday() bool {
	return c.Time.Weekday() == time.Wednesday
}

// IsThursday 是否是周四
func (c SuperTime) IsThursday() bool {
	return c.Time.Weekday() == time.Thursday
}

// IsFriday 是否是周五
func (c SuperTime) IsFriday() bool {
	return c.Time.Weekday() == time.Friday
}

// IsSaturday 是否是周六
func (c SuperTime) IsSaturday() bool {
	return c.Time.Weekday() == time.Saturday
}

// IsSunday 是否是周日
func (c SuperTime) IsSunday() bool {
	return c.Time.Weekday() == time.Sunday
}

// IsWeekday 是否是工作日
func (c SuperTime) IsWeekday() bool {
	return !c.IsSaturday() && !c.IsSunday()
}

// IsWeekend 是否是周末
func (c SuperTime) IsWeekend() bool {
	return c.IsSaturday() || c.IsSunday()
}

// IsYesterday 是否是昨天
func (c SuperTime) IsYesterday() bool {
	yesterday := Now()
	yesterday.SubDay()
	return c.ToDateFormat() == yesterday.ToDateFormat()
}

// IsToday 是否是今天
func (c SuperTime) IsToday() bool {
	return c.ToDateFormat() == Now().ToDateFormat()
}

// IsTomorrow 是否是明天
func (c SuperTime) IsTomorrow() bool {
	tomorrow := Now()
	tomorrow.AddDay()
	return c.ToDateFormat() == tomorrow.ToDateFormat()
}
