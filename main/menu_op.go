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

func FuncMenuPrint(menu *public.StructMenu, share *public.StructShareBase) {
	public.Clear()
	funcPrintProfile(share.Profile)
	fmt.Println("----- ----- ----- ----- -----")
	if menu.Title != "" {
		color.Red.Printf("%s\n", menu.Title)
		fmt.Println("----- ----- ----- ----- -----")
	}
	fmt.Printf("%d. %s\n", 0, "返回上一级")
	funcMenuCheckPermit(share.Profile.Permit, menu)
	if len(menu.MenuNode) != 0 {
		for i, node := range menu.MenuNode {
			if node == nil {
				continue
			}
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

func funcMenuCheckPermit(userPermission int, menu *public.StructMenu) {
	newMenu := &public.StructMenu{}
	for i := range menu.MenuNode {
		if menu.MenuNode[i] == nil {
			continue
		}
		ok := funcIsPermitMenu(userPermission, menu.MenuNode[i].PermitMode, menu.MenuNode[i].PermitList)
		if ok {
			newMenu.MenuNode = append(newMenu.MenuNode, menu.MenuNode[i])
		}
	}
	menu.MenuNode = newMenu.MenuNode
	newMenu.MenuNode = nil
}
func funcIsPermitMenu(userPermit int, mode int, permitList []int) bool {
	if 0 == len(permitList) {
		return false
	}
	if userPermit == 0 {
		return false
	}
	switch mode {
	case public.PermitModeEqual:
		for _, permit := range permitList {
			if permit == userPermit {
				return true
			}
		}
		return false
	case public.PermitModeEqualLess:
		for _, permit := range permitList {
			if userPermit < permit {
				return false
			}
		}
		return true
	case public.PermitModeEqualGreater:
		for _, permit := range permitList {
			if permit < userPermit {
				return false
			}
		}
		return true
	case public.PermitModeLess:
		for _, permit := range permitList {
			if userPermit <= permit {
				return false
			}
		}
		return true
	case public.PermitModeGreater:
		for _, permit := range permitList {
			if permit <= userPermit {
				return false
			}
		}
		return true
	default:
	}
	return false
}

func funcDataLogin(userId string, password string, profile *public.StructProfile) (bool, error) {
	// TODO 验证登陆情况
	if userId == "root" && password == "root" {
		profile.IsLogin = true
		profile.Name = "超级管理员"
		profile.UserId = "root"
		profile.Permit = public.PermitAdministrator
		return true, nil
	} else if userId == "admin" && password == "admin" {
		profile.IsLogin = true
		profile.Name = "最靓的崽"
		profile.UserId = "admin"
		profile.ClassId = "202201234"
		profile.SubjectId = "LyGPC"
		profile.Permit = public.PermitManager
		return true, nil
	} else if userId == "user" && password == "user" {
		profile.IsLogin = true
		profile.Name = "张翼德"
		profile.UserId = "2022070230421"
		profile.ClassId = "202201234"
		profile.Permit = public.PermitUser
		return true, nil
	}
	return false, nil
}
