func longestConsecutive(nums []int) int {
    M := make(map[int]int)
    maxi := 0

    max := func(a, b int) int {
        if a > b {
            return a
        } else {
            return b
        }
    }

    for i := range nums {
        curr := nums[i]
        prev := nums[i]-1
        next := nums[i]+1

        if _, currExists := M[curr]; currExists {
            continue
        }

        prevVal, prevExists := M[prev]
        nextVal, nextExists := M[next]
        newVal := 1

        if prevExists && nextExists {
            newVal += prevVal + nextVal
            M[curr - prevVal] = newVal
            M[curr + nextVal] = newVal
        }else if prevExists{
            newVal += prevVal
            M[curr - prevVal] = newVal
        }else if nextExists{
            newVal += nextVal
            M[curr + nextVal] = newVal
        }

        M[curr] = newVal
        maxi = max(maxi, newVal)
    }
    
    return maxi
}
