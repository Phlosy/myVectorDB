package main

import (
	"fmt"

	"github.com/Phlosy/myVectorDB/internal/dto"
	"github.com/Phlosy/myVectorDB/internal/vector"
)

func main() {
	flatIndex := vector.NewFlatIndex()
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

	// 创建测试向量
	testVector3 := &dto.Vector{
		ID:         "test3",
		VectorData: []float64{7.0, 8.0, 9.0},
	}

	// 创建测试向量
	testVector4 := &dto.Vector{
		ID:         "test4",
		VectorData: []float64{10.0, 11.0, 12.0},
	}

	// 创建测试向量
	testVector5 := &dto.Vector{
		ID:         "test5",
		VectorData: []float64{13.0, 14.0, 15.0},
	}

	// 添加向量
	flatIndex.Add(testVector)
	flatIndex.Add(testVector2)
	flatIndex.Add(testVector3)
	flatIndex.Add(testVector4)
	flatIndex.Add(testVector5)

	// 搜索向量
	results, err := flatIndex.Search(testVector, 0, "euclideandistance")
	if err != nil {
		fmt.Println("搜索失败:", err)
	}
	fmt.Println("搜索结果:", results)
}
