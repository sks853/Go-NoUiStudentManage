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

import (
	"runtime"
)

const SysType = runtime.GOOS

const (
	ProfileDataUid        = 1
	ProfileDataUserId     = 2
	ProfileDataPermit     = 3
	ProfileDataName       = 4
	ProfileDataClassId    = 5
	ProfileDataActiveCode = 6
)

const (
	PasswordDataUid    = 1
	PasswordDataPasswd = 2
)

const (
	Menu00Login         = "登陆账户"
	Menu00Register      = "激活账户"
	Menu00ChangeAccount = "切换账户"
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

const (
	errMsgTypeBase                = "由于`%s`，而无法进行`%s`，辅助信息：`%s`，异常：`%s`"
	errMsgTypeSorSelectColumn     = "排序选取的关键列越界"
	errMsgTypeOffsetGetData       = "无法偏移指针获取数据"
	errMsgTypeOffsetHead          = "文件指针无法偏移到首位"
	errMsgTypeFileNull            = "提供的文件指针为空"
	errMsgTypeCalcLength          = "计算长度过程中出错"
	errMsgTypeArgsLengthDifferent = "传入参数与待比较参数长度不等"
	errMsgTypeCovertInt           = "无法将`%s`转换为数字型"
	errMsgTypeConfigError         = "配置文件参数不足"
	errMsgTypeOpenConfig          = "无法打开目标配置文件"
	errMsgTypeOpenData            = "无法打开目标数据文件"
	errMsgTypeArgsNull            = "传入的参数为空指针"
	errMsgTypeStringEmpty         = "字符串内容为空"
	errMsgTypeListEmpty           = "数组列表为空"
	errMsgTypeListBound           = "提供的参数在数组范围外，已越界"
	errMsgTypeMisMatchLength      = "内容长度不匹配"
	errMsgTypeWriteData           = "无法写入文件"
	errMsgTypeViewError           = "参考异常内容"
	errMsgInvalidArgs             = "提供了非法的参数"
	errMsgTypeReadData            = "文件数据读取失败"
)
