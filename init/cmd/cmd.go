package cmd

import (
	"BACKEND_GO/config"
	"fmt"
)

// repository, service, types, config 항목을 Cmd 항목에 담을 것
type Cmd struct {
	config *config.Config
}

// main 에서 이 NewCmd 함수를 통해 값을 받아서 서버를 구동 할 예정
func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	fmt.Println(c.config.Server.Port)

	return c
}
