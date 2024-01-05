package main

import "time"

// 计算两个日期之间差多少天
func DiffDays(t1, t2 time.Time) int {
	if t1.IsZero() {
		t1 = time.Now()
	}
	
	t1Unix, _ := t1.Unix()
	t2Unix, _ := t2.Unix()

	diff := t2Unix - t1Unix

	return int(diff / (24 * 3600))
}
// 推算几天后的日期


func main() {

}