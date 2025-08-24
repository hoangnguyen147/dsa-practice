package solution

func intersection(nums1 []int, nums2 []int) []int {
    theSet := make(map[int]struct{})
    result := []int{}
    for _, i := range nums1 {
        theSet[i] = struct{}{}
    }
    
    for _, i := range nums2 {
        if _, ok := theSet[i]; ok {
            result = append(result, i)
            delete(theSet, i)
        }
    }
    
    return result
}
