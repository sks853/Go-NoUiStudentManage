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
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
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

func FuncByteSliceToString(bt []byte) string {
	return *(*string)(unsafe.Pointer(&bt))
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

// Clear 	清除控制台屏幕显示
// return 	None
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

// SelectOption		引导用户进行选项选择
// rangeMax:		int		可选的最大范围（默认0为可选最低范围）
// return -1		int		用户选择超出范围
// return ...		int		用户选择的有效选项
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

// Format	格式化输入，提供类似`Sprintf()`的效果，但只允许字符串作为填充便于统一
// format:	string		待格式化的字符串
// strList:	...string	待填充占位符的内容
// return 	string		格式化后的字符串
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

func errF(errMsgType string, errOperation string, otherMsg string, errMsg string) error {
	return errors.New(Format(errMsgTypeBase, errMsgType, errOperation, otherMsg, errMsg))
}

func StrStrip(str string, blockString string) string {
	if str == "" || blockString == "" {
		return str
	}
	index := 0
	block := []byte(blockString)[0]
	for i := 0; i < len(str); i++ {
		if str[i] == block {
			break
		}
		index += 1
	}
	return str[0:index]
}

//func FuncMkDir(path string, dirName string) error {
//	builder := strings.Builder{}
//	builder.Grow(len(path) + len(dirName)*2)
//	builder.WriteString(path)
//	for _, bt := range dirName {
//		builder.WriteByte('/')
//		builder.WriteString(string(bt))
//	}
//	err := os.MkdirAll(builder.Print(), 0755)
//	if err != nil {
//		FuncPrintLog(LogErrs, Format("创建文件夹{%s}时出现异常", builder.Print()), err)
//		return err
//	}
//	return nil
//}

// FuncIsExistFile 	判断文件是否存在 |
// pathDir:			string	文件路径 |
// return 			error	异常 |
// return true		bool	当且仅当 `error` 为空时表示文件存在 |
// return false		bool	当且仅当 `error` 为空时表示文件不存在 |
//func FuncIsExistFile(pathDir string) (error, bool) {
//	_, err := os.Stat(pathDir)
//	if err != nil {
//		fmt.Print(err)
//		return nil, false
//	}
//	if os.IsNotExist(err) {
//		return nil, false
//	}
//	return err, false
//}

// ScoreRecordSum 	求分数项之和 |
// score1:		*StructScoreRecord		第一个待求项 |
// score: 		...*StructScoreRecord		任意个数待求和项 |
// return		error				异常 |
// return		*StructScoreRecord		提供的参数的求和结果 |
func ScoreRecordSum(score1 *StructScoreRecord, score ...*StructScoreRecord) (error, *StructScoreRecord) {
	if score1 == nil {
		return errors.New("传入分数结构为空"), new(StructScoreRecord)
	} else if len(score) <= 0 {
		return errors.New("传入待求和列表为空"), score1
	}
	tmpScore := &StructScoreRecord{
		Score: struct {
			Integer int
			Decimal [DecLength]int
		}{Integer: score1.Score.Integer, Decimal: score1.Score.Decimal},
	}
	for i := 0; i < DecLength; i++ {
		if tmpScore.Score.Decimal[i] < 0 || tmpScore.Score.Decimal[i] > 9 {
			FuncPrintLog(LogErrs, Format("出现小数某一位大于九或小于零的情况，出错数为%s", score1.ToString()))
			return errors.New(Format("出现小数某一位大于九或小于零的情况，出错数为%s", score1.ToString())), tmpScore
		}
	}
	for i := 0; i < len(score); i++ {
		tmpScore.Score.Integer += score[i].Score.Integer
		flagCarry := false
		for j := DecLength - 1; j > -1; j-- {
			if score[i].Score.Decimal[j] < 0 || score[i].Score.Decimal[j] > 9 {
				FuncPrintLog(LogErrs, Format("出现小数某一位大于九或小于零的情况，出错数为%s", score[i].ToString()))
				return errors.New(Format("出现小数某一位大于九或小于零的情况，出错数为%s", score1.ToString())), tmpScore
			}
			tmpScore.Score.Decimal[j] += score[i].Score.Decimal[j]
			if flagCarry {
				tmpScore.Score.Decimal[j] += 1
				flagCarry = false
			}
			if tmpScore.Score.Decimal[j] >= 10 {
				flagCarry = true
				tmpScore.Score.Decimal[j] -= 10
			}
		}
		if flagCarry {
			tmpScore.Score.Integer += 1
		}
	}
	return nil, tmpScore
}

// ScoreRecordAverage 【不建议使用】求分数平均值（警告：因该函数操作存在误差，请勿将该函数用于任何数值比较，仅用于显示不精确的数值） |
// score1:		*StructScoreRecord		第一个求平均的数 |
// score:		...*StructScoreRecord		任意长度的待求平均的数 |
// return		error				异常 |
// return		float64				该列表（含第一个参数）的平均值（强制保留两位作为被除数） |
func ScoreRecordAverage(score1 *StructScoreRecord, score ...*StructScoreRecord) (error, float64) {
	if len(score) <= 0 {
		return nil, score1.ToFloat64(-1)
	}
	flag, tmpScore := ScoreRecordSum(score1, score...)
	return flag, tmpScore.ToFloat64(2) / float64(len(score)+1)
}

// ScoreCompare 	比较两个分数列表总分大小，返回比较结果请以`ScoreRecordSum()`为准 |
// scoreList1: 		[]*StructScoreRecord	第一个待比较列表 |
// scoreList2: 		[]*StructScoreRecord	第二个待比较列表 |
// isOrder:			bool			总分等值情况下是否按顺序比较（若两者总分相等，则按排列项顺序依次比较大小） |
// return 			error 			异常 |
// return 0 		int				前者列表等于后者列表 |
// return -1		int				前者列表小于后者列表 |
// return 1 		int				前者列表大于后者列表 |
func ScoreCompare(scoreList1 []*StructScoreRecord, scoreList2 []*StructScoreRecord, isOrder bool) (error, int) {
	if len(scoreList1) != len(scoreList2) {
		return errors.New("比较分数项数量不同，无法进行比较"), 0
	}
	tmp := new(StructScoreRecord)
	flag1, scoreSum1 := ScoreRecordSum(tmp, scoreList1...)
	flag2, scoreSum2 := ScoreRecordSum(tmp, scoreList2...)
	if flag1 != nil || flag2 != nil {
		return errors.New(Format("在分数比较过程中无法正确处理求和操作，出错信息：%s或%s", flag1.Error(), flag2.Error())), 0
	}
	if isOrder {
		for i := 0; i < len(scoreList1); i++ {
			rs := scoreList1[i].Compare(scoreList2[i])
			if rs != 0 {
				return nil, rs
			}
		}
	}
	return nil, scoreSum1.Compare(scoreSum2)
}

// CompareByteList	比较两个 Byte[] 的大小 |
// s1: 				[]byte	字节列表前者 |
// s2: 				[]byte	字节列表后者 |
// return -1		int		两者长度不一样，无法比较 |
// return 0			int		两者完全相等 |
// return 1			int		前者小于后者 |
// return 2			int		前者大于后者 |
func CompareByteList(s1 []byte, s2 []byte, block byte) int {
	if len(s1) != len(s2) {
		return -1
	}
	for i, bt1 := range s1 {
		if bt1 == block && s2[i] == block {
			break
		} else if bt1 == block && s2[i] != block {
			return 1
		} else if bt1 != block && s2[i] == block {
			return 2
		}
		if bt1 < s2[i] {
			return 1
		} else if bt1 > s2[i] {
			return 2
		}
	}
	return 0
}

// CountRowLength	统计数据每一行的长度，若提供的`keyColumn`有效，则追加计算关键列前和关键列的数据长度 |
// keyColumn:		int		（从第1列始）关键列，若有效则计算该列前的数据长度和该列的数据长度 |
// columnLength:	[]int	每个字段的长度 |
// return			error	异常信息 |
// return			int		数据每一行的最大长度，包含特殊结尾`\r\n` |
// return			int		每一行关键列前的数据长度（若`keyColumn`有效） |
// return			int		关键列的数据长度（若`keyColumn`有效） |
func CountRowLength(keyColumn int, columnLength []int) (error, int, int, int) {
	rowLength := 0
	keyFrontLength := 0
	for i := 0; i < len(columnLength); i++ {
		if i < keyColumn-1 {
			keyFrontLength += columnLength[i] + 1
		}
		rowLength += columnLength[i] + 1
	}
	rowLength += 1
	if rowLength <= 0 {
		return errors.New("待排数据单行长度小于等于0"), 0, 0, 0
	}
	if 0 < keyColumn && keyColumn <= len(columnLength) {
		return nil, rowLength, keyFrontLength, columnLength[keyColumn-1]
	}
	return nil, rowLength, 0, 0
}

// swapRow		交换指定行（注意：操作会改变文件内容） |
// fs			*os.File	文件指针（可写） |
// r1			int64		待交换行 |
// r2			int64		待交换行 |
// rowLength	int64		行长度 |
// return 		None |
func swapRow(fs *os.File, r1 int64, r2 int64, rowLength int64) error {
	btFront := make([]byte, rowLength)
	btTail := make([]byte, rowLength)
	_, err := fs.ReadAt(btFront, rowLength*r1)
	if err != nil {
		return err
	}
	_, err = fs.ReadAt(btTail, rowLength*r2)
	if err != nil {
		return err
	}
	//fmt.Println("---------- ---------- ---------- ----------")
	//fmt.Println("【单行长度】 ", rowLength)
	//fmt.Println("【待交换行】 ", r1, " && ", r2)
	//fmt.Print("【前行内容】  ", *(*string)(unsafe.Pointer(&btFront)))
	//fmt.Print("【后行内容】  ", *(*string)(unsafe.Pointer(&btTail)))
	//fmt.Println("【前行字节】 ", btFront)
	//fmt.Println("【后行字节】 ", btTail)
	//fmt.Println("---------- ---------- ---------- ----------")
	_, err = fs.WriteAt(btTail, rowLength*r1)
	if err != nil {
		return err
	}
	_, err = fs.WriteAt(btFront, rowLength*r2)
	if err != nil {
		return err
	}
	return nil
}

// SortBubble	 	冒泡排序文件指定列，适用于基本有序的序列（注意：操作会改变文件内容） |
// fs:				*os.File	文件指针，可读可写 |
// countRow:		int64		文件总共有效行数 |
// columnLength:	[]int		每个字段的长度 |
// sortField:		int			待排序的列，从`1`开始 |
// return			error		异常 |
func SortBubble(fs *os.File, countRow int64, columnLength []int, sortField int) error {
	const errOperation = "冒泡排序"
	if fs == nil {
		return errF(errMsgTypeFileNull, errOperation, "", "")
	}
	ptrRowStart := int64(0)
	ptrRowEnd := countRow - 1
	flag := false
	if sortField <= 0 || len(columnLength) < sortField {
		return errF(errMsgTypeSorSelectColumn, errOperation, "", "")
	}
	err, rowLength, keyFrontLength, keyColumnLength := CountRowLength(sortField, columnLength)
	if err != nil {
		return err
	}
	for ; ptrRowStart < countRow-1; ptrRowStart++ {
		flag = false
		ptrRowEnd = countRow - 1
		for ; ptrRowEnd > ptrRowStart; ptrRowEnd-- {
			// 待读取的长度，长度为该列长度
			btFront := make([]byte, keyColumnLength)
			btTail := make([]byte, keyColumnLength)
			// 分别读取两行，定位到某一行的前端，根据行长乘以行数定位，并根据偏移（前n列字符和）得到位置
			_, err1 := fs.ReadAt(btFront, int64(rowLength)*(ptrRowEnd-1)+int64(keyFrontLength))
			_, err2 := fs.ReadAt(btTail, int64(rowLength)*(ptrRowEnd)+int64(keyFrontLength))
			if err1 != nil || err2 != nil {
				e1 := strconv.Itoa(rowLength)
				e2 := strconv.FormatInt(ptrRowStart, 10)
				e3 := strconv.FormatInt(ptrRowEnd, 10)
				e4 := strconv.Itoa(keyFrontLength)
				eSum := Format("行长度=`%s`, 循环轮次=`%s`，在处理行=`%s`，关键列前长度=`%s`", e1, e2, e3, e4)
				errs := "无"
				if err1 != nil && err2 != nil {
					errs = "异常1：" + err1.Error() + "，" + "异常2：" + err2.Error()
				} else if err1 != nil {
					errs = err1.Error()
				} else if err2 != nil {
					errs = err2.Error()
				}
				return errF(errMsgTypeOffsetGetData, errOperation, eSum, errs)
			}
			// 比较，前者大于后者便交换
			if 2 == CompareByteList(btFront, btTail, '?') {
				err = swapRow(fs, ptrRowEnd-1, ptrRowEnd, int64(rowLength))
				if err != nil {
					return errF(errMsgTypeWriteData, errOperation, "", err.Error())
				}
				flag = true
			}
		}
		// 如果某一轮基本有序就无需再排
		if !flag {
			return nil
		}
	}
	return nil
}

// SearchBinary 	【必须保证待查列有序】 二分查找指定数据列的数据，返回多个结果，具体参见返回值 |
// fs:				*os.File	文件指针，可读 |
// countRow:		int64		文件总共有效行数 |
// columnLength:	[]int		每个字段的长度 |
// column:			int			待查找的列，从`1`开始 |
// value:			string		待查找的值，空缺位必须包含填充符，长度必须与该列长度相匹配 |
// blockFill:		string		空缺位填充符，仅第一位有效 |
// return			error		异常 |
// return false		bool		当且仅当 `error` == nil 时表示未找到该数据 |
// return true		bool		当且仅当 `error` == nil 时表示成功找到该数据 |
// return 			[]string	当且仅当 `true` 时表示找到数据后的按行分隔的数据 |
// return 			int64		当且仅当 `true` 时表示在该排序下、该列中，该被查找到的元素所在的行数，从`1`开始 |
func SearchBinary(fs *os.File, countRow int64, columnLength []int, column int, value string, blockFill string) (error, bool, []string, int64) {
	const errOperation = "二分查找"
	if fs == nil {
		return errF(errMsgTypeFileNull, errOperation, "", ""), false, []string{}, 0
	}
	_, err := fs.Seek(0, 0)
	if err != nil {
		return errF(errMsgTypeOffsetHead, errOperation, "", err.Error()), false, []string{}, 0
	}
	// 初始化左指针和右指针，即范围，以及中间指针
	var ptrRowL, ptrRowR, ptrRowM int64
	ptrRowL, ptrRowR = 1, countRow
	// 计算单行长度以及该`column`列前数据长度
	err, rowLength, keyFrontLength, keyColumnLength := CountRowLength(column, columnLength)
	if err != nil {
		return errF(errMsgTypeCalcLength, errOperation, "", err.Error()), false, []string{}, 0
	}
	btValue := []byte(value)
	for ptrRowL <= ptrRowR {
		ptrRowM = ptrRowL + (ptrRowR-ptrRowL)/2
		// 已知前n个列共有长度为`keyFrontLength`，第n+1列固定长度为`keyColumnLength`，读取第`ptrRowM`行第n+1列
		bt := make([]byte, keyColumnLength)
		_, err = fs.ReadAt(bt, int64(rowLength)*(ptrRowM-1)+int64(keyFrontLength))
		if err != nil {
			errMsg := Format("当前偏移初始位置是%s", strconv.FormatInt(int64(rowLength)*(ptrRowM-1)+int64(keyFrontLength), 10))
			return errF(errMsgTypeOffsetGetData, errOperation, errMsg, err.Error()), false, []string{}, 0
		}
		rs := CompareByteList(bt, btValue, []byte(blockFill)[0])
		switch rs {
		case 0:
			// 获取该行内容
			btContent := make([]byte, rowLength)
			_, err = fs.ReadAt(btContent, int64(rowLength)*(ptrRowM-1))
			strList := make([]string, len(columnLength))
			queueH, queueT := 0, 0
			for i := 0; i < len(columnLength); i++ {
				queueT += columnLength[i]
				strList[i] = string(btContent[queueH:queueT])
				queueH = queueT + 1
				queueT += 1
			}
			return nil, true, strList, ptrRowM
		case 1:
			ptrRowL = ptrRowM + 1
			break
		case 2:
			ptrRowR = ptrRowM - 1
			break
		default:
			return errF(errMsgTypeArgsLengthDifferent, errOperation, "", ""), false, []string{}, 0
		}
	}
	return nil, false, []string{}, 0
}

func InsertData(df *StructDataFile, fsPath string, fsName string, keyColumn int, keyValue []string) error {
	const errOperation = "标准插入数据"
	// 打开文件并载入`df.File`
	err := df.GetFile(fsPath, fsName, true)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "获取用户信息文件失败，在追加写模式中", err.Error())
	}
	defer func(File *os.File) {
		_ = File.Close()
	}(df.File)
	// 检查参数是否长度符合
	if len(keyValue) != len(df.ColumnLength) {
		_ = df.File.Close()
		return errF(errMsgTypeArgsLengthDifferent, errOperation, "判断传入列表和提供的要求列表长度", "")
	}
	var data string
	// 填充并转化
	for i := 0; i < len(keyValue); i++ {
		var valueKey string
		err, valueKey = FillBlock(keyValue[i], df.ColumnLength[i], df.BlockFill[0])
		if err != nil {
			_ = df.File.Close()
			return errF("填充内容出错", errOperation, "填充用户信息的关键列中", err.Error())
		}
		data += valueKey
		if i == len(keyValue)-1 {
			data += "\r\n"
		} else {
			data += ","
		}
	}
	err = df.DataWrite(-1, -1, data, true)
	if err != nil {
		_ = df.File.Close()
		return errF(errMsgTypeViewError, errOperation, "写入数据文件过程中", err.Error())
	}
	_ = df.File.Close()
	df.DataConfigSetRowCount(df.CountRow + 1)
	err = df.GetFile(fsPath, fsName, false)
	if err != nil {
		_ = df.File.Close()
		return errF(errMsgTypeViewError, errOperation, "获取用户信息文件失败，在标准读写模式中", err.Error())
	}
	err = SortBubble(df.File, df.CountRow, df.ColumnLength, keyColumn)
	if err != nil {
		_ = df.File.Close()
		return errF("", errOperation, "", err.Error())
	}
	_ = df.File.Close()
	return nil
}

func DataFileGet(pathDir string, pathFile string, keyColumn int, keyValue string) (error, bool, []string) {
	const errOperation = "文件数据获取"
	df := new(StructDataFile)
	err := df.GetFile(pathDir, pathFile, false)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "获取用户验证文件失败", err.Error()), false, nil
	}
	defer func(File *os.File) {
		_ = File.Close()
	}(df.File)
	err, valueKey := FillBlock(keyValue, df.ColumnLength[keyColumn-1], df.BlockFill[0])
	if err != nil {
		return errF("填充内容出错", errOperation, "填充关键列中", err.Error()), false, nil
	}
	err = SortBubble(df.File, df.CountRow, df.ColumnLength, keyColumn)
	if err != nil {
		_ = df.File.Close()
		return errF("", errOperation, "", err.Error()), false, nil
	}
	err, flag, data, _ := SearchBinary(df.File, df.CountRow, df.ColumnLength, keyColumn, valueKey, df.BlockFill)
	if err != nil {
		_ = df.File.Close()
		return errF(errMsgTypeViewError, errOperation, "", err.Error()), false, nil
	}
	if !flag {
		_ = df.File.Close()
		return nil, false, nil
	}
	return nil, true, data
}

func DeleteData(df *StructDataFile, fsPath string, fsName string, row int64) error {
	// TODO Delete A Record From Data File
	return nil
}

func DataFileSet(filePath string, fileName string, keyValue string, keyColumn int, setValueArray []string, setColumnArray []int) error {
	const errOperation = "数据文件设定"
	if len(setValueArray) != len(setColumnArray) {
		return errF(errMsgTypeArgsLengthDifferent, errOperation, "数据文件设定比较参数个数过程中", "")
	}
	df := new(StructDataFile)
	err := df.GetFile(filePath, fileName, false)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "获取用户数据文件失败", err.Error())
	}
	err, keyValue = FillBlock(keyValue, df.ColumnLength[keyColumn-1], df.BlockFill[0])
	if err != nil {
		df.FileClose()
		return errF("填充内容出错", errOperation, "填充数据文件的关键列中", err.Error())
	}
	for i := 0; i < len(setValueArray); i++ {
		err, setValueArray[i] = FillBlock(setValueArray[i], df.ColumnLength[setColumnArray[i]-1], df.BlockFill[0])
		if err != nil {
			df.FileClose()
			return errF("填充内容出错", errOperation, "填充数据文件的设定列中", err.Error())
		}
	}
	// 对关键列排序后，使用二分查找指定的内容，获取行数
	err = SortBubble(df.File, df.CountRow, df.ColumnLength, keyColumn)
	if err != nil {
		df.FileClose()
		return errF("冒泡排序出错", errOperation, "排序过程中", err.Error())
	}
	err, flag, _, row := SearchBinary(df.File, df.CountRow, df.ColumnLength, keyColumn, keyValue, df.BlockFill)
	if err != nil {
		df.FileClose()
		return errF("二分查找失败", errOperation, "二分查找数据文件过程中", err.Error())
	}
	if !flag {
		df.FileClose()
		return errF("二分查找未找到结果", errOperation, "查找并修改数据文件中", "")
	}
	for i := 0; i < len(setValueArray); i++ {
		// 将行、列值写入对应位置
		err = df.DataWrite(row, setColumnArray[i], setValueArray[i], false)
		if err != nil {
			df.FileClose()
			return errF(errMsgTypeViewError, errOperation, "写入数据文件过程中", err.Error())
		}
	}
	return nil
}

func FillBlock(value any, lenStr int, block byte) (error, string) {
	var valueStr string
	switch (value).(type) {
	case int:
		valueStr = strconv.Itoa(value.(int))
		break
	case string:
		valueStr = value.(string)
		break
	default:
		return errors.New("提供的待设定参数其数据类型不在可选项中"), ""
	}
	var bt []byte
	if len(valueStr) < lenStr {
		// 先填充符号再进一步处理
		bt = make([]byte, lenStr)
		for i := 0; i < len(valueStr); i++ {
			bt[i] = valueStr[i]
		}
		for i := len(valueStr); i < len(bt); i++ {
			bt[i] = block
		}
	} else if len(valueStr) == lenStr {
		// 长度一致则直接从`string`转`[]byte`
		return nil, valueStr
	} else {
		return errors.New("提供的`value`长度大于实际提供的配置长度要求"), ""
	}
	return nil, FuncByteSliceToString(bt)
}

func GenerateUpperUUID(key string) string {
	var uid string
	if key == "" {
		uid = uuid.NewV4().String()
	} else {
		uid = uuid.NewV5(uuid.NamespaceX500, key).String()
	}
	bt := FuncStringToByteSlice(uid)
	btUid := make([]byte, 32)
	ptrBtUid := 0
	for i := 0; i < len(bt); i++ {
		if bt[i] == '-' {
			continue
		}
		btUid[ptrBtUid] = bt[i]
		ptrBtUid++
	}
	return strings.ToUpper(FuncByteSliceToString(btUid))
}

func (date *StructDate) CheckDate() bool {
	if 2300 < date.Y || date.Y < 1976 {
		return false
	}
	if 12 < date.M || date.M < 1 {
		return false
	}
	if date.M == 1 || date.M == 3 || date.M == 5 || date.M == 7 || date.M == 8 || date.M == 10 || date.M == 12 {
		if 31 < date.D || date.D < 1 {
			return false
		} else {
			return true
		}
	} else if date.M == 4 || date.M == 6 || date.M == 9 || date.M == 11 {
		if 30 < date.D || date.D < 1 {
			return false
		} else {
			return true
		}
	} else if date.M == 2 {
		if date.D == 29 && ((date.Y%100 == 0 && date.Y%400 == 0) || (date.Y%4 == 0)) {
			return true
		} else if date.D < 29 {
			return true
		}
	}
	return false
}
