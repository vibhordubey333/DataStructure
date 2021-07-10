package main

import (
	"strconv"
	"strings"
)

func main() {
	str := "12:00:00PM"
	timeConversion(str)

}
func timeConversion(s string) string {

	timeSlice := strings.Split(s, ":")
	hour, _ := strconv.Atoi(timeSlice[0]) //timeSlice[0]
	min := timeSlice[1]
	sec := timeSlice[2]
	dayNight := sec[3:4]
	sec = sec[:len(sec)-2]

	var exactTime string

	if dayNight != "AM" {
		if hour == 12 { //If hour provided is 12 then it should be retured as 00.
			hour = 12
		} else {
			hour = hour + 12
		}

	}
	exactTime = strconv.Itoa(hour)

	return exactTime + ":" + min + ":" + sec
}
