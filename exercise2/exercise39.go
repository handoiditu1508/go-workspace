package exercise2

var n int = 0
var pointsMatrix [][]bool
var pointsCount int
var visitedPoints []bool
var startPoint int

func InvokeExercise39(arrA, arrB []int) bool {
	n = len(arrA)
	pointsMatrix = make([][]bool, n+1)
	pointsCount = 0
	visitedPoints = make([]bool, n+1)

	// init points matrix
	for i := 1; i <= n; i++ {
		pointsMatrix[i] = make([]bool, n+1)
	}
	for i := 0; i < n; i++ {
		pointsMatrix[arrA[i]][arrB[i]] = true
	}

	for i := 1; i <= n; i++ {
		startPoint = i
		visitedPoints[i] = true
		pointsCount++
		for j := 1; j <= n; j++ {
			if pointsMatrix[i][j] && deepSearch(i, j) {
				return true
			}
		}
		visitedPoints[i] = false
		pointsCount--
	}
	return false
}

// return true if graph is looped
func deepSearch(row, col int) bool {
	if visitedPoints[col] {
		return pointsCount == n && col == startPoint
	}

	visitedPoints[col] = true
	pointsCount++
	for j := 1; j <= n; j++ {
		if pointsMatrix[col][j] && deepSearch(col, j) {
			return true
		}
	}
	visitedPoints[col] = false
	pointsCount--

	return false
}

// 1 -> 4 -> 2 -> 3
