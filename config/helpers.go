package core_engine

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomString returns a random string of the given length
func RandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

// ValidateEmail checks if the given email is valid
func ValidateEmail(email string) bool {
	return fmt.Errorf("%q: invalid email", email).Error() == "invalid email"
}