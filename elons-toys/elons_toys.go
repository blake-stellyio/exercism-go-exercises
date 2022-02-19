package elon

import "fmt"

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.speed < c.battery {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	legs := trackDistance / c.speed
	expectedDrain := legs * c.batteryDrain
	if expectedDrain > c.battery {
		return false
	} else {
		return true
	}
}
