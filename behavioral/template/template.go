package behavioral

type Iotp interface {
	OTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct{
	iotp Iotp
}

func (o *Otp) getAndSenfOTP(len int) error {
	otp := o.iotp.OTP(len)
	o.iotp.saveOTPCache(otp)
	message := o.iotp.getMessage(otp)
	err := o.iotp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}