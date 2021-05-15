package controller

import (
	"MyPostBar/ddsPostBar/dao"
	"MyPostBar/ddsPostBar/model"
	"html/template"
	"net/http"
)

// BarCreate 创建贴吧
func BarCreate(w http.ResponseWriter,r *http.Request){
	//获取当前申请的用户
	_,sess := dao.IsLogin(r)
	user := dao.FindUserByUserID(sess.UserID)
	//获取吧名
	barName := r.PostFormValue("barName")
	//查看该吧名是否存在
	bar := dao.FindBarByBarNameAndKind(barName)
	if bar.BarID > 0{
		//该吧名已经存在
		t := template.Must(template.ParseFiles("views/pages/bar/create_bar.html"))
		_ = t.Execute(w,"该吧名已存在！")
	} else {
		kind := "create"
		//调用数据库将该帖吧存放进数据库
		dao.AddBar(user.UserID,barName,kind)
		//在管理员的信息表中添加申请创建贴吧信息
		dao.AddCreateBarMessage(13,sess.UserID,barName)
		//去往等待审核页面
		t := template.Must(template.ParseFiles("views/pages/bar/create_bar_wait.html"))
		_ = t.Execute(w,"")
	}
}

//CheckBarName 通过Ajax请求验证吧名是否可用
func CheckBarName(w http.ResponseWriter,r *http.Request){
	//获取吧名
	barName := r.PostFormValue("barName")
	//根据吧名调用数据库查询看有无此贴吧
	bar := dao.FindBarByBarNameAndKind(barName)

	if bar.BarID > 0 {
		//此吧存在
		_, _ = w.Write([]byte("该吧名已存在！"))
	} else {
		_, _ = w.Write([]byte("<font >该吧名可用！</font>"))
	}
}

//IsAgreeToCreateBar 是否同意创建贴吧
func IsAgreeToCreateBar(w http.ResponseWriter,r *http.Request){
	//获取吧名和判断
	barName := r.PostFormValue("barName")
	kind := r.PostFormValue("kind")
	//同意创建贴吧,修改贴吧状态
	if kind == "agree" {
		dao.SetBarStatusByBarName(barName, 1)
	} else {
		//不同意则根据吧名删除给贴吧
		dao.DeleteBarByBarName(barName)
	}
	//根据吧名删除申请创建贴吧消息
	dao.DeleteCreateBarMessageByBarName(barName)
}

//GetMyBar 查看我的贴吧
func GetMyBar(w http.ResponseWriter,r * http.Request){
	//获取当前用户
	_,sess:= dao.IsLogin(r)
	//调用数据库查看属于我创建的贴吧和我关注的贴吧
	mybar := dao.GetBarsByUserID(sess.UserID)
	//发送给页面
	t := template.Must(template.ParseFiles("views/pages/bar/mybar.html"))
	_ = t.Execute(w,mybar)
}

//GoToBar 进入贴吧
func GoToBar(w http.ResponseWriter,r *http.Request){
	//查看是否登录
	judge,sess := dao.IsLogin(r)
	//获取吧名
	barName := r.FormValue("barName")
	//调用数据库根据吧名查看吧的信息
	bar := dao.FindBarByBarNameAndKind(barName)
	//获取吧主信息
	barOwner := dao.FindUserByUserID(bar.BarHostID)
	bar.BarHostName = barOwner.UserName
	//获取该吧的公告
	bar.MyNotice = dao.FindNoticesByBarName(bar.BarName)
	//获取该吧的帖子
	posts := dao.FindCreatePostsByBarName(bar.BarName)
	//为了符合page中bar的数据类型，将bar添加到切片中
	var bars []*model.Bar
	bars = append(bars,bar)
	//将所有信息添加到page中
	page := &model.Page{
		Posts: posts,
		Bar: bars,
	}

	//若登录
	if judge == true {
		//获取用户
		user := dao.FindUserByUserID(sess.UserID)

		//判断该用户是否为吧主或系统管理员
		if bar.BarHostName == user.UserName || user.UserName == "admin"  {
			bar.IsBarOwner = true
			//遍历该吧的公告和帖子将isBarOwner设置为true
			for _,v := range bar.MyNotice {
				v.IsBarOwner = true
			}
			for _,v := range posts {
				v.IsBarOwner = true
			}
		}
		//查看当前用户是否关注了该吧
		isLiked := dao.IsLikedByBarNameAndUserID(barName,sess.UserID)

		if isLiked == true {
			bar.IsLiked = true
		}
		page.User = user
		//将访问贴吧的信息存储到历史记录表中
		dao.AddBarHistory(barName,user.UserID)

	} else {
		//若未登录
		user := &model.Users{}
		page.User = user
	}
	//去该吧页面
	t := template.Must(template.ParseFiles("views/pages/bar/look_bar.html"))
	_ = t.Execute(w,page)

}

//IsLikedBar 是否关注贴吧
func IsLikedBar(w http.ResponseWriter,r *http.Request){
	//获取吧名和判断
	barName := r.PostFormValue("barName")
	bar := dao.FindBarByBarNameAndKind(barName)
	kind := r.PostFormValue("kind")
	judge,sess := dao.IsLogin(r)

	//取消关注
	if kind == "disliked" {
		//调用数据库删除该条信息
		dao.DisLikedBar(barName,sess.UserID)
		//修改该把的粉丝信息
		bar.TotalFan--
		dao.SetTotalFanByBarName(bar.TotalFan,bar.BarName)
		_, _ = w.Write([]byte("已取消关注！"))
		return
	}

	//关注
	//判断是否登录了
	//若未登录，则发送给前端未登录的信息
	if judge == false {
		_, _ = w.Write([]byte("请先登录！"))
	} else {
		bar.BarHostID = sess.UserID
		bar.Kind = "liked"
		//调用数据库添加关注bars信息
		dao.LikedBar(bar)
		//修改该把的粉丝信息
		bar.TotalFan++
		dao.SetTotalFanByBarName(bar.TotalFan,bar.BarName)
		_, _ = w.Write([]byte("成功关注！"))
	}
}

//FindBarsAndPosts 搜索贴吧和帖子
func FindBarsAndPosts(w http.ResponseWriter,r *http.Request){
	find := r.FormValue("find")
	//调用数据库查询贴吧和帖子
	bars := dao.FindBars("%"+find+"%")
	posts := dao.FindPostsByFind("%"+find+"%")
	page := &model.Page{
		Posts: posts,
		Bar: bars,
	}
	t := template.Must(template.ParseFiles("views/find_bars_posts.html"))
	_ = t.Execute(w,page)
}

//DeleteBarOrPostHistory 删除贴吧或帖子访问记录
func DeleteBarOrPostHistory(w http.ResponseWriter,r *http.Request){
	//获取当前用户
	_,sess := dao.IsLogin(r)
	barName := r.FormValue("barName")
	time := r.FormValue("time")
	title := r.FormValue("title")
	kind := r.FormValue("kind")

	if kind == "bar" {
		dao.DeleteBarHistory(barName,sess.UserID,time)
	} else {
		dao.DeletePostHistory(barName,title,sess.UserID,time)
	}
	//回到足迹页面
	LookHistory(w,r)
}

//ReportBarOrPost 举报贴吧或帖子
func ReportBarOrPost(w http.ResponseWriter,r *http.Request){
	//获取信息
	barName := r.PostFormValue("barName")
	title := r.PostFormValue("title")
	kind := r.PostFormValue("kind")

	//举报贴吧
	if kind == "bar" {
		//查看该吧是否被举报过
		id := dao.FindReportBarByBarName(barName)
		if id <= 0 {
			//还未被举报则增加举报信息
			dao.AddReportBarMessage(barName)
		}
	} else {
		//查看该帖子是否被举报过
		id := dao.FindReportPostByBarNameAndTitle(barName,title)
		if id <= 0 {
			//还未被举报过
			dao.AddReportPostMessage(barName,title)
		}
	}
	_,_ = w.Write([]byte("举报成功！"))
}

//BannedBar 封禁贴吧
func BannedBar(w http.ResponseWriter,r *http.Request){
	//获取吧名
	barName := r.FormValue("barName")
	//查找该吧下的帖子
	posts := dao.FindCreatePostsByBarName(barName)
	//遍历帖子，将帖子status设置为z0
	for _,v := range posts{
		dao.DeletePostFake(v.BarName,v.PostTitle,0)
	}
	//将该吧的status设置为0
	dao.SetBarStatusByBarName(barName,0)
	GetReportMessage(w,r)
}