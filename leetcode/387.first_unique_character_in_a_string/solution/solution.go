package solution

func firstUniqChar(s string) int {
    if len(s) == 0 {
        return -1
    }
    if len(s) == 1 {
        return 0
    }
    freqMap := make(map[rune]int)
    for _, v := range s {
        if _, ok := freqMap[v]; !ok {
            freqMap[v] = 1
        } else {
            freqMap[v]++
        }
    }
    for i, v := range s {
        if _, ok := freqMap[v]; ok && freqMap[v] == 1 {
            return i
        }
    }
    return -1
}
