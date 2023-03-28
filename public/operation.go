/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-09-07 21:25
 * @FileName     operation.go
 * @Description  None
 * ==================================================
**/

package public

import (
	"strconv"
)

func OpUserGet(pf *StructProfile, userId string, blockFill string) (error, bool) {
	const errOperation = "用户数据获取"
	if userId == "" {
		return errF(errMsgTypeStringEmpty, errOperation, "该字符串由函数参数提供", ""), false
	}
	if pf == nil {
		pf = new(StructProfile)
	}
	err, flag, data := DataFileGet(PathDirUser, FileUser, ProfileDataUserId, userId)
	if err != nil {
		return errF(errMsgTypeReadData, errOperation, "检索指定内容过程中", err.Error()), false
	}
	if !flag {
		return nil, false
	}
	// 解析数据
	pf.UID = StrStrip(data[0], blockFill)
	pf.UserId = StrStrip(data[1], blockFill)
	pf.Permit, err = strconv.Atoi(StrStrip(data[2], blockFill))
	if err != nil {
		return errF(Format(errMsgTypeCovertInt, data[2]), errOperation, "解析用户权限过程中", err.Error()), false
	}
	pf.Name = StrStrip(data[3], blockFill)
	pf.ClassId = StrStrip(data[4], blockFill)
	pf.ActiveCode = StrStrip(data[5], blockFill)
	return nil, true
}

func OpUserSet(pf *StructProfile, df *StructDataFile, setFunc ...StructSetProfile) error {
	const errOperation = "用户数据设定"
	if pf == nil {
		return errF(errMsgTypeArgsNull, errOperation, "该参数为用户信息结构", "")
	} else if df == nil {
		return errF(errMsgTypeArgsNull, errOperation, "该参数为数据文件结构", "")
	}
	// 打开文件
	err := df.GetFile(PathDirUser, FileUser, false)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "获取用户信息文件失败", err.Error())
	}
	// 逐个属性写入
	for i := 0; i < len(setFunc); i++ {
		switch (setFunc[i].Func).(type) {
		case func(*StructDataFile, string) error:
			funcSet := setFunc[i].Func.(func(*StructDataFile, string) error)
			value := setFunc[i].Value.(string)
			err = funcSet(df, value)
			if err != nil {
				df.FileClose()
				return errF(errMsgTypeViewError, errOperation, Format("逐个设置用户信息属性，当前序列为%s", strconv.Itoa(i)), err.Error())
			}
			break
		case func(*StructDataFile, int) error:
			funcSet := setFunc[i].Func.(func(*StructDataFile, int) error)
			value := setFunc[i].Value.(int)
			err = funcSet(df, value)
			if err != nil {
				df.FileClose()
				return errF(errMsgTypeViewError, errOperation, Format("逐个设置用户信息属性，当前序列为%s", strconv.Itoa(i)), err.Error())
			}
			break
		default:
			df.FileClose()
			return errF(errMsgInvalidArgs, errOperation, "列表参数应当为指定的设置用户信息的函数", "")
		}
	}
	df.FileClose()
	return nil
}

func OpUserNew(pf *StructProfile) error {
	const errOperation = "用户数据新增"
	if pf == nil {
		return errF(errMsgTypeArgsNull, errOperation, "该参数为用户信息结构", "")
	}
	data := []string{pf.UID, pf.UserId, strconv.Itoa(pf.Permit), pf.Name, pf.ClassId, pf.ActiveCode}
	err := InsertData(new(StructDataFile), PathDirUser, FileUser, ProfileDataUserId, data)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "插入数据过程中", err.Error())
	}
	return nil
}

func OpUserDel(pf *StructProfile) error {
	const errOperation = "用户数据删除"
	if pf == nil {
		return errF(errMsgTypeArgsNull, errOperation, "该参数为用户信息结构", "")
	}
	// TODO Delete User Data
	err := DeleteData(new(StructDataFile), "", "", -1)
	if err != nil {
		return err
	}
	return nil
}

func OpUserPasswdVerify(uid string, password string, blockFill string) (error, bool) {
	const errOperation = "用户信息验证"
	err, flag, data := DataFileGet(PathDirUser, FilePassword, PasswordDataUid, uid)
	if err != nil {
		return errF(errMsgTypeReadData, errOperation, "检索指定内容过程中", err.Error()), false
	}
	if !flag {
		return nil, false
	}
	// 解析数据
	verifyUid := StrStrip(data[0], blockFill)
	verifyPasswd := StrStrip(data[1], blockFill)
	if verifyUid == uid && verifyPasswd == password {
		return nil, true
	}
	return nil, false
}

func OpUserPasswdNew(uid string, password string) error {
	const errOperation = "用户校验新增"
	data := []string{uid, password}
	err := InsertData(new(StructDataFile), PathDirUser, FilePassword, PasswordDataUid, data)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "插入数据过程中", err.Error())
	}
	return nil
}

func OpUserPasswdSet(uid string, passwordNew string) error {
	const errOperation = "用户校验设定"
	err := DataFileSet(PathDirUser, FilePassword, uid, PasswordDataUid, []string{passwordNew}, []int{PasswordDataPasswd})
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "插入数据过程中", err.Error())
	}
	return nil
}

func OpExamNew(ex *StructExam) error {
	const errOperation = "成绩记录新增"
	if ex == nil {
		return errF(errMsgTypeArgsNull, errOperation, "该参数为用户信息结构", "")
	}
	if ex.ExamId == "" {

	} else if !ex.DateStart.CheckDate() {

	} else if !ex.DateEnd.CheckDate() {

	} else if ex.CountSubject == 0 {

	} else if len(ex.SubjectId) == 0 || len(ex.SubjectAvg) == 0 || len(ex.ExamRecord) == 0 {

	}
	// 读取成绩记录模板

	// 创建新的成绩数据并写入数据
	data := []string{ex.ExamId, ex.DateStart.Format(), ex.DateEnd.Format(), strconv.Itoa(ex.CountSubject)}
	data = append(data, ex.SubjectId...)
	for i := 0; i < len(ex.SubjectAvg); i++ {
		data = append(data, ex.SubjectAvg[i].ToString())
	}
	err := InsertData(new(StructDataFile), "PathDirUser", "FileUser", 0, data)
	if err != nil {
		return errF(errMsgTypeViewError, errOperation, "插入数据过程中", err.Error())
	}
	return nil
}
