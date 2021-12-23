package config

import (
	"testing"
)

func TestConfig(t *testing.T){
	cf := NewConfig()
	josndata,err := Mashal(cf,"./config.yaml")
	if err != nil {
		t.Error(err)
		return
	}
    t.Logf("%s",string(josndata))
	josndata,err = mashal(cf,"./config.ini")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s",string(josndata))
}

func TestDefaultIni(t *testing.T){
	file := "./config.ini"
	d := defaultConfig{}
	jsondata,err:=d.loadini(file)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(jsondata))
}