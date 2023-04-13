package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()

	if err == nil {
		fmt.Println("成功連結", viper.Get("app"))
	} else {
		fmt.Println("failed")
	}
}

func InitMySQL() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info, //級別
			Colorful:      true,
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	// user := models.UserBasic{}
	// DB.Find(&user)
	// fmt.Println(DB)
}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})

	// pong, err := Red.Ping().Result()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("redis init", pong)
}

const (
	PublishKey = "websocket"
)

func PublishToRedis(ctx context.Context, channel string, msg string) error {
	var err error
	Red.Publish(ctx, channel, msg).Err()
	fmt.Println("publish........")
	return err
}

func SubscribeToRedis(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("err", err)
	fmt.Println("msg")
	fmt.Println("SubscribeToRedis............")
	return msg.Payload, err
}
