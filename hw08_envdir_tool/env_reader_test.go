package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func prepare(files, content []string) error {
	err := os.MkdirAll("tests", os.ModePerm)
	if err != nil {
		return fmt.Errorf("os.MkdirAll:%w", err)
	}
	for i := 0; i < len(files); i++ {
		f, err := os.Create(files[i])
		if err != nil {
			return fmt.Errorf("os.Create:%w", err)
		}
		if len(content) >= i {
			_, err = f.WriteString(content[i])
			if err != nil {
				return fmt.Errorf("os.WriteString:%w", err)
			}
		}
		f.Close()
	}
	return nil
}

func clear(dir string) error {
	return os.RemoveAll(dir)
}

func TestReadDir(t *testing.T) {
	type args struct {
		dir     string
		files   []string
		content []string
	}
	tests := []struct {
		name    string
		args    args
		want    Environment
		wantErr bool
	}{
		{
			name: "good reading",
			args: args{
				dir:     "tests",
				files:   []string{"tests/FOO", "tests/BAR"},
				content: []string{"foo", "bar"},
			},
			want: map[string]EnvValue{
				"FOO": {
					Value:      "foo",
					NeedRemove: false,
				},
				"BAR": {
					Value:      "bar",
					NeedRemove: false,
				},
			},
			wantErr: false,
		},
		{
			name: "good reading",
			args: args{
				dir: "some-dir",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "dir does not exists",
			args: args{
				dir:     "tests",
				files:   []string{"tests/FOO", "tests/BAR"},
				content: []string{"foo", ""},
			},
			want: map[string]EnvValue{
				"FOO": {
					Value:      "foo",
					NeedRemove: false,
				},
				"BAR": {
					Value:      "",
					NeedRemove: true,
				},
			},
			wantErr: false,
		},
		{
			name: "dir does not exists",
			args: args{
				dir:     "tests",
				files:   []string{"tests/FOO", "tests/BAR"},
				content: []string{"foo", string([]byte{0x00}) + "new line "},
			},
			want: map[string]EnvValue{
				"FOO": {
					Value:      "foo",
					NeedRemove: false,
				},
				"BAR": {
					Value:      "\nnew line",
					NeedRemove: false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			needWorkWithDir := len(tt.args.files) != 0 && len(tt.args.content) != 0
			if needWorkWithDir {
				err := prepare(tt.args.files, tt.args.content)
				if err != nil {
					t.Errorf("prepare:%v", err)
					return
				}
			}
			got, err := ReadDir(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for key, value := range got {
				if !reflect.DeepEqual(value, tt.want[key]) {
					t.Errorf("ReadDir() got = %v, want %v", value, tt.want[key])
				}
			}
			if needWorkWithDir {
				err := clear("tests")
				if err != nil {
					t.Errorf("clear:%v", err)
				}
			}
		})
	}
}
