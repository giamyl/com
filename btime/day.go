package btime

import "time"

// SubDays 计算 t1 比 t2 大多少天
func SubDays(t1, t2 time.Time) int {

	if t1.Location().String() != t2.Location().String() {
		return -1
	}
	hours := t1.Sub(t2).Hours()

	if hours <= 0 {
		return -1
	}
	// sub hours less than 24
	if hours < 24 {
		// may same day
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)

		if isSameDay {

			return 0
		}
		return 1

	}

	if (hours/24)-float64(int(hours/24)) == 0 { // just 24's times
		return int(hours / 24)
	}
	return int(hours/24) + 1

}
