package controller

import "github.com/gin-gonic/gin"

func getQuery(c *gin.Context) map[string]string {
	query := c.Request.URL.Query()
	params := make(map[string]string)
	for key := range query {
		params[key] = c.Query(key)
	}
	delete(params, "nonce")
	return params
}
