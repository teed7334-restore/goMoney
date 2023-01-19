package controller

import (
	"encoding/json"
	"goMoney/service"
	"log"
	"net/http"
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

		updateAtTs := int64(row["updated_at"].(float64))
		t = time.Unix(updateAtTs, 0)
		row["updatedAt"] = t.Format("2006-01-02 15:04:05")
		response = append(response, row)
	}
	c.JSON(data.StatusCode(), response)
}

func (o *Orders) Create(c *gin.Context) {
	params := getQuery(c)
	data := orders.Create(params)
	result := make(map[string]interface{})
	if err := json.Unmarshal(data.Body(), &result); err != nil {
		log.Println("JSON Decode Error!")
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
		return
	}
	createAtTs := int64(result["created_at"].(float64))
	t := time.Unix(createAtTs, 0)
	result["createdAt"] = t.Format("2006-01-02 15:04:05")

	updateAtTs := int64(result["updated_at"].(float64))
	t = time.Unix(updateAtTs, 0)
	result["updatedAt"] = t.Format("2006-01-02 15:04:05")
	c.JSON(data.StatusCode(), result)
}

func (o *Orders) Delete(c *gin.Context) {
	params := getQuery(c)
	data := orders.Delete(params)
	result := make(map[string]interface{})
	if err := json.Unmarshal(data.Body(), &result); err != nil {
		log.Println("JSON Decode Error!")
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false})
		return
	}
	createAtTs := int64(result["created_at"].(float64))
	t := time.Unix(createAtTs, 0)
	result["createdAt"] = t.Format("2006-01-02 15:04:05")

	updateAtTs := int64(result["updated_at"].(float64))
	t = time.Unix(updateAtTs, 0)
	result["updatedAt"] = t.Format("2006-01-02 15:04:05")
	c.JSON(data.StatusCode(), result)
}

func (o *Orders) Clear(c *gin.Context) {
	params := getQuery(c)
	data := orders.Clear(params)
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

		updateAtTs := int64(row["updated_at"].(float64))
		t = time.Unix(updateAtTs, 0)
		row["updatedAt"] = t.Format("2006-01-02 15:04:05")
		response = append(response, row)
	}
	c.JSON(data.StatusCode(), response)
}
