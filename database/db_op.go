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

import (
	"NoUiStudentManage/public"
	"bufio"
	"os"
)

func FsReadLine(fs *os.File) *[]string {
	var strList []string
	fr := bufio.NewReader(fs)
	for {
		txt, err := fr.ReadString('\n')
		if err != nil {
			break
		}
		strList = append(strList, txt)
	}
	return &strList
}

func FsFormatLine() {

}

func FsReadUserProfile(pathProfile string, profile *public.StructProfile) error {
	fs, err := os.OpenFile("", os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	defer func(fs *os.File) {
		err = fs.Close()
		if err != nil {
		}
	}(fs)

	strList := FsReadLine(fs)
	for _, txt := range *strList {
		if txt == "-" {
			// TODO 过滤用户用户信息
		}
	}

	err = fs.Close()
	if err != nil {
		return err
	}
	return nil
}
