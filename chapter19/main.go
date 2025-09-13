package main

import (
	"fmt"
	"time"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func PrintTime(label string, t *time.Time) {
	Printfln("%s: Day: %v (%v): Month: %v Year: %v", label, t.Day(), t.Weekday(), t.Month(), t.Year())
}

func main() {
	current := time.Now()
	specific := time.Date(1976, time.April, 13, 7, 30, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)

	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("Unix", &unix)
}
