package solution

func containsNearbyDuplicate(nums []int, k int) bool {
    if len(nums) <= 1 {
        return false
    }
    existMap := make(map[int]int)
    for i, v := range nums {
        if prev, ok := existMap[v]; !ok {
            existMap[v] = i
        } else {
            if i - prev <= k {
                return true
            }
            existMap[v] = i
        }
    }
    return false
}
