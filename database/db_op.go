/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-09-07 21:25
 * @FileName     db_op.go
 * @Description  None
 * ==================================================
**/

package database

import "NoUiStudentManage/public"

func FuncDataLogin(userId string, password string, profile *public.StructProfile) (bool, error) {
	// TODO 验证登陆情况
	if userId == "root" && password == "root" {
		return true, nil
	}
	return false, nil
}
