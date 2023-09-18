package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"

	"sokoban_solution/mapdata"
)

type Config struct {
	MapData  *mapdata.Map  `toml:"mapdata"`
	InitStep *mapdata.Step `toml:"initstep"`
}

func (config *Config) InitOnLoad() (err error) {
	err = GConfig.MapData.InitOnLoad()
	if err != nil {
		return
	}
	err = config.InitStep.InitOnLoad()
	return
}

func InitFromFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(data, &GConfig)
	if err != nil {
		panic(err)
	}
	err = GConfig.InitOnLoad()
	if err != nil {
		panic(err)
	}
}

var GConfig Config
