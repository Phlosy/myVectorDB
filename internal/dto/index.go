package dto

// Index struct stores indexing information for all vectors
type Index struct {
	Vectors map[string]Vector `json:"vectors"`
}

// HNSWIndex 使用HNSW索引方式构建的Index
type HNSWIndex struct {
	Index
	MaxLevel    int                 `json:"max_level"`    // 最大层数
	EfConstruct int                 `json:"ef_construct"` // 构建时的邻居数
	M           int                 `json:"m"`            // 每个节点的最大连接数
	Layers      map[int][]string    `json:"layers"`       // 每层的节点ID列表
	Neighbors   map[string][]string `json:"neighbors"`    // 每个节点的邻居节点ID列表
}
