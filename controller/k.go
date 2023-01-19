package controller

import (
	"encoding/json"
	"goMoney/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type K struct{}

var kl = service.K{}.New()

func (k K) New() *K {
	return &k
}

func (k *K) getK(params map[string]string) (int, []map[string]interface{}) {
	data := kl.Index(params)
	var result [][]interface{}
	if err := json.Unmarshal(data.Body(), &result); err != nil {
		log.Println("JSON Decode Error!")
		return data.StatusCode(), nil
	}
	response := make([]map[string]interface{}, 0)
	for _, value := range result {
		row := make(map[string]interface{})
		createAtTs := int64(value[0].(float64))
		t := time.Unix(createAtTs, 0)
		row["createdAt"] = t.Format("2006-01-02 15:04:05")
		open := value[1].(float64)
		high := value[2].(float64)
		low := value[3].(float64)
		close := value[4].(float64)
		volume := value[5].(float64)
		row["open"] = open
		row["high"] = high
		row["low"] = low
		row["close"] = close
		row["volume"] = volume
		row["amplitude"] = (high - low) / low
		row["mean"] = (open + high + low + close) / 4
		response = append(response, row)
	}
	return data.StatusCode(), response
}

func (k *K) Index(c *gin.Context) {
	params := getQuery(c)
	status, response := k.getK(params)
	c.JSON(status, response)
}

func (k *K) SaveToJson(c *gin.Context) {
	params := getQuery(c)
	status, kData := k.getK(params)
	if status != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
		return
	}
	s := kl.SaveToJson(kData)
	if !s {
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
