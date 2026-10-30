/*
start: 4:41pm
idea: binary search
t: O(n log m) s: O(1) where n is the len of piles and m is max of piles because we might find a speed that is any number of bananas from 1 to max of piles
end: 5:10
*/

func minEatingSpeed(piles []int, h int) int {
	if len(piles) == 0 {
		return -1
	}

	sort.Ints(piles)
	minSpeed := piles[len(piles)-1]

	var validEatingSpeed func(speed, limit int) bool
	validEatingSpeed = func(speed, limit int) bool {
		hours := 0
		for i:=0; i<len(piles); i++ {
			bananas := piles[i]
			if bananas <= speed {
				hours += 1 
			} else {
				hours += bananas/speed
				if bananas % speed > 0 {
					hours += 1
				}
			}

			if hours > limit {
				return false
			}
		}

		return true
	} 
	
	l, r := 1, piles[len(piles)-1]
	for l <= r {
		mid := (l+r)/2
		if validEatingSpeed(mid, h) {
			minSpeed = min(minSpeed, mid)
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return minSpeed
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
