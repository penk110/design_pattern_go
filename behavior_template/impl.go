package behavior_template

import "math/rand"

type OtpImpl interface {
	genRandOtp(gl int) string
	saveMsg(ms string) string
	sendNotification(ms string) error
	pubMetrics()
}

const (
	DefaultRandLen = 6
)

var _LETTERS = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type otp struct {
	OtpImpl
}

func (otp *otp) genRandOtp(gl int) string {
	if gl <= 0 {
		gl = DefaultRandLen
	}
	out := make([]rune, gl)
	for r := range out {
		out[r] = _LETTERS[rand.Intn(len(_LETTERS))]
	}
	return string(out)
}

func (otp *otp) saveMsg(ms string) string {
	// impl

	return ""
}

func (otp *otp) sendNotification(ms string) error {
	// impl

	return nil
}

func (otp *otp) pubMetrics() {
	// impl

	return
}
