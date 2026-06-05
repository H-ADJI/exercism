package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (temp TemperatureUnit) String() string {
	return [2]string{"°C", "°F"}[temp]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (temp Temperature) String() string {

	return fmt.Sprintf("%v %v", temp.degree, temp.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (speed SpeedUnit) String() string {
	return [2]string{"km/h", "mph"}[speed]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (speed Speed) String() string {

	return fmt.Sprintf("%v %v", speed.magnitude, speed.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (meteo MeteorologyData) String() string {

	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", meteo.location, meteo.temperature, meteo.windDirection, meteo.windSpeed, meteo.humidity)
}
