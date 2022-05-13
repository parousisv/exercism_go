package cars
import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * float64((successRate / 100))
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
    dec := carsCount / 10
    mon := carsCount % 10
    fmt.Println(dec)
    fmt.Println(mon)
    return uint(dec)*95000+uint(mon)*10000
	panic("CalculateCost not implemented")
}
