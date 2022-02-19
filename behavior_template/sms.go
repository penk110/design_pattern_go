package behavior_template

import (
	"fmt"
	"log"
)

type sms struct {
	otp
}

/*
拓展
复用部分
*/

func (sms *sms) saveMsg(ms string) string {

	fmt.Println("sms save msg to cache: " + ms)
	return ""
}

func (sms *sms) sendNotification(ms string) error {

	log.Println("sms send: " + ms)
	return nil
}
