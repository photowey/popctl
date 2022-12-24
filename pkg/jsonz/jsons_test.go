package jsonz

import (
	"io"
	"strings"
	"testing"
)

// Book
type Book struct {
	Id      string   `json:"id"`      // 图书 ISBN ID
	Name    string   `json:"name"`    // 图书名称
	Authors []string `json:"authors"` // 图书作者
	Press   string   `json:"press"`   // 出版社
}

var jsonData = `{
  "id": "9787111558422",
  "name": "The Go Programming Language",
  "authors": [
    "Alan A.A.Donovan",
    "Brian W. Kergnighan"
  ],
  "press": "Pearson Education"
}`

func TestUnmarshalStructE(t *testing.T) {
	type args struct {
		data   []byte
		target any
	}
	book := &Book{}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test json string to struct(Unmarshal)",
			args: args{
				data:   []byte(jsonData),
				target: book,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnmarshalStructE(tt.args.data, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalStructE() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecodeStruct(t *testing.T) {
	type args struct {
		reader io.Reader
		target any
	}
	book := &Book{}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test json string to struct(Decode)",
			args: args{
				reader: strings.NewReader(jsonData),
				target: book,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DecodeStruct(tt.args.reader, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("DecodeStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var (
	apiErrorBody = `{
  "code": "9787111558422",
  "message": "I'm full message"
}`
	apiErrorShortMessage = `{
  "code": "9787111558422",
  "msg": "I'm short message"
}`
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message" json:"msg"` // 解析失败
}

func TestToStruct(t *testing.T) {
	type args struct {
		data   []byte
		target any
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test deserialize json without error",
			args: args{
				data:   []byte(apiErrorBody),
				target: &APIError{},
			},
		},
		{
			name: "Test deserialize json without error",
			args: args{
				data:   []byte(apiErrorShortMessage),
				target: &APIError{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UnmarshalStruct(tt.args.data, tt.args.target)
		})
	}
}
