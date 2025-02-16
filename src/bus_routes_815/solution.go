package bus_routes_815

import "container/list"

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	stopToBusIds := make(map[int][]int)

	for busId, stops := range routes {
		for _, stop := range stops {
			stopToBusIds[stop] = append(stopToBusIds[stop], busId)
		}
	}

	queue := list.New()
	visited := make(map[int]struct{})

	for _, busId := range stopToBusIds[source] {
		queue.PushBack(busId)
		visited[busId] = struct{}{}
	}

	busCount := 1
	for queue.Len() > 0 {
		size := queue.Len()

		for i := 0; i < size; i++ {
			busIdPtr := queue.Front()
			queue.Remove(busIdPtr)
			busId := busIdPtr.Value.(int)

			for _, stop := range routes[busId] {
				if stop == target {
					return busCount
				}

				for _, nextBusId := range stopToBusIds[stop] {
					if _, seen := visited[nextBusId]; !seen {
						visited[nextBusId] = struct{}{}
						queue.PushBack(nextBusId)
					}
				}
			}
		}
		busCount++
	}

	return -1
}
