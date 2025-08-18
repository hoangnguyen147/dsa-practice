package stack

func dailyTemperatures(temperatures []int) []int {
    if len(temperatures) == 0 {
        return []int{}
    }
    if len(temperatures) == 1 {
        return []int{0}
    }
    stack := make([]int, 0)
    result := make([]int, len(temperatures))
    for i, t := range temperatures {
        for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
            result[stack[len(stack)-1]] = i - stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    
    return result
}
