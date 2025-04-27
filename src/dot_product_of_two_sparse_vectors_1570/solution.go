package dot_product_of_two_sparse_vectors_1570

import "fmt"

type IdxAndValue struct {
	idx   int
	value int
}

type SparseVector struct {
	vector []IdxAndValue
	len    int
}

func Constructor(nums []int) SparseVector {
	vector := []IdxAndValue{}

	for idx, num := range nums {
		if num != 0 {
			vector = append(vector, IdxAndValue{idx, num})
		}
	}

	return SparseVector{vector: vector, len: len(vector)}
}

func (this *SparseVector) dotProduct(other SparseVector) int {
	product := 0

	thisIdx := 0
	otherIdx := 0
	for thisIdx < this.len && otherIdx < other.len {
		thisElement := this.vector[thisIdx]
		otherElement := other.vector[otherIdx]

		if thisElement.idx == otherElement.idx {
			product += thisElement.value * otherElement.value
			thisIdx++
			otherIdx++
		} else if thisElement.idx > otherElement.idx {
			otherIdx++
		} else {
			thisIdx++
		}
	}

	return product
}

func Test() {
	v1 := Constructor([]int{1, 0, 0, 2, 3})
	v2 := Constructor([]int{0, 3, 0, 4, 0})
	fmt.Println(v1.dotProduct(v2))

	v1 = Constructor([]int{0, 1, 0, 0, 0})
	v2 = Constructor([]int{0, 0, 0, 0, 2})
	fmt.Println(v1.dotProduct(v2))

	v1 = Constructor([]int{0, 1, 0, 0, 2, 0, 0})
	v2 = Constructor([]int{1, 0, 0, 0, 3, 0, 4})
	fmt.Println(v1.dotProduct(v2))
}
