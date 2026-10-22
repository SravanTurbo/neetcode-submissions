/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

type event struct {
	time int
	isStart bool
}

func minMeetingRooms(intervals []Interval) int {
	events := []event{}

	for i:=0; i<len(intervals); i++ {
		events = append(events, event{intervals[i].start, true})
		events = append(events, event{intervals[i].end, false})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time < events[j].time {
			return true
		} else if events[i].time == events[j].time {
			return !events[i].isStart
		}
		return false
	})

	count := 0
	maxCount := 0
	
	for i:=0; i<len(events); i++ {
		if events[i].isStart {
			count++
		} else {
			count--
		}
		maxCount = max(maxCount, count)
	}

	return maxCount
}
