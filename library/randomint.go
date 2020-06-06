package random

import (
	"math/rand"
	"time"
)

// Randomint - 랜덤 숫자를 주는 코드 , Ex) Randomint(800) = 0 이상이고 800 이하인 랜덤숫자를 반환
func Randomint(a int) int { //패키지 쓰게 하려면 함수명 대문자로
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(a)

}
