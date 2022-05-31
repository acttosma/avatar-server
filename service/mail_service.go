//
// @File: mail_service.go
// @Version: 1.0.0
// @Date: 2022/5/31 3:01 PM
//
package service

import "net/http"

type MailService struct{}

func (ms *MailService) SendVerifyMail(receiver, verifyCode string) (int, interface{}) {
	return http.StatusOK, nil
}
