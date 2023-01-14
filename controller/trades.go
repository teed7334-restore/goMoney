package controller

import (
	"encoding/json"
	"goMoney/service"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Trades struct{}

var trades = service.Trades{}.New()

func (t Trades) New() *Trades {
	return &t
}

func (t *Trades) My(c *gin.Context) {
	params := getQuery(c)
	data := trades.My(params)
	var result interface{}
	if err := json.Unmarshal(data.Body(), &result); err != nil {
		return
	}
	res, ok := result.([]interface{})
	if !ok {
		log.Println("Response Encode Error!")
	}
	for key, value := range res {
		row := value.(map[string]interface{})
		createAtTs := int64(row["created_at"].(float64))
		loc, _ := time.LoadLocation("Asia/Taipei")
		t := time.Unix(createAtTs, 0).In(loc)
		row["createdAt"] = t.Format("2006-01-02 15:04:05")
		res[key] = row
	}
	c.JSON(data.StatusCode(), res)
}
