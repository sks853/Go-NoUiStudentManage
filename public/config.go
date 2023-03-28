/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2023-02-07 22:03
 * @FileName     config.go
 * @Description  None
 * ==================================================
**/

package public

import "fmt"

var PathDirRoot = "E:\\WorkSpace\\Go\\NoUiStudentManage"

var PathDirClass = PathDirRoot + "/data/class/"
var PathDirExam = PathDirRoot + "/data/exam/"
var PathDirLog = PathDirRoot + "/log/"
var PathDirSubject = PathDirRoot + "/data/subject/"
var PathDirUser = PathDirRoot + "/data/user/"

const FilePassword = "passwd"
const FileUser = "user"

func _() {
	fmt.Println(PathDirRoot, PathDirClass, PathDirExam, PathDirLog, PathDirSubject, PathDirUser, FilePassword, FileUser)
}
