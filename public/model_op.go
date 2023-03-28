/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2023-02-05 17:01
 * @FileName     model_op.go
 * @Description  None
 * ==================================================
**/

package public

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ToString 	将分数项转为字符串并返回 |
// return 		string		转换后的 string 分数 |
func (score *StructScoreRecord) ToString() string {
	strScore := strconv.Itoa(score.Score.Integer) + "."
	for i := 0; i < DecLength; i++ {
		strScore += strconv.Itoa(score.Score.Decimal[i])
	}
	return strScore
}

// ToFloat64 	将分数项转为浮点型64位并返回（存在精度差） |
// keep:		int			保留小数位（<0 为保留原始长度） |
// return 		float64		转换后的 float64 分数 |
func (score *StructScoreRecord) ToFloat64(keep int) float64 {
	tmpScore := float64(score.Score.Integer)
	decRate := 0.1
	for i := 0; i < DecLength; i++ {
		tmpScore += float64(score.Score.Decimal[i]) * decRate
		decRate *= 0.1
	}
	if 0 <= keep {
		valueFormat, _ := strconv.ParseFloat(fmt.Sprintf(fmt.Sprintf("%%.%df", keep), tmpScore), 64)
		return valueFormat
	}
	return tmpScore
}

// Compare		用于单项分数的两者比较，返回该项与另一项的比较结果 |
// numOther:	*StructScoreRecord	待比较的分数项 |
// return 0		int				两者相同 |
// return -1	int				当前项小于待比较项 |
// return 1		int				当前项大于待比较项 |
func (score *StructScoreRecord) Compare(scoreOther *StructScoreRecord) int {
	if score.Score.Integer > scoreOther.Score.Integer {
		return 1
	} else if score.Score.Integer < scoreOther.Score.Integer {
		return -1
	}
	for i := 0; i < DecLength; i++ {
		if score.Score.Decimal[i] > scoreOther.Score.Decimal[i] {
			return 1
		} else if score.Score.Decimal[i] < scoreOther.Score.Decimal[i] {
			return -1
		}
	}
	return 0
}

func (df *StructDataFile) Copy(dfSrc *StructDataFile) {
	df.BlockFill = dfSrc.BlockFill
	df.SortField = dfSrc.SortField
	df.CountRow = dfSrc.CountRow
	df.ColumnLength = dfSrc.ColumnLength
	df.Encoding = dfSrc.Encoding
	df.File = dfSrc.File
	df.FilePath = dfSrc.FilePath
	df.FileName = dfSrc.FileName
	df.IsSort = dfSrc.IsSort
}

func (df *StructDataFile) Print() string {
	blockFill := "Block fill: " + df.BlockFill + "\t"
	sortField := "Sort field: " + strconv.Itoa(df.SortField) + "\t"
	countRow := "Count row: " + strconv.FormatInt(df.CountRow, 10)
	columnLength := "Column length: "
	for i := 0; i < len(df.ColumnLength); i++ {
		columnLength += strconv.Itoa(df.ColumnLength[i]) + " "
	}
	columnLength += "\t"
	encoding := "Encoding: " + strconv.Itoa(df.Encoding) + "\t"
	file := "File: "
	if df.File == nil {
		file += "None"
	} else {
		file += "Ready"
	}
	file += "\t"
	filePath := "File path: " + df.FilePath + "\t"
	fileName := "File name: " + df.FileName + "\t"
	isSort := "Is sorted?: " + strconv.FormatBool(df.IsSort) + "\t"
	return blockFill + sortField + countRow + columnLength + encoding + file + filePath + fileName + isSort
}

// GetFile		获取目标数据文件指针并解析配置文件 |
// path:		string	目标文件路径 |
// filename:	string	目标文件名 |
// isAppend:	bool	是否采用只追加模式 |
// return 		error	异常 |
func (df *StructDataFile) GetFile(path string, filename string, isAppend bool) error {
	const errOperation = "获取目标数据文件"
	errConfig := df.DataConfigRead(path + "/" + filename + "_config")
	df.FilePath = path
	df.FileName = filename
	if errConfig != nil {
		return errConfig
	}
	var fs *os.File
	var err error
	if isAppend {
		fs, err = os.OpenFile(path+"/"+filename, os.O_RDWR|os.O_APPEND, 0644)
	} else {
		fs, err = os.OpenFile(path+"/"+filename, os.O_RDWR, 0644)
	}
	if err != nil {
		return errF(errMsgTypeOpenData, errOperation, Format("目标路径：%s", path+"/"+filename), err.Error())
	}
	df.File = fs
	return nil
}

// DataConfigRead 	解析目标文件的配置
// path:			string	目标文件路径含文件名（完整路径）
// return 			error	异常
func (df *StructDataFile) DataConfigRead(path string) error {
	const errOperation = "解析目标配置文件"
	fs, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return errF(errMsgTypeOpenConfig, errOperation, Format("目标路径：%s", path), err.Error())
	}
	defer func(fs *os.File) {
		_ = fs.Close()
	}(fs)
	buf := bufio.NewScanner(fs)
	buf.Scan()
	config := strings.Split(buf.Text(), "#")
	if len(config) != ConfigLength {
		_ = fs.Close()
		return errF(errMsgTypeConfigError, errOperation, Format("目标路径：%s", path), "")
	}
	df.BlockFill = config[0]
	sortField, err1 := strconv.Atoi(config[1])
	df.SortField = sortField
	listFieldLength := strings.Split(config[2], ",")
	df.ColumnLength = make([]int, len(listFieldLength))
	for i := 0; i < len(listFieldLength); i++ {
		df.ColumnLength[i], err = strconv.Atoi(listFieldLength[i])
		if err != nil {
			_ = fs.Close()
			return errF(Format(errMsgTypeCovertInt, "字段长度"), errOperation, Format("目标路径：%s，出错字段下标：%s", path, strconv.Itoa(i)), err.Error())
		}
	}
	countRow, err2 := strconv.Atoi(config[3])
	df.CountRow = int64(countRow)
	if err1 != nil {
		_ = fs.Close()
		return errF(Format(errMsgTypeCovertInt, "排序列"), errOperation, Format("目标路径：%s", path), err1.Error())
	} else if err2 != nil {
		_ = fs.Close()
		return errF(Format(errMsgTypeCovertInt, "行数"), errOperation, Format("目标路径：%s", path), err2.Error())
	}
	switch config[4] {
	case "UTF-8":
		df.Encoding = 0
		break
	case "GBK":
		df.Encoding = 1
		break
	default:
	}
	_ = fs.Close()
	return nil
}

func (df *StructDataFile) DataConfigNew(path string, dataFileName string) error {
	const errOperation = "新建默认配置文件"
	dfNew := &StructDataFile{
		BlockFill:    "?",
		SortField:    1,
		CountRow:     0,
		ColumnLength: []int{32, 32},
		Encoding:     0,
		File:         nil,
		FilePath:     path,
		FileName:     dataFileName,
	}
	df.Copy(dfNew)
	err := df.DataConfigSet(path, dataFileName)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, Format("目标路径：%s", path+"/"+dataFileName+"_config"), err.Error())
	}
	return nil
}

func (df *StructDataFile) DataConfigSet(path string, dataFileName string) error {
	const errOperation = "设置目标配置文件"
	if path == "" {
		path = df.FilePath
	}
	if dataFileName == "" {
		dataFileName = df.FileName
	}
	fs, err := os.OpenFile(path+"/"+dataFileName+"_config", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return errF(errMsgTypeOpenConfig, errOperation, Format("目标路径：%s", path+"/"+dataFileName+"_config"), err.Error())
	}
	strColumnLength := ""
	for i := 0; i < len(df.ColumnLength); i++ {
		if i == len(df.ColumnLength)-1 {
			strColumnLength += strconv.Itoa(df.ColumnLength[i])
			break
		}
		strColumnLength += strconv.Itoa(df.ColumnLength[i]) + ","
	}
	strEncoding := "UTF-8"
	switch df.Encoding {
	case 0:
		strEncoding = "UTF-8"
		break
	case 1:
		strEncoding = "GBK"
		break
	default:
	}
	configData := df.BlockFill + "#" + strconv.Itoa(df.SortField) + "#" + strColumnLength + "#" + strconv.FormatInt(df.CountRow, 10) + "#" + strEncoding + "\r\n"
	_, err = fs.WriteString(configData)
	if err != nil {
		_ = fs.Close()
		return errF(errMsgTypeWriteData, errOperation, Format("目标路径：%s", path+"/"+dataFileName+"_config"), err.Error())
	}
	_ = fs.Close()
	return nil
}

// DataConfigSetBlockFill	改变配置文件空白填充符号 |
// blockFill:				string				填充空白符号 |
// return 					None |
func (df *StructDataFile) DataConfigSetBlockFill(blockFill string) {
	df.BlockFill = blockFill
	err := df.DataConfigSet("", "")
	if err != nil {
		panic(Format("写入文件配置出现问题，异常：%s", err.Error()))
	}
}

// DataConfigSetSortField	改变配置文件优先排序列 |
// sortField:				int					默认排序列 |
// return 					None |
func (df *StructDataFile) DataConfigSetSortField(sortField int) {
	df.SortField = sortField
	err := df.DataConfigSet("", "")
	if err != nil {
		panic(Format("写入文件配置出现问题，异常：%s", err.Error()))
	}
}

// DataConfigSetFieldLength	改变配置文件字段长度 |
// fieldLength:				[]int				每列长度 |
// return 					None |
func (df *StructDataFile) DataConfigSetFieldLength(fieldLength []int) {
	// TODO Change Data File Length And Fill Block
	df.ColumnLength = fieldLength
	err := df.DataConfigSet("", "")
	if err != nil {
		panic(Format("写入文件配置出现问题，异常：%s", err.Error()))
	}
}

// DataConfigSetRowCount	改变配置文件行总和 |
// countRow:				int64				文件有效总行数 |
// return 					None |
func (df *StructDataFile) DataConfigSetRowCount(countRow int64) {
	df.CountRow = countRow
	err := df.DataConfigSet("", "")
	if err != nil {
		panic(Format("写入文件配置出现问题，异常：%s", err.Error()))
	}
}

// DataConfigSetEncoding	改变配置文件编码格式 |
// dataFile:				*StructDataFile		配置文件结构 |
// encoding:				int					数据文件编码格式 |
// return 					None |
func (df *StructDataFile) DataConfigSetEncoding(encoding int) {
	df.Encoding = encoding
	err := df.DataConfigSet("", "")
	if err != nil {
		panic(Format("写入文件配置出现问题，异常：%s", err.Error()))
	}
}

// DataWrite 	数据文件指定行列写入内容 |
// row:			int64	行【从1开始】 |
// column:		int		列【从1开始】 |
// data:		string	待写入的数据 |
// isAppend:	bool	是否为直接追加操作 |
// return		error	异常 |
func (df *StructDataFile) DataWrite(row int64, column int, data string, isAppend bool) error {
	const errOperation = "数据文件写入"
	// 从`df`打开数据文件，若不能打开则返回错误
	if df.File == nil {
		return errF(errMsgTypeFileNull, errOperation, "", "")
	}
	var err error
	var rowLength, rowLengthFront, columnLength int
	if isAppend {
		bf := bufio.NewWriter(df.File)
		_, err = bf.WriteString(data)
		if err != nil {
			return errF(errMsgTypeWriteData, errOperation, "数据写入过程中", err.Error())
		}
		_ = bf.Flush()
		return nil
	}
	// 从`df`中获得对应列的长度
	if len(df.ColumnLength) == 0 {
		return errF(errMsgTypeListEmpty, errOperation, "", "")
	}
	if column <= 0 || len(df.ColumnLength) < column {
		return errF(errMsgTypeListBound, errOperation, Format("待修改的列：%s", strconv.Itoa(column)), "")
	}
	err, rowLength, rowLengthFront, columnLength = CountRowLength(column, df.ColumnLength)
	if err != nil {
		return errF(errMsgTypeCalcLength, errOperation, "", err.Error())
	}
	// 判断待写入内容长度与字段长度是否一致
	if len(data) != columnLength {
		return errF(errMsgTypeMisMatchLength, errOperation, Format("数据长度：%s，要求列长度：%s", strconv.Itoa(len(data)), strconv.Itoa(columnLength)), "")
	}
	// 转化内容为`[]byte`
	bt := FuncStringToByteSlice(data)
	// 覆盖写入内容
	_, err = df.File.WriteAt(bt, int64(rowLength)*(row-1)+int64(rowLengthFront))
	if err != nil {
		return errF(errMsgTypeWriteData, errOperation, "数据写入过程中", err.Error())
	}
	return nil
}

// ProfileSet 	设置用户信息【df.File不得为空且必须允许写】
// keyValue:	any
// keyColumn:	int
// setValue:	any
// setColumn:	int
// return		error
func (df *StructDataFile) ProfileSet(keyValue string, keyColumn int, setValue any, setColumn int) error {
	const errOperation = "用户信息设定"
	if df.File == nil {
		return errF(errMsgTypeFileNull, errOperation, "判断`df.File`过程中", "")
	}
	err, valueKey := FillBlock(keyValue, df.ColumnLength[keyColumn-1], df.BlockFill[0])
	if err != nil {
		return errF("填充内容出错", errOperation, "填充用户信息的关键列中", err.Error())
	}
	err, valueSet := FillBlock(setValue, df.ColumnLength[setColumn-1], df.BlockFill[0])
	if err != nil {
		return errF("填充内容出错", errOperation, "填充用户信息的设定列中", err.Error())
	}
	// 对关键列排序后，使用二分查找指定的内容，获取行数
	err = SortBubble(df.File, df.CountRow, df.ColumnLength, keyColumn)
	if err != nil {
		return errF("冒泡排序出错", errOperation, "排序过程中", err.Error())
	}
	err, flag, _, row := SearchBinary(df.File, df.CountRow, df.ColumnLength, keyColumn, valueKey, df.BlockFill)
	if err != nil {
		return errF("二分查找失败", errOperation, "二分查找数据文件过程中", err.Error())
	}
	if !flag {
		return errF("二分查找未找到结果", errOperation, "查找并修改个人信息中", "")
	}
	// 将行、列值写入对应位置
	err = df.DataWrite(row, setColumn, valueSet, false)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "写入数据文件过程中", err.Error())
	}
	return nil
}

func (df *StructDataFile) FileClose() {
	if df.File != nil {
		_ = df.File.Close()
	}
}

func (pf *StructProfile) Copy(pfSrc *StructProfile) {
	pf.IsLogin = pfSrc.IsLogin
	pf.UID = pfSrc.UID
	pf.UserId = pfSrc.UserId
	pf.Permit = pfSrc.Permit
	pf.Name = pfSrc.Name
	pf.ClassId = pfSrc.ClassId
	pf.ActiveCode = pfSrc.ActiveCode
}

func (pf *StructProfile) Print() string {
	isLogin := "Is login? " + strconv.FormatBool(pf.IsLogin) + "\t"
	uid := "UID: " + pf.UID + "\t"
	name := "Name: " + pf.Name + "\t"
	userId := "User ID: " + pf.UserId + "\t"
	classId := "Class ID: " + pf.ClassId + "\t"
	permit := "Permit: " + strconv.Itoa(pf.Permit) + "\t"
	activeCode := "ActiveCode: " + pf.ActiveCode + "\t"
	return isLogin + uid + name + userId + classId + permit + activeCode
}

func (pf *StructProfile) ProfileGetIsLogin() bool {
	return pf.IsLogin
}

func (pf *StructProfile) ProfileGetUID() string {
	return pf.UID
}

func (pf *StructProfile) ProfileGetUserId() string {
	return pf.UserId
}

func (pf *StructProfile) ProfileGetPermit() int {
	return pf.Permit
}

func (pf *StructProfile) ProfileGetName() string {
	return pf.Name
}

func (pf *StructProfile) ProfileGetClassId() string {
	return pf.ClassId
}

func (pf *StructProfile) ProfileGetActiveCode() string {
	return pf.ActiveCode
}

func (pf *StructProfile) ProfileSetIsLogin(isLogin bool) {
	pf.IsLogin = isLogin
}

func (pf *StructProfile) ProfileSetUid(df *StructDataFile, uid string) error {
	const errOperation = "修改个人信息的UID"
	if uid == "" {
		return errF(errMsgTypeStringEmpty, errOperation, "", "")
	}
	err := df.ProfileSet(pf.UserId, ProfileDataUserId, uid, ProfileDataUid)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "", err.Error())
	}
	pf.UID = uid
	return nil
}

func (pf *StructProfile) ProfileSetUserId(df *StructDataFile, userId string) error {
	const errOperation = "修改个人信息的UserId"
	if userId == "" {
		return errF(errMsgTypeStringEmpty, errOperation, "", "")
	}
	err := df.ProfileSet(pf.UserId, ProfileDataUserId, userId, ProfileDataUserId)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "", err.Error())
	}
	pf.UserId = userId
	return nil
}

func (pf *StructProfile) ProfileSetPermit(df *StructDataFile, permit int) error {
	const errOperation = "修改个人信息的Permit"
	err := df.ProfileSet(pf.UserId, ProfileDataUserId, permit, ProfileDataPermit)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "", err.Error())
	}
	pf.Permit = permit
	return nil
}

func (pf *StructProfile) ProfileSetName(df *StructDataFile, name string) error {
	const errOperation = "修改个人信息的Name"
	if name == "" {
		return errF(errMsgTypeStringEmpty, errOperation, "", "")
	}
	err := df.ProfileSet(pf.UserId, ProfileDataUserId, name, ProfileDataName)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "", err.Error())
	}
	pf.Name = name
	return nil
}

func (pf *StructProfile) ProfileSetClassId(df *StructDataFile, classId string) error {
	const errOperation = "修改个人信息的ClassId"
	if classId == "" {
		return errF(errMsgTypeStringEmpty, errOperation, "", "")
	}
	err := df.ProfileSet(pf.UserId, ProfileDataUserId, classId, ProfileDataClassId)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "", err.Error())
	}
	pf.ClassId = classId
	return nil
}

func (pf *StructProfile) ProfileSetActiveCode(df *StructDataFile, activeCode string) error {
	const errOperation = "修改个人信息的ActiveCode"
	if activeCode == "" {
		return errF(errMsgTypeStringEmpty, errOperation, "", "")
	}
	err := df.ProfileSet(pf.UserId, ProfileDataUserId, activeCode, ProfileDataActiveCode)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "", err.Error())
	}
	pf.ActiveCode = activeCode
	return nil
}

func (date *StructDate) Format() string {
	if !date.CheckDate() {
		return "19760101"
	}
	fM := strconv.Itoa(date.M)
	fD := strconv.Itoa(date.D)
	if date.M < 10 {
		fM = "0" + fM
	}
	if date.D < 10 {
		fD = "0" + fD
	}
	return strconv.Itoa(date.Y) + fM + fD
}

func (score *StructScore) SetFloat64(value float64) {
	score.Integer = int(value)
	score.Decimal[0] = int(value*10) % 10
	score.Decimal[1] = int(value*100) % 10
	if score.Decimal[0] < 0 {
		score.Decimal[0] = -score.Decimal[0]
	}
	if score.Decimal[1] < 0 {
		score.Decimal[1] = -score.Decimal[1]
	}
}

func (score *StructScore) SetString(value string) bool {
	valueFloat, err := strconv.ParseFloat(value, 10)
	if err != nil {
		return false
	}
	score.SetFloat64(valueFloat)
	return true
}

func (score *StructScore) GetFloat64() float64 {
	if score.Integer < 0 {
		return float64(score.Integer) - float64(score.Decimal[0])*0.1 - float64(score.Decimal[1])*0.01
	}
	return float64(score.Integer) + float64(score.Decimal[0])*0.1 + float64(score.Decimal[1])*0.01
}

func (score *StructScore) GetString(isAbs bool) string {
	if isAbs && score.Integer < 0 {
		return strconv.Itoa(-score.Integer) + "." + strconv.Itoa(score.Decimal[0]) + strconv.Itoa(score.Decimal[1])
	}
	return strconv.Itoa(score.Integer) + "." + strconv.Itoa(score.Decimal[0]) + strconv.Itoa(score.Decimal[1])
}

func (score *StructScore) GetStringFormat(integerBlockLen int) (bool, string) {
	integerLength := len(strconv.Itoa(score.Integer))
	if integerBlockLen < integerLength {
		return false, score.GetString(false)
	}
	if 0 <= score.Integer {
		return true, strings.Repeat("0", integerBlockLen-integerLength) + score.GetString(false)
	} else {
		return true, "-" + strings.Repeat("0", integerBlockLen-integerLength) + score.GetString(true)
	}
}
