package index

import (
	"github.com/Phlosy/myVectorDB/internal/compare"
	"github.com/Phlosy/myVectorDB/internal/dto"
)

type DistanceResult struct {
	Distance float64
	Vector   *dto.Vector
}

// FlatAdd 添加向量
func FlatAdd(vector *dto.Vector, index *dto.Index) error {
	index.Vectors[vector.ID] = *vector
	return nil
}

// FlatSearchByThreshold 通过相似度阈值搜索向量
func FlatSearchByThreshold(index *dto.Index, vector *dto.Vector, threshold float64, compareFunc string) ([]dto.Vector, error) {
	distanceResults := calculateDistances(index, vector, compareFunc)
	results := make([]dto.Vector, 0)
	for i := 0; i < len(distanceResults); i++ {
		if distanceResults[i].Distance > threshold {
			results = append(results, *distanceResults[i].Vector)
		}
	}
	return results, nil
}

// 根据输入的方法计算距离
func calculateDistances(index *dto.Index, vector *dto.Vector, compareFunc string) []DistanceResult {
	results := make([]DistanceResult, 0)
	for _, v := range index.Vectors {
		var distance float64
		switch compareFunc {
		case "euclideandistance":
			distance = compare.EuclideanDistance(vector.VectorData, v.VectorData)
		case "cosinesimilarity":
			distance = compare.CosineSimilarity(vector.VectorData, v.VectorData)
		case "pearsonscorrelation":
			distance = compare.PearsonCorrelation(vector.VectorData, v.VectorData)
		case "manhattandistance":
			distance = compare.ManhattanDistance(vector.VectorData, v.VectorData)
		case "cosinedistance":
			distance = compare.CosineDistance(vector.VectorData, v.VectorData)
		}

		results = append(results, DistanceResult{
			Distance: distance,
			Vector:   &v,
		})
	}
	return results
}

// FlatDelete 删除向量
func FlatDelete(vector *dto.Vector, index *dto.Index) error {
	delete(index.Vectors, vector.ID)
	return nil
}

// FlatUpdate 更新向量
func FlatUpdate(vector *dto.Vector, index *dto.Index) error {
	index.Vectors[vector.ID] = *vector
	return nil
}

// FlatGetAllVectors 获取所有向量
func FlatGetAllVectors(index *dto.Index) ([]dto.Vector, error) {
	vectors := make([]dto.Vector, 0, len(index.Vectors))
	for _, vector := range index.Vectors {
		vectors = append(vectors, vector)
	}
	return vectors, nil
}
