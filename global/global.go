package global

import (
	"wallet-frame/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Config *config.ServerConfig = &config.ServerConfig{}
	DB     *gorm.DB
	RedisC *redis.Client
)

func Initialize(envStr string) {
	Config = config.InitConfig(envStr)

	// // 从配置文件中获取参数
	// DB = dbs.InitDB(Config.DataSource.Host,
	// 	strconv.Itoa(Config.DataSource.Port),
	// 	Config.DataSource.Database,
	// 	Config.DataSource.Username,
	// 	Config.DataSource.Password,
	// 	Config.DataSource.Charset)

	// RedisC = dbs.InitRedis()
}
