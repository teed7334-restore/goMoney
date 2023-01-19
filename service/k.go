package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type K struct{}

func (k K) New() *K {
	return &k
}

func (k *K) Index(params map[string]string) *resty.Response {
	path := "/api/v2/k"
	queryString := fmt.Sprintf("%s%s", host, path)

	nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)
	params["path"] = path
	params["nonce"] = nonce

	req, _ := json.Marshal(params)

	payload, signature := encode(req)

	delete(params, "path")

	resp, err := client.R().
		SetHeader("X-MAX-ACCESSKEY", accessKey).
		SetHeader("X-MAX-PAYLOAD", payload).
		SetHeader("X-MAX-SIGNATURE", signature).
		SetHeader("Content-Type", "application/json").
		SetQueryParams(params).
		Get(queryString)

	writeLog(queryString, req, payload, signature, err, resp)

	return resp
}
