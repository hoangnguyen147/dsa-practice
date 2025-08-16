package open_the_lock

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	deadMap := map[string]struct{}{}
	for _, d := range deadends {
        if d == "0000" || d == target {
            return -1
        }
		deadMap[d] = struct{}{}
	}
	visited := map[string]struct{}{}

	queue := []string{}
	queue = append(queue, "0000")
    visited["0000"] = struct{}{}
    step := 0
    
    getNexts := func(str string) []string {
		nexts := make([]string, 0, 8)
		for i := 0; i < 4; i++ {
			r1 := []rune(str)
			r2 := []rune(str)
			// down
			if r1[i] == '0' {
				r1[i] = '9'
			} else {
				r1[i]--
			}

			// up
			if r2[i] == '9' {
				r2[i] = '0'
			} else {
				r2[i]++
			}

			if _, ok := deadMap[string(r1)]; !ok {
                s1 := string(r1)
				if _, see := visited[string(r1)]; !see {
                    visited[s1] = struct{}{}
					nexts = append(nexts, s1)
				}
			}
			if _, ok := deadMap[string(r2)]; !ok {
                s2 := string(r2)
				if _, see := visited[s2]; !see {
                    visited[s2] = struct{}{}
					nexts = append(nexts, s2)
				}
			}
		}

        return nexts
	}

	for len(queue) > 0 {
        nextQueue := []string{}
        for _, curr := range queue {
            if curr == target {
                return step
            }
            nexts := getNexts(curr)
            nextQueue = append(nextQueue, nexts...)
        }
        step++
        queue = nextQueue
	}

	return -1
}
