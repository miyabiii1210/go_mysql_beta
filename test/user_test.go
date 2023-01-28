package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/go_mysql_beta/go/pkg/database"
	"github.com/miyabiii1210/go_mysql_beta/go/pkg/user"
	"github.com/miyabiii1210/go_mysql_beta/go/pkg/util"
)

func TestSaveUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user user.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "save user test.",
			args: args{
				ctx: context.TODO(),
				user: user.User{
					Uid:       0,
					Name:      "sa",
					Email:     "123abc@gmail.com",
					TelNumber: "050-1234-5678",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := database.ConnectionToDB()
			if err != nil {
				t.Errorf("ConnectionToDB error: %v\n", err)
				return
			}

			lastUid, err := user.SaveUser(db, tt.args.ctx, &tt.args.user)
			if err != nil {
				t.Errorf("SaveUser error: %v\n", err)
				return
			}
			t.Log("last uid:", lastUid)

			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}

func TestGetUserInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		uid uint32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get user info test.",
			args: args{
				ctx: context.TODO(),
				uid: 10000001,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := database.ConnectionToDB()
			if err != nil {
				t.Errorf("ConnectionToDB error: %v\n", err)
				return
			}

			u, err := user.GetUserInfo(db, tt.args.ctx, tt.args.uid)
			if err != nil {
				t.Errorf("GetUserInfo error: %v\n", err)
				return
			}
			t.Log("user info: ", u)

			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
