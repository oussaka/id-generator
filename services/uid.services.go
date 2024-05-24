package services

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

const PREIFX = 729385

func GenerateUid() string {
	var timestamp = time.Now().Unix()

	buf := make([]byte, 4)
	_, err := rand.Read(buf)
	if err != nil {
		log.Fatalf("error while generating random string: %s", err)
	}

	return fmt.Sprintf("%v-%v-%x", PREIFX, timestamp, buf)
}
