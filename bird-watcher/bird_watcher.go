package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    slice := birdsPerDay
    result := 0
 	for _, v := range slice {  
 	 result += v  
 	}  
 	return result 
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	targetDays := (week * 7) - 7
    endTargetdays := targetDays + 7
    slice := birdsPerDay
    var result int
 	for _, v := range slice[targetDays:endTargetdays] {
        result += v
    }
	return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 { 
		birdsPerDay[i] += 1;
    }
	return birdsPerDay
}
