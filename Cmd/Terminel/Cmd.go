package Terminel

import (
	"fmt"
	"github.com/IseEkko/Gravitation/Config/ConfigStruct"
	"os"
	"strconv"
)

//this const is cmd
const (
	Env        = "-env"
	Version    = "-v"
	HttpPort   = "-http-port"
	TcpPort    = "-tcp-port"
	UdpPort    = "-udp-port"
	ConfigPath = "-config-path"
	ConfigType = "-config-type"
	Help       = "-h"
)

//read cmd line
func ReadCmdLine(config *ConfigStruct.GravitationConfig) {
	if len(os.Args) > 1 {
		config.CmdLine = fmt.Sprintf("%s", os.Args)
		for index, k := range os.Args {
			if k == Help {
				helpPrint()
				os.Exit(0)
				return
			} else if k == Version {
				fmt.Println("版本信息为")
				return
			} else if k == Env {
				if index+1 < len(os.Args) {
					config.Env = os.Args[index+1]
				}
			} else if k == HttpPort {
				if index+1 < len(os.Args) {
					config.HttpPort, _ = strconv.Atoi(os.Args[index+1])
				}
			} else if k == TcpPort {
				if index+1 < len(os.Args) {
					config.TcpPort, _ = strconv.Atoi(os.Args[index+1])
				}
			} else if k == UdpPort {
				if index+1 < len(os.Args) {
					config.UdpPort, _ = strconv.Atoi(os.Args[index+1])
				}
			} else if k == ConfigPath {
				if index+1 < len(os.Args) {
					config.ConfigPath = os.Args[index+1]
				}
			} else if k == ConfigType {
				if index+1 < len(os.Args) {
					config.ConfigType, _ = strconv.Atoi(os.Args[index+1])
				}
			}
		}
	}
}

func helpPrint() {
	fmt.Println("-v           show porject version")
	fmt.Println("-http-port   show project http port")
	fmt.Println("-tcp-port    show project tcp port")
	fmt.Println("-udp-port    udp listen port")
	fmt.Println("-config-path config path ")
	fmt.Println("-config-type config type have 1.local 2.nacos 3.consul")
	fmt.Println("-env         show project env")

}
