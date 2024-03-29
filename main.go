package main

import (
	"fmt"
	"goMoney/controller"
	"goMoney/rule"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var rules = rule.Rule{}.New()

var members = controller.Members{}.New()
var orders = controller.Orders{}.New()
var k = controller.K{}.New()
var trades = controller.Trades{}.New()

func main() {
	status := rules.Execute()
	if status <= 0 {
		fmt.Println("Is Expired!!!!")
		os.Exit(status)
	}
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
	route.GET("/members/accounts", members.Accounts)
	route.GET("/orders", orders.Index)
	route.GET("/orders/create", orders.Create)
	route.GET("/orders/delete", orders.Delete)
	route.GET("/orders/clear", orders.Clear)
	route.GET("/k", k.Index)
	route.GET("/k/saveToJson", k.SaveToJson)
	route.GET("/trades/my/of_order", trades.MyOfOrder)
	return route
}
