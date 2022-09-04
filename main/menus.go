package main

import (
	"fmt"
	"github.com/gookit/color"
)

type StructMenuOption struct {
	priSelector       []int
	menuOptionContent string
	functionOption    any
}

type StructMenu struct {
	title    string
	name     string
	userId   string
	option   []string
	selector int
}

func FuncPrintMenu(menu StructMenu) {
	color.Red.Printf("\t\t%s\n", menu.title)
	color.Red.Printf("\t%s\t", menu.userId)
	color.Yellow.Printf("\t%s\n", menu.name)
	fmt.Println("===== ===== ===== ===== ===== ===== ===== ===== =====")
	for i := 0; i < len(menu.option); i++ {
		if menu.selector != i {
			fmt.Printf("  %s\n", menu.option[i])
		} else {
			color.Red.Printf("* %s\n", menu.option[i])
		}
	}
}

func FuncIsPermissionShowMenuOption(userPermission int, menu *StructMenuOption) bool {
	permissionList := menu.priSelector
	if 0 == len(permissionList) {
		return false
	}
	for i := 0; i < len(permissionList); i++ {
		if permissionList[i] < userPermission {
			return false
		}
	}
	return true
}
