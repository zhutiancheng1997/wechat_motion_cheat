package model

import "testing"

func TestZeppLifeUserGetAccess(t *testing.T) {
	user := &ZeppLifeUser{User: "1524495584@qq.com", Password: "aa2505913895"}

	access, err := user.getAccess()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(access)
}

func TestZeppLifeUserLogin(t *testing.T) {
	user := &ZeppLifeUser{User: "1524495584@qq.com", Password: "aa2505913895"}

	accessCode, err := user.getAccess()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(accessCode)
	email, err := user.login(accessCode)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(email)
}

func TestZeppLifeBuildDataJson(t *testing.T) {
	user := NewUser("1524495584@qq.com", "aa2505913895")
	user.Uid = "123"
	_ = user.buildStepData(10023)

}

func TestCheatOnSteps(t *testing.T) {
	user := NewUser("1524495584@qq.com", "aa2505913895")
	err := user.CheatOnSteps(10)
	if err != nil {
		t.Fatal(err)
	}
}
