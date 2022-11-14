package main

import (
	"testing"
)

func TestUserManager(t *testing.T) {
	var man UserManager
	defer man.Disconnect()

	tu := "TestUser"
	tp := "TestPass"
	tm := "TestMail"

	olduser, err := man.FindUser(tu)
	if err != nil {
		t.Fatal(err.Error())
	}
	if olduser != nil {
		man.UnRegisterUser(tu)
	}

	user, err := man.RegisterUser(tu, tp, tm)
	if err != nil || user == nil || user.id < 1 {
		t.Fatal("Failed to register new user")
	}

	if user.name != tu || user.email != tm {
		t.Fatal("new user properties verification failed")
	}

	var token, errt = man.Login(tu, tp)
	if len(token) < 1 || errt != nil {
		t.Fatal("failed to login")
	}
}
