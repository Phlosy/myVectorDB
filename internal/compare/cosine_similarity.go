package compare

import "math"

// CosineSimilarity 计算余弦相似度
func CosineSimilarity(vector1, vector2 []float64) float64 {
	if len(vector1) != len(vector2) {
		panic("Vectors must have the same dimensions")
	}

	var dotProduct, v1Norm, v2Norm float64
	for i := 0; i < len(vector1); i++ {
		dotProduct += vector1[i] * vector2[i]
		v1Norm += math.Pow(vector1[i], 2)
		v2Norm += math.Pow(vector2[i], 2)
	}

	return dotProduct / (math.Sqrt(v1Norm) * math.Sqrt(v2Norm))
}
