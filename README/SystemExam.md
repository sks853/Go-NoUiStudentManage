** == CLI版学生考试成绩管理系统 == **

# 考试系统

## （一）系统介绍

处理考试和成绩相关信息

## （二）通信格式 MsgStruct struct

### 1. 通信确认码 CodeMsg byte

可主动设置或由通信中心系统提供，参考通信系统 - 通信确认码

### 2. 通信用户码 CodeUser string

由通信中心系统提供，返回结果时带上该参数，参考通信系统 - 通信用户码

### 3. 权限码 codePrivilege byte

由通信中心系统提供，参考通信系统 - 权限码

### 4. 功能码 CodeFunc byte

1. `0x00` - 创建考试
2. `0x01` - 取消考试
3. `0x02` - 保留（修改考试）
4. `0x03` - 查询考试
5. `0x10` - 录入成绩
6. `0x11` - 删除成绩
7. `0x12` - 查询成绩
8. `0x13` - 保留（修改成绩）

### 5. 数据信息 DataInfo any

参考各功能数据信息格式部分

## （三）各功能数据信息格式

### 1. 创建考试功能

#### 考试日期时间 ExamDatetime string

#### 创建人 ExamAdministrator string

#### 考试科目数目 ExamSubjectCount int

#### 考试科目编号 ExamSubjectId []string

#### 参加考试人数 ExamPeopleCount int

#### 参加考试名单 ExamPeopleId []string

### 2. 取消考试功能

#### 考试编号 ExamId string

### 3. 修改考试功能（保留）

#### 考试编号 ExamId string

#### 考试状态 ExamStatus bool

#### 考试时间 ExamDatetime string

### 4. 查询考试功能

#### 考试编号 ExamId string

#### 考试创建者 ExamAdministrator string

#### 考生学工号 ExamPersonId string

#### 考试起始日期 ExamDatetimeStart string

#### 考试截至日期 ExamDatetimeEnd string

### 5. 录入成绩功能

#### 考试日期时间 ExamDatetime string

#### 创建人 ExamAdministrator string

#### 考试编号 ExamId string

#### 考试科目数目 ExamSubjectCount int

#### 考试科目编号 ExamSubjectId []string

#### 参加考试人数 ExamPeopleCount int

#### 总平均成绩 ExamAverageTotal double .2

#### 各科平均成绩 ExamAverageSubject []double .2

#### 个人成绩 ExamPersonScore []struct

##### 排名 ExamPersonRank int

##### 学工号 ExamPersonId string

##### 姓名 ExamPersonName string

##### 总成绩 ExamPersonScoreTotal double .2

##### 各科成绩 ExamPersonScoreSubject []double .2

### 6. 删除成绩功能

#### 考试编号 ExamId string

#### 考试创建者 ExamAdministrator string

#### 考试起始日期 ExamDatetimeStart string

#### 考试截至日期 ExamDatetimeEnd string

### 7. 查询成绩功能

#### 考试编号 ExamId string

#### 考生学工号 ExamPersonId string

#### 考试创建者 ExamAdministrator string

#### 考试起始日期 ExamDatetimeStart string

#### 考试截至日期 ExamDatetimeEnd string

### 8. 修改成绩功能（保留）

#### 考试编号 ExamId string

#### 考生学工号 ExamPersonId string

#### 考试科目编号 ExamSubjectId string

#### 考试科目修改后成绩 ExamScoreFix double .2

## （四）简明流程

### 功能入口

[0] 通信结构体 MsgStruct struct

[1] 通信确认码 CodeMsg byte

[2] 通信用户码 CodeUser string

[3] 权限码 codePrivilege byte

[4] 功能码 CodeFunc byte

[5] 数据信息 DataInfo struct

def func(args) -> struct {
	if `权限码` == `学生` and `功能码` != `查询考试` and `功能码` != `查询成绩` {
		return new struct `通信格式`.set(`通信确认码`=`失败`, `通信用户码`=`通信用户码`)
	}
	if `权限码` == `教师` and `功能码` == ``
}

### 创建考试功能

[1] 考试日期时间 ExamDatetime string

[2] 创建人 ExamAdministrator string

[3] 考试科目数目 ExamSubjectCount int

[4] 考试科目编号 ExamSubjectId []string

[5] 参加考试人数 ExamPeopleCount int

[6] 参加考试名单 ExamPeopleId []string

def func(args) -> bool {
	create `考试编号` as `ExamId`
	with open `考试目录->日期映射->日期映射.csv` as f {
		f.write({`考试状态`=`等待`, `考试日期`, `考试编号`, `创建者账号`})
	}
	with open `考试目录->日期映射->年->月->日->考试编号.csv` as f {
		f.write(`参见成绩结构无成绩格式`)
	}
	for in range `参加考试名单账号` as `ExamPeopleId` {
		with open `考试目录->学号映射->参加考试名单账号` as f {
			add `当前记录`
		}
		f.write(`新增该记录后的文件内容`)
	}
}

### 取消考试功能

[1] 考试编号 ExamId string

def func(args) -> bool {
	with open `考试目录->日期映射->日期映射.csv` as f {
		`记录` = new list()
		while record = f.readline() != EOF {
			if exist `考试编号` in record {
				contine
			}
			`记录`.append({`考试编号`, `考试日期`})
		}
		f.write(`记录`)
	}
	with open `考试目录->日期映射->年->月->日->考试编号.csv` as f {
		get `考试人数` as `ExamPeopleCount`
		for in range `考试人数` {
			get `参加考试名单账号` as `ExamPeopleId[]`
		}
	}
	del `考试编号.csv`
	for in range `参加考试名单账号` as `ExamPeopleId` {
		with open `考试目录->学号映射->参加考试名单账号` as f {
			if `考试编号` == `行子元素内容` {
				del `当前记录`
			}
		}
		f.write(`删除该记录后的文件内容`)
	}
	return True
}

### 修改考试功能（保留）

[1] 考试编号 ExamId string

[2] 考试状态 ExamStatus bool

[3] 考试时间 ExamDatetime string

def func(args) -> bool {

}

### 查询考试功能

[1] 考试编号 ExamId string

[2] 考试创建者 ExamAdministrator string

[3] 考生学工号 ExamPersonId string

[4] 考试起始日期 ExamDatetimeStart string

[5] 考试截至日期 ExamDatetimeEnd string

def func(args) -> bool {

}

### 录入成绩功能

[1] 考试日期时间 ExamDatetime string

[2] 创建人 ExamAdministrator string

[3] 考试编号 ExamId string

[4] 考试科目数目 ExamSubjectCount int

[5] 考试科目编号 ExamSubjectId []string

[6] 参加考试人数 ExamPeopleCount int

[7] 总平均成绩 ExamAverageTotal double .2

[8] 各科平均成绩 ExamAverageSubject []double .2

[9] 个人成绩 ExamPersonScore []struct

[9][1] 排名 ExamPersonRank int

[9][2] 学工号 ExamPersonId string

[9][3] 姓名 ExamPersonName string

[9][4] 总成绩 ExamPersonScoreTotal double .2

[9][5] 各科成绩 ExamPersonScoreSubject []double .2

def func(args) -> bool {
	with open `考试目录->日期映射->日期映射.csv` as f {
		if `考试编号` == `行子元素内容` {
				set `考试状态`=`未批改`
			}
		f.write(`写回`)
	}
	with open `考试目录->日期映射->年->月->日->考试编号.csv` as f {
		f.write(`参见成绩结构格式`)
	}
}

### 删除成绩功能

[1] 考试编号 ExamId string

[2] 考试创建者 ExamAdministrator string

[3] 考试起始日期 ExamDatetimeStart string

[4] 考试截至日期 ExamDatetimeEnd string

def func(args) -> bool {

}

### 查询成绩功能

[1] 考试编号 ExamId string

[2] 考生学工号 ExamPersonId string

[3] 考试创建者 ExamAdministrator string

[4] 考试起始日期 ExamDatetimeStart string

[5] 考试截至日期 ExamDatetimeEnd string

def func(args) -> bool {

}

### 修改成绩功能

[1] 考试编号 ExamId string

[2] 考生学工号 ExamPersonId string

[3] 考试科目编号 ExamSubjectId string

[4] 考试科目修改后成绩 ExamScoreFix double .2

def func(args) -> bool {

}
