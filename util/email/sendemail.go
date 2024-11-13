package email

import (

	"github.com/gokins/gokins/comm"
	"gopkg.in/gomail.v2"
)
func SendEmail(to []string, subject string, body string) error {

  if comm.Cfg.Email.Host == "" {
    return nil
  }
  if comm.Cfg.Email.Sender == "" {
    return nil
  }
  if comm.Cfg.Email.Pass == "" {
    return nil
  }
  if comm.Cfg.Email.Port == 0 {
    return nil
  }

  from := comm.Cfg.Email.Sender
  password := comm.Cfg.Email.Pass

  smtpHost := comm.Cfg.Email.Host
  smtpPort := comm.Cfg.Email.Port



  m := gomail.NewMessage()
  m.SetHeader("From", from)
  m.SetHeader("To", to...)
  m.SetHeader("Subject", subject)
  m.SetBody("text/html", body)

  d := gomail.NewDialer(smtpHost, smtpPort, from,password)

  // Send the email to Bob, Cora and Dan.
  if err := d.DialAndSend(m); err != nil {
   return err
  }
  return nil
}
