package exTime

// DiffInYears 相差多少年
func (c SuperTime) DiffInYears(end SuperTime) int64 {
	return c.DiffInMonths(end) / 12
}

// DiffInYearsWithAbs 相差多少年(绝对值)
func (c SuperTime) DiffInYearsWithAbs(end SuperTime) int64 {
	return GetAbsValue(c.DiffInYears(end))
}

// DiffInMonths 相差多少月
func (c SuperTime) DiffInMonths(end SuperTime) int64 {
	dy, dm, dd := end.Year()-c.Year(), end.Month()-c.Month(), end.Day()-c.Day()

	if dd < 0 {
		dm = dm - 1
	}
	if dy == 0 && dm == 0 {
		return 0
	}
	if dy == 0 && dm != 0 && dd != 0 {
		if int(end.DiffInHoursWithAbs(c)) < c.DaysInMonth()*HoursPerDay {
			return 0
		}
		return int64(dm)
	}

	return int64(dy*MonthsPerYear + dm)
}

// DiffInMonthsWithAbs 相差多少月(绝对值)
func (c SuperTime) DiffInMonthsWithAbs(end SuperTime) int64 {
	return GetAbsValue(c.DiffInMonths(end))
}

// DiffInWeeks 相差多少周
func (c SuperTime) DiffInWeeks(end SuperTime) int64 {
	return c.DiffInDays(end) / DaysPerWeek
}

// DiffInWeeksWithAbs 相差多少周(绝对值)
func (c SuperTime) DiffInWeeksWithAbs(end SuperTime) int64 {
	return GetAbsValue(c.DiffInWeeks(end))
}

// DiffInDays 相差多少天
func (c SuperTime) DiffInDays(end SuperTime) int64 {
	return c.DiffInSeconds(end) / SecondsPerDay
}

// DiffInDaysWithAbs 相差多少天(绝对值)
func (c SuperTime) DiffInDaysWithAbs(end SuperTime) int64 {
	return GetAbsValue(c.DiffInDays(end))
}

// DiffInHours 相差多少小时
func (c SuperTime) DiffInHours(end SuperTime) int64 {
	return c.DiffInSeconds(end) / SecondsPerHour
}

// DiffInHoursWithAbs 相差多少小时(绝对值)
func (c SuperTime) DiffInHoursWithAbs(end SuperTime) int64 {
	return GetAbsValue(c.DiffInHours(end))
}

// DiffInMinutes 相差多少分钟
func (c SuperTime) DiffInMinutes(end SuperTime) int64 {
	return c.DiffInSeconds(end) / SecondsPerMinute
}

// DiffInMinutesWithAbs 相差多少分钟(绝对值)
func (c SuperTime) DiffInMinutesWithAbs(end SuperTime) int64 {
	return GetAbsValue(c.DiffInMinutes(end))
}

// DiffInSeconds 相差多少秒
func (c SuperTime) DiffInSeconds(end SuperTime) int64 {
	return end.ToSecond() - c.ToSecond()
}

// DiffInSecondsWithAbs 相差多少秒(绝对值)
func (c SuperTime) DiffInSecondsWithAbs(end SuperTime) int64 {
	return GetAbsValue(c.DiffInSeconds(end))
}

// DiffInMillisecond 相差多少毫秒
func (c SuperTime) DiffInMillisecond(end SuperTime) int64 {
	return end.ToMillisecond() - c.ToMillisecond()
}

// DiffInMicrosecond 相关多少微秒
func (c SuperTime) DiffInMicrosecond(end SuperTime) int64 {
	return end.ToMicrosecond() - c.ToMicrosecond()
}

// DiffINanosecond 相关多少纳秒
func (c SuperTime) DiffINanosecond(end SuperTime) int64 {
	return end.ToNanosecond() - c.ToNanosecond()
}
