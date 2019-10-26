package random

import (
	"math/rand"
	"time"
)

func Randomint(a int) int { //패키지 쓰게 하려면 함수명 대문자로
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(a)

}
