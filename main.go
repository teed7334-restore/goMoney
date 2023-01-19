package main

import (
	"fmt"
	"goMoney/controller"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var members = controller.Members{}.New()
var orders = controller.Orders{}.New()
var k = controller.K{}.New()

func main() {
	webService()
}

func webService() {
	api := doRegister()
	runAt := fmt.Sprintf("127.0.0.1:%s", os.Getenv("port"))
	api.Run(runAt)
}

func doRegister() *gin.Engine {
	route := gin.Default()
	route.GET("/members/vipList", members.VipList)
	route.GET("/orders", orders.Index)
	route.GET("/orders/create", orders.Create)
	route.GET("/orders/delete", orders.Delete)
	route.GET("/orders/clear", orders.Clear)
	route.GET("/k", k.Index)
	route.GET("/k/saveToJson", k.SaveToJson)
	return route
}
