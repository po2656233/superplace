package exTime

import (
	"time"

	cstring "github.com/po2656233/superplace/extend/string"
)

// ToSecond 输出秒级时间戳
func (c SuperTime) ToSecond() int64 {
	return c.Unix()
}

// ToMillisecond 输出毫秒级时间戳
func (c SuperTime) ToMillisecond() int64 {
	return c.Time.UnixNano() / int64(time.Millisecond)
}

func (c SuperTime) ToMillisecondString() string {
	t := c.ToMillisecond()
	return cstring.ToString(t)
}

// ToMicrosecond 输出微秒级时间戳
func (c SuperTime) ToMicrosecond() int64 {
	return c.UnixNano() / int64(time.Microsecond)
}

// ToNanosecond 输出纳秒级时间戳
func (c SuperTime) ToNanosecond() int64 {
	return c.UnixNano()
}

// ToDateMillisecondFormat 2023-04-10 12:26:57.420
func (c SuperTime) ToDateMillisecondFormat() string {
	return c.Format(DateTimeMillisecondFormat)
}

// ToDateTimeFormat 2006-01-02 15:04:05
func (c SuperTime) ToDateTimeFormat() string {
	return c.Format(DateTimeFormat)
}

// ToDateFormat 2006-01-02
func (c SuperTime) ToDateFormat() string {
	return c.Format(DateFormat)
}

// ToTimeFormat 15:04:05
func (c SuperTime) ToTimeFormat() string {
	return c.Format(TimeFormat)
}

// ToShortDateTimeFormat 20060102150405
func (c SuperTime) ToShortDateTimeFormat() string {
	return c.Format(ShortDateTimeFormat)
}

// ToShortDateFormat 20060102
func (c SuperTime) ToShortDateFormat() string {
	return c.Format(ShortDateFormat)
}

// ToShortIntDateFormat 20060102
func (c SuperTime) ToShortIntDateFormat() int32 {
	strDate := c.ToShortDateFormat()
	intDate, _ := cstring.ToInt32(strDate)
	return intDate
}

// ToShortTimeFormat 150405
func (c SuperTime) ToShortTimeFormat() string {
	return c.Format(ShortTimeFormat)
}
