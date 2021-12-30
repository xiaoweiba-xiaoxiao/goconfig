package config_util

import (
	"errors"
	"io/ioutil"
	"strings"
)

var (
	Errfilesyntax     error = errors.New("syntax erorr")
	Errstyle          error = errors.New("this is a unkonw type of the config file")
	Erremptyinterface error = errors.New("the interface config is nil")
	Erremptyfile      error = errors.New("the config file is nil")
)

func ReadFile(file string) ([]byte, error) {
	cfdata, err := ioutil.ReadFile(file)
	return cfdata, err
}

/*parse section*/
func ParseSection(line string) (string, error) {
	if line == "" {
		return "", Errfilesyntax
	}
	if line[0] != '[' || (line[0] == '[' && line[len(line)-1] != ']') {
		return "", Errfilesyntax
	}
	line = strings.Replace(line, "[", "", 1)
	line = strings.Replace(line, "]", "", 1)
	return line, nil
}

/*pasre line not section*/
func ParseItem(line string) (key string, val string, err error) {
	if line == "" {
		err = Errfilesyntax
		return
	}

	eqindex := strings.Index(line, "=")
	if eqindex == -1 {
		err = Errfilesyntax
		return
	}

	key = line[:eqindex]
	if len(key) == 0 {
		err = Errfilesyntax
		return
	}

	key = strings.ReplaceAll(key, "\"", "")
	val = ""
	if eqindex < len(line)-1 {
		val = line[eqindex+1:]
		val = strings.ReplaceAll(val, "\"", "")
	}
	return
}


