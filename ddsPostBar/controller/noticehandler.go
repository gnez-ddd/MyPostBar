package controller

import (
	"MyPostBar/ddsPostBar/dao"
	"html/template"
	"net/http"
)

//ToSendNoticePage 去发布公告页面
func ToSendNoticePage(w http.ResponseWriter,r *http.Request){
	//获取吧名
	barName := r.FormValue("barName")
	t := template.Must(template.ParseFiles("views/pages/notice/send_notice.html"))
	_ = t.Execute(w,barName)
}


//SendNotice 通过Ajax发布公告
func SendNotice(w http.ResponseWriter,r *http.Request){
	//获取吧名
	barName := r.PostFormValue("barName")
	//获取公告内容
	content := r.PostFormValue("content")
	//向公告表添加公告
	dao.AddNotice(barName,content)
	//发布完成后向前端返回发布成功信息
	_,_ = w.Write([]byte("发布公告成功！"))

}

//DeleteNotice 删除公告
func DeleteNotice(w http.ResponseWriter,r *http.Request){
	//获取吧名和时间
	barName := r.FormValue("barName")
	time := r.FormValue("time")
	//调用数据库将该公告删除
	dao.DeleteNoticeByBarNameAndTime(barName,time)
	GoToBar(w,r)
}