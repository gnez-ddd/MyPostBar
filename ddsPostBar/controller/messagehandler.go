package controller

import (
	"MyPostBar/ddsPostBar/dao"
	"MyPostBar/ddsPostBar/model"
	"html/template"
	"net/http"
)

//GetCreateBarMessages 查看收到的申请创建贴吧的信息
func GetCreateBarMessages(w http.ResponseWriter,r *http.Request){
	//调用数据库查看当前用户所拥有申请创建贴吧的信息
	createBarMessages := dao.GetCreateBarMessageByReceiverID(13)
	//发送给前端
	t := template.Must(template.ParseFiles("views/pages/message/createBarMessages.html"))
	_ = t.Execute(w,createBarMessages)
}

//ToGetMessagesPage 去查看信息页面
func ToGetMessagesPage(w http.ResponseWriter,r *http.Request){
	//获取当前用户
	_,sess := dao.IsLogin(r)
	var isAdmin bool
	//若为系统管理员
	if sess.UserID == 13 {
		isAdmin = true
	}
	t := template.Must(template.ParseFiles("views/pages/message/myMessages.html"))
	_ = t.Execute(w,isAdmin)
}

//GetPostApplicationToRecover 得到申请恢复帖子消息
func GetPostApplicationToRecover(w http.ResponseWriter,r *http.Request){
	//获取当前用户
	_,sess := dao.IsLogin(r)
	//调用数据库中查看其拥有的申请恢复帖子消息
	messages := dao.FindPostApplicationToRecoverMessage(sess.UserID)
	//查看该帖子是否已经被恢复，若已经被恢复则删除帖子再查找
	for _,v := range messages{
		//调用数据库查看该帖子状态
		status := dao.FindStatusByBarNameAndTitleAndKind(v.BarName,v.PostTitle)
		if status == true {
			//该帖子已经被恢复，删除该信息
			dao.DeletePostApplicationToRecoverMessage(v.BarName,v.PostTitle)
		}
	}
	//重新查找
	messages = dao.FindPostApplicationToRecoverMessage(sess.UserID)
	t := template.Must(template.ParseFiles("views/pages/message/postApplication.html"))
	_ = t.Execute(w,messages)
}

//GetReportMessage 获取举报消息
func GetReportMessage(w http.ResponseWriter,r *http.Request){
	//获取吧的举报消息
	barMessages := dao.FindAllBarReportMessage()
	//遍历吧举报消息获取吧主和吧的状态
	for _,v := range barMessages{
		//获取吧在状态
		v.BarStatus = dao.FindBarStatusByBarNameAndKind(v.BarName)
		//获取吧主的状态
		ownerID := dao.FindBarOwnerIDByBarNameAndKind(v.BarName)
		owner := dao.FindUserByUserID(ownerID)
		if owner.Status == 0 {
			v.UserStatus = false
		} else {
			v.UserStatus = true
		}
	}

	//获取帖子的举报消息
	postMessages := dao.FindAllPostReportMessage()
	//遍历帖子举报消息获取发帖者和帖子的状态
	for _,v := range postMessages{
		//获取发帖者
		ownerID := dao.FindOwnerIDByBarNameAndTitle(v.BarName,v.PostTitle)
		owner := dao.FindUserByUserID(ownerID)
		if owner.Status == 0{
			v.UserStatus = false
		} else {
			v.UserStatus = true
		}
		//获取帖子状态
		v.PostStatus = dao.FindStatusByBarNameAndTitleAndKind(v.BarName,v.PostTitle)
	}

	messages := &model.ReportMessages{
		BarMessages: barMessages,
		PostMessages: postMessages,
	}

	t := template.Must(template.ParseFiles("views/pages/message/reportMessage.html"))
	_ = t.Execute(w,messages)
}

//DeleteReport 删除举报消息
func DeleteReport(w http.ResponseWriter,r *http.Request){
	//获取吧名和帖子
	barName := r.FormValue("barName")
	title := r.FormValue("title")

	//若为删除贴吧举报消息
	if title == "nil"{
		dao.DeleteBarReport(barName)
	} else {
		dao.DeletePostReport(barName,title)
	}
	GetReportMessage(w,r)
}

//GetMakeFriendMessage 得到申请成为好友消息
func GetMakeFriendMessage(w http.ResponseWriter,r *http.Request){
	//获取当前用户
	_,sess := dao.IsLogin(r)
	//调用数据库进行查询
	users := dao.FindMakeFriendUserIDByReceiverID(sess.UserID)
	//通过遍历查找请求者用户名
	for _,v := range users{
		user := dao.FindUserByUserID(v.UserID)
		v.UserName = user.UserName
	}
	t := template.Must(template.ParseFiles("views/pages/message/makeFriendMessage.html"))
	_ = t.Execute(w,users)

}