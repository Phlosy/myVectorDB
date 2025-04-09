package vector

import (
	"log"

	"github.com/Phlosy/myVectorDB/internal/dto"
	"github.com/Phlosy/myVectorDB/internal/index"
)

type KeyIndex struct {
	dto.Index
}

func NewKeyIndex() *KeyIndex {
	return &KeyIndex{
		Index: dto.Index{
			Vectors: make(map[string]dto.Vector),
		},
	}
}

// 添加向量
func (i *KeyIndex) Add(key string, vector *dto.Vector) {
	err := index.AddByKeys(key, vector, &i.Index)
	if err != nil {
		log.Fatal("添加向量失败:", err)
	}

}

// 根据Key搜索向量
func (i *KeyIndex) Search(key string) *dto.Vector {
	vector, ok := i.Index.Vectors[key]
	if !ok {
		log.Fatal("搜索向量失败：向量不存在")
	}
	return &vector
}

// 根据Key删除向量
func (i *KeyIndex) Delete(key string) {
	err := index.DeleteByKeys(key, &i.Index)
	if err != nil {
		log.Fatal("删除向量失败:", err)
	}
}

// 根据Key更新向量
func (i *KeyIndex) Update(key string, vector *dto.Vector) {
	err := index.UpdateByKeys(key, vector, &i.Index)
	if err != nil {
		log.Fatal("更新向量失败:", err)
	}
}

// 获取所有向量
func (i *KeyIndex) GetAllVectors() []dto.Vector {
	vectors, err := index.GetAllVectors(&i.Index)
	if err != nil {
		log.Fatal("获取所有向量失败:", err)
	}
	return vectors
}
