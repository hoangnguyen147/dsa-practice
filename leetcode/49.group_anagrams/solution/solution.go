package solution

import "strconv"

func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 {
        return [][]string{{""}}
    }
    if len(strs) == 1 {
        return [][]string{{strs[0]}}
    }
    res := make([][]string, 0)
    alphabets := []rune{}
    for i := 'a'; i <= 'z'; i++ {
        alphabets = append(alphabets, i)
    }
    strMap := make(map[string][]string)
    
    for _, c := range strs {
        key := getKey(alphabets, c)
        if _, ok := strMap[key]; !ok {
            strMap[key] = []string{c}
        } else {
            strMap[key] = append(strMap[key], c)
        }
    }
    
    for _, arr := range strMap {
        res = append(res, arr)
    }
    
    return res
}

func getKey(alphabets []rune, str string) string {
    freqMap := make(map[rune]int)
    var key string
    for _, i := range str {
        if _, ok := freqMap[i]; !ok {
            freqMap[i] = 1
        } else {
            freqMap[i]++
        }
    }
    for _, letter := range alphabets {
        if k, ok := freqMap[letter]; ok {
            kstr := strconv.Itoa(k)
            key += string(letter) + kstr
        }
    }
    
    return key
}
