package app

import (
	"fmt"
	"github.com/go-redis/redis"
)

func ConnectRedis()(*redis.Client,error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "" + Config.Hostname + ":" + Config.Port + "",
		Password: "" + Config.Password + "",
		DB:       Config.Database,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return client,fmt.Errorf("ping error[%s]\n", err.Error())
	}
	return client,nil

}
