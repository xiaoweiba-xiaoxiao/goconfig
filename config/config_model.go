package config

/*the interface load config yaml file,return json and error*/
type Config interface {
	//	Loadyaml(file string) (jsonByte []byte, err error)
	LoadConfig(file string) (jsonByte []byte, err error)
	//Loadini(file string) (jsonbyte []byte, err error)
}

/*
the default config
you can implement config
if you do not implement,will use default config
that will be used for Configuration Center in the future
*/
type defaultConfig map[string]interface{}

/*
this is used for just read the yaml file and pasre json
*/
type defaultSlice []defaultConfig
