package test

import (
	"NoUiStudentManage/public"
	"NoUiStudentManage/strings"
	"fmt"
)

func testExceptionFunc() {
	// 预定义一个`defer`标志的函数，函数内抛出异常后会跳转到这里，通过`recover`使得程序继续执行
	defer func() {
		if err := recover(); err != nil {
			public.FuncPrintLog(strings.LogErrs, "Error in test", err)
		}
	}()
	var x interface{} = "Hello"
	// 断言产生的异常
	value := x.(int)
	// 不会执行的后续部分
	fmt.Println(value)
	fmt.Println("Next test")
}

func Test() {
	testExceptionFunc()
	fmt.Println("Finish test")
}
