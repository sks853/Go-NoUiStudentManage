/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2022-02-11 21:52
 * @FileName     log_system.go
 * @Description  None
 * ==================================================
**/

package main

import (
	"NoUiStudentManage/public"
	"net"
	"os"
	"sync"
	"time"
)

var mutex sync.Mutex
var log *os.File

func ServerRun() {
	log = getLogFile("./log")
	server, err := net.ResolveUDPAddr("udp4", "127.0.0.1:32111")
	if err != nil {
		public.TipWait("开启日志系统时出现错误，即将关闭系统")
		os.Exit(0)
	}
	conn, err := net.ListenUDP("udp", server)
	for {
		disposeSocket(conn)
	}
}

func disposeSocket(conn *net.UDPConn) {
	var dataByte [1024]byte
	count, _, err := conn.ReadFromUDP(dataByte[:])
	if err != nil {
		public.TipWait("接收日志操作请求时出现错误")
		return
	}
	go logWrite(dataByte[:count])
}

func logWrite(data []byte) {
	// TODO 尚未解决跨日问题，可以通过在此调用查询文件名是否匹配当天日期而解决
	if log == nil {
		return
	}
	mutex.Lock()
	timeStr := time.Now().Format("[2006-01-02 15:04:05] ")
	_, err := log.WriteString(timeStr)
	_, err = log.Write(data)
	_, err = log.Write([]byte{'\r', '\n'})
	err = log.Sync()
	if err != nil {
		mutex.Unlock()
		return
	}
	mutex.Unlock()
}

func isFileExist(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	return false
}

func getLogFile(path string) *os.File {
	var logFile *os.File
	fileName := time.Now().Format("2006-01-02") + ".log"
	filePath := path + "/" + fileName
	flag := isFileExist(filePath)
	if flag {
		logFileTemp, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil
		}
		logFile = logFileTemp
	} else {
		logFileTemp, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil
		}
		logFile = logFileTemp
	}

	return logFile
}
