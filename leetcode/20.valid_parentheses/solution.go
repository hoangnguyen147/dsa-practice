package valid_parentheses

func isValid(s string) bool {
    if len(s) == 0 || len(s) % 2 != 0 {
        return false
    }
    
    stack := make([]rune, 0)
    for _, rc := range s {
        c := string(rc)
        if c == "(" || c == "{" || c == "[" {
            stack = append(stack, rc)
        } else {
            if len(stack) == 0 {
                return false
            }
            i := string(stack[len(stack)-1])
            if (i == "(" && c == ")") || (i == "{" && c == "}") || (i == "[" && c == "]") {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
    }
    
    return len(stack) == 0
}
