// Package weather : this package is for weather forcasting.
package weather

// CurrentCondition :weather conditions.
var CurrentCondition string

// CurrentLocation : location where we wanna forcast the weather.
var CurrentLocation string

// Forecast give the weather conditions of a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
