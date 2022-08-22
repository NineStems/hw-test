package main

import (
	"bytes"
	"io"
	"testing"
)

func TestCopy(t *testing.T) {
	type args struct {
		fromPath string
		toPath   string
		offset   int64
		limit    int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty in file",
			args: args{
				fromPath: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Copy(tt.args.fromPath, tt.args.toPath, tt.args.offset, tt.args.limit)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				return
			}
		})
	}
}

func Test_copyReader(t *testing.T) {
	type args struct {
		from   io.ReaderAt
		offset int64
		limit  int64
		size   int64
	}
	tests := []struct {
		name    string
		args    args
		wantTo  string
		wantErr bool
	}{
		{
			name: "good copy offset=0 limit=0",
			args: args{
				from:   bytes.NewReader([]byte("1234567890")),
				offset: 0,
				limit:  0,
				size:   int64(len([]byte("1234567890"))),
			},
			wantTo:  "1234567890",
			wantErr: false,
		},
		{
			name: "good copy offset=0 limit=5",
			args: args{
				from:   bytes.NewReader([]byte("1234567890")),
				offset: 0,
				limit:  5,
				size:   int64(len([]byte("1234567890"))),
			},
			wantTo:  "12345",
			wantErr: false,
		},
		{
			name: "good copy offset=0 limit=15",
			args: args{
				from:   bytes.NewReader([]byte("1234567890")),
				offset: 0,
				limit:  15,
				size:   int64(len([]byte("1234567890"))),
			},
			wantTo:  "1234567890",
			wantErr: false,
		},
		{
			name: "good copy offset=5 limit=5",
			args: args{
				from:   bytes.NewReader([]byte("1234567890")),
				offset: 5,
				limit:  5,
				size:   int64(len([]byte("1234567890"))),
			},
			wantTo:  "67890",
			wantErr: false,
		},
		{
			name: "good copy offset=8 limit=5",
			args: args{
				from:   bytes.NewReader([]byte("1234567890")),
				offset: 8,
				limit:  5,
				size:   int64(len([]byte("1234567890"))),
			},
			wantTo:  "90",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			to := &bytes.Buffer{}
			err := copyReader(tt.args.from, to, tt.args.offset, tt.args.limit, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("copyReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTo := to.String(); gotTo != tt.wantTo {
				t.Errorf("copyReader() gotTo = %v, want %v", gotTo, tt.wantTo)
			}
		})
	}
}
