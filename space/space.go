/*
Given an age in seconds, calculate how old someone would be on:

Mercury: orbital period 0.2408467 Earth years
Venus: orbital period 0.61519726 Earth years
Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
Mars: orbital period 1.8808158 Earth years
Jupiter: orbital period 11.862615 Earth years
Saturn: orbital period 29.447498 Earth years
Uranus: orbital period 84.016846 Earth years
Neptune: orbital period 164.79132 Earth years
So if you were told someone were 1,000,000,000 seconds old, you should be able to say that they're 31.69 Earth-years old.
*/
// Package space given an age in seconds ouputs how old a planet is in Earth years
package main

import "fmt"

/* Iteration #1
func Age(seconds float64, planet string) float64 {
	var planetAge float64
	var planets = map[string]float64{
		"Mercury": 0.2408467,
		"Venus": 0.61519726,
		"Earth":  1.0,
		"Mars": 1.8808158,
		"Jupiter": 11.862615,
		"Saturn": 29.447498,
		"Uranus": 84.016846,
		"Neptune": 164.79132,
	}
	// planetConversion = (planets[planet])*(365.25/1)*(24/1)*(60/1)*(60/1)
	// planetAge = (seconds)*(1/60)*(1/60)*(1/24)(1/planetEarthYears)
	planetAge = (seconds)*(1/60)*(1/60)*(1/24)*(1/365.25)(planets[planet]/1)
	return planetAge
}

func main() {
	fmt.Println(Age(1000000000, "Earth"))
}
*/

// Iteration #2
// Solution idea from Kevin-DL https://exercism.io/tracks/go/exercises/space-age/solutions/6c85d4da8e574f32ae17f46e4ae35004
package space

// given the planet the person is on we can find their age
type Planet string

// The duration of one complete revolution of earth in seconds
const earthRevolution = 31557600

// Age calculates the age relative to a planet rotation cycle
func Age(ageInSeconds float64, planet Planet) float64 {
	// Calulate the associationg between the planets and one complete revolution of Earth in seconds
	Planets := map[Planet]float64{
		"Mercury": earthRevolution * 0.2408467,
		"Venus": earthRevolution * 0.61519726,
		"Earth":  earthRevolution,
		"Mars": earthRevolution * 1.8808158,
		"Jupiter": earthRevolution * 11.862615,
		"Saturn": earthRevolution * 29.447498,
		"Uranus": earthRevolution * 84.016846,
		"Neptune": earthRevolution * 164.79132,
	}
	return calculateAge(Planets[planet], ageInSeconds)
}

func calculateAge(planetRevolution float64, ageInSeconds float64) float64 {
	return ageInSeconds / planetRevolution
}