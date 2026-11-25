// Package weather provides tools to forecast the weather in a city.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string
	// CurrentLocation represents the city for the current forecast.
	CurrentLocation string
)

// Forecast updates the current location and condition, and returns a
// string describing the forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
