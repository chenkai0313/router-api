package app

import (
	"fmt"
	"github.com/go-redis/redis"
)

func ConnectRedis() *redis.Client {
	fmt.Printf("standalon_redis_test")

	client := redis.NewClient(&redis.Options{
		Addr:     "" + Config.Hostname + ":" + Config.Port + "",
		Password: "" + Config.Password + "",
		DB:       Config.Database,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("ping error[%s]\n", err.Error())
		err_handler(err)
	}
	fmt.Printf("ping result: %s\n", pong)
	return client

}
func err_handler(err error) {
	fmt.Printf("err_handler, error:%s\n", err.Error())
	panic(err.Error())
}
