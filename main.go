package main

import (
	"fmt"
	"gingin/config"
	"gingin/types"
	"github.com/gin-gonic/gin"
	_ "go.etcd.io/etcd/client/v3"
	"strconv"
)

func main() {

}

func bootStrat() {
	gin.SetMode(config.ConfigMap.Mode)
	r := gin.New()
	addr := "0.0.0.0:"+ strconv.Itoa(config.ConfigMap.Port)
	e := new(types.ConfigErr)
	fmt.Println(e.S)
	/**
	// 中间件【钩子函数】全局生效
	r.Use(middleware.Cors())
	//加路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200,"{\"message\":\"success\"}")
	})

	 */
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, struct {
			Name string `json:"name"`
		}{Name: "hello"})
	})
	fmt.Println("start go"+addr)
	_ = r.Run(addr)

}