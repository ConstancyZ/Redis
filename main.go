package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

//测试redis链接
func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	// 商品存入redis的list类型
	//for i := 0; i < 10; i++ {
	//	results := client.LPush("goods", i)
	//	if results != nil {
	//		fmt.Println("results:", results)
	//	}
	//}

	// 使用redis的list队列操作模仿秒杀，因为redis的操作是原子性的,30个人抢redis list中10件商品
	for i := 0; i < 50; i++ {
		go getGoods(client,i)
	}

	time.Sleep(10*time.Second)
	defer client.Close()
}

func getGoods(client *redis.Client, i int) {
	cmd := client.LPop("goods")
	goodID := cmd.Val()
	if goodID != "" || len(goodID) != 0 {
		fmt.Println(i, "顾客抢到", goodID, "商品")
	} else {
		fmt.Println(i, "顾客没有抢到商品")
	}
}
