package compare

import (
	"math"
)

// PearsonCorrelation 计算皮尔逊相关系数
func PearsonCorrelation(v1, v2 []float64) float64 {
	if len(v1) != len(v2) {
		panic("Vectors must have the same dimensions")
	}

	var sumX, sumY, sumXY, sumX2, sumY2 float64
	n := float64(len(v1))

	for i := 0; i < len(v1); i++ {
		sumX += v1[i]
		sumY += v2[i]
		sumXY += v1[i] * v2[i]
		sumX2 += v1[i] * v1[i]
		sumY2 += v2[i] * v2[i]
	}

	numerator := sumXY - (sumX * sumY / n)
	denominator := math.Sqrt((sumX2 - sumX*sumX/n) * (sumY2 - sumY*sumY/n))

	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}
