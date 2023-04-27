package ConfigStruct

type GravitationConfig struct {
	Env        string //环境参数
	HttpPort   int    //http启动端口
	TcpPort    int    //tcp监听端口
	UdpPort    int    //udp监听端口
	ConfigType int    //配置类型
	ConfigPath string //配置中心路径
	CmdLine    string // 程序启动命令行信息
}

var GravitConfig GravitationConfig

//return GravitationConfig is sigle
func NewGravitationConfig() *GravitationConfig {
	return &GravitConfig
}
