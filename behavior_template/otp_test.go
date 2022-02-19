package behavior_template

import "testing"

func TestSms(t *testing.T) {
	var _otp OtpImpl
	_otp = &sms{}

	msg := "sms test"
	randStr := _otp.genRandOtp(8)
	t.Log("sms rand str: " + randStr)

	_otp.saveMsg(msg)
	_ = _otp.sendNotification(msg)
}

func TestEmail(t *testing.T) {

}
