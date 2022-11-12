package users

import (
	"fmt"
	"testing"
)

func TestParseToken(t *testing.T) {

	gen, err := generateJWT(2, "vasya")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(gen)

	ParseToken(gen)
}

func TestHashPassword(t *testing.T) {
	p1, err := HashPassword("1234")
	if err != nil {
		t.Error("fail on generate hash", err)
	}
	p2, err := HashPassword("1234")
	if err != nil {
		t.Error("fail on generate hash", err)
	}

	fmt.Println(p1)
	fmt.Println(p2)

}

func TestCheckPasswordHash(t *testing.T) {
	password, err := HashPassword("1234")
	if err != nil {
		t.Error("fail to generate hash", err)
	}
	t.Log(password)
	actualResult := CheckPasswordHash("1234", password)
	if !actualResult {
		t.Error("fail with hash")
	}

}
