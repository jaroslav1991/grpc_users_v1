package posts

import "testing"

func TestEmptyTitle(t *testing.T) {
	badTitle := ""
	actualTitle := "testTitle"
	badRes := EmptyTitle(badTitle)

	if badRes == true {
		t.Error("invalid title")
	}
	t.Log(badRes)

	actualRes := EmptyTitle(actualTitle)
	if actualRes != true {
		t.Error("valid title failed")
	}
	t.Log(actualRes)
}

func TestEmptyMessage(t *testing.T) {
	badMessage := ""
	actualMessage := "test"

	badRes := EmptyMessage(badMessage)
	if badRes == true {
		t.Error("invalid message")
	}
	t.Log(badRes)

	actualRes := EmptyTitle(actualMessage)
	if actualRes != true {
		t.Error("valid message failed")
	}
	t.Log(actualRes)
}
