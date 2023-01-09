package service

import (
	"encoding/json"
	"fmt"
	"goMoney/bean"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type Members struct{}

func (m Members) New() *Members {
	return &m
}

func (m *Members) VipLevel() *resty.Response {

	path := "/api/v2/members/vip_level"
	queryString := fmt.Sprintf("%s%s", host, path)

	var request = &bean.VipLevelRequest{}
	request.Path = path
	request.Params.Nonce = time.Now().UnixMilli()

	req, _ := json.Marshal(request)

	payload, signature := encode(req, secretKey)

	values, _ := query.Values(request.Params)
	values.Add("arrayFormat", "brackets")

	resp, err := client.R().
		SetHeader("X-MAX-ACCESSKEY", accessKey).
		SetHeader("X-MAX-PAYLOAD", payload).
		SetHeader("X-MAX-SIGNATURE", signature).
		SetHeader("Content-Type", "application/json").
		SetQueryString(values.Encode()).
		SetAuthToken(signature).
		Get(queryString)

	log.Println("Path: ", queryString+"?"+values.Encode())
	log.Println("params: ", string(req))
	log.Println("payload: ", payload)
	log.Println("signature: ", signature)
	writeLog(err, resp)
	return resp
}
