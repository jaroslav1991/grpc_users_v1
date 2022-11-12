package validators

import (
	"database/sql"
	"strings"
)

const (
	getEmail = `select email from users where email=$1`
)

func ExistEmail(db *sql.DB, email string) error {
	err := db.QueryRow(getEmail, email).Scan(&email)
	if err != nil {
		return err
	}
	return nil
}

func ValidateDomain(email string) bool {
	validEmail := []string{"@mail.ru", "@gmail.com", "@yandex.ru", "@mail.com"}
	for _, v := range validEmail {
		if strings.HasSuffix(email, v) == true {
			return true
		}
	}
	return false
}

func ValidateCountSymbol(email string) bool {
	var counter string
	for _, v := range email {
		if string(v) == string('@') {
			counter += string('@')
		}
	}
	if len(counter) == 1 {
		return true
	}
	return false
}
