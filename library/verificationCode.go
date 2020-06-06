package lib

import (
	"crypto/rand"
	"io"
)

// VerificationCode - x자리 숫자를 반환해 주는 함수 Ex) VerificationCode(6) - 6자리 숫자를 랜덤으로 생성해서 반환
func VerificationCode(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
