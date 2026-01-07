package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {

	var birdCount int

	for _ , value := range birdsPerDay {
		birdCount = birdCount + value
	}
	return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	start := (week - 1) * 7
	end := start + 7

	var birdsInWeek int
	for i := start; i < end; i++ {
		birdsInWeek += birdsPerDay[i]
	}

	return birdsInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	
	for  i := 0 ; i < len(birdsPerDay) ; i++ {

		if(i % 2 == 0) {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}
	return  birdsPerDay
}
