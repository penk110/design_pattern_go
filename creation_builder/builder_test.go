package creation_builder

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuilder_1(t *testing.T) {
	/*
		TODO:
			测试类型，如对象存在参数校验

		1.name: "MaxSize is zero"
		2.name: "MaxIdle < MinIdle"

	*/
	tests := []struct {
		name    string //
		builder *HttpClientBuilder
		want    *HttpClient
		wantErr bool
	}{
		{
			name: "MaxSize is zero",
			builder: &HttpClientBuilder{
				MaxSize:    -1,
				MaxIdle:    time.Second * 2,
				MinIdle:    time.Second,
				MaxTimeout: time.Second,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "MaxIdle < MinIdle",
			builder: &HttpClientBuilder{
				MaxSize:    1024,
				MaxIdle:    time.Second,
				MinIdle:    time.Second * 2,
				MaxTimeout: time.Second,
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "success",
			builder: &HttpClientBuilder{
			},
			want: &HttpClient{
				MaxSize:    4096,
				MaxIdle:    time.Second * 2,
				MinIdle:    time.Second,
				MaxTimeout: time.Second,
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := test.builder.Builder()
			require.Equalf(t, test.wantErr, err != nil, "Build() error = %v, wantErr %v", err, test.wantErr)
			assert.Equal(t, test.want, client)
		})
	}
}

func TestBuilder_2(t *testing.T) {
	type args struct {
		optFuncList []HttpConfigOptionFunc
	}
	tests := []struct {
		name    string
		args    args
		want    *HttpClient
		wantErr bool
	}{
		{
			name:    "MaxIdle < MinIdle",
			want:    nil,
			wantErr: true,
			args: args{
				optFuncList: []HttpConfigOptionFunc{
					func(option *HttpConfigOption) {
						option.MaxSize = 1024
						option.MaxIdle = time.Second
						option.MinIdle = time.Second * 2
						option.MaxTimeout = time.Second
					},
				},
			},
		},
		{
			name: "success",
			want: &HttpClient{
				MaxSize:    4096,
				MaxIdle:    time.Second * 2,
				MinIdle:    time.Second,
				MaxTimeout: time.Second,
			},
			wantErr: false,
			args: args{
				optFuncList: []HttpConfigOptionFunc{
					func(option *HttpConfigOption) {
						option.MaxSize = 4096
						option.MaxIdle = time.Second * 2
						option.MinIdle = time.Second
						option.MaxTimeout = time.Second
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := NewHttpClientBuilder(test.args.optFuncList...)
			require.Equalf(t, test.wantErr, err != nil, "error = %v, wantErr %v", err, test.wantErr)
			assert.Equal(t, test.want, client)
		})
	}
}
