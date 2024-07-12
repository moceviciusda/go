// Package weather provides tools to forecast current
// weather conditions of various cities in Goblinocus.
package weather

// CurrentCondition represents a certain weather condition.
var CurrentCondition string
// CurrentLocation represents a certain city. 
var CurrentLocation string

// Forecast returns a string representing current weather conditions in a certain city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
