/**
 * ==================================================
 * @Author       HDM
 * @Copyright    HDM
 * @Blogs        https://blog.csdn.net/qq_43634664
 * @QQ           1695692119
 * @Version      0.0.1
 * @Time         2023-02-15 21:48
 * @FileName     model_op_test.go
 * @Description  None
 * ==================================================
**/

package public

import (
	"os"
	"testing"
)

func TestStructDataFile_DataConfigSet(t *testing.T) {
	path := t.TempDir()
	fileName := "tmp"
	tmpFile, err := os.OpenFile(path+"/"+fileName+"_config", os.O_CREATE, 0644)
	if err != nil {
		t.Fatalf("Cannot create new temp file. Err: %s", err.Error())
	}
	_ = tmpFile.Close()
	type fields struct {
		BlockFill    string
		SortField    int
		CountRow     int64
		ColumnLength []int
		Encoding     int
		File         *os.File
		FilePath     string
		FileName     string
		IsSort       bool
	}
	type args struct {
		path         string
		dataFileName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "标准写入配置",
			fields: fields{
				BlockFill:    "?",
				SortField:    2,
				ColumnLength: []int{5, 4, 3, 2, 1},
				CountRow:     1154689,
				Encoding:     0,
			},
			args: args{
				path:         path,
				dataFileName: fileName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := &StructDataFile{
				BlockFill:    tt.fields.BlockFill,
				SortField:    tt.fields.SortField,
				CountRow:     tt.fields.CountRow,
				ColumnLength: tt.fields.ColumnLength,
				Encoding:     tt.fields.Encoding,
				File:         tt.fields.File,
				FilePath:     tt.fields.FilePath,
				FileName:     tt.fields.FileName,
				IsSort:       tt.fields.IsSort,
			}
			if err = df.DataConfigSet(tt.args.path, tt.args.dataFileName); (err != nil) != tt.wantErr {
				t.Errorf("DataConfigSet() error = %v, wantErr %v", err, tt.wantErr)
			}
			dfs := new(StructDataFile)
			err = dfs.DataConfigRead(path + "/" + fileName + "_config")
			if err != nil {
				t.Fatal(err)
			}
			t.Log(dfs.Print())
		})
	}
}

func TestStructScore_SetValue(t *testing.T) {
	type fields struct {
		Integer int
		Decimal [DecLength]int
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "浮点正数",
			args: args{value: "12.3"},
			want: true,
		},
		{
			name: "浮点负数",
			args: args{value: "-12.3"},
			want: true,
		},
		{
			name: "小数为0的浮点正数",
			args: args{value: "12.0"},
			want: true,
		},
		{
			name: "正整数多带几个前置0",
			args: args{value: "000012.3"},
			want: true,
		},
		{
			name: "负整数多带几个前置0且负号在0前",
			args: args{value: "000012.3"},
			want: true,
		},
		{
			name: "负整数多带几个前置0且负号在0后",
			args: args{value: "0000-12.3"},
			want: false,
		},
		{
			name: "非法数",
			args: args{value: "1.2.3"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := &StructScore{
				Integer: tt.fields.Integer,
				Decimal: tt.fields.Decimal,
			}
			if got := score.SetString(tt.args.value); got != tt.want {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
			t.Log(score.Integer, ".", score.Decimal[0], score.Decimal[1])
		})
	}
}

func TestStructScore_GetStringFormat(t *testing.T) {
	type fields struct {
		Integer int
		Decimal [DecLength]int
	}
	type args struct {
		integerBlockLen int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  string
	}{
		{
			name:   "正数：提供长度等于整数长度",
			fields: fields{Integer: 53, Decimal: [DecLength]int{1, 2}},
			args:   args{integerBlockLen: 2},
			want:   true,
			want1:  "53.12",
		},
		{
			name:   "正数：提供长度大于整数长度",
			fields: fields{Integer: 5, Decimal: [DecLength]int{1, 2}},
			args:   args{integerBlockLen: 4},
			want:   true,
			want1:  "0005.12",
		},
		{
			name:   "正数：提供长度大于整数长度且整数为0",
			fields: fields{Integer: 0, Decimal: [DecLength]int{1, 2}},
			args:   args{integerBlockLen: 3},
			want:   true,
			want1:  "000.12",
		},
		{
			name:   "正数：提供长度小于整数长度",
			fields: fields{Integer: 546, Decimal: [DecLength]int{1, 2}},
			args:   args{integerBlockLen: 2},
			want:   false,
			want1:  "546.12",
		},
		{
			name:   "负数：提供长度等于整数长度",
			fields: fields{Integer: -53, Decimal: [DecLength]int{1, 2}},
			args:   args{integerBlockLen: 3},
			want:   true,
			want1:  "-53.12",
		},
		{
			name:   "负数：提供长度大于整数长度",
			fields: fields{Integer: -5, Decimal: [DecLength]int{1, 2}},
			args:   args{integerBlockLen: 4},
			want:   true,
			want1:  "-005.12",
		},
		{
			name:   "负数：提供长度小于整数长度",
			fields: fields{Integer: -546, Decimal: [DecLength]int{1, 2}},
			args:   args{integerBlockLen: 2},
			want:   false,
			want1:  "-546.12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := &StructScore{
				Integer: tt.fields.Integer,
				Decimal: tt.fields.Decimal,
			}
			got, got1 := score.GetStringFormat(tt.args.integerBlockLen)
			if got != tt.want {
				t.Errorf("GetStringFormat() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetStringFormat() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStructDataFile_DataConfigNew(t *testing.T) {
	type fields struct {
		BlockFill    string
		SortField    int
		CountRow     int64
		ColumnLength []int
		Encoding     int
		File         *os.File
		FilePath     string
		FileName     string
		IsSort       bool
	}
	type args struct {
		path         string
		dataFileName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "在能访问的目录下创建新的默认配置文件",
			args: args{
				path:         "E:\\WorkSpace\\Go\\NoUiStudentManage\\test\\",
				dataFileName: "Temp",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := &StructDataFile{
				BlockFill:    tt.fields.BlockFill,
				SortField:    tt.fields.SortField,
				CountRow:     tt.fields.CountRow,
				ColumnLength: tt.fields.ColumnLength,
				Encoding:     tt.fields.Encoding,
				File:         tt.fields.File,
				FilePath:     tt.fields.FilePath,
				FileName:     tt.fields.FileName,
				IsSort:       tt.fields.IsSort,
			}
			if err := df.DataConfigNew(tt.args.path, tt.args.dataFileName); (err != nil) != tt.wantErr {
				t.Errorf("DataConfigNew() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
