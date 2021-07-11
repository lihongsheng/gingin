package main

import (
	"fmt"
	"gingin/Util"
	"gingin/config"
	"gingin/types"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	//初始化DB，后期单独放到bootstrat.go里【启动类】
	 //model.InitDB()
	//userModel := new(model.UserModel)
	//user,err := userModel.FindMany("he",1,10)
	//fmt.Println(user)
	//fmt.Println(err)
	h := Util.NewHashMap(Util.HashCrc)
	ipList := []string{"192.168.0.1","192.168.0.2","192.168.0.3","192.168.0.4","192.168.0.1#1","192.168.0.2#2","192.168.0.3#3","192.168.0.4#4"}
	for _, v := range ipList{
		h.Add(v)
	}
	fmt.Println(h.Get("10.10.10.1"))

	for _, v := range h.HashValue{
		fmt.Println("v:%s",v)
	}
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