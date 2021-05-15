package controller

import (
	"MyPostBar/ddsPostBar/dao"
	"MyPostBar/ddsPostBar/model"
	"net/http"
)

//SendEmailByUserName 根据用户名发送密码
func SendEmailByUserName(w http.ResponseWriter,r *http.Request){
	//获取用户名
	userName := r.PostFormValue("userName")
	//根据用户名查找用户
	user := dao.FindUserByUserName(userName)
	//若用户不存在
	if user.UserID <= 0 {
		_,_ = w.Write([]byte("发送失败"))
		return
	} else {
		//获取收件人邮箱
		address := user.Email
		code := &model.Code{}
		//生成随机数
		code.CreateCode()
		err := code.SendMail(address)
		if err != nil {
			//邮箱有误，发送失败
			_,_ = w.Write([]byte("发送失败"))
		} else {
			_,_ = w.Write([]byte("<font>成功发送！</font>"))
			//将该验证码与邮箱存储进数据库中
			//查找原先有无该邮箱
			id := dao.FindCodeIdByEmail(address)
			if id > 0 {
				dao.SetCodeByID(id,code.CodeNum)
			} else {
				dao.AddCode(address,code.CodeNum)
			}
		}
	}
}

//SendEmail 发送验证码
func SendEmail(w http.ResponseWriter,r *http.Request){
	//获取收件人邮箱
	address := r.PostFormValue("email")

	code := &model.Code{}
	//生成随机数
	code.CreateCode()
	err := code.SendMail(address)
	if err != nil {
		//邮箱有误，发送失败
		_,_ = w.Write([]byte("发送失败"))
	} else {
		_,_ = w.Write([]byte("<font>成功发送！</font>"))
		//将该验证码与邮箱存储进数据库中
		//查找原先有无该邮箱
		id := dao.FindCodeIdByEmail(address)
		if id > 0 {
			dao.SetCodeByID(id,code.CodeNum)
		} else {
			dao.AddCode(address,code.CodeNum)
		}
	}
}

//CheckCode 通过发送Ajax验证验证码与邮箱是否同时正确
func CheckCode(w http.ResponseWriter, r *http.Request) {

	//获取用户输入的验证码与邮箱
	address := r.PostFormValue("email")
	codeInput := r.PostFormValue("codeInput")
	//根据邮箱查找验证码
	code := dao.FindCodeNumByEmail(address)
	if codeInput != code {
		//验证码错误
		_, _ = w.Write([]byte("验证码错误！"))
	} else {
		_, _ = w.Write([]byte("<font >验证码正确！</font>"))
	}
}

//CheckCodeByUserName 根据用户名验证验证码是否正确
func CheckCodeByUserName(w http.ResponseWriter,r *http.Request){
	//获取验证码与用户名
	userName := r.PostFormValue("userName")
	codeInput := r.PostFormValue("codeInput")
	//根据用户名查找邮箱
	user := dao.FindUserByUserName(userName)
	//根据邮箱查找验证码
	code := dao.FindCodeNumByEmail(user.Email)
	if codeInput != code {
		//验证码错误
		_, _ = w.Write([]byte("验证码错误！"))
	} else {
		_, _ = w.Write([]byte("<font >验证码正确！</font>"))
	}
}

