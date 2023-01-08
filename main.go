package main

import (
	"fmt"
	"goMoney/controller"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var members = controller.Members{}.New()

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
	return route
}
