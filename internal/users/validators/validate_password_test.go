package validators

import "testing"

func TestValidatePassword(t *testing.T) {
	badPassword := "1234"
	goodPassword := "12345"
	actualRes := ValidatePassword(goodPassword)
	badRes := ValidatePassword(badPassword)

	if badRes == true {
		t.Error("invalid password, need 5 or more symbols")
	}
	t.Log(badRes)

	if actualRes != true {
		t.Error("valid password failed")
	}
	t.Log(actualRes)
}
