package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建时间
	now := time.Now()
	epoch := now.Unix()

	fmt.Println("Now: ", now)
	fmt.Println("Unix Time: ", epoch)

	// 指定的参考输出格式
	fmt.Println(now.Format("Mon, Jan 2, 2006 at 3:04pm"))
	fmt.Println(now.Format("06年-Jan月-2日 3:04pm"))

	// 年月日标准输出
	fmt.Println("Year: ", now.Year())
	fmt.Println("Month: ", now.Month())

	// 预定义时间输出格式
	fmt.Println(now.Format(time.RFC850))
	fmt.Println(now.Format(time.RFC1123))

	// 指定日期
	// 1:指定日期之前必须指定时区
	// 你可以使用LoadLocation构建时区，也可以使用time.UTC常量
	est, _ := time.LoadLocation("EST")
	july4 := time.Date(1776, 7, 4, 12, 15, 0, 0, est)

	// 另一种格式化输出日期的方式
	fmt.Println("July 4, 1776 was a ", july4.Format("Monday"))

	// 可以使用Before() After() Equal()来进行日期比较
	if july4.Before(now) {
		fmt.Println("July 4 is before Now ")
	}

	// 允许加减法的简便处理
	diff := now.Sub(july4)
	days := int(diff.Hours() / 24)
	fmt.Printf("July 4 was about %d days ago \n", days)

	twodaysDiff := time.Hour * 24 * 2
	twodays := now.Add(twodaysDiff)
	fmt.Println("Two Days: ", twodays.Format(time.ANSIC))

	// 指定参考输出格式
	input_form := "1/2/2006 3:04pm"
	user_str := "4/16/2014 11:38am"
	user_date, err := time.Parse(input_form, user_str)
	if err != nil {
		fmt.Println(">>> Error parsing date string")
	}
	fmt.Println("User Date: ", user_date.Format("Jan 2, 2006 @ 3:04pm"))

	fmt.Println("============================")
	// 常见time操作和函数
	CurTime := time.Now()
	fmt.Println(CurTime)
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(then)
	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Hour())
	fmt.Println(then.Minute())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Location())
	fmt.Println(then.Weekday())
	fmt.Println(then.Before(CurTime))
	fmt.Println(then.After(CurTime))
	fmt.Println(then.Equal(CurTime))
	diffentTime := CurTime.Sub(then)
	fmt.Println(diffentTime)
	fmt.Println(diffentTime.Hours())
	fmt.Println(diffentTime.Minutes())
	fmt.Println(diffentTime.Seconds())
	fmt.Println(diffentTime.Nanoseconds())
	fmt.Println(then.Add(diffentTime))
	fmt.Println(then.Add(-diffentTime))

	fmt.Println("============================")
	// 格式化Time输出
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	fmt.Println(t1)
	fmt.Println(t.Format("3:04PM"))
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	fmt.Println(t2)
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	fmt.Println(e)

	fmt.Println("============================")
	// 获取从Unix纪元开始的秒数，毫秒数，纳秒数
	now2 := time.Now()
	secs2 := now2.Unix()
	nanos2 := now2.UnixNano()
	fmt.Println(now2)
	millis2 := nanos2 / 1000000
	fmt.Println(secs2)
	fmt.Println(millis2)
	fmt.Println(nanos2)
	fmt.Println(time.Unix(secs2, 0))
	fmt.Println(time.Unix(0, nanos2))
}
