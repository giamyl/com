package btime

import (
	"time"
)

// MyTime 时间的别名， 方便转化
type MyTime time.Time

// ToTime 转化成 go 的时间格式
func (mt MyTime) ToTime() time.Time {
	return time.Time(mt)
}

// ToDateString 转化成字符串格式的时间： 2006-01-02
func (mt MyTime) ToDateString() string {
	return mt.ToTime().Format("2006-01-02")
}

// ToDateTimeString 转化成字符串格式的时间： 2006-01-02 15:03:04
func (mt MyTime) ToDateTimeString() string {
	return mt.ToTime().Format("2006-01-02 15:04:05")
}

// GetDateOfMonday 获取指定时间对应的周一的时间
// 不传默认是当前时间
func GetDateOfMonday(ts ...time.Time) MyTime {
	now := time.Now()
	if len(ts) > 0 {
		now = ts[0]
	}

	// 因为日期是 0,1……6 周日到周一， 所以计算偏移量的时候如果正数，
	// 证明当前时间是周日， 所以向前推 6 天（只有周日会是大于 0的）
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset -= 6
	}

	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).
		AddDate(0, 0, offset)

	return MyTime(t)
}

// GetDateOfWeekMonday 获取指定周周一的时间
// 如当前时间所在周 传 0， 上周 -1， 下周 1
func GetDateOfWeekMonday(week int) MyTime {
	thisDate := GetDateOfMonday().ToTime()
	return MyTime(thisDate.AddDate(0, 0, week*7))
}

// GetDateOfSunday 获取指定时间对应的周日的时间
func GetDateOfSunday(ts ...time.Time) MyTime {
	now := time.Now()
	if len(ts) > 0 {
		now = ts[0]
	}

	offset := int(time.Sunday - now.Weekday())
	if offset < 0 {
		offset += 7
	}

	t := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).
		AddDate(0, 0, offset)

	return MyTime(t)
}

// GetDateOfWeekSunday 获取指定周周日的时间
// 如当前时间所在周 传 0， 上周 -1， 下周 1
func GetDateOfWeekSunday(week int) MyTime {
	thisDate := GetDateOfSunday().ToTime()
	return MyTime(thisDate.AddDate(0, 0, week*7))
}
