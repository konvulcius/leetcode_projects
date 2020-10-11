package main

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	var (
		out = make([][]int, 0, len(intervals)+1)
		cur = 0
	)

	for i := range intervals {
		if newInterval[0] <= intervals[i][0] {
			out = append(out, newInterval)
			out = append(out, intervals[i:]...)
			break
		}
		out = append(out, intervals[i])
	}
	if len(intervals) == len(out) {
		out = append(out, newInterval)
	}

	for i := range out {
		if out[i][0] <= out[cur][1] {
			out[cur][1] = func(i, j int) int {
				if i > j {
					return i
				}
				return j
			}(out[cur][1], out[i][1])
		} else {
			cur++
			out[cur] = out[i]
		}
	}
	return out[:cur+1]
}
