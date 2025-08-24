package solution

func containsDuplicate(nums []int) bool {
    theSet := make(map[int]struct{})
    for _, i := range nums {
        if _, exist := theSet[i]; exist {
            return true
        }
        theSet[i] = struct{}{}
    }
    return false
}
