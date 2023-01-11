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
	payload := base64.StdEncoding.EncodeToString(params)
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(payload))
	result := h.Sum(nil)
	signature := hex.EncodeToString(result)
	return payload, signature
}

func writeLog(err error, resp *resty.Response) {
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
