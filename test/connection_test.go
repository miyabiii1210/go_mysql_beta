package test

import (
	"context"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/miyabiii1210/go_mysql_beta/go/pkg/database"
	"github.com/miyabiii1210/go_mysql_beta/go/pkg/util"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func TestConnectionToDB(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "connection to database test.",
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
			defer db.Close()

			if err = db.Ping(); err != nil {
				t.Errorf("database connection failed. %v\n", err)
				return
			}
			t.Log("database connection success.")

			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
