package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// 可用字符列表
	var characters = []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+'}

	var result = make([]rune, 16) // 结果切片

	for i := 0; i < len(result); i++ { // 填充切片
		var index, err = rand.Int(rand.Reader, big.NewInt(int64(len(characters)))) // 安全随机一个 index
		if err != nil {
			fmt.Println(err)
			return
		}

		result[i] = characters[int(index.Int64())] // 将结果取出并填充到 result
	}

	fmt.Println(string(result))
}
