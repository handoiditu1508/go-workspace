package exercise2

func InvokePalindrome(S string) int {
	var (
		charCounts = make(map[byte]int)
		oddCount   int
	)

	for i := 0; i < len(S); i++ {
		var char = S[i]
		charCounts[char]++
	}

	for _, charCount := range charCounts {
		if charCount%2 != 0 {
			oddCount++
		}
	}

	if oddCount != 0 {
		return oddCount - 1
	}
	return 0
}
