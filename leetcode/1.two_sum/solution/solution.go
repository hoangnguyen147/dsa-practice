package solution

func twoSum(nums []int, target int) []int {
    theMap := make(map[int]int)
    for i, v := range nums {
        want := target - v
        if iw, ok := theMap[want]; ok {
            return []int{iw, i}
        }
        if _, ok := theMap[v]; !ok {
            theMap[v] = i
        }
    }
    
    return []int{}
}
