package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkEndpoints(t *testing.T) {
	type args struct {
		endpoints []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test empty",
			args:    args{},
			wantErr: true,
		},
		{
			name: "test one endpoint",
			args: args{
				endpoints: []string{"192.168.1.5:6379"},
			},
			wantErr: false,
		},
		{
			name: "test more endpoint",
			args: args{
				endpoints: []string{"127.0.0.1:6379", "192.168.1.5:6378"},
			},
			wantErr: false,
		},
		{
			name: "test one no port",
			args: args{
				endpoints: []string{"127.0.0.1"},
			},
			wantErr: false,
		},
		{
			name: "test more endpoint no port",
			args: args{
				endpoints: []string{"192.168.1.5", "127.0.0.1:6378", "10.88.1.5"},
			},
			wantErr: false,
		},
		{
			name: "test use domain",
			args: args{
				endpoints: []string{"redis.localhost.com:6378", "redis.localhost.com:6379"},
			},
			wantErr: false,
		},
		{
			name: "test bad domain",
			args: args{
				endpoints: []string{"bad-====:331"},
			},
			wantErr: true,
		},
		{
			name: "test bad port",
			args: args{
				endpoints: []string{"192.168.1.3:abc123"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkEndpoints(tt.args.endpoints...)
			if err != nil {
				assert.True(t, tt.wantErr)
			} else {
				assert.False(t, tt.wantErr)
			}
		})
	}
}
