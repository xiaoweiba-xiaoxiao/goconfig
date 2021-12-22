package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	yaml "gopkg.in/yaml.v3"
)

func readFile(file string) ([]byte,error){
	cfdata, err := ioutil.ReadFile(file)
	return cfdata,err
} 

func (dc *defaultConfig)parseJosn()([]byte,error){
	return json.Marshal(dc)
}

func (dc *defaultConfig)load(file string)(jsonByte []byte,err error){
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

func (dc *defaultConfig)Load(file string)(jsonByte []byte,err error){
	return dc.load(file)
}

/*
how to mashel config
if you define a type of data that have method Load,but the method Load has not details
just return []byte and error
*/
func mashal(c Config,file string)(jsonstr []byte,err error){
	if c == nil {
		err = errors.New("the interface config is nil")
		return 
	}
	if file == "" {
		err = errors.New("the config file is nil")
		return
	}
	if jsonstr,err = c.Load(file);len(jsonstr)==0 && err == nil{ //if you 
		dc := &defaultConfig{}
		c = dc
		return c.Load(file)
	}
	return 
}

/*
mashal api has two parameters the Config interface and a yaml file
read file and return the data of the file in josn.  
*/
func Mashal(c Config,file string)(jsonstr []byte,err error){
	return mashal(c,file)
}

/*
new default config interface
*/
func NewConfig()(Config){
	return &defaultConfig{}
}

