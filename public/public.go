package public

import (
	"fmt"
	"time"
)

/***********************************************************************
 *	判斷值是否在Array裡
 ***********************************************************************/
func Contains(v int, a []int) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}

func Println(msg string) {
	t := time.Now()

	now := fmt.Sprintf("[%d-%02d-%02d %02d:%02d:%02d]", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Printf("%s %s \n", now, msg)
}

/***********************************************************************
 *	取得立春的陽曆
 *
 *	立春节气一般是从2月4日或5日开始，到2月19或20日结束。有时在农历的腊月，有时在农历的正月。
 *	计算公式：[Y*D+C]-L
 *	公式解读：年数的后 2 位乘 0.2422 加 3.87 取整数减闰年数。21世纪 C 值 = 3.87，22 世纪 C 值 = 4.15 。
 *	举例说明：2058年立春日期的计算步骤[58×0.2422+3.87]-[(58-1)/4]=17-14=3，则2月3日立春。
 *
 *	2017-02-03
 *	2018-02-04
 *	2019-02-04
 *	2020-02-04
 *	2021-02-03
 *	2022-02-04
 *	2023-02-04
 *	2024-02-04
 *	2025-02-03
 *	2026-02-04
 *	2027-02-04
 *	2028-02-04
 *	2029-02-03
 *	2030-02-04
 ***********************************************************************/

func GetDateForSpringBeginning() bool {
	t := time.Now()

	Year := t.Year()
	// now := fmt.Sprintf("[%d-%02d-%02d %02d:%02d:%02d]", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// fmt.Printf("%s %s \n", now, msg)

	// date := 0.0
	GDFSB_C := 0.0

	if Year > 2100 {
		GDFSB_C = 4.15
	} else {
		GDFSB_C = 3.87
	}

	GDFSB_N := 0.2422

	YearF := float64(Year % 2000)
	date := int64(YearF*float64(GDFSB_N)+float64(GDFSB_C)) - int64((YearF-1)/4)
	loc, _ := time.LoadLocation("PRC")

	t1, _ := time.ParseInLocation("2006-01-02", fmt.Sprintf("20%d-02-0%d", int(YearF), date), loc)
	// fmt.Println("date  t1 ", t1)

	// fmt.Println("GetDateForSpringBeginning   ", t1.Unix())
	// fmt.Println("GetDateForSpringBeginning   ", time.Now().Unix())
	return time.Now().Unix() < t1.Unix()

	//"20" + fmt.Sprintf("%f", YearF) + "-02-0" + fmt.Sprintf("%f", YearF)
}
