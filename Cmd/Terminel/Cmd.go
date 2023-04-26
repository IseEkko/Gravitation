package Terminel

import "fmt"

//this const is cmd
const (
	Env      = "env"
	HttpPort = "http-port"
	TcpPort  = "tcp-port"
	Help     = "help"
)

func ReadCmdLine() {
   
}

func helpPrint() {
	fmt.Println("-v           show porject version")
	fmt.Println("-http-port   show project http port")
	fmt.Println("-tcp-port    show project tcp port")
	fmt.Println("-env         show project env")
}
