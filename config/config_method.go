package config

import (
	"encoding/json"
	"fmt"
	"strings"
	yaml "gopkg.in/yaml.v3"
)



func (dc *defaultConfig)loadini(file string)(jsonbyte []byte,err error){
	jsonbyte,err = readFile(file)
	if err != nil {
		return 
	}
	return dc.parseini(jsonbyte)
}

func (dc *defaultConfig)parseini(configdatas []byte)(jsonByte []byte,err error){
	configstr := string(configdatas)
	lineslice := strings.Split(configstr,"\n")	
	dcmap := *dc
	var section string 	
	
	for i,linestr := range lineslice {
		
		line := strings.Trim(linestr," ") //
		
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
		line = strings.Split(line,"#")[0]
		/*
		if there is some space bwteen comment string and item,
		trim those space.
		*/ 
		line = strings.Trim(line," ")
		
		/*
		if this line has prefix "[",
		*/
		if line[0] == '[' {
		    section,err = parseSection(line)
			if err != nil{
				err = fmt.Errorf("the %d line %v",i,err)
				return 
			}
			continue			 
		}
        /*
		if this line is just a item
		*/
		k,v,perr := parseItem(line)
		if perr != nil {
		   err = fmt.Errorf("the %d line %v",i,perr)
		   return
		}
		
		if section != "" { // if section is not ""			
			_,ok := dcmap[section] //check section include map
			
			if !ok {  //if section not include map
				item := map[string]interface{}{}
				item[k] = v
				dcmap[section] = item
				continue
			}
			
			val,ok := dcmap[section].(map[string]interface{})
		    if ok {
				val[k] = v
				dcmap[section] = val
				continue
			}

		}
		dcmap[k] = v		
	}
    return dc.parseJosn()
}


func (dc *defaultConfig)Loadini(file string)(jsonByte []byte,err error){
	return dc.loadini(file)
}

func (dc *defaultConfig)parseJosn()([]byte,error){
	return json.Marshal(dc)
}

func (dc *defaultConfig)loadyaml(file string)(jsonByte []byte,err error){
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

func (dc *defaultConfig)Loadyaml(file string)(jsonByte []byte,err error){
	return dc.loadyaml(file)
}



