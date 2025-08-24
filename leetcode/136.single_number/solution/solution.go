package solution

func singleNumber(nums []int) int {
    theSet := make(map[int]struct{})
    
    for _, i := range nums {
        if _, exist := theSet[i]; !exist {
            theSet[i] = struct{}{}
        } else {
            delete(theSet, i)
        }
    }
    
    for k := range theSet {
        return k
    }
    
    return -1
}
