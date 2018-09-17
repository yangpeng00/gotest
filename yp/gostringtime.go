package main

import (
	"fmt"
	"time"
)

const (
	date        = "2006-01-02"
	shortdate   = "06-01-02"
	times       = "15:04:02"
	shorttime   = "15:04"
	datetime    = "2006-01-02 15:04:02"
	newdatetime = "2006/01/02 15~04~02"
	newtime     = "15~04~02"
)

func main() {
	str_time := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(time.Unix(time.Now().Unix(), 0))
	fmt.Println(time.Now().Unix())
	fmt.Println(str_time)

	thisdate := "2014-03-17 14:55:06"
	timeformatdate, _ := time.Parse(datetime, thisdate)
	fmt.Println(timeformatdate)
	convdate := timeformatdate.Format(date)
	convshortdate := timeformatdate.Format(shortdate)
	convtime := timeformatdate.Format(times)
	convshorttime := timeformatdate.Format(shorttime)
	convnewdatetime := timeformatdate.Format(newdatetime)
	convnewtime := timeformatdate.Format(newtime)
	fmt.Println(convdate)
	fmt.Println(convshortdate)
	fmt.Println(convtime)
	fmt.Println(convshorttime)
	fmt.Println(convnewdatetime)
	fmt.Println(convnewtime)

}
