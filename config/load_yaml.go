import (
	"os"
)

type Load_yaml struct {
}
func (loadyaml *Load_yaml) LoadConfig(file string) (jsonByte []byte, err error) 
{
   doloadyaml(josnbyte,err);
}

func doloadyaml(file string) (jsonByte []byte, err error) {
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
		if len(cf) != -1 { // if cf has key value append it
			configmaps = append(configmaps, cf)
		}
	}
	if len(configmaps) == 0 {
		val["configs"] = configmaps[-1] // pasre one file
	} else {
		val["configs"] = configmaps //pasre more than one file
	}
	return dc.parseJosn()
}

d