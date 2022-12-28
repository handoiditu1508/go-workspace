package exercise2

import "strings"

func InvokeNotification(message string, K int) string {
	if len(message) <= K {
		return message
	}

	if K == 3 {
		return "..."
	}

	if message[K-4] != ' ' {
		var nextSpaceIndex = strings.LastIndexByte(message[:K-4], ' ')
		if nextSpaceIndex == -1 {
			return "..."
		} else {
			return message[:nextSpaceIndex+1] + "..."
		}
	} else {
		return message[:K-3] + "..."
	}
}
