package solution

func intersect(nums1 []int, nums2 []int) []int {
    freqMap := make(map[int]int)
    for _, n := range nums1 {
        if _, ok := freqMap[n]; !ok {
            freqMap[n] = 1
        } else {
            freqMap[n]++
        }
    }
    res := []int{}
    for _, n := range nums2 {
        if _, ok := freqMap[n]; ok {
            res = append(res, n)
            if freqMap[n] == 1 {
                delete(freqMap, n)
            } else {
                freqMap[n] --
            }
        }
    }
    
    return res
}
