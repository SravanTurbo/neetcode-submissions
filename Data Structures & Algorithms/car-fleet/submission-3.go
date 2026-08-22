/*
Time: O(nlogn) for sorting
Space: O(n) for map
*/

func carFleet(target int, position []int, speed []int) int {
	eta := make(map[int]float64)

	for i:=0; i<len(position); i++ {
		currETA := float64(target-position[i])/float64(speed[i])
		eta[position[i]] = roundFloat(currETA, 2)
	}

	sort.Ints(position)

	fleet := 0
	prevFleetETA := 0.0
	for i:=len(position)-1; i>=0; i-- {
		currFleetETA := eta[position[i]]

		if currFleetETA > prevFleetETA {
			fleet += 1
		}

		prevFleetETA = max(prevFleetETA, currFleetETA)
	}

	return fleet
}

func roundFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val * ratio) / ratio
}
