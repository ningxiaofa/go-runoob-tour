// https://www.runoob.com/go/go-slice.html -- Go 语言切片(Slice)

package main

import "fmt"

func sliceFunc() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
// numbers == [0 1 2 3 4 5 6 7 8]
// numbers[1:4] == [1 2 3]
// numbers[:3] == [0 1 2]
// numbers[4:] == [4 5 6 7 8]
// len=0 cap=5 slice=[]
// len=2 cap=9 slice=[0 1]
// len=3 cap=7 slice=[2 3 4]
// 可以看出，len有cap的区别，如果是基于已有数组/切片创建，那么len就是新切片的元素个数，cap就是从第一个元素下标开始到最后一个已有数组/切片的最后一个下标元素的个数

// append() 和 copy() 函数
// 如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。 -- 动态扩容，很多编程语言的底层实现都是如此
// 下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法
func dynamicSlice() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}
