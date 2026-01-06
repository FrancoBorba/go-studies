package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	carPerHour := successRate/100 * float64(productionRate)
	return  carPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsInHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(carsInHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	
	if carsCount < 10{
		return  uint(carsCount) * 10000
	}

	groupOfTen := carsCount / 10;
	rest := carsCount %10
	return  uint(groupOfTen) * 95000 + uint(rest) * 10000
}
