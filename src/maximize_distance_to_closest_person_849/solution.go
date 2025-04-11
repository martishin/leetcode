package maximize_distance_to_closest_person_849

import (
	"fmt"
	"math"
)

func maxDistToClosest(seats []int) int {
	resultLen := math.MinInt32
	lastSeen := -1
	n := len(seats)

	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			if lastSeen != -1 {
				resultLen = max(resultLen, (i-lastSeen)/2)
			} else {
				resultLen = i
			}
			lastSeen = i
		}
	}

	resultLen = max(resultLen, n-lastSeen-1)

	return resultLen
}

func Test() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1}))
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))
}
