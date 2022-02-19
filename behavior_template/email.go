package behavior_template

import (
	"fmt"
	"log"
)

type email struct {
	otp
}

func (email *email) saveMsg(ms string) string {

	fmt.Println("email save msg to cache: " + ms)
	return ""
}

func (email *email) sendNotification(ms string) error {

	log.Println("email send: " + ms)
	return nil
}
