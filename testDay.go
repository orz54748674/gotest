package main

import (
	"fmt"
	"time"
)

//Sunday
//Monday
//Tuesday
//Wednesday
//Thursday
//Friday
//Saturday
func main() {
	fmt.Println(IsThisWeek(1573887670))
}

func IsThisWeek(tm int64) bool {
	tn := time.Now()
	var dayStart, dayEnd int64
	//计算周一5点的时间戳: dayStart
	wd := int(tn.Weekday()) //周日为0，周一为1...周六为6
	day := tn.Day()
	if wd != 1 {
		if wd == 0 {
			wd = 7
		}
		day -= wd - 1
	}
	dayStart = time.Date(tn.Year(), tn.Month(), day, 5, 0, 0, 0, tn.Location()).Unix()
	dayEnd = dayStart + 7*24*3600 - 1
	return tm >= dayStart && tm <= dayEnd
}
