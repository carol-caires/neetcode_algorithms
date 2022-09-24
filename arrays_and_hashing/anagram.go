package arrays_and_hashing

// There is another solution involving sorting the strings, but we never know the time complexity of the sorting algorithm

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
func isAnagram(s string, t string) bool {
	var sRunes, tRunes rune
	for _, runeValue := range s {
		sRunes += runeValue
	}
	for _, runeValue := range t {
		tRunes += runeValue
	}
	if sRunes == tRunes {
		return true
	}
	return false
}

// Other Solutions

// do the same thing, counting characters occurrences with hash maps
func isAnagram2(s string, t string) bool {
	countS := map[uint8]int{}
	countT := map[uint8]int{}

	for i := 0; i < len(s); i++ {
		// if t[i] does not exist, consider 0
		// s[i] will always exist as we are using len(s) as maximum size
		var ti uint8
		if len(t) > i {
			ti = t[i]
		}
		countS[s[i]] = countS[s[i]] + 1
		countT[ti] = countT[ti] + 1
	}

	for index := range countS {
		if countS[index] != countT[index] {
			return false
		}
	}
	return true
}
