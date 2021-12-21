package config

import (
	"testing"
	"encoding/json"
)

func TestConfig(t *testing.T){
	var (
		ssh *SshConf = &SshConf{}
	    cf Config
	)
	cf = ssh 
	josndata,err := Mashal(cf,"./config.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	err = json.Unmarshal(josndata,ssh)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v",ssh)
}