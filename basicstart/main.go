package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	dateTime := "2022-09-11T16:35:04+07:00"
	// layout := "2022-11-11T16:35:04+07:00"
	// location, _ := time.LoadLocation("Asia/Bangkok")
	// t, _ := time.ParseInLocation(layout, dateTime, location)
	// println(t.Month())

	t1, _ := time.Parse(time.RFC3339, dateTime)
	t11 := strings.Split(t1.String(), " ")
	fmt.Printf("%v %v\n", t11[0], t11[1])
	fmt.Printf("%v/%v/%v %v:%v:%v\n", t1.Year(), int(t1.Month()), t1.Day(), t1.Hour(), t1.Minute(), t1.Second())

	t2, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	fmt.Println(t2)

	// dateTime := "2010-01-23 11:44:20"
	// dt, _ := time.Parse("2006-01-2 15:04:05", dateTime)
	// dtstr2 := dt.Format("2006-01-2 15:04:05")
	// fmt.Println(dtstr2)
	// fmt.Println(time.Now().Format("2006/01/02 15:04:21"))
	// fmt.Println(time.Now().Format("yyyy/MM/ss HH:mm:ss"))

}
