/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-09-05 12:02
 * @FileName     menu_op.go
 * @Description  None
 * ==================================================
**/

package main

import (
	"NoUiStudentManage/public"
	"fmt"
	"github.com/gookit/color"
)

func FuncPrintMenu(menu public.StructMenu, share *public.StructShareBase) {
	public.Clear()
	funcPrintProfile(share.Profile)
	fmt.Println("----- ----- ----- ----- -----")
	if menu.Title != "" {
		color.Red.Printf("%s\n", menu.Title)
		fmt.Println("----- ----- ----- ----- -----")
	}
	fmt.Printf("%d. %s\n", 0, "返回上一级")
	if len(menu.MenuNode) != 0 {
		for i, node := range menu.MenuNode {
			fmt.Printf("%d. %s\n", i+1, node.Text)
		}
	}
	fmt.Println("----- ----- ----- ----- -----")
}

func funcPrintProfile(profile *public.StructProfile) {
	color.Blue.Printf("%s\n", "欢迎使用教学管理系统")
	if profile.IsLogin {
		color.Red.Printf("%s\t", profile.UserId)
		color.Yellow.Printf("%s\n", profile.Name)
	}
}

func funcIsPermitShowMenu(userPermission int, menu *public.StructMenu) bool {
	if 0 == len(menu.PermitList) {
		return false
	}
	for permit := range menu.PermitList {
		if permit == userPermission {
			return true
		}
	}
	return false
}
