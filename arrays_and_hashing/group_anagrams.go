package arrays_and_hashing

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// O(n^2) solution
func groupAnagrams(strs []string) [][]string {
	anagramsMap := map[int32][]string{}
	for key := range strs {
		var charsSum rune
		for _, char := range strs[key] {
			charsSum += char
		}
		anagramsMap[charsSum] = append(anagramsMap[charsSum], strs[key])
	}

	resp := [][]string{}
	for _, val := range anagramsMap {
		resp = append(resp, val)
	}
	return resp
}
