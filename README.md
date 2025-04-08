# myVectorDB 🚀
一个简单的向量数据库学习项目，用于理解向量数据库的基本原理和实现。

├── Makefile
├── README.md
├── api
├── bin
├── cmd
│   └── main.go
├── config
├── go.mod
├── internal
│   ├── dto # 声明结构体
│   ├── index # 声明索引方式，例如扁平索引、HNSW索引等
│   ├── model # 声明模型
│   ├── storage # 声明数据持久化方式
│   ├── utils # 声明工具
│   └── vector # 声明对向量数据的操作
├── pkg
└── test
    └── benchmark