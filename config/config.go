package config

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

// DataSourceConfig 数据库的结构体
type DataSourceConfig struct {
	DriverName string `mapstructure:"driverName"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Charset    string `mapstructure:"charset"`
}

// ServerConfig 整个项目的配置
type ServerConfig struct {
	Port       int              `mapstructure:"port"`
	Salt       string           `mapstructure:"salt"`
	DataSource DataSourceConfig `mapstructure:"datasource"`
}

func InitConfig(envString string) *ServerConfig {
	workDir, _ := os.Getwd()
	configFileName := path.Join(workDir, fmt.Sprintf("./config/config.%s.yml", envString))
	fmt.Println(configFileName, "文件")
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	var conf = &ServerConfig{}
	err := v.Unmarshal(conf)
	if err != nil {
		fmt.Println("读取配置失败")
		panic(err)
	}
	fmt.Println(conf)
	return conf
}
