package helper

import (
	"log"

	"github.com/cpartogi/izyai/internal/account"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(password string, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err
}

func HashPassword(password string) (hashedPassword string, err error) {
	// hash
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error generate password hash:", err)
		return hashedPassword, errors.Wrap(err, account.ErrRegisteringUser.Error())
	}
	hashedPassword = string(hash)

	return hashedPassword, err
}
