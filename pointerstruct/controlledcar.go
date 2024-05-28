package pointerstruct

import "fmt"

type Car struct {
	speed   int
	battery int
}

func NewCar(speed, battery int) *Car {
	return &Car{speed: speed, battery: battery}
}
func GetSpeed(car *Car) int {
	return car.speed
}
func GetBattery(car *Car) int {
	return car.battery
}
func ChargeCar(car *Car, minutes int) {
	car.battery += minutes / 2
	if car.battery > 100 {
		car.battery = 100
	}
}
func TryFinish(car *Car, distance int) string {
	batteryUsage := distance / 2
	if batteryUsage > car.battery {
		car.battery = 0
		return ""
	}
	car.battery = car.battery - batteryUsage
	time := float32(distance) / float32(car.speed)
	return fmt.Sprintf("%.2f", time)
}
