/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-09-05 12:02
 * @FileName     text_exception.go
 * @Description  None
 * ==================================================
**/

package main

import (
	"NoUiStudentManage/public"
	"fmt"
)

func testExceptionFunc() {
	// 预定义一个`defer`标志的函数，函数内抛出异常后会跳转到这里，通过`recover`使得程序继续执行
	defer func() {
		if err := recover(); err != nil {
			public.FuncPrintLog(public.LogErrs, "Error in test", nil)
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
