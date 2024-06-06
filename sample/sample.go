package main

import (
	"fmt"
	"github.com/i31pc17/zerowidth"
)

func main() {
	zw := zerowidth.NewZeroWidth()

	text := "한글 테스트 중(​)입니다.❤️ 공백도 테스트 해️보겠습니다."
	result, err := zw.Remove(text)

	fmt.Println("제거전 : ", text)
	if err == nil {
		fmt.Println("제거후 : ", result)
	}
}
