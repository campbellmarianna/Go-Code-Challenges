package main

// import path for the time package from the standard library
import (
	"time"
	"fmt"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	// turn time into seconds 
	// add seconds to a billion
	// format seconds into format as input

	// if given year month day
	// return given year month day and hour mintue second
	return t
}

func main(){
	var time = "2011-04-25"
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println(AddGigasecond(start))
}