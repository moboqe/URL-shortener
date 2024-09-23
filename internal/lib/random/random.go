package random

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	chars := []rune("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUyVvWwXxYyZz" + "0123456789" + "!@#$%^&*()")

	res := make([]rune, size)

	for i := range res {
		res[i] = chars[rnd.Intn(len(chars))]
	}
	return string(res)
}
