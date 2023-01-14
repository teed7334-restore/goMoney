package controller

import (
	"encoding/json"
	"goMoney/service"
	"log"

	"github.com/gin-gonic/gin"
)

type Members struct{}

var members = service.Members{}.New()

func (m Members) New() *Members {
	return &m
}

func (m *Members) VipList(c *gin.Context) {
	data := members.VipLevel()
	var result interface{}
	if err := json.Unmarshal(data.Body(), &result); err != nil {
		log.Println("JSON Decode Error")
		return
	}
	c.JSON(data.StatusCode(), result)
}
