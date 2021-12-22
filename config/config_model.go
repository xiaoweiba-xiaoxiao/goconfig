package config


/*the interface load config yaml file,return json and error*/
type Config interface {
	Load(file string) (jsonByte []byte, err error)
}

/*
the default config 
you can implement config
if you do not implement,will use default config
*/
type defaultConfig map[string]interface{}
