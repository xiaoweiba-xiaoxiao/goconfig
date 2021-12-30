package config

/*
 deal json file
*/
type Load_json struct {
	*defaultConfig
}

func (load * Load_json)Load_perperties(file string)(jsonByte []byte, err error){
	return load.loadyaml(file)
}