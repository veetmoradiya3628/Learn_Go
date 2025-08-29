// Package weather represents a weather status and forecase information.
package weather

// CurrentCondition represents a current weather conditions.
var CurrentCondition string
// CurrentLocation represents a current location.
var CurrentLocation string

// Forecast returns an string value that shows current location and current conditions in reading format.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
