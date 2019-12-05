// Package main - given a time - provides the time duration after a gigasecond 
package main

// import path for the time package from the standard library
import (
  "time"
  "fmt"
)

// AddGigasecond returns a given time plus 10 billion seconds
func AddGigasecond(t time.Time) time.Time {

  t = t.Add(time.Second * 1000000000)

  return t
}

func main() {
  // Parse a time value from a string in the standard Unix format.
  t, err := time.Parse(time.UnixDate, "2015-01-24T22:00:00")
  if err != nil { // Always check errors even if they should not happen.
    panic(err)
  }

  // time.Time's Stringer method is useful without any format.
  // fmt.Println("default format:", t)
  fmt.Println(AddGigasecond(t))
}

// Initial Pseudocode
	// turn time into seconds 
	// add seconds to a billion
	// format seconds into format as input

	// if given year month day
	// return given year month day and hour mintue second