package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) (car Car) {
	car = Car{
		speed: speed, batteryDrain: batteryDrain, battery: 100,
	}
	return
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	track := Track{distance}
	return track
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if battery_lvl := car.battery - car.batteryDrain; battery_lvl >= 0 {
		car.battery = battery_lvl
		car.distance = car.distance + car.speed
		return car
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	remainingDistance := track.distance
	for car.battery >= car.batteryDrain {
		car = Drive(car)
		remainingDistance -= car.speed
		if remainingDistance <= 0{
			return true
		}
	}
	return false
}
