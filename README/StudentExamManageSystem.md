** == CLI版学生考试成绩管理系统 == **

# 一、全局初始化

##################################################

# 二、通信中心系统
	第二优先启动顺序，负责各通道的数据交换

## （一）通信格式（与用户操作）

### 通信确认码 CodeMsg byte
1. `0x00` - 确认 / 准备完成 / 查询但存在结果 / 查询但不存在结果
2. `0x01` - 无法操作 / 出现错误 / 抛出异常 / 无法读取 / 无法写入 / 文件缺失
3. `0x02` - 请求退出 / 关闭子线程

### 通信类型 MsgType byte
1. `0x01` - 日志系统
2. `0x02` - 鉴权系统
3. `0x03` - 行政系统
4. `0x04` - 考试系统
5. `0x05` - 邮件系统

### 通信内容 MsgContent any
参考子系统 - 通信格式的相关内容参数

## （二）简明流程
1. 等待各子系统完成基本初始化接收指令
2. 告知客户端允许请求登录，并等待客户端通道
3. 收到客户端登录请求后转发给鉴权系统，并等待鉴权系统通道
4. 验证通过后生成唯一标识权限码（便于后期结合Redis拓展）
5. 将权限码和用户标识号绑定，并返回相应权限菜单的组合，等待客户端通道
6. 无限循环等待通道消息并解析：
	--  判断客户端通道消息的通信类型
	--	if 退出：
			依次通知（`CodeMsg`==`0x02`）并等待各子系统退出（`CodeMsg`==`0x00`）
		elif 子系统：
			判断权限是否存在于所属权限变量中(可选，当前阶段无需考虑)
			随机产生通信用户码，并和用户通道绑定（便于日后拓展成多用户登录）
			整理格式，将数据直接传递给对应的子系统
	--  等待子系统通道，并返回结果数据给客户端

##################################################

# 三、日志系统
	日志输出等记录（保证最优先启动顺序，倘若启动失败则直接终止程序）

## （一）通信格式

### 通信确认码 CodeMsg byte
参考通信系统 - 通信确认码

### 日志类型 LogType byte
1. `0x00` - 调试
2. `0x01` - 普通
3. `0x02` - 警告
4. `0x03` - 出错

### 日志内容 LogMsg string

### 报错内容 LogErr string
当且仅当`LogType` == `警告` or `出错` 时有效
	
## （二）简明流程
1. 检测日志目录是否存在、能否访问
2. 判断当天日期的日志文件是否存在
	-- 存在则打开追加
	-- 不存在则创建并打开
3. 无限循环等待通道消息
	--  判断消息类型：退出、普通、警告、出错
	--	if 退出：
			跳出循环
		else：
			写入日志文件
4. 关闭文件
5. 通知通信中心并设置通信确认码为`0x00`、退出函数

##################################################

# 四、鉴权系统
	负责登录校验、注册账户、修改账户信息等
	
## （一）通信格式

### 通信确认码 CodeMsg byte
参考通信系统 - 通信确认码

### 通信用户码 CodeUser string
由通信中心系统提供，返回结果时带上该参数

### 功能码 CodeFunc byte
1. `0x00` - 登录验证
2. `0x01` - 注册新用户
3. `0x02` - 修改账户信息
4. `0x03` - 修改密码

### 数据信息 DataInfo struct
参考各功能数据信息格式部分

## （二）各功能数据信息格式

### 1. 登录验证功能

#### 账号/学工号 DataUser string

#### 密码 DataPassword string

### 2. 注册新用户功能

#### 是否覆盖 IsCover bool

#### 账号/学工号 DataUser string

#### 密码 DataPassword string

#### 权限 DataAccess byte

#### 姓名 DataName string

#### 性别 DataGender byte

#### 专业号 DataMajor int

#### 班级号 DataClass int

### 3. 修改账户信息功能

#### 账号/学工号 DataUser string

#### 姓名 DataName string

#### 性别 DataGender byte

#### 专业号 DataMajor int

#### 班级号 DataClass int

### 4. 修改密码功能

#### 账号/学工号 DataUser string

#### 密码 DataPassword string

#### 权限 DataAccess byte

## （三）简明流程
	依次检查通信确认码、通信用户码、功能码从而确定执行的功能，并从data里面提取相关参数。

### 登录验证功能
1. 将账号哈希摘要后得到字符数组
2. 根据字符数组逐层检索账号目录，直至最底层
	if 某一层不存在：
		return 空结果
	else：
		直到底层的文件夹后，且存在相同的文件，判断摘要密码是否匹配文件名
		if 密码不匹配：
			return 空结果
		else：
			打开文件，获取权限值
	return 结果
3. 功能执行完毕

### 注册新用户功能

#### 注册机密信息
1. 将账号哈希摘要后得到字符数组
2. 依次检索机密目录下，逐层字符文件夹，直至最底层
	if 某一层不存在：
		创建后续层数的文件夹
	else：
		直到底层的文件夹后，且存在相同的文件，判断是否允许覆盖
		if 不允许覆盖：
			return 错误和空结果
		else：
			next
3. 覆写密码文件（不存在自动创建），文件名为哈希后的密码（考虑RSA加密但需要解码后再比对），文件内容为权限值
		
#### 注册用户信息
4. 根据账号得到字符数组
5. 根据字符数组逐层检索账号目录，直至最底层
	if 某一层不存在：
		创建后续层数的文件夹
	else：
		直到底层的文件夹后，且存在相同的文件，判断是否允许覆盖
		if 不允许覆盖：
			return 错误和空结果
		else：
			next
6. 覆写用户信息文件，文件名为账号，文件内容为`个人信息`格式的内容
7. 功能执行完毕

### 修改账户信息功能
1. 根据账号得到字符数组
2. 根据字符数组逐层检索账号目录，直至最底层
	if 某一层不存在：
		return 错误和空结果
	else：
		直到底层的文件夹后，无论是否存在该账号文件，根据格式直接覆写传递内容
3. 功能执行完毕

### 修改密码功能
1. 根据账号哈希摘要得到字符数组、分别哈希摘要新旧密码后得到摘要密码 (<=128)
2. 依次检索机密目录下，逐层字符文件夹，直至最底层
	if 某一层不存在：
		return 错误和空结果
	else：
		直到底层的文件夹后，且存在该命名为旧摘要密码的文件
		if 不存在：
			return 错误和空结果
		打开该文件获取权限值并关闭
		重命名该文件加上前缀`双下划线`以便和新密码文件区分
		创建新文件命名为新的摘要密码并写入权限值
3. 功能执行完毕

##################################################

# 五、行政系统
	处理班集体、专业、科目等信息

## （一）通信格式

### 通信确认码 CodeMsg byte
参考通信系统 - 通信确认码

## （二）简明流程

##################################################

# 六、考试系统
	处理考试和成绩相关信息
	
## （一）通信格式

### 通信确认码 CodeMsg byte
参考通信系统 - 通信确认码

### 通信用户码 CodeUser string
由通信中心系统提供，返回结果时带上该参数

### 功能码 CodeFunc byte
1. `0x00` - 创建考试
2. `0x01` - 取消考试
3. `0x02` - 保留（修改考试）
4. `0x03` - 查询考试
5. `0x10` - 录入成绩
6. `0x11` - 删除成绩
7. `0x12` - 查询成绩
8. `0x13` - 保留（修改成绩）

### 数据信息 DataInfo struct
参考各功能数据信息格式部分

## （二）各功能数据信息格式

### 1. 创建考试功能

#### 考试日期时间 ExamDatetime string

#### 创建人 ExamAdministrator string

#### 考试科目数目 ExamSubjectCount int

#### 考试科目编号 ExamSubjectId []string

#### 参加考试人数 ExamPeopleCount int

#### 参加考试名单 ExamPeopleId []string

### 2. 取消考试

#### 考试编号 ExamId string

### 3. 修改考试（保留）

#### 考试编号 ExamId string

#### 考试状态 ExamStatus bool

#### 考试时间 ExamDatetime string

### 4. 查询考试

#### 考试编号 ExamId string

#### 考试创建者 ExamAdministrator string

#### 考生学工号 ExamPersonId string

#### 考试起始日期 ExamDatetimeStart string

#### 考试截至日期 ExamDatetimeEnd string

### 5. 录入成绩

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

### 6. 删除成绩

#### 考试编号 ExamId string

#### 考试创建者 ExamAdministrator string

#### 考试起始日期 ExamDatetimeStart string

#### 考试截至日期 ExamDatetimeEnd string

### 7. 查询成绩

#### 考试编号 ExamId string

#### 考生学工号 ExamPersonId string

#### 考试创建者 ExamAdministrator string

#### 考试起始日期 ExamDatetimeStart string

#### 考试截至日期 ExamDatetimeEnd string

### 8. 修改成绩（保留）

#### 考试编号 ExamId string

#### 考生学工号 ExamPersonId string

#### 考试科目编号 ExamSubjectId string

#### 考试科目修改后成绩 ExamScoreFix double .2

## （三）简明流程

##################################################

# 七、邮件系统
	
## （一）通信格式

### 通信确认码 CodeMsg byte
参考通信系统 - 通信确认码

### 通信用户码 CodeUser string
由通信中心系统提供，返回结果时带上该参数

## （二）简明流程

##################################################

# 八、客户端系统
	交互式客户端操作

## （一）通信格式

### 通信确认码 CodeMsg byte
参考通信系统 - 通信确认码

### 通信用户码 CodeUser string
由通信中心系统提供，返回结果时带上该参数

## （二）简明流程