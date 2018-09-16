package main

import (
	"github.com/spf13/pflag"
	"xingej-go/Apiserver-go/demo02/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"xingej-go/Apiserver-go/demo02/router"
	"time"
	"errors"
	"fmt"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

func main() {
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	// 测试热加载的
	go func() {
		for {
			fmt.Println(viper.GetString("runmode"))
			time.Sleep(4*time.Second)
		}
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(":"+viper.GetString("addr"), g).Error())

}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		log.Println("----->:\t",viper.GetString("url"))
		resp, err := http.Get("http://"+viper.GetString("url") + "/sd/health")

		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
