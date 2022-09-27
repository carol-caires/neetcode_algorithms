package arrays_and_hashing

// Given an array of integers nums and an integer target,
//return indices of the two numbers such that they add up to target.
// video: https://www.youtube.com/watch?v=KLlXCFG5TnA

// first solution O(n^2) complexity (brute force)
func twoSum(nums []int, target int) []int {
	for key := range nums {
		for key2 := range nums {
			if key2 <= key {
				continue
			}
			if nums[key]+nums[key2] == target {
				return []int{key, key2}
			}
		}
	}
	return []int{}
}

// it does the same thing, but with time complexity of O(n). Explanation: https://youtu.be/KLlXCFG5TnA?t=110
func twoSum2(nums []int, target int) []int {
	hashSet := map[int]int{}
	for key, value := range nums {
		res := target - value
		if key2, ok := hashSet[res]; ok {
			return []int{key2, key}
		}
		hashSet[value] = key
	}
	return []int{}
}
