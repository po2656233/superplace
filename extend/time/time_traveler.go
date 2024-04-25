package exTime

import (
	"time"
)

// AddDuration 按照持续时间字符串增加时间
// 支持整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
func (c *SuperTime) AddDuration(duration string) error {
	td, err := ParseByDuration(duration)
	if err != nil {
		return err
	}

	c.Time = c.Add(td)
	return nil
}

// SubDuration 按照持续时间字符串减少时间
// 支持整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
func (c *SuperTime) SubDuration(duration string) error {
	return c.AddDuration("-" + duration)
}

// AddCenturies N世纪后
func (c *SuperTime) AddCenturies(centuries int) {
	c.AddYears(YearsPerCentury * centuries)
}

// AddCenturiesNoOverflow N世纪后(月份不溢出)
func (c *SuperTime) AddCenturiesNoOverflow(centuries int) {
	c.AddYearsNoOverflow(centuries * YearsPerCentury)
}

// AddCentury 1世纪后
func (c *SuperTime) AddCentury() {
	c.AddCenturies(1)
}

// AddCenturyNoOverflow 1世纪后(月份不溢出)
func (c *SuperTime) AddCenturyNoOverflow() {
	c.AddCenturiesNoOverflow(1)
}

// SubCenturies N世纪前
func (c *SuperTime) SubCenturies(centuries int) {
	c.SubYears(YearsPerCentury * centuries)
}

// SubCenturiesNoOverflow N世纪前(月份不溢出)
func (c *SuperTime) SubCenturiesNoOverflow(centuries int) {
	c.SubYearsNoOverflow(centuries * YearsPerCentury)
}

// SubCentury 1世纪前
func (c *SuperTime) SubCentury() {
	c.SubCenturies(1)
}

// SubCenturyNoOverflow 1世纪前(月份不溢出)
func (c *SuperTime) SubCenturyNoOverflow() {
	c.SubCenturiesNoOverflow(1)
}

// AddYears N年后
func (c *SuperTime) AddYears(years int) {
	c.Time = c.Time.AddDate(years, 0, 0)
}

// AddYearsNoOverflow N年后(月份不溢出)
func (c *SuperTime) AddYearsNoOverflow(years int) {
	// 获取N年后本月的最后一天
	last := time.Date(c.Year()+years, c.Time.Month(), 1, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location()).AddDate(0, 1, -1)

	day := c.Day()
	if c.Day() > last.Day() {
		day = last.Day()
	}

	c.Time = time.Date(last.Year(), last.Month(), day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// AddYear 1年后
func (c *SuperTime) AddYear() {
	c.AddYears(1)
}

// AddYearNoOverflow 1年后(月份不溢出)
func (c *SuperTime) AddYearNoOverflow() {
	c.AddYearsNoOverflow(1)
}

// SubYears N年前
func (c *SuperTime) SubYears(years int) {
	c.AddYears(-years)
}

// SubYearsNoOverflow N年前(月份不溢出)
func (c *SuperTime) SubYearsNoOverflow(years int) {
	c.AddYearsNoOverflow(-years)
}

// SubYear 1年前
func (c *SuperTime) SubYear() {
	c.SubYears(1)
}

// SubYearNoOverflow 1年前(月份不溢出)
func (c *SuperTime) SubYearNoOverflow() {
	c.SubYearsNoOverflow(1)
}

// AddQuarters N季度后
func (c *SuperTime) AddQuarters(quarters int) {
	c.AddMonths(quarters * MonthsPerQuarter)
}

// AddQuartersNoOverflow N季度后(月份不溢出)
func (c *SuperTime) AddQuartersNoOverflow(quarters int) {
	c.AddMonthsNoOverflow(quarters * MonthsPerQuarter)
}

// AddQuarter 1季度后
func (c *SuperTime) AddQuarter() {
	c.AddQuarters(1)
}

// AddQuarterNoOverflow 1季度后(月份不溢出)
func (c *SuperTime) AddQuarterNoOverflow() {
	c.AddQuartersNoOverflow(1)
}

// SubQuarters N季度前
func (c *SuperTime) SubQuarters(quarters int) {
	c.AddQuarters(-quarters)
}

// SubQuartersNoOverflow N季度前(月份不溢出)
func (c *SuperTime) SubQuartersNoOverflow(quarters int) {
	c.AddMonthsNoOverflow(-quarters * MonthsPerQuarter)
}

// SubQuarter 1季度前
func (c *SuperTime) SubQuarter() {
	c.SubQuarters(1)
}

// SubQuarterNoOverflow 1季度前(月份不溢出)
func (c *SuperTime) SubQuarterNoOverflow() {
	c.SubQuartersNoOverflow(1)
}

// AddMonths N月后
func (c *SuperTime) AddMonths(months int) {
	c.Time = c.Time.AddDate(0, months, 0)
}

// AddMonthsNoOverflow N月后(月份不溢出)
func (c *SuperTime) AddMonthsNoOverflow(months int) {
	month := c.Time.Month() + time.Month(months)

	// 获取N月后的最后一天
	last := time.Date(c.Year(), month, 1, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location()).AddDate(0, 1, -1)

	day := c.Day()
	if c.Day() > last.Day() {
		day = last.Day()
	}

	c.Time = time.Date(last.Year(), last.Month(), day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

// AddMonth 1月后
func (c *SuperTime) AddMonth() {
	c.AddMonths(1)
}

// AddMonthNoOverflow 1月后(月份不溢出)
func (c *SuperTime) AddMonthNoOverflow() {
	c.AddMonthsNoOverflow(1)
}

// SubMonths N月前
func (c *SuperTime) SubMonths(months int) {
	c.AddMonths(-months)
}

// SubMonthsNoOverflow N月前(月份不溢出)
func (c *SuperTime) SubMonthsNoOverflow(months int) {
	c.AddMonthsNoOverflow(-months)
}

// SubMonth 1月前
func (c *SuperTime) SubMonth() {
	c.SubMonths(1)
}

// SubMonthNoOverflow 1月前(月份不溢出)
func (c *SuperTime) SubMonthNoOverflow() {
	c.SubMonthsNoOverflow(1)
}

// AddWeeks N周后
func (c *SuperTime) AddWeeks(weeks int) {
	c.AddDays(weeks * DaysPerWeek)
}

// AddWeek 1天后
func (c *SuperTime) AddWeek() {
	c.AddWeeks(1)
}

// SubWeeks N周后
func (c *SuperTime) SubWeeks(weeks int) {
	c.SubDays(weeks * DaysPerWeek)
}

// SubWeek 1天后
func (c *SuperTime) SubWeek() {
	c.SubWeeks(1)
}

// AddDays N天后
func (c *SuperTime) AddDays(days int) {
	c.Time = c.Time.AddDate(0, 0, days)
}

// AddDay 1天后
func (c *SuperTime) AddDay() {
	c.AddDays(1)
}

// SubDays N天前
func (c *SuperTime) SubDays(days int) {
	c.AddDays(-days)
}

// SubDay 1天前
func (c *SuperTime) SubDay() {
	c.SubDays(1)
}

// AddHours N小时后
func (c *SuperTime) AddHours(hours int) {
	td := time.Duration(hours) * time.Hour
	c.Time = c.Time.Add(td)
}

// AddHour 1小时后
func (c *SuperTime) AddHour() {
	c.AddHours(1)
}

// SubHours N小时前
func (c *SuperTime) SubHours(hours int) {
	c.AddHours(-hours)
}

// SubHour 1小时前
func (c *SuperTime) SubHour() {
	c.SubHours(1)
}

// AddMinutes N分钟后
func (c *SuperTime) AddMinutes(minutes int) {
	td := time.Duration(minutes) * time.Minute
	c.Time = c.Time.Add(td)
}

// AddMinute 1分钟后
func (c *SuperTime) AddMinute() {
	c.AddMinutes(1)
}

// SubMinutes N分钟前
func (c *SuperTime) SubMinutes(minutes int) {
	c.AddMinutes(-minutes)
}

// SubMinute 1分钟前
func (c *SuperTime) SubMinute() {
	c.SubMinutes(1)
}

// AddSeconds N秒钟后
func (c *SuperTime) AddSeconds(seconds int) {
	td := time.Duration(seconds) * time.Second
	c.Time = c.Time.Add(td)
}

// AddSecond 1秒钟后
func (c *SuperTime) AddSecond() {
	c.AddSeconds(1)
}

// SubSeconds N秒钟前
func (c *SuperTime) SubSeconds(seconds int) {
	c.AddSeconds(-seconds)
}

// SubSecond 1秒钟前
func (c *SuperTime) SubSecond() {
	c.SubSeconds(1)
}
