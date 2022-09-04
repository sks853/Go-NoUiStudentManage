** == CLI版学生考试成绩管理系统 == **

# 信息系统

## （一）系统介绍

负责登录校验、注册账户、修改账户信息等
	
## （二）通信格式 MsgStruct struct

### 1. 通信确认码 CodeMsg byte

可主动设置或由通信中心系统提供，参考通信系统 - 通信确认码

### 2. 通信用户码 CodeUser string

由通信中心系统提供，返回结果时带上该参数，参考通信系统 - 通信用户码

### 3. 权限码 codePrivilege byte

由通信中心系统提供，参考通信系统 - 权限码

### 4. 功能码 CodeFunc byte

1. `0x00` - 登录验证
2. `0x01` - 注册用户
3. `0x02` - 修改信息
4. `0x03` - 修改密码
5. `0x04` - 查询用户

### 5. 数据信息 DataInfo struct

#### 操作用户 OperationUser string

由通信中心系统提供，用于判断当前操作用户是否合法

#### 其他数据 OtherData any

参考各功能数据信息格式部分

## （三）各功能数据信息格式

### 1. 登录验证功能

#### 账号/学工号 InfoUser string

#### 密码 InfoPassword string

### 2. 注册用户功能

#### 账号/学工号 InfoUser string

#### 密码 InfoPassword string

#### 权限 InfoAccess byte

#### 姓名 InfoName string

#### 性别 InfoGender byte

#### 专业号 InfoMajor int

#### 班级号 InfoClass int

### 3. 修改信息功能

#### 账号/学工号 InfoUser string

#### 姓名 InfoName string

#### 性别 InfoGender byte

#### 专业号 InfoMajor int

#### 班级号 InfoClass int

### 4. 修改密码功能

#### 账号/学工号 InfoUser string

#### 原密码 InfoPasswordOld string

#### 新密码 InfoPasswordNew string

### 5. 查询用户功能

#### 账号/学工号 InfoUser string

## （四）简明流程

依次检查通信确认码、通信用户码、功能码从而确定执行的功能，并从data里面提取相关参数。

### 功能入口

[0] 通信结构体 MsgStruct struct

[1] 通信确认码 CodeMsg byte

[2] 通信用户码 CodeUser string

[3] 权限码 codePrivilege byte

[4] 功能码 CodeFunc byte

[5] 数据信息 DataInfo struct

[5][1] 操作用户 OperationUser string

[5][2] 其他数据 OtherData any


def func(args) -> struct {
1.	判断`操作用户`的`权限`和`被操作用户`的`权限`
2.	如果`被操作用户`的`权限`大于等于`操作用户`的`权限`则拒绝操作
3.	如果`操作用户`不为`管理员`且`功能码`为`修改信息`或`注册用户`或`查询用户`则拒绝操作
4.	如果`操作用户`不为`管理员`且`功能码`为`修改密码`且`被操作用户`不为自己的则拒绝操作
5.	`通信格式` = new MsgStruct struct
5.	if `功能码` == `登录验证` {
		执行对应功能
	} else if `功能码` == `注册用户` {
		执行对应功能
	} else if `功能码` == `修改信息` {
		执行对应功能
	} else if `功能码` == `修改密码` {
		执行对应功能
	} else if `功能码` == `查询用户` {
		执行对应功能
	}
}

### 登录验证功能

[1] 学工号 InfoUser string

[2] 密码 InfoPassword string

def func(args) -> byte {

1.  将`学工号`哈希摘要后得到`字符数组`的`Hash@学工号`，将`密码`哈希后得到`Hash@密码`
2.  逐字符检索机密目录下的文件夹名并进入，直至`字符数组`结束
	for x in `Hash@学工号`{
		if 任一层名为`x`的文件夹`不存在` or `无法打开` {
			return `0x03`
		}
		cd dir(x)
	}
	if `Hash@密码` != `该文件夹下已存在的文件名` {
		return `0x03`
	} else {
		打开`密码文件`，获取`权限`值
	}
	return `权限`
}

### 注册用户功能

[1] 学工号 InfoUser string

[2] 密码 InfoPassword string

[3] 权限 InfoAccess byte

[4] 姓名 InfoName string

[5] 性别 InfoGender byte

[6] 专业号 InfoMajor int

[7] 班级号 InfoClass int

def func(args) -> void {

	//注册机密信息
1.  将`学工号`哈希后得到`字符数组`的`Hash@学工号`，将`密码`哈希后得到`Hash@密码`
2.  逐字符检索`机密目录`下的文件夹名并进入，直至`字符数组`结束
	for x in `Hash@学工号` {
		if 任一层名为`x`的文件夹`不存在` {
			mkdir dir(x)
		}
		cd dir(x)
	}
3.  create `随机名字`文件
4.  write `权限` and save
5.	rename `随机名字`文件为`Hash@密码`文件
		
	//注册用户信息
6.  from `学工号` get `字符数组`
7.  逐字符检索`账号目录`下的文件夹名并进入，直至`字符数组`结束
	for x in `字符数组` {
		if 任一层名为`x`的文件夹`不存在` {
			mkdir dir(x)
		}
		cd dir(x)
	}
8.  create `随机名字`文件
9.  write `学工号`, `姓名`, `性别`, `专业号`, `班级号` and save
10.	rename `随机名字`文件为`学工号`文件
}

### 修改信息功能

[1] 学工号 InfoUser string

[2] 姓名 InfoName string

[3] 性别 InfoGender byte

[4] 专业号 InfoMajor int

[5] 班级号 InfoClass int

def func(args) -> bool {

1.  from `学工号` get `字符数组`
2.  逐字符检索`账号目录`下的文件夹名并进入，直至`字符数组`结束
	for x in `字符数组` {
		if 任一层名为`x`的文件夹`不存在` {
			return False
		}
		cd dir(x)
	}
3.  create `随机名字`文件
4.	write `学工号`, `姓名`, `性别`, `专业号`, `班级号` and save
5.	rename `随机名字`文件为`学工号`文件
}

### 修改密码功能

[1] 学工号 InfoUser string

[2] 原密码 InfoPasswordOld string

[3] 新密码 InfoPasswordNew string

def func(args) -> bool {

1.  将`学工号`哈希后得到`字符数组`的`Hash@学工号`，将`原密码`和`新密码`哈希后得到`Hash@原密码`和`Hash@新密码`
2.  逐字符检索`机密目录`下的文件夹名并进入，直至`字符数组`结束
	for x in `字符数组` {
		if 任一层名为`x`的文件夹`不存在` {
			return False
		}
		cd dir(x)
	}
	if `Hash@密码` != `该文件夹下已存在的文件名` {
		return False
	}
3.	打开`Hash@原密码`文件，获取`权限`值后关闭
4. 	rename `原密码文件`加上前缀`双下划线`
5.	create `随机名字`文件
6.	write `权限` and save
7.	rename `随机名字`文件为`Hash@新密码`文件
8.	return True
}

### 查询用户功能

[1] 学工号 InfoUser string

def func(args) -> struct {
1.	valueStruct = new struct
2.	将`学工号`哈希后得到`字符数组`的`Hash@学工号`
3.	逐字符检索`账号目录`下的文件夹名并进入，直至`字符数组`结束
	for x in `字符数组` {
		if 任一层名为`x`的文件夹`不存在` {
			raise Exception
		}
		cd dir(x)
	}
4.  复制、重命名为`随机字符串`并打开`学工号`文件，找不到则`raise Exception`
5.	根据`一定的格式`读取用户信息，存入`valueStruct`
6.	关闭文件，删除复制后的文件
7.	逐字符检索`机密目录`下的文件夹名并进入，直至`字符数组`结束
	for x in `字符数组` {
		if 任一层名为`x`的文件夹`不存在` {
			raise Exception
		}
		cd dir(x)
	}
8.	遍历文件夹下每一个文件名，找到第一个不含`双下划线`前缀的文件
9.	复制、重命名为`随机字符串`并打开该文件，找不到则抛出异常
10.	根据`一定的格式`读取`权限`，存入`valueStruct`
11.	关闭文件，删除复制后的文件
12.	return `valueStruct`
}
