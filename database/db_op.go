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
		profile.IsLogin = true
		profile.Name = "超级管理员"
		profile.UserId = "root"
		profile.Permit = public.PermitAdministrator
		return true, nil
	} else if userId == "admin" && password == "admin" {
		profile.IsLogin = true
		profile.Name = "最靓的崽"
		profile.UserId = "admin"
		profile.ClassId = "202201234"
		profile.SubjectId = "LyGPC"
		profile.Permit = public.PermitManager
		return true, nil
	} else if userId == "user" && password == "user" {
		profile.IsLogin = true
		profile.Name = "张翼德"
		profile.UserId = "2022070230421"
		profile.ClassId = "202201234"
		profile.Permit = public.PermitUser
		return true, nil
	}
	return false, nil
}
