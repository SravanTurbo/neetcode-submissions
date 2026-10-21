/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

/*
start: 12:28
idea: intervals: sort by startTime 
t:O() s:O()
end:
*/

func canAttendMeetings(intervals []Interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	prevEnd := -1
	for i:=0; i<len(intervals); i++ {
		if intervals[i].start < prevEnd {
			return false
		}
		prevEnd = intervals[i].end
	}

	return true
}
