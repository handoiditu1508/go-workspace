package exercise2

import (
	"sort"
	"strconv"
	"strings"
)

var weekDays map[string]int = map[string]int{"Mon": 0, "Tue": 1, "Wed": 2, "Thu": 3, "Fri": 4, "Sat": 5, "Sun": 6}
var maxMinutes = 24 * 60 * 7

type timePeriod struct {
	Start int
	End   int
}

func InvokeExercise3(input string) int {
	// splite input
	splitedInput := strings.Split(input, "\n")

	// convert string to struct
	meetings := make([]timePeriod, 0, 100)
	for _, value := range splitedInput {
		dayOffSet := weekDays[value[:3]] * 24 * 60
		period := timePeriod{dayOffSet + toMinutes(value[4:9]), dayOffSet + toMinutes(value[10:15])}
		meetings = append(meetings, period)
	}

	// sort
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].Start < meetings[j].Start
	})

	// find longest time
	longestTime := meetings[0].Start
	for i := 0; i < len(meetings)-1; i++ {
		newTime := meetings[i+1].Start - meetings[i].End
		if newTime > longestTime {
			longestTime = newTime
		}
	}

	newTime := maxMinutes - meetings[len(meetings)-1].End
	if newTime > longestTime {
		longestTime = newTime
	}

	return longestTime
}

func toMinutes(time string) int {
	hours, _ := strconv.ParseInt(time[0:2], 0, 8)
	minutes, _ := strconv.ParseInt(time[3:5], 0, 8)
	return int(hours)*60 + int(minutes)
}
