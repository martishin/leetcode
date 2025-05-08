package optimize_water_distribution_in_a_village_1168

import (
	"fmt"
	"sort"
)

type UnionFind struct {
	group []int
	rank  []int
}

func NewUnionFind(size int) *UnionFind {
	group := make([]int, size+1)
	rank := make([]int, size+1)

	for i := 0; i <= size; i++ {
		group[i] = i
		rank[i] = 0
	}

	return &UnionFind{group: group, rank: rank}
}

func (uf *UnionFind) Find(element int) int {
	if uf.group[element] != element {
		uf.group[element] = uf.Find(uf.group[element])
	}

	return uf.group[element]
}

func (uf *UnionFind) Union(element1 int, element2 int) bool {
	group1 := uf.Find(element1)
	group2 := uf.Find(element2)

	if group1 == group2 {
		return false
	}

	if uf.rank[group1] < uf.rank[group2] {
		uf.group[group1] = group2
	} else if uf.rank[group1] > uf.rank[group2] {
		uf.group[group2] = group1
	} else {
		uf.group[group2] = group1
		uf.rank[group1]++
	}

	return true
}

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	edgesByCost := make([][]int, 0, len(wells)+len(pipes))

	for i := 0; i < n; i++ {
		edgesByCost = append(edgesByCost, []int{0, i + 1, wells[i]})
	}

	for _, pipe := range pipes {
		edgesByCost = append(edgesByCost, pipe)
	}

	sort.Slice(edgesByCost, func(i int, j int) bool {
		return edgesByCost[i][2] < edgesByCost[j][2]
	})

	uf := NewUnionFind(n)
	minCost := 0
	for _, edge := range edgesByCost {
		house1, house2, cost := edge[0], edge[1], edge[2]

		if uf.Union(house1, house2) {
			minCost += cost
		}
	}

	return minCost
}

func Test() {
	fmt.Println(minCostToSupplyWater(3, []int{1, 2, 2}, [][]int{{1, 2, 1}, {2, 3, 1}}))
	fmt.Println(minCostToSupplyWater(2, []int{1, 1}, [][]int{{1, 2, 1}, {1, 2, 2}}))
}
