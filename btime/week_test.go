package btime

import (
	"fmt"
	"testing"
	"time"
)

func TestWeek(t *testing.T) {
	// 获取当前时间周一的时间
	fmt.Println(GetDateOfMonday().ToTime().Format("2006-01-02"))
	fmt.Println(GetDateOfWeekMonday(0).ToDateTimeString())
	// 获取上周周一的时间
	fmt.Println(GetDateOfWeekMonday(-1).ToDateString())
	// 获取下周周一的时间
	fmt.Println(GetDateOfWeekMonday(1).ToDateString())

	// 获取当前时间周日的时间
	fmt.Println(GetDateOfSunday().ToDateString())
	// 获取指定时间周日的时间
	tm := time.Date(2019, 7, 14, 0, 0, 0, 0, time.Local)
	fmt.Println(GetDateOfSunday(tm).ToDateString())
	tm = time.Date(2019, 11, 11, 0, 0, 0, 0, time.Local)
	fmt.Println(GetDateOfSunday(tm).ToDateString())

	// 获取上周周日
	fmt.Println(GetDateOfWeekSunday(-1).ToDateString())
	// 下周周日
	fmt.Println(GetDateOfWeekSunday(1).ToDateString())
	// 本周周日
	fmt.Println(GetDateOfWeekSunday(0).ToDateString())
}
