package compare

// CosineDistance 计算余弦距离
func CosineDistance(vector1, vector2 []float64) float64 {
	return 1 - CosineSimilarity(vector1, vector2)
}
