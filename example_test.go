package sentence_test

import (
	"fmt"
	"github.com/engigu/go-random-sentence"
)

// ExampleRandom demonstrates how to use the Random function.
func ExampleRandom() {
	// 获取随机诗词
	// 结果是一个 map，包含 json 中的字段，如 "name", "from" 等
	data, err := sentence.Random()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印输出 (这里因为是随机的，实际输出会变化)
	// 为了演示，我们只打印 key 是否存在
	if name, ok := data["name"]; ok {
		fmt.Printf("Get sentence: %v\n", name)
	}
	
	// Output:
	// (Note: The strict Output comment is omitted because randomness makes strict matching impossible for 'go test', 
	// but this function demonstrates usage.)
}

func ExampleRandom_print() {
	// 简单调用示例
	data, _ := sentence.Random()
	fmt.Printf("句子: %s\n来源: %s\n", data["name"], data["from"])
}
