package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	yaml "gopkg.in/yaml.v3"
)



func (ssh *SshConf)Load(file string)(jsonByte []byte,err error){
	return jsonByte,err
}

func readFile(file string) ([]byte,error){
	cfdata, err := ioutil.ReadFile(file)
	return cfdata,err
} 

func (dc *defaultConfig)parseJosn()([]byte,error){
	return json.Marshal(dc)
}

func (dc *defaultConfig)Load(file string)(jsonByte []byte,err error){
	jsonByte,err = readFile(file)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(jsonByte,dc)
	if err != nil {
		return
	}
	return dc.parseJosn()
}

/*mashel config*/
func mashal(c Config,file string)(jsonstr []byte,err error){
	if c == nil {
		err = errors.New("the inferce config is nil")
		return 
	}
	if file == "" {
		err = errors.New("the config file is nil")
		return
	}
	if jsonstr,err = c.Load(file);len(jsonstr)==0 && err == nil{
		dc := &defaultConfig{}
		c = dc
		return c.Load(file)
	}
	return 
}

/*
mashal api 
*/
func Mashal(c Config,file string)(jsonstr []byte,err error){
	return mashal(c,file)
}

func NewConfig()(Config){
	return &defaultConfig{}
}

