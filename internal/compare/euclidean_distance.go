package compare

import "math"

// EuclideanDistance 计算两个向量之间的欧几里得距离
func EuclideanDistance(vector1, vector2 []float64) float64 {
	if len(vector1) != len(vector2) {
		panic("Vectors must have the same dimensions")
	}
	var sum float64
	for i := 0; i < len(vector1); i++ {
		sum += math.Pow(vector1[i]-vector2[i], 2)
	}
	return math.Sqrt(sum)
}
