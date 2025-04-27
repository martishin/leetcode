package dot_product_of_two_sparse_vectors_1570

type SparseVectorNegative struct {
	len    int
	vector []int
}

func ConstructorNegative(nums []int) SparseVectorNegative {
	vector := []int{}
	for _, num := range nums {
		if num > 0 {
			vector = append(vector, num)
		} else {
			if len(vector) == 0 {
				vector = append(vector, -1)
				continue
			}
			if vector[len(vector)-1] > 0 {
				vector = append(vector, 0)
			}
			vector[len(vector)-1]--
		}
	}

	return SparseVectorNegative{len: len(vector), vector: vector}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVectorNegative) dotProduct(vec SparseVectorNegative) int {
	thisIdx := 0
	otherIdx := 0

	result := 0
	for thisIdx < this.len && otherIdx < vec.len {
		if this.vector[thisIdx] == 0 {
			thisIdx++
		}

		if vec.vector[otherIdx] == 0 {
			otherIdx++
		}

		if this.vector[thisIdx] > 0 && vec.vector[otherIdx] > 0 {
			result += this.vector[thisIdx] * vec.vector[otherIdx]
			thisIdx++
			otherIdx++
		} else if this.vector[thisIdx] > 0 {
			thisIdx++
			vec.vector[otherIdx]++
		} else if vec.vector[otherIdx] > 0 {
			otherIdx++
			this.vector[thisIdx]++
		} else {
			if this.vector[thisIdx] == vec.vector[otherIdx] {
				thisIdx++
				otherIdx++
			} else if this.vector[thisIdx] < vec.vector[otherIdx] {
				this.vector[thisIdx] -= vec.vector[otherIdx]
				otherIdx++
			} else {
				vec.vector[otherIdx] -= this.vector[thisIdx]
				thisIdx++
			}
		}
	}

	return result
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
