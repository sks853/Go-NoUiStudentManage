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
	"time"
)

func main() {
	initMain()
	startMain()
}

func initMain() {
	go ServerRun()
	time.Sleep(2 * time.Second)
	public.InitUdpClient()
	public.FuncChangeModeDebug(false)
	public.FuncPrintLog(public.LogInfo, "程序开始运行...")
}

func startMain() {
	profile := public.StructProfile{IsLogin: false, Permit: public.PermitGuest}
	share := public.StructShareBase{Profile: &profile}
	menu := InitTreeMenu(&share)
	nodeList := &public.StructMenuLink{Node: menu}
	FuncMenuPrint(menu, &share)
	for {
		option := public.SelectOption(len(menu.MenuNode))
		if option < 0 {
			continue
		}
		if option == 0 {
			// 若上一级不为空则将上一级的链点赋给当前链点，并传递菜单属性，否则退出并结束程序
			if nodeList.NodeLast != nil {
				nodeList = nodeList.NodeLast
				menu = nodeList.Node
			} else {
				public.FuncPrintLog(public.LogInfo, "程序正在退出...")
				public.Clear()
				public.TipWait("正在退出")
				public.Clear()
				return
			}
		} else {
			// 判断下一链点是否为空，如果为空则创建链点，并将下一链点的`NodeLast`指向为当前链点，同时`nodeList`指针指向下一链点
			if nodeList.NodeNext == nil {
				node := &public.StructMenuLink{NodeLast: nodeList}
				nodeList = node
			}
			// 对当前菜单指向用户所选菜单选项，并将指向后的菜单属性赋予`nodeList`链点的菜单属性`Node`
			menu = menu.MenuNode[option-1]
			nodeList.Node = menu
			// 执行当前菜单附带的操作函数，若需要验证则传入验证操作量，通过后允许进入菜单，否则返回上一级菜单，或者保持当前菜单
			if menu.Func != nil {
				if menu.HasVerifier {
					vfResult := false
					menu.Func(&share, &vfResult)
					if !vfResult {
						nodeList = nodeList.NodeLast
						menu = nodeList.Node
					}
				} else if menu.IsKeepMenu {
					menu.Func(&share)
					nodeList = nodeList.NodeLast
					menu = nodeList.Node
				} else {
					menu.Func(&share)
				}
			}
		}
		FuncMenuPrint(menu, &share)
	}
}
