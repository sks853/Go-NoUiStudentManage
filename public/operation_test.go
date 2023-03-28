/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2023-02-08 10:47
 * @FileName     operation_test.go
 * @Description  None
 * ==================================================
**/

package public

import (
	"errors"
	"github.com/ser163/WordBot/generate"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestOpUserGet(t *testing.T) {
	type args struct {
		pf     *StructProfile
		df     *StructDataFile
		userId string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  error
		wantFlag bool
	}{
		{
			name: "获取已存在的用户信息",
			args: args{
				pf:     new(StructProfile),
				df:     new(StructDataFile),
				userId: "bLukYUqrNUXjO",
			},
			wantErr:  nil,
			wantFlag: true,
		},
		{
			name: "获取不存在的用户信息",
			args: args{
				pf:     new(StructProfile),
				df:     new(StructDataFile),
				userId: "abc",
			},
			wantErr:  nil,
			wantFlag: false,
		},
		{
			name: "获取空参数用户信息",
			args: args{
				pf:     new(StructProfile),
				df:     new(StructDataFile),
				userId: "",
			},
			wantErr:  errors.New(""),
			wantFlag: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("待查询用户编号：%s", tt.args.userId)
			gotErr, gotFlag := OpUserGet(tt.args.pf, tt.args.userId, tt.args.df.BlockFill)
			if tt.wantErr == nil && gotErr != nil {
				t.Errorf("验证失败，操作失败，异常信息：%s", gotErr.Error())
			} else if tt.wantErr != nil && gotErr == nil {
				t.Errorf("验证失败，操作成功，请逐步调试排查。")
			} else if tt.wantErr != nil && gotErr != nil {
				t.Logf("验证成功，操作失败，异常：%s", gotErr)
			} else if tt.wantErr == nil && gotErr == nil && tt.wantFlag == gotFlag {
				if gotFlag {
					t.Logf("验证成功，操作成功，找到用户信息：%s", tt.args.pf.Print())
				} else {
					t.Logf("验证成功，操作失败，未找到用户信息。")
				}
			} else if tt.wantErr == nil && gotErr == nil && tt.wantFlag != gotFlag {
				if gotFlag {
					t.Errorf("验证失败，操作成功，期望无法找到但却能找到用户信息：%s", tt.args.pf.Print())
				} else {
					t.Errorf("验证失败，操作失败，期望能找到但未能找到用户信息。")
				}
			}
		})
	}
}

func TestOpUserSet(t *testing.T) {
	pf := new(StructProfile)

	type args struct {
		pf      []*StructProfile
		df      *StructDataFile
		setFunc []StructSetProfile
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "修改用户编号",
			args: args{
				pf:      []*StructProfile{pf, {UserId: "rootRootRootRoot"}},
				df:      new(StructDataFile),
				setFunc: []StructSetProfile{{pf.ProfileSetUserId, "root"}},
			},
			wantErr: true,
		},
		{
			name: "无内容修改用户编号",
			args: args{
				pf:      []*StructProfile{pf, {UserId: "root"}},
				df:      new(StructDataFile),
				setFunc: []StructSetProfile{{pf.ProfileSetUserId, ""}},
			},
			wantErr: false,
		},
		{
			name: "越长度限制修改用户编号",
			args: args{
				pf:      []*StructProfile{pf, {UserId: "root"}},
				df:      new(StructDataFile),
				setFunc: []StructSetProfile{{pf.ProfileSetUserId, "rootRootRootRootRoot"}},
			},
			wantErr: false,
		},
		{
			name: "恰好长度限制修改用户编号",
			args: args{
				pf:      []*StructProfile{pf, {UserId: "root"}},
				df:      new(StructDataFile),
				setFunc: []StructSetProfile{{pf.ProfileSetUserId, "rootRootRootRoot"}},
			},
			wantErr: true,
		},
		{
			name: "修改用户唯一标识码",
			args: args{
				pf:      []*StructProfile{pf, {UserId: "rootRootRootRoot"}},
				df:      new(StructDataFile),
				setFunc: []StructSetProfile{{pf.ProfileSetUid, "B6795C555DC1303D8DEC64647C99CD1E"}},
			},
			wantErr: true,
		},
		{
			name: "修改用户的权限",
			args: args{
				pf:      []*StructProfile{pf, {UserId: "rootRootRootRoot"}},
				df:      new(StructDataFile),
				setFunc: []StructSetProfile{{pf.ProfileSetPermit, 3}},
			},
			wantErr: true,
		},
		{
			name: "修改用户姓名",
			args: args{
				pf:      []*StructProfile{pf, {UserId: "rootRootRootRoot"}},
				df:      new(StructDataFile),
				setFunc: []StructSetProfile{{pf.ProfileSetName, "刘玄德"}},
			},
			wantErr: true,
		},
		{
			name: "同时修改用户姓名和权限",
			args: args{
				pf:      []*StructProfile{pf, {UserId: "username"}},
				df:      new(StructDataFile),
				setFunc: []StructSetProfile{{pf.ProfileSetName, "曹操"}, {pf.ProfileSetPermit, 1}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.pf[0].Copy(tt.args.pf[1])
			err := OpUserSet(tt.args.pf[0], tt.args.df, tt.args.setFunc...)
			if tt.wantErr && err == nil {
				t.Logf("验证成功，操作成功，当前用户信息：%s", tt.args.pf[0].Print())
			} else if !tt.wantErr && err != nil {
				t.Logf("验证成功，操作失败，异常信息：%s", err.Error())
			} else if tt.wantErr && err != nil {
				t.Errorf("验证失败，操作失败，异常信息：%s", err.Error())
			} else if !tt.wantErr && err == nil {
				t.Errorf("验证失败，操作成功，请逐步调试排查。")
			}
		})
	}
}

func TestOpUserNew(t *testing.T) {
	type args struct {
		pf *StructProfile
		df *StructDataFile
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "标准新建用户信息",
			args: args{
				pf: &StructProfile{
					UID:        "",
					UserId:     GenerateWord(4, 16),
					Name:       GenerateName(false),
					Permit:     rand.Intn(3),
					ClassId:    GenerateWord(4, 8),
					ActiveCode: "",
				},
				df: new(StructDataFile),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.pf.UID = GenerateUpperUUID(tt.args.pf.UserId)
			tt.args.pf.ActiveCode = GenerateUpperUUID("")
			if err := OpUserNew(tt.args.pf); (err != nil) != tt.wantErr {
				t.Errorf("OpUserNew() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func GenerateName(isMale bool) string {
	lastName := []string{
		"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "褚", "卫", "蒋", "沈", "韩", "杨", "朱", "秦", "尤", "许", "何", "吕", "施",
		"张", "孔", "曹", "严", "华", "金", "魏", "陶", "姜", "戚", "谢", "邹", "喻", "柏", "水", "窦", "章", "云", "苏", "潘", "葛", "奚", "范",
		"彭", "郎", "鲁", "韦", "昌", "马", "苗", "凤", "花", "方", "俞", "任", "袁", "柳", "酆", "鲍", "史", "唐", "费", "廉", "岑", "薛", "雷",
		"贺", "倪", "汤", "滕", "殷", "罗", "毕", "郝", "邬", "安", "常", "乐", "于", "时", "傅", "皮", "卞", "齐", "康", "伍", "余", "元", "卜",
		"顾", "孟", "平", "黄", "和", "穆", "萧", "尹", "姚", "邵", "堪", "汪", "祁", "毛", "禹", "狄", "米", "贝", "明", "臧", "计", "伏", "成",
		"戴", "谈", "宋", "茅", "庞", "熊", "纪", "舒", "屈", "项", "祝", "董", "梁",
	}
	male := "辰,士,以,建,家,致,树,炎,德,行,时,泰,盛,雄,琛,钧,冠,策,铭,腾,伟,刚,勇,毅,俊,峰,强,军,平,保,东,文,辉,力,明,永,健,世,广," +
		"志,义,兴,良,海,山,仁,波,宁,贵,福,生,龙,元,全,国,胜,学,祥,才,发,成,康,星,光,天,达,安,岩,中,茂,武,新,利,清,飞,彬,富,顺," +
		"信,子,杰,楠,榕,风,航,弘,鼎,启"

	female := "嘉,琼,桂,娣,叶,璧,璐,娅,琦,晶,妍,茜,秋,珊,莎,锦,黛,青,倩,婷,姣,婉,娴,瑾,颖,露,瑶,怡,婵,雁,蓓,纨,仪,荷,丹,蓉,眉,君," +
		"琴,蕊,薇,菁,梦,岚,苑,婕,馨,瑗,琰,韵,融,园,艺,咏,卿,聪,澜,纯,毓,悦,昭,冰,爽,琬,茗,羽,希,宁,欣,飘,育,滢,馥,筠,柔,竹," +
		"霭,凝,晓,欢,霄,枫,芸,菲,寒,伊,亚,宜,可,姬,舒,影,荔,枝,思,丽,秀,娟,英,华,慧,巧,美,娜,静,淑,惠,珠,翠,雅,芝,玉,萍,红," +
		"娥,玲,芬,芳,燕,彩,春,菊,勤,珍,贞,莉,兰,凤,洁,梅,琳,素,云,莲,真,环,雪,荣,爱,妹,霞,香,月,莺,媛,艳,瑞,凡,佳,轩"
	lsNameFirstMale := strings.Split(male, ",")
	lsNameFirstFemale := strings.Split(female, ",")
	rand.Seed(time.Now().UnixNano())
	nameLength := rand.Intn(2) + 1
	name := lastName[rand.Intn(len(lastName))]
	for i := 0; i < nameLength; i++ {
		if isMale {
			name += lsNameFirstMale[rand.Intn(len(lsNameFirstMale))]
		} else {
			name += lsNameFirstFemale[rand.Intn(len(lsNameFirstFemale))]
		}
	}
	time.Sleep(1)
	return name
}

func GenerateWord(minLen int, maxLen int) string {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(1)
	userId, _ := generate.GenRandomMix(rand.Intn(maxLen-minLen) + minLen)
	return userId.Word
}

func TestOpUserDel(t *testing.T) {
	type args struct {
		pf *StructProfile
		df *StructDataFile
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OpUserDel(tt.args.pf); (err != nil) != tt.wantErr {
				t.Errorf("OpUserDel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOpUserPasswdVerify(t *testing.T) {
	type args struct {
		uid       string
		password  string
		blockFill string
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 bool
	}{
		{
			name: "已存在的正确信息判断",
			args: args{
				uid:       "982C871A1DC358D2B5E9309FE16ED6E3",
				password:  "MtneJACaIP",
				blockFill: "?",
			},
			want:  nil,
			want1: true,
		},
		{
			name: "未存在的正确信息判断",
			args: args{
				uid:       "982C871A1DC358D2B5E9309FE16ED6E3",
				password:  "abc",
				blockFill: "?",
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OpUserPasswdVerify(tt.args.uid, tt.args.password, tt.args.blockFill)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpUserPasswdVerify() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("OpUserPasswdVerify() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOpUserPasswdNew(t *testing.T) {
	type args struct {
		uid      string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "标准用户校验新增",
			args: args{
				uid:      "982C871A1DC358D2B5E9309FE16ED6E3",
				password: GenerateWord(4, 12),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OpUserPasswdNew(tt.args.uid, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("OpUserPasswdNew() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOpUserPasswdSet(t *testing.T) {
	type args struct {
		uid         string
		passwordNew string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "修改已存在UID的密码",
			args: args{uid: "982C871A1DC358D2B5E9309FE16ED6E3", passwordNew: "12345"},
		},
		{
			name:    "修改不存在UID的密码",
			args:    args{uid: "782C871A1DC358D2B5E9309FE16ED6E3", passwordNew: "12345"},
			wantErr: true,
		},
		{
			name:    "修改已存在UID的超过长度限制的密码",
			args:    args{uid: "982C871A1DC358D2B5E9309FE16ED6E3", passwordNew: "12345678901234567890"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OpUserPasswdSet(tt.args.uid, tt.args.passwordNew); (err != nil) != tt.wantErr {
				t.Errorf("OpUserPasswdSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
