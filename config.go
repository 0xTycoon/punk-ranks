package punksranking

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	SQLDriver string `json:"sql_driver"`
	SQLDSN    string `json:"sql_dsn"`
	RPCURL    string `json:"rpc_url"`
}

func LoadConfig(path string) (*Config, error) {
	c := &Config{}
	if path == "" {
		path = "./config.json"
	}
	if data, err := ioutil.ReadFile(path); err == nil {
		if err := json.Unmarshal(data, c); err == nil {
			return c, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

}
