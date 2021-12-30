package config

type Load_perperties struct {
	*defaultConfig
}

func (load * Load_perperties)Load_perperties(file string)(jsonByte []byte, err error){
	return load.loadini(file)
}