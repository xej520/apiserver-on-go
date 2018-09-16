package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"xingej-go/Apiserver-go/demo01/router"
)

// 启动一个最简单的RESTful API服务器

func main() {
	//不带中间件的路由
	g := gin.New()
	middlewares := []gin.HandlerFunc{}

	//	加载路由
	// 实际上对 g的初始化
	router.Load(
		//Cores,
		g,
		//Middlwares.
		middlewares...,
	)
	// Ping the server to make sure the router is working.

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())

}

func pingServer() error {
	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to '/health'
		//HTTP GET 请求
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
