/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-09-06 14:50
 * @FileName     menu_route.go
 * @Description  None
 * ==================================================
**/

package main

import (
	"NoUiStudentManage/database"
	"NoUiStudentManage/public"
)

func funcLogin(share *public.StructShareBase, option ...any) {
	public.FuncPrintLog(public.LogInfo, "用户正在尝试登陆")
	count := 2
	if len(option) == 0 {
		panic("请设置参数位置[0]且为布尔型变量的地址，当前参数数量为0")
		return
	}
	switch option[0].(type) {
	case *bool:
		break
	default:
		panic("请设置参数位置[0]且为布尔型变量的地址，当前参数类型非布尔型指针地址")
		return
	}
	vfResult := option[0].(*bool)
	for {
		userId, err := public.TipInput(public.TipInputUserId)
		if err != nil {
			public.FuncPrintLog(public.LogErrs, "登陆时输入用户编号出现异常", err)
			continue
		}
		password, err := public.TipInput(public.TipInputPassword)
		if err != nil {
			public.FuncPrintLog(public.LogErrs, "登陆时输入密码出现异常", err)
			continue
		}
		ok, err := database.FuncDataLogin(userId, password, share.Profile)
		if err != nil {
			public.FuncPrintLog(public.LogErrs, "登陆时校验账户出现异常", err)
			continue
		}
		if ok {
			public.FuncPrintLog(public.LogInfo, "用户登陆成功")
			public.TipWait("登陆成功")
			*vfResult = true
			return
		}
		if count == 0 {
			public.TipWait("密码输入错误过多")
			public.Clear()
			*vfResult = false
			return
		}
		public.TipWait("密码不匹配，请重新输入")
		public.FuncPrintLog(public.LogWarn, "用户尝试登陆但失败")
		count--
	}
}
func funcRegister(value *public.StructShareBase, option ...any) {
	public.FuncPrintLog(public.LogInfo, "用户尝试注册和激活")
	public.TipWait("Register")
}
func funcProfileSelfCheck(share *public.StructShareBase, option ...any)     {}
func funcProfileSelfModify(share *public.StructShareBase, option ...any)    {}
func funcProfileOtherCheck(share *public.StructShareBase, option ...any)    {}
func funcProfileOtherEdit(share *public.StructShareBase, option ...any)     {}
func funcSubjectAdd(share *public.StructShareBase, option ...any)           {}
func funcSubjectDel(share *public.StructShareBase, option ...any)           {}
func funcSubjectCheck(share *public.StructShareBase, option ...any)         {}
func funcSubjectEdit(share *public.StructShareBase, option ...any)          {}
func funcClassAdd(share *public.StructShareBase, option ...any)             {}
func funcClassDel(share *public.StructShareBase, option ...any)             {}
func funcClassCheck(share *public.StructShareBase, option ...any)           {}
func funcClassEdit(share *public.StructShareBase, option ...any)            {}
func funcClassMemberAdd(share *public.StructShareBase, option ...any)       {}
func funcClassMemberDel(share *public.StructShareBase, option ...any)       {}
func funcExamAdd(share *public.StructShareBase, option ...any)              {}
func funcExamDel(share *public.StructShareBase, option ...any)              {}
func funcExamCheck(share *public.StructShareBase, option ...any)            {}
func funcExamEdit(share *public.StructShareBase, option ...any)             {}
func funcReportAdd(share *public.StructShareBase, option ...any)            {}
func funcReportEdit(share *public.StructShareBase, option ...any)           {}
func funcReportCheckDate(share *public.StructShareBase, option ...any)      {}
func funcReportCheckId(share *public.StructShareBase, option ...any)        {}
func funcMailboxSend(share *public.StructShareBase, option ...any)          {}
func funcMailboxReceive(share *public.StructShareBase, option ...any)       {}
func funcToolboxAddActiveCode(share *public.StructShareBase, option ...any) {}
func funcToolboxDelAccount(share *public.StructShareBase, option ...any)    {}
