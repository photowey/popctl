package alphabet

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test CamelCase()-Snake",
			args: args{
				src: "hello_world",
			},
			want: "hello_world",
		},
		{
			name: "Test CamelCase()-Pascal",
			args: args{
				src: "HelloWorld",
			},
			want: "helloWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.args.src); got != tt.want {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPascalCase(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test PascalCase()-Snake",
			args: args{
				src: "hello_world",
			},
			want: "Hello_world",
		},
		{
			name: "Test PascalCase()-Pascal",
			args: args{
				src: "helloWorld",
			},
			want: "HelloWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PascalCase(tt.args.src); got != tt.want {
				t.Errorf("PascalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test SnakeCase()-Pascal",
			args: args{
				src: "HelloWorld",
			},
			want: "hello_world",
		},
		{
			name: "Test SnakeCase()-Camel",
			args: args{
				src: "helloWorld",
			},
			want: "hello_world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCase(tt.args.src); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnake2Pascal(t *testing.T) {
	type args struct {
		snakeCase string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Snake2Pascal()",
			args: args{
				snakeCase: "hello_world",
			},
			want: "HelloWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Snake2Pascal(tt.args.snakeCase); got != tt.want {
				t.Errorf("Snake2Pascal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnake2Camel(t *testing.T) {
	type args struct {
		snakeCase string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Snake2Camel()",
			args: args{
				snakeCase: "hello_world",
			},
			want: "helloWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Snake2Camel(tt.args.snakeCase); got != tt.want {
				t.Errorf("Snake2Camel() = %v, want %v", got, tt.want)
			}
		})
	}
}
