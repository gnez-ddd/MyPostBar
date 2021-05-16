// Package controller 用户处理器
package controller

import (
	"MyPostBar/ddsPostBar/dao"
	"MyPostBar/ddsPostBar/model"
	"MyPostBar/ddsPostBar/utils"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// UserRegist 用户注册处理器
func UserRegist(w http.ResponseWriter,r *http.Request){
	//获取用户名密码邮箱验证码
	username := r.PostFormValue("username")
	psw := r.PostFormValue("password")
	email := r.PostFormValue("email")
	codeInput := r.PostFormValue("codeInput")

	//调用数据库检测该用户名邮箱是否注册
	user1 := dao.FindUserByUserName(username)
	user2 := dao.FindUserByEmail(email)
	//判断验证码是否正确
	code := dao.FindCodeNumByEmail(email)
	//若用户名存在或邮箱存在或验证码不正确
	if user1.UserID > 0 || user2.UserID > 0 || codeInput != code{
		//去注册页面发送用户名已存在的信息
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		_ = t.Execute(w,"用户名或邮箱已存在或验证码不正确！")
		return
	}

	//若一切都可以则将验证码表中对应邮箱验证码设置为空
	dao.SetCodeNumNil(email)
	//注册该用户，将信息保存在数据库中,并去往注册成功页面
	//获取盐和密文
	salt := model.GetSalt()
	PSW := model.GetPBKDF2(psw,salt)
	dao.AddUser(username,PSW,email,salt)
	t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
	_ = t.Execute(w,"成功注册！")
}

// UserLogin 用户登录处理器
func UserLogin(w http.ResponseWriter,r *http.Request){
	//获取用户名密码验证码
	username := r.PostFormValue("username")
	psw := r.PostFormValue("password")
	codeInput := r.PostFormValue("codeInput")

	//查找有无该用户
	user := dao.FindUserByUserName(username)
	//若用户不存在或用户被封禁
	if user.UserID <= 0 || user.Status == 0 {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		_ = t.Execute(w, "用户名或密码错误！")
		return
	}

	//用户存在
	//判断验证码账号密码是否正确
	code := dao.FindCodeNumByEmail(user.Email)
	//利用用户的盐获取用户输入密码的密文
	PSW := model.GetPBKDF2(psw,user.Salt)
	//判断
	if code != codeInput || PSW != user.PassWord{
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		_ = t.Execute(w,"验证码或密码错误！")
	} else {
		//若一切都可以则将验证码表中对应邮箱验证码设置为空
		dao.SetCodeNumNil(user.Email)
		//生成UUID作为session的id
		uuid := utils.CreateUUID()
		//建立一个session
		sess := &model.Session{
			SessionID: uuid,
			UserName: user.UserName,
			UserID: user.UserID,
		}
		//将session保存在数据库中
		dao.AddSession(sess)
		//创建一个Cookie与Session相关联
		cookie := http.Cookie{
			Name:"userhandler",
			Value:uuid,
			HttpOnly: true,
		}
		//将cookie发送给浏览器并去往首页
		http.SetCookie(w,&cookie)
		//去登录成功页面
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		_ = t.Execute(w,"")
	}
}

// FindPassWord 用户密码找回处理器
func FindPassWord(w http.ResponseWriter,r *http.Request){
	//获取邮箱及验证码
	email := r.PostFormValue("email")
	codeInput := r.PostFormValue("codeInput")
	code := dao.FindCodeNumByEmail(email)
	//根据邮箱查找用户
	user := dao.FindUserByEmail(email)

	//如果用户不存在或验证码输入错误
	if user.UserID <= 0 || user.Status == 0 || codeInput != code {
		t := template.Must(template.ParseFiles("views/pages/user/find_password.html"))
		_ = t.Execute(w,"验证码错误或该用户不存在！")
	} else {
		//若一切都可以则将验证码表中对应邮箱验证码设置为空
		dao.SetCodeNumNil(email)
		//获取密码
		password := r.PostFormValue("password")
		//获取盐值
		salt := model.GetSalt()
		//获取密文
		PSW := model.GetPBKDF2(password,salt)

		//调用数据库修改密码和盐值
		dao.SetPassword(PSW,email,salt)
		t := template.Must(template.ParseFiles("views/pages/user/find_password_success.html"))
		_ = t.Execute(w,"")
	}
}

// ShowUserInformation 显示用户信息
func ShowUserInformation(w http.ResponseWriter,r *http.Request){
	//获取当前登录的用户
	_,sess := dao.IsLogin(r)
	//根据sessid查找当前用户
	user := dao.FindUserByUserID(sess.UserID)
	t := template.Must(template.ParseFiles("views/pages/user/set_information.html"))
	_ = t.Execute(w,user)
}

//ToProveUserPage 去用户身份验证页面
func ToProveUserPage(w http.ResponseWriter,r *http.Request){
	//获取当前登录的用户
	_,sess := dao.IsLogin(r)
	//根据sessid查找当前用户
	user := dao.FindUserByUserID(sess.UserID)
	t := template.Must(template.ParseFiles("views/pages/user/prove_user.html"))
	_ = t.Execute(w,user)
}

// ProveUser 用户身份验证
func ProveUser(w http.ResponseWriter,r *http.Request){
	//获取当前登录的用户
	_,sess := dao.IsLogin(r)
	//根据sessid查找当前用户
	user := dao.FindUserByUserID(sess.UserID)

	//获取用户输入密码
	password := r.PostFormValue("password")
	//根据用户的盐值生成密文查看是否与用户密文相同
	psw := model.GetPBKDF2(password,user.Salt)
	//若密码码正确
	if psw == user.PassWord  {
		//去修改页面
		t := template.Must(template.ParseFiles("views/pages/user/set_information.html"))
		_ = t.Execute(w,user)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/prove_user.html"))
		_ = t.Execute(w,user)
	}
}

//SetUserInformation 用户信息修改
func SetUserInformation(w http.ResponseWriter,r *http.Request){
	//获取当前登录的用户
	_,sess := dao.IsLogin(r)
	//获取用户输入的验证码与邮箱
	codeInput := r.PostFormValue("codeInput")
	email := r.PostFormValue("newEmail")
	code := dao.FindCodeNumByEmail(email)
	//根据sessid查找当前用户
	user := dao.FindUserByUserID(sess.UserID)

	if email != "" && codeInput != "" {
		//判断用户名与邮箱是否正确
		if codeInput != code {
			t := template.Must(template.ParseFiles("views/pages/user/set_information.html"))
			_ = t.Execute(w,user)
			return
		}
	}
	//若邮箱有修改
	if email != ""{
		user.Email = email
	}

	//获取头像并上传到文件中
	if r.Method == "GET" {
		data,_ := ioutil.ReadFile("views/pages/user/set_information.html")

		_,_ = io.WriteString(w,string(data))

	} else if r.Method == "POST" {
		file,head,_ := r.FormFile("image")

		//若有头像
		if file != nil {
			//修改用户头像地址
			user.HeadPath = "../../static/img/head/" + strconv.FormatInt(user.UserID,10) + head.Filename
			defer  file.Close()
			//在本地创建一个新文件去接受上传的文件
			newFile,_ := os.OpenFile("views/static/img/head/" + strconv.FormatInt(user.UserID,10) + head.Filename,os.O_CREATE | os.O_APPEND,0666)
			defer newFile.Close()
			_,_ = io.Copy(newFile,file)
		}
	}
	//获取用户名
	username := r.PostFormValue("newUsername")
	//若用户名有修改
	if username != ""{
		user.UserName = username
	}

	//获取密码
	password := r.PostFormValue("newPSW")
	//若密码有修改
	if password != ""{
		//获取盐值
		user.Salt = model.GetSalt()
		//获取密文
		user.PassWord = model.GetPBKDF2(password,user.Salt)
	}

	//调用数据库修改用户信息
	dao.SetUserInformation(user.UserName,user.Email,user.HeadPath,user.UserID,user.Salt,user.PassWord)
	t := template.Must(template.ParseFiles("views/pages/user/set_information_success.html"))
	_ = t.Execute(w,"")
}

//ToIndexPage 去首页
func ToIndexPage(w http.ResponseWriter,r *http.Request){
	//创建page变量
	page := model.Page{}
	//获取贴吧
	page.Bar = dao.GetAllBars()
	//获取帖子
	page.Posts = dao.GetAllCreatePostOrderByExperience()
	//判断是否已经登录
	judge,sess := dao.IsLogin(r)

	//若已经登录,则获取用户
	if judge == true {
		page.IsLogin = true
		page.User = dao.FindUserByUserID(sess.UserID)
		page.User.IsLogin = true

		//判断用户是否已经签到了
		//获取当前时间
		tm := time.Now()
		now := tm.Format("2006-01-02")
		if page.User.LastSign == now {
			page.User.IsSignIn = true
		}

		t := template.Must(template.ParseFiles("views/userIndex.html"))
		_ = t.Execute(w,page)

	} else {
		//若为游客前往首页
		t := template.Must(template.ParseFiles("views/userIndex.html"))
		_ = t.Execute(w,page)

	}
}

// UserLogout 用户退出登录
func UserLogout(w http.ResponseWriter,r *http.Request){
	//获取cookie
	cookie,_ := r.Cookie("userhandler")
	//获取cookie的value值
	cookieValue := cookie.Value
	//将数据库中与cookieValue对应的session删除
	dao.DeleteSession(cookieValue)
	//设置cookie失效
	cookie.MaxAge = -1
	//将修改之后的cookie发送给浏览器告诉浏览器该cookie已失效
	http.SetCookie(w, cookie)
	//注销成功后去往首页
	ToIndexPage(w,r)
}

//CheckUserNameOrEmail 通过发送Ajax验证用户名或验证码是否可用
func CheckUserNameOrEmail(w http.ResponseWriter, r *http.Request) {
	//获取用户输入的要检查的信息
	check := r.PostFormValue("check")
	kind := r.PostFormValue("kind")
	var user *model.Users
	//查询有无用户
	if kind == "userName" {
		user = dao.FindUserByUserName(check)
	} else {
		user = dao.FindUserByEmail(check)
	}

	//若存在
	if user.UserID > 0 {
		//已存在
		_, _ = w.Write([]byte("已存在！"))
	} else {
		//用户名可用
		_, _ = w.Write([]byte("<font >可用！</font>"))
	}
}

//LookHistory 访问历史记录
func LookHistory(w http.ResponseWriter,r *http.Request){
	//获取当前用户
	_,sess := dao.IsLogin(r)
	//调用数据库查看其访问过的贴吧
	bars := dao.FindHistoryBarsByUserID(sess.UserID)
	//调用数据库查看其访问过的帖子
	posts := dao.FindHistoryPostsByUserID(sess.UserID)
	page := &model.Page{
		Bar: bars,
		Posts: posts,
	}
	t := template.Must(template.ParseFiles("views/pages/historyLook/history_look.html"))
	_ = t.Execute(w,page)
}

//SignIn 签到
func SignIn(w http.ResponseWriter,r *http.Request){
	//获取当前用户
	_,sess := dao.IsLogin(r)
	user := dao.FindUserByUserID(sess.UserID)
	user.Experience++
	//获取当前时间
	tm := time.Now()
	now := tm.Format("2006-01-02")
	//修改用户信息
	dao.UserSignIn(user.UserName,now,user.Experience)
	_,_ = w.Write([]byte("签到成功！"))
}

//BannedUser 封禁用户
func BannedUser(w http.ResponseWriter,r *http.Request){
	//获取吧名和帖子
	barName := r.FormValue("barName")
	title := r.FormValue("title")

	//若为封禁吧主
	if title == "nil"{
		//查找吧主
		ownerID := dao.FindBarOwnerIDByBarNameAndKind(barName)
		//根据id修改用户status
		dao.SetUserStatusByID(0,ownerID)
	} else {
		//封禁发帖者
		ownerID := dao.FindOwnerIDByBarNameAndTitle(barName,title)
		//根据id修改用户status
		dao.SetUserStatusByID(0,ownerID)
	}
	GetReportMessage(w,r)
}

//LookUser 查看用户
func LookUser(w http.ResponseWriter,r *http.Request){
	//当前用户是否登录
	judge,sess := dao.IsLogin(r)
	userName := r.FormValue("userName")
	user := dao.FindUserByUserName(userName)
	//如果用户不存在
	if user.Status == 0 {
		t := template.Must(template.ParseFiles("views/pages/user/look_user_failed.html"))
		_ = t.Execute(w,"")
		return
	}
	//如果用户存在
	//如果用户有登录
	if judge == true {
		//判断是否为好友
		user.IsFriend = dao.FindIsFriend(user.UserID,sess.UserID)
		//判断是否关注了
		user.IsLiked = dao.FindIsLiked(sess.UserID,user.UserID)
		//若查看的用户为用户本人
		if user.UserID == sess.UserID {
			user.IsLiked = true
			user.IsFriend = true
		}
	}


	t := template.Must(template.ParseFiles("views/pages/user/look_user.html"))
	_ = t.Execute(w,user)
}

