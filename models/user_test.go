package models

import (
	"fmt"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	u := new(User)
	u.SetSalt()
	u.SetPassword("abcde")
	if !u.ValidatePassword("abcde") {
		fmt.Printf("%#v\n", u)
		t.Fatal("cannot validate password")
	}
}
