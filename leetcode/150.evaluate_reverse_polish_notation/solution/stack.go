package stack

import "strconv"

func evalRPN(tokens []string) int {
    stack := make([]int64, 0)
    for _, i := range tokens {
        if i == "+" || i == "-" || i == "*" || i == "/" {
            if len(stack) < 2 {
                return 0
            }
            v := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if i == "+" {
                stack[len(stack)-1] += v
            }
            if i == "-" {
                stack[len(stack)-1] -= v
            }
            if i == "*" {
                stack[len(stack)-1] *= v
            }
            if i == "/" {
                stack[len(stack)-1] = int64(stack[len(stack)-1]/v)
            }
            continue
        }
        num, _ := strconv.ParseInt(i, 10, 64)
        stack = append(stack, num)
    }
    
    if len(stack) != 1 {
        return 0
    }
    
    return int(stack[0])
}