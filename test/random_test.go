package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/go_mysql_beta/go/pkg/util"
)

func TestGenerateRandomNumber(t *testing.T) {
	type args struct {
		ctx   context.Context
		digit int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "generate random number test",
			args: args{
				ctx:   context.TODO(),
				digit: 100, // 0以上、N未満の範囲で乱数を生成
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("result:", util.GenerateRandomNumber(tt.args.digit))
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}

func TestGenerateMultipleRandomNumber(t *testing.T) {
	type args struct {
		ctx   context.Context
		digit int
		count int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "generate multiple random number test",
			args: args{
				ctx:   context.TODO(),
				digit: 100, // 0以上、N未満の範囲で乱数を生成
				count: 5,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numArr := util.GenerateMultipleRandomNumber(tt.args.digit, tt.args.count)
			for i, v := range numArr {
				t.Logf("[%d] %d\n", i, v)
			}
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		ctx        context.Context
		digit      int
		randomType string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "generate random string test",
			args: args{
				ctx:        context.TODO(),
				digit:      8,
				randomType: util.BaseStr,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str, err := util.GenerateRandomString(tt.args.digit, tt.args.randomType)
			if err != nil {
				t.Errorf("GenerateRandomString err: %v\n", err)
			}
			t.Log("result:", str)

			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}

func TestGenerateMultipleRandomString(t *testing.T) {
	type args struct {
		ctx        context.Context
		digit      int    // 桁数
		randomType string // 作成する個数
		count      int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "generate multiple random string test",
			args: args{
				ctx:        context.TODO(),
				digit:      64,
				count:      4,
				randomType: util.BaseStr,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("============================ [ BaseStr ] ============================")
			str, err := util.GenerateMultipleRandomString(tt.args.digit, tt.args.count, tt.args.randomType)
			if err != nil {
				t.Errorf("GenerateRandomString err: %v\n", err)
			}
			for i, v := range str {
				t.Logf("[%d] %s\n", i, v)
			}

			t.Log("============================ [ LowercaseStrOnly ] ============================")
			tt.args.randomType = util.LowercaseStrOnly
			str, err = util.GenerateMultipleRandomString(tt.args.digit, tt.args.count, tt.args.randomType)
			if err != nil {
				t.Errorf("GenerateRandomString err: %v\n", err)
			}
			for i, v := range str {
				t.Logf("[%d] %s\n", i, v)
			}

			t.Log("============================ [ UppercaseStrOnly ] ============================")
			tt.args.randomType = util.UppercaseStrOnly
			str, err = util.GenerateMultipleRandomString(tt.args.digit, tt.args.count, tt.args.randomType)
			if err != nil {
				t.Errorf("GenerateRandomString err: %v\n", err)
			}
			for i, v := range str {
				t.Logf("[%d] %s\n", i, v)
			}

			t.Log("============================ [ LowercaseStrAndNumMix ] ============================")
			tt.args.randomType = util.LowercaseStrAndNumMix
			str, err = util.GenerateMultipleRandomString(tt.args.digit, tt.args.count, tt.args.randomType)
			if err != nil {
				t.Errorf("GenerateRandomString err: %v\n", err)
			}
			for i, v := range str {
				t.Logf("[%d] %s\n", i, v)
			}

			t.Log("============================ [ UppercaseStrAndNumMix ] ============================")
			tt.args.randomType = util.UppercaseStrAndNumMix
			str, err = util.GenerateMultipleRandomString(tt.args.digit, tt.args.count, tt.args.randomType)
			if err != nil {
				t.Errorf("GenerateRandomString err: %v\n", err)
			}
			for i, v := range str {
				t.Logf("[%d] %s\n", i, v)
			}

			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}

func TestDigitCompensating(t *testing.T) {
	type args struct {
		ctx        context.Context
		digit      int
		convertNum int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "digit compensating test",
			args: args{
				ctx:        context.TODO(),
				digit:      8,
				convertNum: 4,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("result:", util.DigitCompensating(tt.args.digit, tt.args.convertNum))
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}

func TestMultipleDigitCompensating(t *testing.T) {
	type args struct {
		ctx       context.Context
		digit     int
		count     int
		compDigit int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "digit compensating test",
			args: args{
				ctx:       context.TODO(),
				digit:     9999,
				count:     10,
				compDigit: 4,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numArr := util.GenerateMultipleRandomNumber(tt.args.digit, tt.args.count)
			for i, v := range numArr {
				t.Logf("[%d] %d\n", i, v)
			}

			numStr := util.MultipleDigitCompensating(numArr, tt.args.compDigit)
			t.Log("result:", numStr)
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
