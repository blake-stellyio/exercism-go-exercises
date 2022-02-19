
              

package speed

type Car struct {
	battery int
	batteryDrain int
	speed int
	distance int
}       

func NewCar(speed, batteryDrain int) Car {
	return Car {
		battery: 100,
		batteryDrain: batteryDrain,
		speed: speed,
		distance: 0,      
	}   

}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	return Track{distance};    

}
    
func Drive(car Car) Car {
	if car.battery > car.batteryDrain {
		car.battery -= car.batteryDrain;
		car.distance += car.speed;
	}
	return car;
}

func CanFinish(car Car, track Track) bool {
	batteryUsage := (track.distance / car.speed) * car.batteryDrain;
	return batteryUsage <= car.battery;
}

        

          