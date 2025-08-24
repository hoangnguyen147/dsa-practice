package solution

func getNext(n int) int {
    next := 0
    for n > 0 {
        d := n % 10
        next += d*d
        n /= 10
    }
    
    return next
}

func isHappy(n int) bool {
    theSet := make(map[int]struct{})
    theSet[n] = struct{}{}
    for n > 1 {
        n = getNext(n)
        if _, exist := theSet[n]; exist {
            return false
        }
        theSet[n] = struct{}{}
    }
    
    return true
}
