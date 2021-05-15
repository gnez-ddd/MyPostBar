package model

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

type Code struct {
	Address string
	CodeNum string
}

const(
	// SmtpMailHost 邮件服务器地址
	SmtpMailHost = "smtp.qq.com"
	// SmtpMailPort 端口
	SmtpMailPort = "587"
	// SmtpMailUser 发送邮件用户账号
	SmtpMailUser = "519364808@qq.com"
	// SmtpMailPwd 授权密码
	SmtpMailPwd = "dhbgwetsgkiybjdi"
	// SmtpMailNickname 发送邮件昵称
	SmtpMailNickname = "dd"
)

//CreateCode 生成随机数
func (code *Code)CreateCode(){
	var num [6]int
	var Num string
	//随机种子
	rand.Seed(time.Now().UnixNano())
	for i := 0;i < 6;i++ {
		num[i] = rand.Intn(10)
		Num += strconv.FormatInt(int64(num[i]),10)
	}
	code.CodeNum = Num
}


//SendMail 发送邮箱验证码
func (code *Code)SendMail(address string)error{
	code.Address = address
	//邮件主题
	subject := "dd's贴吧"
	//邮件内容
	body := "验证码为：" + code.CodeNum

	//认证，content-type设置
	auth := smtp.PlainAuth("", SmtpMailUser, SmtpMailPwd,SmtpMailHost)
	contentType := "Content-Type:text/html;charset=UTF-8"
	s := fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n\r\n%s",address, SmtpMailNickname,SmtpMailUser,subject,contentType,body)
	msg := []byte(s)
	addr := fmt.Sprintf("%s:%s",SmtpMailHost, SmtpMailPort)
	err := smtp.SendMail(addr,auth,SmtpMailUser,[]string{address},msg)
	return err
}

//JudgeCode 验证验证码与邮箱是否同时正确
func (code *Code)JudgeCode(codeInput string,address string)bool{
	if codeInput != code.CodeNum || address != code.Address {
		return false
	} else {
		return true
	}
}


