package controller

import (
	"MyPostBar/ddsPostBar/dao"
	"MyPostBar/ddsPostBar/model"
	"net/http"
)

//SendEmail 发送验证码
func SendEmail(w http.ResponseWriter,r *http.Request) {
	//获取信息
	kind := r.PostFormValue("kind")
	var address string
	//若为根据用户名获取邮箱
	if kind == "userName" {
		//获取用户
		userName := r.PostFormValue("send")
		user := dao.FindUserByUserName(userName)
		address = user.Email
	} else {
		address = r.PostFormValue("send")
	}

	code := &model.Code{}
	//生成随机数
	code.CreateCode()
	err := code.SendMail(address)
	if err != nil {
		//邮箱有误，发送失败
		_, _ = w.Write([]byte("发送失败"))
		return
	} else {
		_, _ = w.Write([]byte("<font>成功发送！</font>"))
	}
	//若发送成功则将该验证码与邮箱存储进数据库中
	//查找原先有无该邮箱
	id := dao.FindCodeIdByEmail(address)
	if id > 0 {
		dao.SetCodeByID(id, code.CodeNum)
	} else {
		dao.AddCode(address, code.CodeNum)
	}
}

//CheckCode 通过发送Ajax验证验证码与邮箱是否同时正确
func CheckCode(w http.ResponseWriter, r *http.Request) {
	//获取信息
	kind := r.PostFormValue("kind")
	address := r.PostFormValue("address")
	codeInput := r.PostFormValue("codeInput")

	//若通过用户名发送
	if kind == "userName" {
		user := dao.FindUserByUserName(address)
		address = user.Email
	}
	//根据邮箱查找验证码
	code := dao.FindCodeNumByEmail(address)
	if codeInput != code {
		//验证码错误
		_, _ = w.Write([]byte("验证码错误！"))
	} else {
		_, _ = w.Write([]byte("<font >验证码正确！</font>"))
	}
}


