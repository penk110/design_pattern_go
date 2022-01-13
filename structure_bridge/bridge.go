package structure_bridge

import "log"

/*
	将抽象和实现分别抽离，各自实现
*/

type MessageSenderImp interface {
	Sender(msg string, options ...interface{}) error
}

type EmailSender struct {
	emails []string
}

func (email *EmailSender) Sender(msg string, options ...interface{}) error {

	log.Printf("[EmailSender.Sender] send message: %s\n", msg)
	return nil
}

// NewEmailSender new email email sender
func NewEmailSender(emails []string) *EmailSender {
	// init logic

	return &EmailSender{emails: emails}
}

// 告警 和 告警消息拆开为两个接口，各自扩招各自功能

type NotificationImp interface {
	Notify(msg string, options interface{}) error
}

type ErrorNotification struct {
	sender MessageSenderImp
}

func (ef *ErrorNotification) Notify(msg string, options interface{}) error {
	err := ef.sender.Sender("OOM alert ...")
	if err != nil {
		return err
	}
	return nil
}

// NewErrorNotification new error notify
func NewErrorNotification(sender MessageSenderImp) *ErrorNotification {
	// init logic

	return &ErrorNotification{sender: sender}
}
