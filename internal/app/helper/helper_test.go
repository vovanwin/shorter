package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat2builder(t *testing.T) {
	type args struct {
		http string
		x    string
		z    string
		y    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Позитивный тест",
			args: args{"http://", "localhost", "/", "123"},
			want: "http://localhost/123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Concat2builder(tt.args.http, tt.args.x, tt.args.z, tt.args.y)
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestIsURL(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Негативный тест",
			args: args{"127.0.0.1:8080"},
			want: false,
		},
		{
			name: "Позитивный тест",
			args: args{"http://127.0.0.1:8080"},
			want: true,
		},
		{
			name: "http:// Не ссылка",
			args: args{"http://"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := IsURL(tt.args.str)
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestNewCode(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Проверка работоспособности",
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewCode()
			assert.Equal(t, tt.want, len(v))
		})
	}
}
