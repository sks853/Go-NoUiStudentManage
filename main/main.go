/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-02-11 15:42
 * @FileName     main.go
 * @Description  None
 * ==================================================
**/

package main

import (
	"NoUiStudentManage/public"
	"NoUiStudentManage/strings"
)

func main() {
	public.FuncChangeModeDebug(true)
	public.FuncPrintLog(strings.LogInfo, "Start running...", nil)
	var content = []string{"这是第一条很长很长的文字", "这是第二条很长很长的文字", "这是第三条", "This is the forth"}
	var menu = StructMenu{
		title:    "欢迎使用教学管理系统",
		name:     "张翼德",
		userId:   "2022070230421",
		option:   content,
		selector: 2,
	}
	FuncPrintMenu(menu)
}
