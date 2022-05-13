package hash

import (
	"unicode"

	"github.com/eifzed/ares/lib/common/commonerr"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	if !isValidPassword(password) {
		return "", commonerr.ErrorBadRequest("Password must be at least 7 characters, with at least one uppercase, one special character, and one number")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func IsCorrectPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func isValidPassword(s string) bool {
	letters := 0
	var number, upper, special, sevenOrMore bool
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
			letters++
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			letters++
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			return false
		}
	}
	sevenOrMore = letters >= 7
	return number && upper && special && sevenOrMore
}
