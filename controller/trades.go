package controller

import (
	"encoding/json"
	"goMoney/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Trades struct{}

var trades = service.Trades{}.New()

func (t Trades) New() *Trades {
	return &t
}

func (t *Trades) MyOfOrder(c *gin.Context) {
	params := getQuery(c)
	data := trades.MyOfOrder(params)
	result := make([]interface{}, 0)
	if err := json.Unmarshal(data.Body(), &result); err != nil {
		log.Println("JSON Decode Error!")
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
		return
	}
	response := make([]map[string]interface{}, 0)
	for _, value := range result {
		row := value.(map[string]interface{})

		createAtTs := int64(row["created_at"].(float64))
		t := time.Unix(createAtTs, 0)
		row["createdAt"] = t.Format("2006-01-02 15:04:05")

		response = append(response, row)
	}
	c.JSON(data.StatusCode(), response)
}
