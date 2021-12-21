package config



type Host struct{
	Addrr  string `yaml:"addrr" json:"addrr"`
	Port   int `yaml:"port" json:"port"`
	Auth    `yaml:"Auth" json:"auth"`
	PubKey  `yaml:"pubkey" json:"pubkey"`
	Kind  string  `yaml:"kind" json:"kind"`
}

type PubKey struct{
	KeyFile  string  `yaml:"keyfile" json:"keyfile"`
}

type Auth  struct{
	User  string `yaml:"User" json:"user"`
	PassWord string `yaml:"password" json:"password"`
}

type SshConf struct{
	Hosts []Host  `yaml:"hosts" json:"hosts"`
	Kind  string  `yaml:"kind" json:"kind"`
}

/*the interface load config yaml file,return json and err*/
type Config interface{
	Load(file string) (jsonByte []byte,err error)
}

/*
the default config
you can implements config
if you not implements,the will User default config
*/
type defaultConfig map[string]interface{}

