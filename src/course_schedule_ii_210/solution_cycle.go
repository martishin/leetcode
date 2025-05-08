package course_schedule_ii_210

type Color int

const (
	White Color = iota
	Grey
	Black
)

func dfs(graph map[int][]int, course int, colors []Color, order *[]int) bool {
	colors[course] = Grey

	for _, nb := range graph[course] {
		if colors[nb] == Grey {
			return false
		} else if colors[nb] == White {
			if !dfs(graph, nb, colors, order) {
				return false
			}
		}
	}

	colors[course] = Black
	*order = append(*order, course)

	return true
}

func findOrderCycle(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)

	for i := range prerequisites {
		src := prerequisites[i][1]
		dst := prerequisites[i][0]

		graph[src] = append(graph[src], dst)
	}

	colors := make([]Color, numCourses)

	order := []int{}
	for i := range numCourses {
		if colors[i] == White && !dfs(graph, i, colors, &order) {
			return []int{}
		}
	}

	for i := 0; i < len(order)/2; i++ {
		order[i], order[len(order)-i-1] = order[len(order)-i-1], order[i]
	}

	return order
}
