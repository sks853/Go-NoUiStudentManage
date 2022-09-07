/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-09-06 14:50
 * @FileName     menu_tree.go
 * @Description  None
 * ==================================================
**/

package main

import (
	"NoUiStudentManage/public"
)

// ---------- ---------- ---------- ---------- ---------- 02

var nodeProfileSelfCheck = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ProfileSelfCheck,
	Func:   nil,
}

var nodeProfileSelfModify = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ProfileSelfModify,
	Func:   nil,
}

var nodeProfileOtherCheck = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ProfileOtherCheck,
	Func:   nil,
}

var nodeProfileOtherEdit = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ProfileOtherEdit,
	Func:   nil,
}

var nodeSubjectAdd = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02SubjectAdd,
	Func:   nil,
}

var nodeSubjectDel = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02SubjectDel,
	Func:   nil,
}

var nodeSubjectCheck = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02SubjectCheck,
	Func:   nil,
}

var nodeSubjectEdit = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02SubjectEdit,
	Func:   nil,
}

var nodeClassAdd = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ClassAdd,
	Func:   nil,
}

var nodeClassDel = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ClassDel,
	Func:   nil,
}

var nodeClassCheck = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ClassCheck,
	Func:   nil,
}

var nodeClassEdit = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ClassEdit,
	Func:   nil,
}

var nodeClassMemberAdd = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ClassMemberAdd,
	Func:   nil,
}

var nodeClassMemberDel = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ClassMemberDel,
	Func:   nil,
}

var nodeExamAdd = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ExamAdd,
	Func:   nil,
}

var nodeExamDel = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ExamDel,
	Func:   nil,
}

var nodeExamCheck = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ExamCheck,
	Func:   nil,
}

var nodeExamEdit = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ExamEdit,
	Func:   nil,
}

var nodeReportAdd = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ReportAdd,
	Func:   nil,
}

var nodeReportEdit = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ReportEdit,
	Func:   nil,
}

var nodeReportCheckDate = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ReportCheckDate,
	Func:   nil,
}

var nodeReportCheckId = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ReportCheckId,
	Func:   nil,
}

var nodeMailboxSend = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02MailboxSend,
	Func:   nil,
}

var nodeMailboxReceive = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02MailboxReceive,
	Func:   nil,
}

var nodeToolboxAddActiveCode = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ToolboxAddActiveCode,
	Func:   nil,
}

var nodeToolboxDelAccount = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu02ToolboxDelAccount,
	Func:   nil,
}

// ---------- ---------- ---------- ---------- ---------- 01

var nodeProfile = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu01Profile,
	MenuNode: []*public.StructMenu{
		&nodeProfileSelfCheck,
		&nodeProfileSelfModify,
		&nodeProfileOtherCheck,
		&nodeProfileOtherEdit,
	},
	Func: nil,
}

var nodeSubject = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu01Subject,
	MenuNode: []*public.StructMenu{
		&nodeSubjectAdd,
		&nodeSubjectDel,
		&nodeSubjectCheck,
		&nodeSubjectEdit,
	},
	Func: nil,
}

var nodeClass = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu01Class,
	MenuNode: []*public.StructMenu{
		&nodeClassAdd,
		&nodeClassDel,
		&nodeClassCheck,
		&nodeClassEdit,
		&nodeClassMemberAdd,
		&nodeClassMemberDel,
	},
	Func: nil,
}

var nodeExam = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu01Exam,
	MenuNode: []*public.StructMenu{
		&nodeExamAdd,
		&nodeExamDel,
		&nodeExamCheck,
		&nodeExamEdit,
	},
	Func: nil,
}

var nodeReport = public.StructMenu{
	Permit: []int{},
	Text:   public.Menu01Report,
	MenuNode: []*public.StructMenu{
		&nodeReportAdd,
		&nodeReportEdit,
		&nodeReportCheckDate,
		&nodeReportCheckId,
	},
	Func: nil,
}

var nodeMailbox = public.StructMenu{
	Permit: []int{public.PermitManager, public.PermitUser},
	Text:   public.Menu01Mailbox,
	MenuNode: []*public.StructMenu{
		&nodeMailboxSend,
		&nodeMailboxReceive,
	},
	Func: nil,
}

var nodeToolbox = public.StructMenu{
	Permit: []int{public.PermitAdministrator},
	Text:   public.Menu01Toolbox,
	MenuNode: []*public.StructMenu{
		&nodeToolboxAddActiveCode,
		&nodeToolboxDelAccount,
	},
	Func: nil,
}

// ---------- ---------- ---------- ---------- ---------- 00

var nodeLogin = public.StructMenu{
	Permit: []int{public.PermitGuest},
	Text:   public.Menu00Login,
	MenuNode: []*public.StructMenu{
		&nodeProfile,
		&nodeSubject,
		&nodeClass,
		&nodeExam,
		&nodeReport,
		&nodeMailbox,
		&nodeToolbox,
	},
	Func: nil,
}

var nodeRegister = public.StructMenu{
	Permit: []int{public.PermitGuest},
	Text:   public.Menu00Register,
	Func:   nil,
}

// ---------- ---------- ---------- ---------- ----------

func InitTreeMenu() *public.StructMenu {
	return &public.StructMenu{
		Permit: []int{public.PermitGuest},
		Text:   public.Menu00Login,
		MenuNode: []*public.StructMenu{
			&nodeLogin,
			&nodeRegister,
		},
		Func: nil,
	}
}
