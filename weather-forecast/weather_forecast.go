// Package weather is the package that containts all the info about the weather.
package weather

// CurrentCondition is a string var that contains the current condition of the weather.
var CurrentCondition string
// CurrentLocation is a string var that contains the current location.
var CurrentLocation string 

/*Forecast function returns the location given and its current weather condition.*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
