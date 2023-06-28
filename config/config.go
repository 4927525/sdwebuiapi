package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Config *viper.Viper
var DB *gorm.DB
var RDB *redis.Client

func Init() {
	InitConfig()
	InitDB()
	InitRDB()
}

// InitRDB 读取redis配置
func InitRDB() {
	addr := Config.GetString("redis.addr")
	pwd := Config.GetString("redis.pwd")
	dbname := Config.GetString("redis.dbname")
	db, _ := strconv.Atoi(dbname)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RDB = client
}

// InitDB 读取数据库配置
func InitDB() {
	user := Config.GetString("database.user")
	pwd := Config.GetString("database.pwd")
	host := Config.GetString("database.host")
	port := Config.GetString("database.port")
	dbname := Config.GetString("database.dbname")
	dsn := strings.Join([]string{user, ":", pwd, "@tcp(", host, ":", port, ")/", dbname, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true, // 开启事务
		//Logger:                 logger.Default.LogMode(logger.Info), // 打印sql
	}
	if Config.GetString("server.app") == "prod" {
		gormConfig = &gorm.Config{
			SkipDefaultTransaction: true, // 开启事务
		}
	}
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		panic(err)
	}
	DB = db
}

// InitConfig 读取配置文件
func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := viper.New()
	config.AddConfigPath(path)
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	Config = config
}
