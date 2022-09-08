/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-09-06 12:51
 * @FileName     struct_.go
 * @Description  None
 * ==================================================
**/

package public

type StructProfile struct {
	IsLogin   bool
	Name      string
	UserId    string
	ClassId   string
	SubjectID string
}

type StructMenu struct {
	Title       string
	Profile     *StructProfile
	PermitMode  int
	PermitList  []int
	Text        string
	HasVerifier bool
	MenuNode    []*StructMenu
	Func        func(*StructShareBase, ...any)
}

type StructMenuLink struct {
	NodeLast *StructMenuLink
	NodeNext *StructMenuLink
	Node     *StructMenu
}

type StructShareBase struct {
	Profile *StructProfile
}
