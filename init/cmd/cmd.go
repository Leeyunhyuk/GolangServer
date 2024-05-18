package cmd

import (
	"GOLANGSERVER/config"
	"GOLANGSERVER/network"
	"fmt"
)

// repository, service, types, config 항목을 Cmd 항목에 담을 것
type Cmd struct {
	config  *config.Config
	network *network.Network
}

// main 에서 이 NewCmd 함수를 통해 값을 받아서 서버를 구동 할 예정
func NewCmd(filePath string) *Cmd {

	fmt.Println("여기는 뜹니다")

	c := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}
	//이 아래는 동작을 안함

	fmt.Println(c.config.Server.Port)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
