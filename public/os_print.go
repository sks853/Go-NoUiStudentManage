/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-02-11 19:11
 * @FileName     os_print.go
 * @Description  None
 * ==================================================
**/

package public

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var openDebug = true

func FuncChangeModeDebug(isDebug bool) {
	openDebug = isDebug
}

func FuncPrintLog(level byte, info string, args ...any) {
	var stringErr = "Undefined EOF..."
	if 0 < len(args) && args[0] != nil {
		stringErr = fmt.Sprintf("%v", args[0])
	}
	if openDebug {
		if 0 == len(args) {
			fmt.Printf("[DEBUG] [INFO] %s\n", info)
		} else {
			fmt.Printf("[DEBUG] [ERRS] %s --> %s\n", info, stringErr)
		}
	}
	switch level {
	case LogErrs:
		break
	case LogWarn:
		break
	case LogInfo:
		break
	default:

	}
}

func Tip(str string) {
	fmt.Print(str)
}

func TipWait(str string) {
	fmt.Print(str)
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}

func Clear() {
	var cmd *exec.Cmd
	switch SysType {
	case "linux":
		cmd = exec.Command("clear")
		break
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
		break
	default:
		FuncPrintLog(LogWarn, "没有对应的系统类型，无法清屏")
		return
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		FuncPrintLog(LogWarn, "操作系统无法完全匹配，无法清屏")
	}
}

func SelectOption(rangeMax int) int {
	var option int
	Tip("请选择菜单选项：")
	_, err := fmt.Scanln(&option)
	if err != nil {
		Tip(fmt.Sprintf(TipOpFail, "输入内容非法！请重新输入："))
		return -1
	}
	if rangeMax < option || option < 0 {
		Tip(fmt.Sprintf(TipOpFail, "请选择范围内的选项："))
		return -1
	}
	return option
}
