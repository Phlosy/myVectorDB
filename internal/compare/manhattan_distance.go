package compare

import "math"

// ManhattanDistance 计算曼哈顿距离
func ManhattanDistance(vector1, vector2 []float64) float64 {
	if len(vector1) != len(vector2) {
		panic("Vectors must have the same dimensions")
	}
	var sum float64
	for i := 0; i < len(vector1); i++ {
		sum += math.Abs(vector1[i] - vector2[i])
	}
	return sum
}
