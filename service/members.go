package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type Members struct{}

func (m Members) New() *Members {
	return &m
}

func (m *Members) VipLevel() *resty.Response {

	path := "/api/v2/members/vip_level"
	queryString := fmt.Sprintf("%s%s", host, path)

	nonce := time.Now().UnixMilli()
	params := map[string]interface{}{
		"path":  path,
		"nonce": nonce,
	}

	req, _ := json.Marshal(params)

	payload, signature := encode(req)

	resp, err := client.R().
		SetHeader("X-MAX-ACCESSKEY", accessKey).
		SetHeader("X-MAX-PAYLOAD", payload).
		SetHeader("X-MAX-SIGNATURE", signature).
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{"nonce": strconv.FormatInt(nonce, 10)}).
		Get(queryString)

	writeLog(queryString, req, payload, signature, err, resp)
	return resp
}
