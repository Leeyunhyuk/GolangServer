package config

import (
	"os"

	"github.com/naoina/toml"
)

//환경 변수 파일 => toml 형식

type Config struct {
	Server struct {
		Port string
	}
}

func NewConfig(filePath string) *Config {
	c := new(Config)
	//var c *Config

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
