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
	"NoUiStudentManage/public"
)

func checkArgs(option []any) bool {
	if len(option) == 0 {
		panic("请设置参数位置[0]且为布尔型变量的地址，当前参数数量为0")
		return false
	}
	switch option[0].(type) {
	case *bool:
		break
	default:
		panic("请设置参数位置[0]且为布尔型变量的地址，当前参数类型非布尔型指针地址")
		return false
	}
	return true
}

func funcLogin(share *public.StructShareBase, option ...any) {
	public.FuncPrintLog(public.LogInfo, "用户正在尝试登陆")
	count := 2
	if !checkArgs(option) {
		return
	}
	vfResult := option[0].(*bool)
	for {
		if share.Profile.IsLogin {
			public.TipWait("当前已登陆，即将进入主菜单", 1)
			*vfResult = true
			return
		}
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
		ok, err := funcDataLogin(userId, password, share.Profile)
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
		public.FuncPrintLog(public.LogWarn, public.Format("未登录用户尝试登陆{%s}但失败", userId))
		count--
	}
}
func funcRegister(share *public.StructShareBase, _ ...any) {
	public.FuncPrintLog(public.LogInfo, "未登录用户尝试注册和激活")
	public.TipWait("Register")
}
func funcChangeAccount(share *public.StructShareBase, _ ...any) {
	if share.Profile == nil || !share.Profile.IsLogin {
		public.TipWait(public.Format(public.TipOpFail, "当前尚未登陆，请先进行登陆"))
		return
	}
	share.Profile.IsLogin = false
	public.TipWait(public.TipOpSuccess)
	public.FuncPrintLog(public.LogInfo, public.Format("用户{%s}切换账户", share.Profile.UserId))
}
func funcProfileSelfCheck(share *public.StructShareBase, _ ...any)     {}
func funcProfileSelfModify(share *public.StructShareBase, _ ...any)    {}
func funcProfileOtherCheck(share *public.StructShareBase, _ ...any)    {}
func funcProfileOtherEdit(share *public.StructShareBase, _ ...any)     {}
func funcSubjectAdd(share *public.StructShareBase, _ ...any)           {}
func funcSubjectDel(share *public.StructShareBase, _ ...any)           {}
func funcSubjectCheck(share *public.StructShareBase, _ ...any)         {}
func funcSubjectEdit(share *public.StructShareBase, _ ...any)          {}
func funcClassAdd(share *public.StructShareBase, _ ...any)             {}
func funcClassDel(share *public.StructShareBase, _ ...any)             {}
func funcClassCheck(share *public.StructShareBase, _ ...any)           {}
func funcClassEdit(share *public.StructShareBase, _ ...any)            {}
func funcClassMemberAdd(share *public.StructShareBase, _ ...any)       {}
func funcClassMemberDel(share *public.StructShareBase, _ ...any)       {}
func funcExamAdd(share *public.StructShareBase, _ ...any)              {}
func funcExamDel(share *public.StructShareBase, _ ...any)              {}
func funcExamCheck(share *public.StructShareBase, _ ...any)            {}
func funcExamEdit(share *public.StructShareBase, _ ...any)             {}
func funcReportAdd(share *public.StructShareBase, _ ...any)            {}
func funcReportEdit(share *public.StructShareBase, _ ...any)           {}
func funcReportCheckDate(share *public.StructShareBase, _ ...any)      {}
func funcReportCheckId(share *public.StructShareBase, _ ...any)        {}
func funcMailboxSend(share *public.StructShareBase, _ ...any)          {}
func funcMailboxReceive(share *public.StructShareBase, _ ...any)       {}
func funcToolboxAddActiveCode(share *public.StructShareBase, _ ...any) {}
func funcToolboxDelAccount(share *public.StructShareBase, _ ...any)    {}
