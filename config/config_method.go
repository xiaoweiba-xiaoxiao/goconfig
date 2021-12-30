package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	u "github.com/xiaoweiba-xiaoxiao/goconfig/config_util"
	yaml "gopkg.in/yaml.v3"
)


func (dc *defaultConfig)LoadConfig(file string)(jsonbyte []byte,err error){
	if file == "" {
		err = u.Erremptyfile
		return
	}

	if strings.HasSuffix(file, ".ini") {
		return dc.Loadini(file)
	}

	if strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml") {
		return dc.Loadyaml(file)
	}
	err = u.Errstyle
	return 
}

/*
the method read the config file return the json byte and error
*/
func (dc *defaultConfig) loadini(file string) (jsonbyte []byte, err error) {
	jsonbyte, err = u.ReadFile(file)
	if err != nil {
		return
	}
	return dc.parseini(jsonbyte)
}

/*
the pasre the config datas from file ,return json byte and error
*/
func (dc *defaultConfig) parseini(configdatas []byte) (jsonByte []byte, err error) {
	configstr := string(configdatas)
	lineslice := strings.Split(configstr, "\n")
	val := *dc
	dcmap := defaultConfig{}
	var section string

	for i, linestr := range lineslice {

		line := strings.Trim(linestr, " ") //

		if line == "" { //if this line is line.
			section = line
			continue
		}
		/*
			if this line has prefix "#" or ",",
			this is a comment line.
		*/
		if line[0] == '#' || line[0] == ',' { //if this line has prefix "#" or "," this is a comment line
			continue
		}

		/*
			delete the comment string after the item
		*/
		line = strings.Split(line, "#")[0]
		/*
			if there is some space bwteen comment string and item,
			trim those space.
		*/
		line = strings.Trim(line, " ")

		/*
			if this line has prefix "[",
		*/
		if line[0] == '[' {
			section, err = u.ParseSection(line)
			if err != nil {
				err = fmt.Errorf("the %d line %v", i, err)
				return
			}
			continue
		}
		/*
			if this line is just a item
		*/
		k, v, perr := u.ParseItem(line)
		if perr != nil {
			err = fmt.Errorf("the %d line %v", i, perr)
			return
		}

		if section != "" { // if section is not ""
			_, ok := dcmap[section] //check section include map

			if !ok { //if section not include map
				item := map[string]interface{}{}
				item[k] = v
				dcmap[section] = item
				continue
			}

			val, ok := dcmap[section].(map[string]interface{})
			if ok {
				val[k] = v
				dcmap[section] = val
				continue
			}

		}
		dcmap[k] = v
	}
	val["configs"] = dcmap
	return dc.parseJosn()
}

func (dc *defaultConfig) Loadini(file string) (jsonByte []byte, err error) {
	return dc.loadini(file)
}

/*pasre dc to json*/
func (dc *defaultConfig) parseJosn() ([]byte, error) {
	return json.Marshal(dc)
}

/*
load yaml config
return json data and error
*/
func (dc *defaultConfig) loadyaml(file string) (jsonByte []byte, err error) {
	confile, ok := os.Open(file)
	err = ok
	if err != nil {
		return
	}
	defer confile.Close()

	val := *dc
	configmaps := []defaultConfig{}
	dec := yaml.NewDecoder(confile)

	for err == nil { //decode yaml file if err != nil
		cf := defaultConfig{}
		err = dec.Decode(cf)
		if err != nil && err != io.EOF {
			return
		}
		if len(cf) != 0 { // if cf has key value append it
			configmaps = append(configmaps, cf)
		}
	}
	if len(configmaps) == 0 {
		val["configs"] = configmaps[0] // pasre one file
	} else {
		val["configs"] = configmaps //pasre more than one file
	}
	return dc.parseJosn()
}

func (dc *defaultConfig) Loadyaml(file string) (jsonByte []byte, err error) {
	return dc.loadyaml(file)
}

func (dcs *defaultSlice)loadyaml(file string) (jsonByte []byte, err error) {
	confile, ok := os.Open(file)
	err = ok
	if err != nil {
		return
	}

	dec := yaml.NewDecoder(confile)

	for err == nil { //decode yaml file if err != nil
		dc := defaultConfig{}
		err = dec.Decode(&dc)
		if err != nil && err != io.EOF {
			return
		}
		if len(dc) != 0 { // if cf has key value append it
			*dcs = append(*dcs, dc)
		}
	}
	if len(*dcs) == 1 {
		dcsvalue := *dcs
		return dcsvalue[0].parseJosn()
	} // pasre one file
	return dcs.parseJosn()
}

func (dcs *defaultSlice) Loadini(file string) (jsonByte []byte, err error) {
	return
}

func (dcs *defaultSlice) parseJosn() (jsonByte []byte, err error) {
	return json.Marshal(dcs)
}

func (dcs *defaultSlice) Loadyaml(file string) (jsonByte []byte, err error) {
	return dcs.loadyaml(file)
}
