package course_schedule_ii_210

func findOrderKahn(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	indegree := make([]int, numCourses)
	order := []int{}

	for i := range prerequisites {
		src := prerequisites[i][1]
		dst := prerequisites[i][0]

		graph[src] = append(graph[src], dst)
		indegree[dst]++
	}

	queue := []int{}
	for i := range numCourses {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]

		order = append(order, course)

		for _, nb := range graph[course] {
			indegree[nb]--

			if indegree[nb] == 0 {
				queue = append(queue, nb)
			}
		}
	}

	if len(order) != numCourses {
		return []int{}
	}

	return order
}
