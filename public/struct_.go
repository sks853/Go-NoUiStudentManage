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
	Name   string
	UserId string
}

type StructMenu struct {
	Title    string
	Profile  StructProfile
	Permit   []int
	Text     string
	MenuNode []*StructMenu
	Func     any
}

type StructMenuLink struct {
	NodeLast *StructMenuLink
	NodeNext *StructMenuLink
	Node     *StructMenu
}
