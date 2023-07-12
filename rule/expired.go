package rule

import (
	"log"
	"time"
)

const expired = ""

func (r *Rule) validateExpired() int {
	now := time.Now()
	checkDate := now.Format("2006-01-02")

	if expired == "" {
		return -1
	}

	checkDateTs, err := time.Parse("2006-01-02", checkDate)
	if err != nil {
		log.Println(err.Error())
		return -2
	}

	expiredTs, err := time.Parse("2006-01-02", expired)
	if err != nil {
		log.Println(err.Error())
		return -2
	}

	if expiredTs.Before(checkDateTs) {
		return 0
	}

	return 1
}
