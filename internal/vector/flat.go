package vector

import (
	"github.com/Phlosy/myVectorDB/internal/dto"
)

// FlatIndex 使用扁平索引方式构建的Index
type FlatIndex struct {
	dto.Index
	// 扁平索引不需要额外的字段,直接使用基础Index即可
}

func NewFlatIndex() *FlatIndex {
	return &FlatIndex{
		Index: dto.Index{},
	}
}

// Add adds a vector to the database
func (i *FlatIndex) FlatIndexAdd(vector *dto.Vector) error {
	i.Index.Vectors[vector.ID] = *vector
	return nil
}

// Search searches for vectors in the database
func (i *FlatIndex) FlatIndexSearch(vector *dto.Vector) ([]dto.Vector, error) {
	results := make([]dto.Vector, 0)
	for _, v := range i.Index.Vectors {
		results = append(results, v)
	}
	return results, nil
}

// Delete deletes a vector from the database
func (i *FlatIndex) FlatIndexDelete(vector *dto.Vector) error {
	return nil
}

// Update updates a vector in the database
func (i *FlatIndex) FlatIndexUpdate(vector *dto.Vector) error {
	return nil
}

// GetAllVectors gets all vectors from the database
func (i *FlatIndex) FlatIndexGetAllVectors() ([]dto.Vector, error) {
	return nil, nil
}
