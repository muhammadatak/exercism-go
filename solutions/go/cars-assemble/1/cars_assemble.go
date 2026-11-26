package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	result := float64(productionRate) * (successRate / 100)
	return result
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	result := CalculateWorkingCarsPerHour(productionRate, successRate) / 60
	return int(result)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	cost := (carsCount % 10) * 10000
	if carsCount > 9 {
		front := carsCount / 10
		cost += front * 95000
	}
	return uint(cost)
}
