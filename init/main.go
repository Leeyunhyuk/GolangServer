package main

import (
	"GOLANGSERVER/init/cmd"
	"flag"
)

// *configPathFlag는 Default값인 "./config.toml"이다.
var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
