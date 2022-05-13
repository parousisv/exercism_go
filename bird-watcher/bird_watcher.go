package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var sum int
    for i := 0; i < len(birdsPerDay); i++{
        sum += birdsPerDay[i]
    }
	return sum
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    var sum int
    if len(birdsPerDay) >= 7*week {
        for i := (7*week) - 7; i < 7*week; i++{
			sum += birdsPerDay[i]
        }
    }
	return sum
	panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i = i + 2{
        birdsPerDay[i] = birdsPerDay[i] + 1
    }
	return birdsPerDay
	panic("Please implement the FixBirdCountLog() function")
}
