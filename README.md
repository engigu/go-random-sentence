# Go Random Sentence

这是一个 Go 语言包，用于获取随机诗词/句子。数据已嵌入包中，无需外部文件依赖。

## 功能特点

1. **高性能**：初始化时预处理行索引，随机读取时无大内存分配。
2. **零依赖/自包含**：利用 Go `embed` 特性将数据打包进二进制文件以及库中。
3. **API 简单**：只需调用 `Random()` 即可获得结果。

## 安装

```bash
go get github.com/engigu/go-random-sentence
```

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/engigu/go-random-sentence"
)

func main() {
	// 获取随机诗词
	// 返回格式为 map[string]interface{}
	// 数据包含: name (句子), from (来源), _id (ID)
	data, err := sentence.Random()
	if err != nil {
		panic(err)
	}

	fmt.Printf("句子: %v\n", data["name"])
	fmt.Printf("来源: %v\n", data["from"])
}
```
