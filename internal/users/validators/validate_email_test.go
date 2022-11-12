package validators

import (
	"fmt"
	"testing"
)

func TestValidateDomain(t *testing.T) {
	email1 := "jopa@ggmail.coom"
	email2 := "jopa@yandex.ru"
	badRes := ValidateDomain(email1)
	if badRes == true {
		t.Error("invalid email")
	}
	t.Log(badRes)
	actualRes := ValidateDomain(email2)
	if actualRes != true {
		t.Error("valid email failed")
	}
	t.Log(actualRes)
}

func TestValidateCountSymbol(t *testing.T) {
	email1 := "vasya@m@ail.com"
	email2 := "vasya@mail.com"
	badEmail := ValidateCountSymbol(email1)
	if badEmail == true {
		t.Error("invalid email")
	}
	fmt.Println(badEmail)
	actualEmail := ValidateCountSymbol(email2)
	if actualEmail != true {
		t.Error("valid email failed")
	}
	fmt.Println(actualEmail)
}
