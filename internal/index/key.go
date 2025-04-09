package index

import (
	"fmt"

	"github.com/Phlosy/myVectorDB/internal/dto"
)

// AddByKeys 添加向量
func AddByKeys(keys string, vector *dto.Vector, index *dto.Index) error {
	index.Vectors[keys] = *vector
	return nil
}

// SearchByKeys 搜索向量
func SearchByKeys(keys string, index *dto.Index) (*dto.Vector, error) {
	vector, ok := index.Vectors[keys]
	if !ok {
		return nil, fmt.Errorf("vector not found")
	}
	return &vector, nil
}

// DeleteByKeys 删除向量
func DeleteByKeys(keys string, index *dto.Index) error {
	delete(index.Vectors, keys)
	return nil
}

// UpdateByKeys 更新向量
func UpdateByKeys(keys string, vector *dto.Vector, index *dto.Index) error {
	index.Vectors[keys] = *vector
	return nil
}

// GetAllVectors 获取所有向量
func GetAllVectors(index *dto.Index) ([]dto.Vector, error) {
	vectors := make([]dto.Vector, 0, len(index.Vectors))
	for _, vector := range index.Vectors {
		vectors = append(vectors, vector)
	}
	return vectors, nil
}
