package controller

import (
	"encoding/json"
	"goMoney/service"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Orders struct{}

var orders = service.Orders{}.New()

func (o Orders) New() *Orders {
	return &o
}

func (o *Orders) Index(c *gin.Context) {
	params := getQuery(c)
	data := orders.Index(params)
	var result interface{}
	if err := json.Unmarshal(data.Body(), &result); err != nil {
		log.Println("JSON Decode Error!")
		return
	}
	res, ok := result.([]interface{})
	if !ok {
		log.Println("Response Decode Error!")
		c.JSON(data.StatusCode(), result)
		return
	}
	for key, value := range res {
		row := value.(map[string]interface{})

		createAtTs := int64(row["created_at"].(float64))
		t := time.Unix(createAtTs, 0)
		row["createdAt"] = t.Format("2006-01-02 15:04:05")

		updateAtTs := int64(row["updated_at"].(float64))
		t = time.Unix(updateAtTs, 0)
		row["updatedAt"] = t.Format("2006-01-02 15:04:05")
		res[key] = row
	}
	c.JSON(data.StatusCode(), res)
}

func (o *Orders) Create(c *gin.Context) {
	params := getQuery(c)
	data := orders.Create(params)
	var result interface{}
	if err := json.Unmarshal(data.Body(), &result); err != nil {
		log.Println("JSON Decode Error!")
		return
	}
	res, ok := result.(map[string]interface{})
	if !ok {
		log.Println("Response Decode Error!")
		c.JSON(data.StatusCode(), result)
		return
	}
	createAtTs := int64(res["created_at"].(float64))
	t := time.Unix(createAtTs, 0)
	res["createdAt"] = t.Format("2006-01-02 15:04:05")

	updateAtTs := int64(res["updated_at"].(float64))
	t = time.Unix(updateAtTs, 0)
	res["updatedAt"] = t.Format("2006-01-02 15:04:05")
	c.JSON(data.StatusCode(), res)
}

func (o *Orders) Delete(c *gin.Context) {
	params := getQuery(c)
	data := orders.Delete(params)
	var result interface{}
	if err := json.Unmarshal(data.Body(), &result); err != nil {
		log.Println("JSON Decode Error!")
		return
	}
	res, ok := result.(map[string]interface{})
	if !ok {
		log.Println("Response Decode Error!")
		c.JSON(data.StatusCode(), result)
		return
	}
	createAtTs := int64(res["created_at"].(float64))
	t := time.Unix(createAtTs, 0)
	res["createdAt"] = t.Format("2006-01-02 15:04:05")

	updateAtTs := int64(res["updated_at"].(float64))
	t = time.Unix(updateAtTs, 0)
	res["updatedAt"] = t.Format("2006-01-02 15:04:05")
	c.JSON(data.StatusCode(), res)
}
