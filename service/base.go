package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	_ "github.com/joho/godotenv/autoload"
)

var client = resty.New()
var host = os.Getenv("url")
var accessKey = os.Getenv("accessKey")
var secretKey = os.Getenv("secretKey")

func encode(params []byte) (string, string) {
	payload := base64.URLEncoding.EncodeToString(params)
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(payload))
	result := h.Sum(nil)
	signature := hex.EncodeToString(result)
	return payload, signature
}

func writeLog(url string, params []byte, payload, signature string, err error, resp *resty.Response) {
	log.Println("Request Info:")
	log.Println("  URL        :", url)
	log.Println("  Params     :\n", string(params))
	log.Println("  Payload    :", payload)
	log.Println("  Signature  :", signature)
	log.Println()
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
}
