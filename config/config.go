package config

import (
	"flag"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"testing"
	"wxGameWebSocket/utils/logger"
)

func init() {
	buildFlags()
	LoadConfig()
	buildServerConfig()
}

var log = logger.New("config")

// ServerConfig 文件配置参数
var ServerConfig serverConfig

// ServerStartupFlags 启动自定义参数
var ServerStartupFlags serverStartupFlags

type serverConfig struct {
	Host            string
	Port            int
	AppId           string
	Secret          string
	WxLoginURL      string
	GroupCodeSecret string
	SecretKey       string
	ShowSQL         bool
	RunMode         string // debug release test

	WebSocketPort string
	RPCPort       string
}
type serverStartupFlags struct {
	Host        string
	Port        int
	Environment string
}

// LoadConfig 加载配置文件
func LoadConfig() {
	// 设置配置文件名
	var configPrefix = "FarmConfig"
	configName := fmt.Sprintf("%s-%s", configPrefix, ServerStartupFlags.Environment)
	viper.SetConfigName(configName)
	// 设置配置文件路径
	viper.AddConfigPath("conf")
	// 测试时使用路径
	//viper.AddConfigPath("../../conf")
	// 解析配置
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Error("Fatal error config file", "err", err.Error())
	}
}

// GetDBConfig 获取db配置
func GetDBConfig() map[string]interface{} {
	return viper.GetStringMap("db")
}

// GetServerConfig 获取服务器配置
func GetServerConfig() map[string]interface{} {
	return viper.GetStringMap("server")
}

// buildServerConfig 构建文件服务器配置
func buildServerConfig() {
	cfg := GetServerConfig()
	ServerConfig = serverConfig{
		Port:            cast.ToInt(cfg["port"]),
		AppId:           cast.ToString(cfg["appid"]),
		Secret:          cast.ToString(cfg["secret"]),
		WxLoginURL:      cast.ToString(cfg["wxloginurl"]),
		GroupCodeSecret: cast.ToString(cfg["groupcodesecret"]),
		SecretKey:       cast.ToString(cfg["secretkey"]),
		ShowSQL:         cast.ToBool(cfg["showsql"]),
		RunMode:         cast.ToString(cfg["runmode"]),
		WebSocketPort:   cast.ToString(cfg["websocketport"]),
		RPCPort:         cast.ToString(cfg["rpcport"]),
	}
	ServerConfig.Port = ServerStartupFlags.Port
	ServerConfig.Host = ServerStartupFlags.Host
}

// buildFlags 构建启动时参数配置
func buildFlags() {
	testing.Init()
	flag.StringVar(&ServerStartupFlags.Host, "host", "127.0.0.1", "listening host")
	flag.IntVar(&ServerStartupFlags.Port, "port", 7500, "listening port")
	flag.StringVar(&ServerStartupFlags.Environment, "env", "dev", "run time environment")
	if !flag.Parsed() {
		flag.Parse()
	}
}
