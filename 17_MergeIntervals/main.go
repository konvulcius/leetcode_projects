package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	var cur = 0
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[cur][1] {
			intervals[cur][1] = func(i, j int) int {
				if i > j {
					return i
				}
				return j
			}(intervals[cur][1], intervals[i][1])
		} else {
			cur++
			intervals[cur] = intervals[i]
		}
	}
	return intervals[:cur+1]
}
