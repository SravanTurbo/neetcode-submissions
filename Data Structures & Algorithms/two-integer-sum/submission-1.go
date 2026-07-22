func twoSum(nums []int, target int) []int {
    keyIndexMap := make(map[int]int)
    i := 0
    var idx int
    var exists bool
    for i<len(nums) {
        idx, exists = keyIndexMap[target-nums[i]]
        if exists {
            break
        } else {
            keyIndexMap[nums[i]] = i
            i++
        }
    }

    return []int{idx, i}
}
