package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/go_mysql_beta/go/pkg/util"
)

func Test(t *testing.T) {
	type args struct {
		ctx context.Context
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
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
