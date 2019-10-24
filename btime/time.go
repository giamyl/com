package btime

import (
	"fmt"
	"time"
)

/**
下面功能是为了实现 json 序列化的时候指定时间格式
示例:
	t := time.Now()
	res := json.Marshal(t)
	// 这个时间的时间格式会是 2019-09-19T18:12:28.605295515+08:00 UTC 格式
	通过 JSONTime 包装下
	t := JSONTime(time.Now())
	res := json.Marshal(t)
	// 结果就会变成： 2019-09-19 18:12:28
	通过 JSONDate 包装下
	t := JSONTime(time.Now())
	res := json.Marshal(t)
	// 结果就会变成： 2019-09-19
*/

// json 序列化时候的时间格式
const (
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

// JSONTime 实现 json 序列化的时间格式
type JSONTime time.Time

// MarshalJSON 实现它的json序列化方法
func (j JSONTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("%q", time.Time(j).Format(TimeFormat))
	return []byte(stamp), nil
}

// UnmarshalJSON 实现 json 反序列化的时间格式
func (j *JSONTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*j = JSONTime(now)
	return
}

// JSONDate 实现 json 序列化的时间格式
type JSONDate time.Time

// MarshalJSON 实现它的json序列化方法
func (j JSONDate) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("%q", time.Time(j).Format(DateFormat))
	return []byte(stamp), nil
}

// UnmarshalJSON 实现 json 反序列化的时间格式
func (j *JSONDate) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+DateFormat+`"`, string(data), time.Local)
	*j = JSONDate(now)
	return
}
