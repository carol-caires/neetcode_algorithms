package arrays_and_hashing

// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.
func containsDuplicate(nums []int) bool {
	count := map[int]struct{}{} // empty structs occupy 0 memory
	for _, v := range nums {
		if _, ok := count[v]; ok {
			return true
		}
		count[v] = struct{}{}
	}
	return false
}
