package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	err := Load("../")
	if err != nil {
		t.Fatal(err)
		return
	}

	if configObj.DbName == "" {
		t.Errorf("DbName Err")
	} else {
		t.Logf("DbName = %s", configObj.DbName)
	}

	if configObj.DbPort == 0 {
		t.Errorf("DbPort Err")
	} else {
		t.Logf("DbPort = %d", configObj.DbPort)
	}

	if configObj.DbIp == "" {
		t.Errorf("DbIp Err")
	} else {
		t.Logf("DbIp = %s", configObj.DbIp)
	}

	if configObj.DbPw == "" {
		t.Errorf("DbPw Err")
	} else {
		t.Logf("DbPw = %s", configObj.DbPw)
	}

	if configObj.DbUser == "" {
		t.Errorf("DbUser Err")
	} else {
		t.Logf("DbUser = %s", configObj.DbUser)
	}
}
