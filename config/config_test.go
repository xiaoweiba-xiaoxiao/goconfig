package config

import (
	"io"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestConfig(t *testing.T){
	cf := NewConfig()
	josndata,err := ConfigLoad(cf,"./config.yaml")
	if err != nil {
		t.Error(err)
		return
	}
    t.Logf("%s",string(josndata))
	josndata,err = ConfigLoad(cf,"./config.ini")
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

func TestManyConfig(t *testing.T){
	f,err := os.Open("./config.yaml")
	if err != nil {
		t.Log(err)
		return
	}
	dec := yaml.NewDecoder(f)
	for {
		dc := defaultConfig{}
		err = dec.Decode(dc)
		if err == io.EOF {
			break
		}
		t.Logf("%#v",dc)
	}
}