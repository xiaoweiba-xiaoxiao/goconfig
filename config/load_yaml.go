package config

// import (
// 	"os"
// )

type Load_yaml struct {
	*defaultConfig
}
func (loadyaml *Load_yaml) LoadConfig(file string) (jsonByte []byte, err error){
    return loadyaml.loadyaml(file)
}
