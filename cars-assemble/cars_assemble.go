package cars

const productionRate = 1000
const successRate = 98.8
const costOftenCars = 95000
const costOfidvCars = 10000
// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(CalculateWorkingCarsPerHour(productionRate, successRate) / float64(60))
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
    var groupsOften int = carsCount / 10
    var remainders int = carsCount % 10
	return uint((groupsOften * costOftenCars) + (remainders * costOfidvCars))
}
