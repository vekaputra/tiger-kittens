package hash

import (
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
	"golang.org/x/crypto/bcrypt"
)

func BCrypt(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(b), pkgerr.ErrWithStackTrace(err)
}

func CheckBCrypt(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
