package arrays_and_hashing

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
func topKFrequent(nums []int, k int) []int {
	count := map[int]int{} // index: nums[i]; value: count
	for _, n := range nums {
		count[n] = count[n] + 1
	}

	freq := make([][]int, len(nums)+1)
	for num, c := range count {
		freq[c] = append(freq[c], num)
	}

	res := []int{}
	for i := len(freq) - 1; i >= 0; i-- {
		for _, n := range freq[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
