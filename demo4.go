package main

import "fmt"

func pointerFunc() {
	var ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr)
	fmt.Printf("ptr 的值为 : %x\n", &ptr)

	if nil != ptr {
		fmt.Printf("ptr为empty pointer \n")
	}

	if nil == ptr {
		fmt.Printf("ptr为empty pointer \n")
	}

	fmt.Println(nil)
}

// ptr 的值为 : 0
// ptr 的值为 : c00000e028
// ptr为empty pointer
// <nil>

const MAX int = 3

func ptrArr() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
		fmt.Printf("a[%d] = %d\n", i, ptr[i])
	}
}

// a[0] = 10
// a[0] = 824634466304
// a[1] = 100
// a[1] = 824634466312
// a[2] = 200
// a[2] = 824634466320

// https://www.runoob.com/go/go-pointer-to-pointer.html -- Go 语言指向指针的指针
func ptrPtrVar() {

	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

// 变量 a = 3000
// 指针变量 *ptr = 3000
// 指向指针的指针变量 **pptr = 3000

func multiPtrFunc() {
	var a int = 1
	var ptr1 *int = &a
	var ptr2 **int = &ptr1
	var ptr3 **(*int) = &ptr2 // 也可以写作：var ptr3 ***int = &ptr2
	// 依次类推
	fmt.Println("a:", a)
	fmt.Println("ptr1", ptr1)
	fmt.Println("ptr2", ptr2)
	fmt.Println("ptr3", ptr3)
	fmt.Println("*ptr1", *ptr1)
	fmt.Println("**ptr2", **ptr2)
	fmt.Println("**(*ptr3)", **(*ptr3)) // 也可以写作：***ptr3
}

// a: 1
// ptr1 0xc0000140a0
// ptr2 0xc00000e028
// ptr3 0xc00000e030
// *ptr1 1
// **ptr2 1
// **(*ptr3) 1

// https://www.runoob.com/go/go-passing-pointers-to-functions.html -- Go 语言指针作为函数参数
func ptrParaFunc() {
	a := 100
	b := 200

	fmt.Printf("交换前a的value: %d\n", a)
	fmt.Printf("交换前b的value: %d\n", b)

	swapVal(&a, &b)
	fmt.Printf("交换后a的value: %d\n", a)
	fmt.Printf("交换后b的value: %d\n", b)

	a, b = swapVal2(&a, &b)

	fmt.Printf("交换后a的value: %d\n", a)
	fmt.Printf("交换后b的value: %d\n", b)

	// 交互内存地址方式 -- 不允许
	// swapVal3(a, b)
}

// 输出结果
// 交换前a的value: 100
// 交换前b的value: 200

// 交换后a的value: 200 -- 第一次交互
// 交换后b的value: 100

// 交换后a的value: 100 -- 第二次交互
// 交换后b的value: 200

// 方式一: 通过交换value
func swapVal(a, b *int) {
	// temp := *a
	// *a = *b
	// *b = temp

	// 简写方式
	*a, *b = *b, *a
}

// 交换前a的value: 100
// 交换前b的value: 200
// 交换后a的value: 200
// 交换后b的value: 100

// 方式二：还是通过交互value达到目的, 这里的返回值为了辅助完成重新赋值
func swapVal2(a, b *int) (int, int) {
	/**
	// fmt.Printf("交换前a的内存地址: %x\n", a)
	// fmt.Printf("交换前b的内存地址: %x\n", b)
	a, b = b, a
	// fmt.Printf("交换后a的内存地址: %x\n", a)
	// fmt.Printf("交换后b的内存地址: %x\n", b)
	// 交换后a的内存地址: c0000b2008
	// 交换后b的内存地址: c0000b2010
	// 交换后a的内存地址: c0000b2010
	// 交换后b的内存地址: c0000b2008

	fmt.Printf("交换后a的value: %d\n", *a)
	fmt.Printf("交换后b的value: %d\n", *b)
	// 交换后a的value: 200
	// 交换后b的value: 100
	// 可以看到确实可以通过交互内存地址达到目的

	return *a, *b
	*/

	// 简写方式
	return *b, *a
}

//
func swapVal3(a, b int) {
	// https://pkg.go.dev/golang.org/x/tools/internal/typesinternal#UnassignableOperand
	// 看样子是不允许这样操作 -- 交换内存地址
	// &a, &b = &b, &a
}
