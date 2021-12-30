package config
import (
	_"sync"
	u "github.com/xiaoweiba-xiaoxiao/goconfig/config_util"
)


/*
how to mashel config and the interface Config must be not nil
if you define a type of data that have method Load,but the method Load has not details
just return []byte and error
*/
func configLoad(c Config, file string) (jsonstr []byte, err error) {
	if c == nil {
		err = u.Erremptyinterface
		return
	}
	return c.LoadConfig(file)
}

/*
mashal api has two parameters the Config interface and a yaml file
read file and return the data of the file in josn.
*/
func ConfigLoad(c Config, file string) (jsonstr []byte, err error) {
	return configLoad(c,file)
}

/*
new default config interface
*/
func NewConfig() Config {
	return &defaultConfig{}
}

/*
just read one file like this
---
a:
---
b:
can load
*/

func loadyaml(file string) (josnbyte []byte, err error) {
	ds := &defaultSlice{}
	return ds.Loadyaml(file)
}

func LoadYaml(file string) (josnbyte []byte, err error) {
	return loadyaml(file)
}