package exTime

import "time"

func (c *SuperTime) SetTimezone(timezone string) error {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}

	c.Time = c.Time.In(loc)
	return nil
}

// SetYear 设置年
func (c SuperTime) SetYear(year int) SuperTime {
	c.Time = time.Date(year, c.Time.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// SetMonth 设置月
func (c SuperTime) SetMonth(month int) SuperTime {
	c.Time = time.Date(c.Year(), time.Month(month), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// SetDay 设置日
func (c SuperTime) SetDay(day int) SuperTime {
	c.Time = time.Date(c.Year(), c.Time.Month(), day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// SetHour 设置时
func (c SuperTime) SetHour(hour int) SuperTime {
	c.Time = time.Date(c.Year(), c.Time.Month(), c.Day(), hour, c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// SetMinute 设置分
func (c SuperTime) SetMinute(minute int) SuperTime {
	c.Time = time.Date(c.Year(), c.Time.Month(), c.Day(), c.Hour(), minute, c.Second(), c.Nanosecond(), c.Location())
	return c
}

// SetSecond 设置秒
func (c SuperTime) SetSecond(second int) SuperTime {
	c.Time = time.Date(c.Year(), c.Time.Month(), c.Day(), c.Hour(), c.Minute(), second, c.Nanosecond(), c.Location())
	return c
}

// SetNanoSecond 设置纳秒
func (c SuperTime) SetNanoSecond(nanoSecond int) SuperTime {
	c.Time = time.Date(c.Year(), c.Time.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), nanoSecond, c.Location())
	return c
}
