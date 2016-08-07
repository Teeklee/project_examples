package main

import (
	"regexp"
	"time"
	"net/url"
)

const c1 = "12345"
const c2 = "1.2.3.4.5" 	//Month.Day.Hour.Minute.Second
const c3 = "02012006"	//DayMonthYear

func main() {
	regexp.Compile(`(?i)(\b[^aeiou][a-z]*(\s)*)`)
	regexp.Compile(`(?i)\b[^aeiou][a-z]*(\s)*)`)
	
	time.Parse("12345", "")
	time.Parse(c1, "")
	time.Parse(c2, "")
	time.Parse(c3, "")
	
	url.Parse("https://max:muster@www.example.com:8080/index.html?p1=A&p2=B#ressource")
	url.Parse("https://golang org")
}