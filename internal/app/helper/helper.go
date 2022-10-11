package helper

import (
	"math/rand"
	"net/url"
	"strings"
	"time"
)

const (
	alphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length     = len(alphabet)
	CodeLength = 6
)

type Helper struct {
}

func IsURL(str string) bool {
	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

func Concat2builder(http, x, z, y string) string {
	var builder strings.Builder
	builder.Grow(len(http) + len(x) + len(z) + len(y)) // Только эта строка выделяет память
	builder.WriteString(http)
	builder.WriteString(x)
	builder.WriteString(z)
	builder.WriteString(y)
	return builder.String()
}

func NewCode() string {
	var letters = []rune(alphabet)

	code := make([]rune, CodeLength)
	for i := range code {
		rand.Seed(time.Now().UnixNano())
		code[i] = letters[rand.Intn(length)]
	}

	return string(code)
}
