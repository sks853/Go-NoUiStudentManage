/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-02-11 19:11
 * @FileName     tool.go
 * @Description  None
 * ==================================================
**/

package public

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
	"unsafe"
)

var openDebug = false
var socket net.PacketConn
var socketDst *net.UDPAddr

func FuncChangeModeDebug(isDebug bool) {
	openDebug = isDebug
}

func InitUdpClient() {
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	dst, err := net.ResolveUDPAddr("udp", "127.0.0.1:32111")
	if err != nil {
		log.Fatal(err)
	}
	socketDst = dst
	socket = conn
}

func FuncStringToByteSlice(s string) []byte {
	tmp1 := (*[2]uintptr)(unsafe.Pointer(&s))
	tmp2 := [3]uintptr{tmp1[0], tmp1[1], tmp1[1]}
	return *(*[]byte)(unsafe.Pointer(&tmp2))

}

func FuncPrintLog(level byte, info string, args ...error) {
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
		time.Sleep(3 * time.Second)
	}
	if socket == nil || socketDst == nil {
		fmt.Printf("\n[ERROR] 发送数据时建立链接失败，套接字对象为空，数据内容：%s\n", info)
		time.Sleep(10 * time.Second)
		return
	}
	switch level {
	case LogErrs:
		data := "[ERROR] " + info
		if 0 != len(args) {
			data = data + args[0].Error()
		}
		_, err := socket.WriteTo(FuncStringToByteSlice(data), socketDst)
		if err != nil {
			fmt.Printf("\n发送数据失败，原始错误信息：%s，日志无法被记录，err:%v\n", stringErr, err)
			time.Sleep(10 * time.Second)
			return
		}
		break
	case LogWarn:
		data := "[WARN] " + info
		_, err := socket.WriteTo(FuncStringToByteSlice(data), socketDst)
		if err != nil {
			fmt.Printf("\n发送数据失败，日志无法被记录，err:%v\n", err)
			time.Sleep(10 * time.Second)
			return
		}
		break
	case LogInfo:
		data := "[INFO] " + info
		_, err := socket.WriteTo(FuncStringToByteSlice(data), socketDst)
		if err != nil {
			fmt.Printf("\n发送数据失败，日志无法被记录，err:%v\n", err)
			time.Sleep(10 * time.Second)
			return
		}
		break
	default:
	}

}

func Tip(str string) {
	fmt.Print(str)
}

func TipWait(str string, timeSleep ...int) {
	fmt.Print(str)
	timer := 3
	if len(timeSleep) != 0 && 0 < timeSleep[0] {
		timer = timeSleep[0]
	}
	for i := 0; i < timer; i++ {
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
		Tip(Format(TipOpFail, "输入内容非法！请重新输入："))
		return -1
	}
	if rangeMax < option || option < 0 {
		Tip(Format(TipOpFail, "请选择范围内的选项："))
		return -1
	}
	return option
}

func TipInput(tip string, verifier ...func(string) (bool, string)) (string, error) {
	var input string
	for {
		Clear()
		fmt.Println(tip)
		_, err := fmt.Scanln(&input)
		if err != nil {
			TipWait(Format(TipOpFail, "输入内容非法！请重新输入！"))
			return "", err
		}
		if len(verifier) == 0 || verifier[0] == nil {
			return input, nil
		}
		vf := verifier[0]
		vfFlag, tips := vf(input)
		if vfFlag {
			break
		}
		TipWait(tips)
	}
	return input, nil
}

func Format(format string, strList ...string) string {
	var builder strings.Builder
	countLengthStrList := 0
	lenList := len(strList)
	for _, str := range strList {
		countLengthStrList += len(str)
	}
	// 计算容量并预分配
	builder.Grow(countLengthStrList + len(format) - lenList*2)
	count := 0
	start := 0
	lastByte := int32(format[0])
	for i, bt := range format {
		if i == 0 {
			continue
		}
		if count == lenList {
			builder.WriteString(format[i:])
			break
		}
		if lastByte == '%' && bt == 's' {
			for j := start; j < i-1; j++ {
				builder.WriteByte(format[j])
			}
			//builder.WriteString(format[start : i-1])
			builder.WriteString(strList[count])
			count++
			start = i + 1
		}
		lastByte = bt
	}
	return builder.String()
}

func FuncMkDir(path string, dirName string) error {
	builder := strings.Builder{}
	builder.Grow(len(path) + len(dirName)*2)
	builder.WriteString(path)
	for _, bt := range dirName {
		builder.WriteByte('/')
		builder.WriteString(string(bt))
	}
	err := os.MkdirAll(builder.String(), 0755)
	if err != nil {
		FuncPrintLog(LogErrs, Format("创建文件夹{%s}时出现异常", builder.String()), err)
		return err
	}
	return nil
}

func FuncIsExistFile(pathDir string) (bool, error) {
	_, err := os.Stat(pathDir)
	if err != nil {
		fmt.Print(err)
		return false, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
