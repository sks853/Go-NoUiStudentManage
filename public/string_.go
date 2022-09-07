/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-09-05 12:02
 * @FileName     string_.go
 * @Description  None
 * ==================================================
**/

package public

import "runtime"

const SysType = runtime.GOOS

const (
	Menu00Login    = "登陆账户"
	Menu00Register = "激活账户"
)

const (
	Menu01Profile = "账户信息"
	Menu01Subject = "学科管理"
	Menu01Class   = "班级管理"
	Menu01Exam    = "考试管理"
	Menu01Report  = "成绩管理"
	Menu01Mailbox = "邮件箱"
	Menu01Toolbox = "工具箱"
)

const (
	Menu02ProfileSelfCheck  = "查看当前用户信息"
	Menu02ProfileSelfModify = "修改当前用户信息"
	Menu02ProfileOtherCheck = "查看其他用户信息"
	Menu02ProfileOtherEdit  = "修改其他用户信息"

	Menu02SubjectAdd   = "新增科目"
	Menu02SubjectDel   = "删除科目"
	Menu02SubjectCheck = "科目详情"
	Menu02SubjectEdit  = "编辑科目"

	Menu02ClassAdd       = "创建班级"
	Menu02ClassDel       = "删除班级"
	Menu02ClassCheck     = "班级详情"
	Menu02ClassEdit      = "编辑班级"
	Menu02ClassMemberAdd = "添加班级成员"
	Menu02ClassMemberDel = "移除班级成员"

	Menu02ExamAdd   = "安排考试"
	Menu02ExamDel   = "取消考试"
	Menu02ExamCheck = "考试列表"
	Menu02ExamEdit  = "编辑考试"

	Menu02ReportAdd       = "成绩录入"
	Menu02ReportEdit      = "修改成绩"
	Menu02ReportCheckDate = "根据考试日期查询成绩"
	Menu02ReportCheckId   = "根据考试编号查询成绩"

	Menu02MailboxSend    = "发送邮件"
	Menu02MailboxReceive = "查看收件箱"

	Menu02ToolboxAddActiveCode = "创建注册授权码"
	Menu02ToolboxDelAccount    = "销毁账户"
)

const (
	TipSelectOptionEdit = "请选择需要修改的部分：\n"
	TipInputContentEdit = "请输入修改后的内容：\n"
	TipInputAgain       = "请重新进行操作！\n"
	TipConfirmResult    = "请确认当前操作无误：\n"
	TipConfirmInputNext = "是否继续%s？\n"
	TipOpSuccess        = "操作成功！\n"
	TipOpCancel         = "操作取消！\n"
	TipOpFail           = "操作失败！%s\n"

	TipInputActiveCode        = "请输入授权码："
	TipInputClassId           = "请输入行政班级号："
	TipInputClassManager      = "请输入班级管理员的用户编号："
	TipInputEmailText         = "请输入邮件内容：\n"
	TipInputName              = "请输入用户姓名："
	TipInputPassword          = "请输入用户密码："
	TipInputSubjectId         = "请输入科目编号："
	TipInputSubjectName       = "请输入科目名称："
	TipInputUserId            = "请输入用户编号："
	TipInputExamDatetimeStart = "请输入考试起始日期时间："
	TipInputExamDatetimeEnd   = "请输入考试截至日期时间："
	TipInputExamAdministrator = "请输入考试管理员用户编号："
	TipSelectExamColony       = "请选择考试群体：\n"
	TipSelectPermission       = "请选择用户权限：\n"
)
