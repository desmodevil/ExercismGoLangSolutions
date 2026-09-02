// Package weather provides function to get weather forecast based on city and current weather condition.
package weather

var (
    // CurrentCondition represents a current weather condition in the place that forecast is planned to 
    // get for(e.g sunny, rainy, etc.).
	CurrentCondition string
    // CurrentLocation represents a city that user needs to get forecast for.
	CurrentLocation  string
)
// Forecast returns string with weather forecast for current city and its current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
