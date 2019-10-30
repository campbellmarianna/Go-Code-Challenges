package space

import "fmt"

//
func Age(seconds float32, planet string) float32 {
	var planetAge float32
	var planetConversion float64
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
	planetConversion = (planets[planet])*(365.25/1)*(24/1)*(60/1)*(60/1)
	planetAge = (seconds)*(1/60)*(1/60)*(1/24)(1/planetEarthYears)
	return planetAge
}

func main() {
	fmt.Println(Age())
}