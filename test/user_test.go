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
					Uid:       10000001,
					Name:      "sa123",
					Email:     "12345abcde@gmail.com",
					TelNumber: "080-1234-5678",
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

func TestGetAllUsersInfo(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get all users info",
			args: args{
				ctx: context.TODO(),
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

			users, err := user.GetAllUsersInfo(db, tt.args.ctx)
			if err != nil {
				t.Errorf("GetAllUsersInfo error: %v\n", err)
				return
			}

			if len(users) <= 0 {
				t.Log("user does not exist.")
				return
			}

			for i, user := range users {
				t.Logf("[%d] user: %v\n", i, user)
			}

			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}

func TestDeleteUser(t *testing.T) {
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
			name: "test format",
			args: args{
				ctx: context.TODO(),
				uid: 10000003,
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

			if err = user.DeleteUser(db, tt.args.ctx, tt.args.uid); err != nil {
				t.Errorf("DeleteUser error: %v\n", err)
				return
			}

			u, err := user.GetUserInfo(db, tt.args.ctx, tt.args.uid)
			if err != nil {
				t.Logf("GetUserInfo error: %v\n", err)
			}
			t.Log("delete process completed. ", u)

			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
