package main

import (
	"fmt"

	"github.com/wuxiaoxiaoshen/holidays"
)

func main() {

	h := holidays.FetchByYear(2021)
	fmt.Println(h)

	h2 := holidays.FetchMonthHolidayCount(2021, 5)
	fmt.Println(h2)

	h3 := holidays.IsHoliday("2021-05-08")
	fmt.Println(h3)
}
