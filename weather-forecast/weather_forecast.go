// Package weather provides information about the area's weather at any given time.
package weather

// CurrentCondition describes the area's current weather condition.
var CurrentCondition string
// CurrentLocation describes the area's current weather condition.
var CurrentLocation string

// Forecast returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
