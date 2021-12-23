package config

import (
	"errors"
	"io/ioutil"
	"strings"
)

var(
	errfilesyntax error = errors.New("syntax erorr")
	errstyle error = errors.New("this is a unkonw type of the config file")
	erremptyinterface error = errors.New("the interface config is nil")
	erremptyfile error = errors.New("the config file is nil")
)

func readFile(file string) ([]byte,error){
	cfdata, err := ioutil.ReadFile(file)
	return cfdata,err
}



/*parse section*/
func parseSection(line string)(string,error){
	if line == "" {		
		return "",errfilesyntax
	}
	if line[0] != '[' || (line[0]=='[' && line[len(line)-1] != ']') {
		return "",errfilesyntax
	}
	line = strings.Replace(line,"[","",1)
	line = strings.Replace(line,"]","",1)
	return line,nil
}

/*pasre line not section*/
func parseItem(line string)(key string,val string,err error){
	if line == "" {
		err = errfilesyntax
		return 
	}

    eqindex := strings.Index(line,"=")
	if eqindex == -1 {
		err = errfilesyntax
		return 
	}

	key = line[:eqindex]
	if len(key) == 0 {
		err = errfilesyntax
		return 
	}
    
	key = strings.ReplaceAll(key,"\"","")
	val = ""
	if eqindex < len(line) - 1{
		val = line[eqindex+1:]
		val = strings.ReplaceAll(val,"\"","") 
	}
	return   	 
}


/*
how to mashel config and the interface Config must be not nil
if you define a type of data that have method Load,but the method Load has not details
just return []byte and error
*/
func mashal(c Config,file string)(jsonstr []byte,err error){
    
    if c == nil {
		err = erremptyinterface
		return 
	}
	
	if file == "" {
		err = erremptyfile
		return
	}
		
	if strings.HasSuffix(file,".ini") {
		if jsonstr,err = c.Loadini(file);len(jsonstr)==0 && err == nil{ //if you 
			dc := &defaultConfig{}
			c = dc
		}
		return c.Loadini(file) 
	}
	
	if strings.HasSuffix(file,".yaml") || strings.HasSuffix(file,".yml") {
		if jsonstr,err = c.Loadyaml(file);len(jsonstr)==0 && err == nil{ //if you 
			dc := &defaultConfig{}
			c = dc			
		}
		return c.Loadyaml(file)
    } 
    		
	err = errstyle
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