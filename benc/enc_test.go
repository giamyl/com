package benc

import (
	"fmt"
	"testing"
)

func TestPwdSlat(t *testing.T) {
	pwd := "helloworld"
	salt := "default"
	sPwd := PwdSlat(pwd, "")
	fmt.Println("加密后的密码", sPwd)

	if sPwd != PwdSlat(pwd, salt) {
		t.Fatalf("两次密码不一致")
	}
}
