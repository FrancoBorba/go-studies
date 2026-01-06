// Package weather allows you to see the weather forecast in Goblinoculos.
package weather


var (
	// CurrentCondition represents the current temperature.
	CurrentCondition string
	// CurrentLocation represents the currente Location.
	CurrentLocation  string
)

//Forecast returns a string  reporting the temperature of a specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
