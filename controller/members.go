package controller

import (
	"encoding/json"
	"goMoney/bean"
	"goMoney/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Members struct{}

var members = service.Members{}.New()

func (m Members) New() *Members {
	return &m
}

func (m *Members) VipList(c *gin.Context) {
	res := members.VipLevel()
	if res.Status() != "200" {
		var responseError = &bean.ResponseError{}
		json.Unmarshal(res.Body(), responseError)
		c.JSON(res.RawResponse.StatusCode, responseError)
		return
	}
	var result = &bean.VipLevelResponse{}
	json.Unmarshal(res.Body(), result)
	c.JSON(http.StatusOK, result)
}
