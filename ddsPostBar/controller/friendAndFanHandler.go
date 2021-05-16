package controller

import (
	"MyPostBar/ddsPostBar/dao"
	"html/template"
	"net/http"
	"strconv"
)

//IsAgreeToMakeFriend 是否同意成为好友
func IsAgreeToMakeFriend(w http.ResponseWriter,r *http.Request){
	friendID := r.FormValue("friendID")
	FriendID,_ := strconv.ParseInt(friendID,10,64)
	kind := r.FormValue("kind")
	_,sess := dao.IsLogin(r)

	if kind == "agree"{
		//同意成为好友，修改status
		dao.AgreeToMakeFriend(FriendID,sess.UserID)
		//同时将我申请添加对方为好友的消息删除
		dao.DeleteMakeFriend(sess.UserID,FriendID)
	} else {
		//将该记录删除
		dao.DeleteMakeFriend(FriendID,sess.UserID)
	}
	GetMakeFriendMessage(w,r)
}

//GetMyFriend 查看我的好友
func GetMyFriend(w http.ResponseWriter,r *http.Request){
	//获取当前用户
	_,sess := dao.IsLogin(r)
	users := dao.GetFriendByUserID(sess.UserID)
	//遍历切片查看用户姓名
	for _,v := range users{
		user := dao.FindUserByUserID(v.UserID)
		v.UserName = user.UserName
	}
	t := template.Must(template.ParseFiles("views/pages/user/myFriends.html"))
	_ = t.Execute(w,users)
}

//DeleteFriend 删除好友
func DeleteFriend(w http.ResponseWriter,r *http.Request){
	_,sess := dao.IsLogin(r)
	friendID := r.FormValue("friendID")
	FriendID,_ := strconv.ParseInt(friendID,10,64)
	dao.DeleteFriend(sess.UserID,FriendID)
	GetMyFriend(w,r)
}

//AskToMakeFriend 申请添加好友
func AskToMakeFriend(w http.ResponseWriter,r *http.Request){
	//是否登录当前用户
	judge,sess := dao.IsLogin(r)

	if judge == false {
		_,_ = w.Write([]byte("请先登录！"))
		return
	}

	friendID := r.PostFormValue("friendID")
	FriendID,_ := strconv.ParseInt(friendID,10,64)
	//查询添加信息是否发送过
	judge = dao.IsAskToMakeFriend(sess.UserID,FriendID)

	//若没发送过,则增加好友申请信息
	if judge == false {
		dao.AddMakeFriend(sess.UserID,FriendID)
	}

	_,_ = w.Write([]byte("已向对方申请添加为好友！"))
}

//LikedUser 关注用户
func LikedUser(w http.ResponseWriter,r *http.Request){
	//是否登录当前用户
	judge,sess := dao.IsLogin(r)

	if judge == false {
		_,_ = w.Write([]byte("请先登录！"))
		return
	}
	id := r.PostFormValue("id")
	ID,_ := strconv.ParseInt(id,10,64)
	dao.AddLikedUser(ID,sess.UserID)
	_,_ = w.Write([]byte("成功关注！"))
}

//DisLikedUser 取消关注
func DisLikedUser(w http.ResponseWriter,r *http.Request){
	_,sess := dao.IsLogin(r)
	id := r.FormValue("id")
	ID,_ := strconv.ParseInt(id,10,64)
	dao.DeleteLikedUser(ID,sess.UserID)
	GetMyLikedUser(w,r)
}

//GetMyLikedUser 查看我关注的用户
func GetMyLikedUser(w http.ResponseWriter,r *http.Request){
	_,sess := dao.IsLogin(r)
	users := dao.FindLikedUserBySenderID(sess.UserID)
	//查找用户姓名
	for _,v := range users{
		user := dao.FindUserByUserID(v.UserID)
		v.UserName = user.UserName
	}
	t := template.Must(template.ParseFiles("views/pages/user/myLikedUsers.html"))
	_ = t.Execute(w,users)
}

//GetMyFan 查看我的粉丝
func GetMyFan(w http.ResponseWriter,r *http.Request){
	_,sess := dao.IsLogin(r)
	users := dao.FindMyFan(sess.UserID)
	for _,v := range users{
		user := dao.FindUserByUserID(v.UserID)
		v.UserName = user.UserName
		//查看我是否也关注了对方
		v.IsLiked = dao.FindIsLiked(sess.UserID,v.UserID)
	}
	t := template.Must(template.ParseFiles("views/pages/user/myFan.html"))
	_ = t.Execute(w,users)
}