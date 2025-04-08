# myVectorDB 🚀

一个简单的向量数据库学习项目，用于理解向量数据库的基本原理和实现。

## 项目结构

```
├── Makefile          # 项目构建和管理
├── README.md         # 项目文档
├── api              # API 接口定义
├── bin              # 编译后的二进制文件
├── cmd              # 主程序入口
│   └── main.go
├── config           # 配置文件
├── go.mod           # Go 模块定义
├── internal         # 内部包
│   ├── dto         # 数据传输对象
│   ├── index       # 索引实现（扁平索引、HNSW等）
│   ├── model       # 数据模型
│   ├── storage     # 数据持久化
│   ├── utils       # 工具函数
│   └── vector      # 向量操作
├── pkg             # 可供外部使用的包
└── test            # 测试相关
    └── benchmark   # 性能测试
```

## 功能特性

- 支持多种索引类型
  - 扁平索引（暴力搜索）
  - HNSW 索引（近似最近邻）
- 向量操作
  - 添加/删除向量
  - 向量搜索
  - 向量更新
- 数据持久化
- 并发处理

## 快速开始

### 安装

```bash
go get github.com/Phlosy/myVectorDB
```

### 使用示例

```go
package main

import (
    "github.com/Phlosy/myVectorDB/internal/vector"
)

func main() {
    // 创建索引
    index := vector.NewFlatIndex()
    
    // TODO: 添加使用示例
}
```

## 开发计划

- [ ] 实现基础向量操作
- [ ] 完善 HNSW 索引
- [ ] 添加持久化支持
- [ ] 优化并发性能
- [ ] 添加 API 接口

## 贡献指南

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License