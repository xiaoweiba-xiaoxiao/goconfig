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
}