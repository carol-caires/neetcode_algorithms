package arrays_and_hashing

/*
	Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
	The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
	You must write an algorithm that runs in O(n) time and without using the division operation.
*/

// time O(n), space O(n)
func productExceptSelf(nums []int) []int {
	out := []int{}
	prefix := []int{}
	postfix := []int{}
	for i := 0; i < len(nums); i++ {
		postfix = append(postfix, 0)
		if i == 0 {
			prefix = append(prefix, nums[i])
		}
		if i-1 < 0 {
			continue
		}
		prefix = append(prefix, nums[i]*prefix[i-1])
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			postfix[i] = nums[i]
			continue
		}
		postfix[i] = nums[i] * postfix[i+1]
	}

	for i := 0; i < len(nums); i++ {
		var timesPrefix = 1
		if i-1 >= 0 {
			timesPrefix = prefix[i-1]
		}
		var timesPostfix = 1
		if i+1 < len(nums) {
			timesPostfix = postfix[i+1]
		}
		out = append(out, timesPrefix*timesPostfix)
	}
	return out
}

// time O(n), space O(1)
func productExceptSelf2(nums []int) []int {
	out := []int{}
	for range nums {
		out = append(out, 0)
	}

	prefix := 1
	for i := 0; i < len(nums); i++ {
		out[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		out[i] *= postfix
		postfix *= nums[i]
	}

	return out
}
