package redisConnect

import (
	"encoding/json"
	"fmt"
	"getReslut/config"
	"getReslut/public"

	"github.com/go-redis/redis"
)

func GetBlockChain(key string) []interface{} {
	redisCongig := config.GetRedisConfig()
	// t := time.Now()
	// now := t.Unix()
	// nowUnix := t.Unix()
	// now := fmt.Sprintf("%4d-%02d-%02d:%d", t.Year(), t.Month(), t.Day(), t.Unix())
	public.Println(fmt.Sprint("redisCongig -------> redisCongig  ", redisCongig))
	public.Println(fmt.Sprint("redisCongig -------> password  ", redisCongig["password"].(string)))

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

	val, err := client.LRange(key, 0, 1000).Result()
	//val, err := client.HGet("BE_AgentLogin", "8f985a9d2d0441a79295d08e5780f71d").Result()
	if err != nil {
		panic(err)
	}

	// fmt.Println("key     ", reflect.TypeOf(val))
	// fmt.Println("key    \n", val)
	var availableChain []interface{} //可用來調整權重的Index  小於目標RTP的獎號
	for _, v := range val {
		//fmt.Printf("%s \n", v)
		in := []byte(v)
		var raw map[string]interface{}
		if err := json.Unmarshal(in, &raw); err != nil {
			panic(err)
		}

		// out, _ := json.Marshal(raw)
		//fmt.Println("raw     ", raw)
		//fmt.Println("out     ", reflect.TypeOf(out))
		availableChain = append(availableChain, raw)
		//println(string(out))
	}
	// fmt.Println("availableChain -------> now  ", len(availableChain))
	// for _, v := range availableChain {
	// 	fmt.Println("now -------> now  ", v.(map[string]interface{}))

	// }
	// public.Println(fmt.Sprint("now -------> now  ", availableChain))

	return availableChain
}
