// https://www.runoob.com/go/go-type-casting.html -- Go 语言类型转换

package main

import "fmt"

func dataTypeConversion() {
	var sum int = 17
	var count int = 5
	// var mean float32
	// mean = float32(sum) / float32(count)

	mean := float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}

// mean 的值为: 3.400000
