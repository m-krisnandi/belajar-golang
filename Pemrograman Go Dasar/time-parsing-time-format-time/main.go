package main

import (
	"fmt"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)

	// method milik time.Time
	var now = time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month())

	// parsing dari string ke time.Time
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2019-12-24 10:20:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2019 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String()) 

	// predefined layout format untuk parsing time
	var date2, _ = time.Parse(time.RFC822, "02 Sep 19 15:04 WIB")
	fmt.Println(date2.String())

	// format time.Time ke string
	var dateS1 = date2.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1:", dateS1)

	var dateS2 = date2.Format(time.RFC3339)
	fmt.Println("dateS2:", dateS2)

	// handling error parsing
	var date3, err = time.Parse("06 Jan 15", "02 Sep 19 08:00 WIB") // error
	// var date3, err = time.Parse("02 Jan 15 03:04 MST", "02 Sep 19 08:00 WIB") // ok

	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(date3)
}