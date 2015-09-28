package conf
import (
	"os"
	"encoding/json"
//	"fmt"
)

type Config struct {
	Redis []string
	Aerospike struct {
			  Host string
			  Port int
		  }
	Logs string
}

func (conf *Config) LoadFromJson(configFile string) (err error) {
	r, err := os.Open("conf.json")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	decoder := json.NewDecoder(r)
	err = decoder.Decode(conf)

	return err
}