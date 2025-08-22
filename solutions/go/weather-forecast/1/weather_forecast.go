// Package weather an good package.
package weather

// CurrentCondition an good variable.
var CurrentCondition string
// CurrentLocation other good variable.
var CurrentLocation string

// Forecast an good function.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
