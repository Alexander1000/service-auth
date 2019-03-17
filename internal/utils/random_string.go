package utils

import (
	"math/rand"
	"time"
)

const (
	abc = "qwertyuiopasdfghjklzxcvbnm.0123456789-QWERTYUIOPASDFGHJKLZXCVBNM="
)

func RandomString(n int) string {
	str := make([]byte, 0, n)
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	for i := 0; i < n; i++ {
		num := rnd.Intn(len(abc))
		str = append(str, abc[num])
	}

	return string(str)
}
