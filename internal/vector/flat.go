package vector

import (
	"log"

	"github.com/Phlosy/myVectorDB/internal/dto"
	"github.com/Phlosy/myVectorDB/internal/index"
)

// FlatIndex 使用扁平索引方式构建的Index
type FlatIndex struct {
	dto.Index
}

func NewFlatIndex() *FlatIndex {
	return &FlatIndex{
		Index: dto.Index{
			Vectors: make(map[string]dto.Vector),
		},
	}
}

// 添加向量
func (i *FlatIndex) Add(vector *dto.Vector) {
	err := index.AddByKeys(vector.ID, vector, &i.Index)
	if err != nil {
		log.Fatal("添加向量失败:", err)
	}

}

func (i *FlatIndex) Search(vector *dto.Vector, threshold float64, compareFunc string) ([]dto.Vector, error) {
	searchresult, err := index.FlatSearchByThreshold(&i.Index, vector, threshold, compareFunc)
	if err != nil {
		log.Fatal("搜索向量失败:", err)
	}
	return searchresult, nil
}

// 根据Key删除向量
func (i *FlatIndex) Delete(key string) {
	err := index.DeleteByKeys(key, &i.Index)
	if err != nil {
		log.Fatal("删除向量失败:", err)
	}
}

// 根据Key更新向量
func (i *FlatIndex) Update(key string, vector *dto.Vector) {
	err := index.UpdateByKeys(key, vector, &i.Index)
	if err != nil {
		log.Fatal("更新向量失败:", err)
	}
}

// 获取所有向量
func (i *FlatIndex) GetAllVectors() []dto.Vector {
	vectors, err := index.GetAllVectors(&i.Index)
	if err != nil {
		log.Fatal("获取所有向量失败:", err)
	}
	return vectors
}
