/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-02-11 15:42
 * @FileName     main.go
 * @Description  Start main function
 * ==================================================
**/

package main

import (
	"NoUiStudentManage/public"
)

func main() {
	public.FuncChangeModeDebug(true)
	public.FuncPrintLog(public.LogInfo, "Start running...")
	menu := InitTreeMenu()
	nodeList := &public.StructMenuLink{Node: menu}
	FuncPrintMenu(*menu)
	for {
		option := public.SelectOption(len(menu.MenuNode))
		if option < 0 {
			continue
		}
		if option == 0 {
			if nodeList.NodeLast != nil {
				nodeList = nodeList.NodeLast
				menu = nodeList.Node
			} else {
				public.Clear()
				public.TipWait("正在退出")
				public.Clear()
				return
			}
		} else {
			if nodeList.NodeNext == nil {
				node := &public.StructMenuLink{NodeLast: nodeList}
				nodeList = node
			}
			menu = menu.MenuNode[option-1]
			nodeList.Node = menu
		}
		FuncPrintMenu(*menu)
	}
}
