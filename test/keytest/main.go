package main

import (
	"fmt"

	"github.com/Phlosy/myVectorDB/internal/dto"
	"github.com/Phlosy/myVectorDB/internal/vector"
)

func main() {
	// 创建索引实例
	keyindex := vector.NewKeyIndex()

	// 创建测试向量
	testVector := &dto.Vector{
		ID:         "test1",
		VectorData: []float64{1.0, 2.0, 3.0},
	}

	// 创建测试向量
	testVector2 := &dto.Vector{
		ID:         "test2",
		VectorData: []float64{4.0, 5.0, 6.0},
	}

	// 通过key添加向量
	keyindex.Add("test1", testVector)

	// 通过key添加向量
	keyindex.Add("test2", testVector2)

	// 通过key更新向量
	keyindex.Update("test1", testVector)

	// 通过key搜索向量
	found := keyindex.Search("test1")
	fmt.Printf("找到向量: ID=%s, Vector=%v\n", found.ID, found.VectorData)

	// 通过key搜索向量
	found = keyindex.Search("test2")
	fmt.Printf("找到向量: ID=%s, Vector=%v\n", found.ID, found.VectorData)

	// 获取所有向量
	vectors := keyindex.GetAllVectors()
	fmt.Printf("所有向量: %v\n", vectors)

	// 通过key删除向量
	keyindex.Delete("test1")

	// 通过key搜索向量
	found = keyindex.Search("test1")

}
