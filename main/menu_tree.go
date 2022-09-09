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
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager, public.PermitUser},
	Text:       public.Menu02ProfileSelfCheck,
	Func:       funcProfileSelfCheck,
}

var nodeProfileSelfModify = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager, public.PermitUser},
	Text:       public.Menu02ProfileSelfModify,
	Func:       funcProfileSelfModify,
}

var nodeProfileOtherCheck = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02ProfileOtherCheck,
	Func:       funcProfileOtherCheck,
}

var nodeProfileOtherEdit = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02ProfileOtherEdit,
	Func:       funcProfileOtherEdit,
}

var nodeSubjectAdd = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02SubjectAdd,
	Func:       funcSubjectAdd,
}

var nodeSubjectDel = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02SubjectDel,
	Func:       funcSubjectDel,
}

var nodeSubjectCheck = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager, public.PermitUser},
	Text:       public.Menu02SubjectCheck,
	Func:       funcSubjectCheck,
}

var nodeSubjectEdit = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02SubjectEdit,
	Func:       funcSubjectEdit,
}

var nodeClassAdd = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02ClassAdd,
	Func:       funcClassAdd,
}

var nodeClassDel = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02ClassDel,
	Func:       funcClassDel,
}

var nodeClassCheck = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager, public.PermitUser},
	Text:       public.Menu02ClassCheck,
	Func:       funcClassCheck,
}

var nodeClassEdit = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02ClassEdit,
	Func:       funcClassEdit,
}

var nodeClassMemberAdd = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02ClassMemberAdd,
	Func:       funcClassMemberAdd,
}

var nodeClassMemberDel = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02ClassMemberDel,
	Func:       funcClassMemberDel,
}

var nodeExamAdd = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager},
	Text:       public.Menu02ExamAdd,
	Func:       funcExamAdd,
}

var nodeExamDel = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager},
	Text:       public.Menu02ExamDel,
	Func:       funcExamDel,
}

var nodeExamCheck = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager},
	Text:       public.Menu02ExamCheck,
	Func:       funcExamCheck,
}

var nodeExamEdit = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager},
	Text:       public.Menu02ExamEdit,
	Func:       funcExamEdit,
}

var nodeReportAdd = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager},
	Text:       public.Menu02ReportAdd,
	Func:       funcReportAdd,
}

var nodeReportEdit = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02ReportEdit,
	Func:       funcReportEdit,
}

var nodeReportCheckDate = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager},
	Text:       public.Menu02ReportCheckDate,
	Func:       funcReportCheckDate,
}

var nodeReportCheckId = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager},
	Text:       public.Menu02ReportCheckId,
	Func:       funcReportCheckId,
}

var nodeMailboxSend = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager, public.PermitUser},
	Text:       public.Menu02MailboxSend,
	Func:       funcMailboxSend,
}

var nodeMailboxReceive = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager, public.PermitUser},
	Text:       public.Menu02MailboxReceive,
	Func:       funcMailboxReceive,
}

var nodeToolboxAddActiveCode = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02ToolboxAddActiveCode,
	Func:       funcToolboxAddActiveCode,
}

var nodeToolboxDelAccount = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu02ToolboxDelAccount,
	Func:       funcToolboxDelAccount,
}

// ---------- ---------- ---------- ---------- ---------- 01

var nodeProfile = public.StructMenu{
	PermitMode: public.PermitModeGreater,
	PermitList: []int{public.PermitGuest},
	Text:       public.Menu01Profile,
	MenuNode: []*public.StructMenu{
		&nodeProfileSelfCheck,
		&nodeProfileSelfModify,
		&nodeProfileOtherCheck,
		&nodeProfileOtherEdit,
	},
}

var nodeSubject = public.StructMenu{
	PermitMode: public.PermitModeGreater,
	PermitList: []int{public.PermitUser},
	Text:       public.Menu01Subject,
	MenuNode: []*public.StructMenu{
		&nodeSubjectAdd,
		&nodeSubjectDel,
		&nodeSubjectCheck,
		&nodeSubjectEdit,
	},
}

var nodeClass = public.StructMenu{
	PermitMode: public.PermitModeGreater,
	PermitList: []int{public.PermitGuest},
	Text:       public.Menu01Class,
	MenuNode: []*public.StructMenu{
		&nodeClassAdd,
		&nodeClassDel,
		&nodeClassCheck,
		&nodeClassEdit,
		&nodeClassMemberAdd,
		&nodeClassMemberDel,
	},
}

var nodeExam = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator, public.PermitManager},
	Text:       public.Menu01Exam,
	MenuNode: []*public.StructMenu{
		&nodeExamAdd,
		&nodeExamDel,
		&nodeExamCheck,
		&nodeExamEdit,
	},
}

var nodeReport = public.StructMenu{
	PermitMode: public.PermitModeGreater,
	PermitList: []int{public.PermitGuest},
	Text:       public.Menu01Report,
	MenuNode: []*public.StructMenu{
		&nodeReportAdd,
		&nodeReportEdit,
		&nodeReportCheckDate,
		&nodeReportCheckId,
	},
}

var nodeMailbox = public.StructMenu{
	PermitMode: public.PermitModeGreater,
	PermitList: []int{public.PermitGuest},
	Text:       public.Menu01Mailbox,
	MenuNode: []*public.StructMenu{
		&nodeMailboxSend,
		&nodeMailboxReceive,
	},
}

var nodeToolbox = public.StructMenu{
	PermitMode: public.PermitModeEqual,
	PermitList: []int{public.PermitAdministrator},
	Text:       public.Menu01Toolbox,
	MenuNode: []*public.StructMenu{
		&nodeToolboxAddActiveCode,
		&nodeToolboxDelAccount,
	},
}

// ---------- ---------- ---------- ---------- ---------- 00

var nodeLogin = public.StructMenu{
	PermitMode:  public.PermitModeEqualGreater,
	PermitList:  []int{public.PermitGuest},
	Text:        public.Menu00Login,
	HasVerifier: true,
	MenuNode: []*public.StructMenu{
		&nodeProfile,
		&nodeSubject,
		&nodeClass,
		&nodeExam,
		&nodeReport,
		&nodeMailbox,
		&nodeToolbox,
	},
	Func: funcLogin,
}

var nodeRegister = public.StructMenu{
	PermitMode: public.PermitModeEqualGreater,
	PermitList: []int{public.PermitGuest},
	Text:       public.Menu00Register,
	Func:       funcRegister,
}

var nodeChangeAccount = public.StructMenu{
	PermitMode:  public.PermitModeEqualGreater,
	PermitList:  []int{public.PermitGuest},
	HasVerifier: true,
	IsKeepMenu:  true,
	Text:        public.Menu00ChangeAccount,
	Func:        funcChangeAccount,
}

// ---------- ---------- ---------- ---------- ----------

func InitTreeMenu(share *public.StructShareBase) *public.StructMenu {
	return &public.StructMenu{
		Profile:    share.Profile,
		PermitList: []int{public.PermitGuest},
		Text:       public.Menu00Login,
		MenuNode: []*public.StructMenu{
			&nodeLogin,
			&nodeRegister,
			&nodeChangeAccount,
		},
	}
}
