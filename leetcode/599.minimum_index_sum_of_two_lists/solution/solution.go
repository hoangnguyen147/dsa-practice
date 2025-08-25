package solution

func findRestaurant(list1 []string, list2 []string) []string {
    m := make(map[string]int)
    res := []string{}
    minI := len(list1) + len(list2)
    for i1, s1 := range list1 {
        if _, ok := m[s1]; !ok {
            m[s1] = i1
        }
    }
    for i2, s2 := range list2 {
        if _, ok := m[s2]; ok {
            if i2 + m[s2] < minI {
                res = []string{s2}
                minI = i2 + m[s2]
            } else if i2 + m[s2] == minI {
                res = append(res, s2)
            }
        }
    }
    
    return res
}
