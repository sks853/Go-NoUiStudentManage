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
	"NoUiStudentManage/strings"
	"fmt"
)

var openDebug = true

func FuncChangeModeDebug(isDebug bool) {
	openDebug = isDebug
}

func FuncPrintLog(level byte, info string, err any) {
	var stringErr = "None"
	if err != nil {
		stringErr = fmt.Sprintf("%s", err)
	}
	if openDebug {
		if err == nil {
			fmt.Printf("[DEBUG] [INFO] %s\n", info)
		} else {
			fmt.Printf("[DEBUG] [ERRS] %s --> %s\n", info, stringErr)
		}
	}
	switch level {
	case strings.LogErrs:
		break
	case strings.LogWarn:
		break
	case strings.LogInfo:
		break
	default:

	}
}
