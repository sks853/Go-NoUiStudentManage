/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-09-06 12:51
 * @FileName     model.go
 * @Description  None
 * ==================================================
**/

package public

import (
	"os"
)

const DecLength = 2
const ConfigLength = 5

// StructProfile 用户信息
type StructProfile struct {
	IsLogin    bool
	UID        string
	Name       string
	UserId     string
	ClassId    string
	Permit     int
	ActiveCode string
}

// StructScore 分数结构
type StructScore struct {
	Integer int
	Decimal [DecLength]int
}

// StructScoreRecord 单一科目成绩
type StructScoreRecord struct {
	UserId    string
	SubjectId string
	Score     StructScore
}

// StructExamRecord 对于某场考试的该考生所有科目成绩和排名
type StructExamRecord struct {
	Rank         int
	UserId       string
	SubjectScore StructScoreRecord
}

// StructDate 日期格式
type StructDate struct {
	Y int
	M int
	D int
}

// StructExam 详细的考试记录
type StructExam struct {
	ExamId       string
	DateStart    StructDate
	DateEnd      StructDate
	CountSubject int
	SubjectId    []string
	SubjectAvg   []StructScoreRecord
	ExamRecord   []StructExamRecord
}

// StructMenu 菜单某结点的结构
type StructMenu struct {
	Title       string                         // 菜单结点标题
	Profile     *StructProfile                 // 菜单结点用户信息
	PermitMode  int                            // 权限模式
	PermitList  []int                          // 权限列表
	Text        string                         // 显示的文本
	HasVerifier bool                           // 是否需要校验器
	IsKeepMenu  bool                           // 是否保持停留这个页面
	MenuNode    []*StructMenu                  // 菜单结点数据结构
	Func        func(*StructShareBase, ...any) // 当前结点执行的函数，如果有
}

// StructMenuLink 菜单结点链式结构
type StructMenuLink struct {
	NodeLast *StructMenuLink
	NodeNext *StructMenuLink
	Node     *StructMenu
}

// StructShareBase 共享变量的结构
type StructShareBase struct {
	Profile *StructProfile // 用户信息数据结构
}

// StructDataFile 某个文件的配置结构
type StructDataFile struct {
	BlockFill    string   // 当字段不满时填充的字符（不得和内容冲突）
	SortField    int      // 排序优先列
	CountRow     int64    // 文件总行数
	ColumnLength []int    // 每个字段的长度
	Encoding     int      // 文件编码格式，默认为`UTF-8`，0: UTF-8, 1: GBK
	File         *os.File // 数据文件指针
	FilePath     string   // 数据文件路径（不含文件名）
	FileName     string   // 数据文件名
	IsSort       bool     // 是否已经排序
}

// StructSetProfile 设置用户信息的结构
type StructSetProfile struct {
	Func  any // 设置用户信息的函数
	Value any // 设置用户信息函数所需传递的形式参数
}
