package redisConnect

import (
	"encoding/json"
	"fmt"
	"getReslut/config"

	"github.com/go-redis/redis"
)

func GetBlockChain(key string) []interface{} {
	redisCongig := config.GetRedisConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     redisCongig["ip"].(string) + ":" + redisCongig["port"].(string),
		Password: "",                            // no password set
		DB:       redisCongig["database"].(int), // use default DB
	})

	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println(pong, err)
	}

	defer client.Close()

	val, err := client.SMembers(key).Result()
	if err != nil {
		panic(err)
	}
	var availableChain []interface{} //可用來調整權重的Index  小於目標RTP的獎號
	for _, v := range val {
		in := []byte(v)
		var raw map[string]interface{}
		if err := json.Unmarshal(in, &raw); err != nil {
			panic(err)
		}

		availableChain = append(availableChain, raw)
	}

	return availableChain
}
