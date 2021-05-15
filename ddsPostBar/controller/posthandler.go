package controller

import (
	"MyPostBar/ddsPostBar/dao"
	"MyPostBar/ddsPostBar/model"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

//PostCreate 写帖子
func PostCreate(w http.ResponseWriter,r *http.Request){
	judge,sess := dao.IsLogin(r)
	//若未登录
	if judge == false {
		_, _ = w.Write([]byte("请先登录！"))
	} else {

		//获取贴吧名当前用户内容标题
		barName := r.PostFormValue("barName")
		content := r.PostFormValue("content")
		title := r.PostFormValue("title")
		user := dao.FindUserByUserID(sess.UserID)


		//验证标题是否存在
		post := dao.FindCreatePostByBarNameAndTitle(barName,title)

		//若存在
		if post.PostID > 0 {
			_, _ = w.Write([]byte("该标题已存在！"))

		} else {
			//将内容存放进文件夹中
			postPath := "views/static/post/" + barName + title + ".txt"
			file,_ := os.OpenFile("views/static/post/" + barName + title + ".txt" ,os.O_CREATE | os.O_APPEND,0666)

			defer file.Close()
			_,_ = file.WriteString(content)

			//将吧名当前用户时间内容路径标题存放进数据库中
			dao.AddPost(barName,user.UserID,title,postPath)

			_, _ = w.Write([]byte("发贴成功！"))

		}
	}
}

//GetMyPosts 查看创建和关注和删除和被删除的贴子
func GetMyPosts(w http.ResponseWriter,r *http.Request){
	//获取用户
	_,sess := dao.IsLogin(r)

	user := dao.FindUserByUserID(sess.UserID)

	//调用数据库根据用户名和种类查看创建和关注和删除的帖子
	myPosts := dao.FindPostsByUserIDAndKind(sess.UserID)


	//若该用户为系统管理员，则查找所有被删除的帖子
	if user.UserName == "admin" {
		myPosts.PostDelete = dao.FindAllDeletePosts()

	}

	t := template.Must(template.ParseFiles("views/pages/post/my_post.html"))
	_ = t.Execute(w,myPosts)
}

/*
//DeletePost 删除帖子真正
func DeletePost(w http.ResponseWriter,r *http.Request){
	//获取吧名和帖子标题
	barName := r.FormValue("barName")
	postTitle := r.FormValue("postTitle")
	now := r.FormValue("now")

	//查找帖子回复内容
	postReplys := dao.FindReplyByBarNameAndTitleOrderByTime(barName,postTitle)
	//查找帖子回复的用户回复并删除并删除帖子回复内容
	for _,v := range postReplys{
		userReplys := dao.FindUserReplyByPostReplyID(v.ID)
		//通过遍历将用户回复内容删除
		for _,userReply := range userReplys{
			//根据用户回复id查找用户回复内容路径
			userReply.Content = dao.FindContentByUserReplyID(userReply.UserReplyID)
			_ = os.Remove(userReply.Content)
			//调用数据库将该用户回复删除
			dao.DeleteUserReplyByID(userReply.UserReplyID)
		}
		//查找帖子回复内容路径
		v.Content = dao.FindPostReplyContentByID(v.ID)
		_ = os.Remove(v.Content)
		//调用数据库根据帖子回复id将帖子回复删除
		dao.DeletePostReplyByID(v.ID)
	}

	//查找帖子内容地址
	post := dao.FindPostByBarNameAndTitleAndKind(barName,postTitle)
	post.PostContent = dao.FindPostContentByBarNameAndTitleAndKind(post.BarName,post.PostTitle)

	//删除帖子内容
	_ = os.Remove(post.PostContent)
	//根本吧名和标题删除帖子
	dao.DeletePostsByBarNameAndPostTitle(barName,postTitle)
	if now == "LookBar" {
		//去往该贴吧
		GoToBar(w,r)
	 } else {
	 	//去我的帖子
	 	GetMyPosts(w,r)
	}

}

 */

//DeletePostFake 删除帖子
func DeletePostFake(w http.ResponseWriter,r *http.Request){
	//获取吧名和帖子标题
	barName := r.FormValue("barName")
	postTitle := r.FormValue("postTitle")
	now := r.FormValue("now")
	//获取当前用户
	_,sess := dao.IsLogin(r)
	//修改删除帖子信息
	dao.DeletePostFake(barName,postTitle,sess.UserID)
	if now == "LookBar" {
		//去往该贴吧
		GoToBar(w,r)
	} else {
		//去我的帖子
		GetMyPosts(w,r)

	}
}

//DisDeletePost 恢复帖子
func DisDeletePost(w http.ResponseWriter,r *http.Request){
	//获取吧名和帖子标题
	barName := r.PostFormValue("barName")
	postTitle := r.PostFormValue("title")
	//获取给吧的状态
	status := dao.FindBarStatusByBarNameAndKind(barName)

	//若该吧存在
	if status == true {
		//修改信息
		dao.DisDeletePostByBarNameAndTitle(barName,postTitle)
		_,_ = w.Write([]byte("成功恢复帖子！"))
	} else {
		_,_ = w.Write([]byte("恢复失败！"))
	}

}

//GoToPost 去看帖子
func GoToPost(w http.ResponseWriter,r *http.Request){
	//查看是否登录了
	judge,sess := dao.IsLogin(r)
	//获取吧名帖子标题
	barName := r.FormValue("barName")
	postTitle := r.FormValue("postTitle")
	//根据吧名和帖子标题搜索帖子
	post := dao.FindCreatePostByBarNameAndTitle(barName,postTitle)
	//根据发帖者id查看发帖者项目
	owner := dao.FindUserByUserID(post.PostHostID)
	post.PostHostName = owner.UserName

	//获取帖子内容
	bytes,_ := ioutil.ReadFile(post.PostContent)

	post.PostContent = string(bytes)

	//根据吧名和帖子标题获取帖子的回复内容按时间降序排序
	post.Reply = dao.FindReplyByBarNameAndTitleOrderByTime(barName,postTitle)

	//帖子回复获取内容
	for k,v := range post.Reply{
		bytes,_ := ioutil.ReadFile(v.Content)

		v.Content = string(bytes)
		postUser := dao.FindUserByUserID(v.UserID)
		v.UserName = postUser.UserName
		//获得用户回复内容并获取内容
		v.UserReply = dao.FindUserReplyByPostReplyID(v.ID)
		for _,value := range v.UserReply{
			value.UserRank1 = strconv.FormatInt(int64(k)+1,10) + value.UserRank1
			value.UserRank2 = strconv.FormatInt(int64(k)+1,10) + value.UserRank2
			userUser := dao.FindUserByUserID(value.UserID)
			value.UserName = userUser.UserName
			userBytes,_ := ioutil.ReadFile(value.Content)
			value.Content = string(userBytes)
			v.TotalUserReply++
		}
		v.TotalUserReplyRank = "id" + strconv.FormatInt(int64(k)+1,10)
	}

	//若用户已经登录
	if judge == true {
		user := dao.FindUserByUserID(sess.UserID)

		//判断帖子是否点赞了
		post.IsThumb = dao.IsPostThumbByBarNameAndUserIDAndKind(barName,user.UserID,postTitle)

		//判断帖子是否收藏了
		post.IsLiked = dao.IsPostLikedByBarNameAndUserID(barName,user.UserID,postTitle)

		//将访问帖子信息存储到历史信息表中
		dao.AddPostHistory(barName,postTitle,user.UserID)
	}

	t := template.Must(template.ParseFiles("views/pages/post/look_post.html"))
	_ = t.Execute(w,post)
}

//FindPosts 模糊搜索帖子
func FindPosts(w http.ResponseWriter,r *http.Request){
	barName := r.PostFormValue("bar")
	find := r.PostFormValue("find")
	//模糊搜索帖子
	posts := dao.FindPostsByBarNameAndFind(barName,"%"+find+"%")
	page := &model.Page{
		Posts: posts,
		BarNow: barName,
		Find: find,
	}
	t := template.Must(template.ParseFiles("views/pages/post/find_posts.html"))
	_ = t.Execute(w,page)
}

//PostsOrderByTimeInFindPosts 在搜索帖子内按时间排序
func PostsOrderByTimeInFindPosts (w http.ResponseWriter,r *http.Request){
	//获取吧名和搜索的信息
	barName := r.PostFormValue("orderBarName")
	find := r.PostFormValue("orderFind")
	//调用数据库根据吧名和搜索信息按时间排序搜索帖子
	posts := dao.FindPostsByBarNameAndFindOrderByTime(barName,"%"+find+"%")
	page := &model.Page{
		Posts: posts,
		BarNow:barName,
		Find: find,
	}
	t := template.Must(template.ParseFiles("views/pages/post/find_posts.html"))
	_ = t.Execute(w,page)
}

//GetPostReply 写帖子回复内容
func GetPostReply(w http.ResponseWriter,r *http.Request){
	//是否登录
	judge,sess  := dao.IsLogin(r)

	//若未登录
	if judge == false {
		_,_ = w.Write([]byte("请先登录！"))
		return
	} else {
		//获取贴吧名当前用户内容标题
		barName := r.PostFormValue("barName")
		content := r.PostFormValue("content")
		title := r.PostFormValue("title")
		user := dao.FindUserByUserID(sess.UserID)

		//先将信息存储到数据库中并获取该帖子回复id
		id := dao.AddPostReply(barName,title,user.UserID)

		//将内容放进文件夹中
		postPath := "views/static/postReply/" + barName + title + "id" + strconv.FormatInt(id,10) + ".txt"
		file,_ := os.OpenFile(postPath ,os.O_CREATE | os.O_APPEND,0666)
		defer file.Close()
		file.WriteString(content)

		//修改数据库中该帖子回复的内容路径
		dao.SetContentByID(id,postPath)


		//修改帖子的回复量
		post := dao.FindCreatePostByBarNameAndTitle(barName,title)
		post.TotalReply++
		dao.SetTotalPostReply(barName,title,post.TotalReply)

		_, _ = w.Write([]byte("回复成功！"))
	}
}

//ToReplyUserInPostPage 去写帖子回复内容的回复内容页面
func ToReplyUserInPostPage(w http.ResponseWriter,r *http.Request){
	//获取帖子id
	postID := r.FormValue("postReplyID")

	//去页面
	t := template.Must(template.ParseFiles("views/pages/post/write_reply_user.html"))
	_ = t.Execute(w,postID)
}

//GetUserReply 获取回复他人回复内容
func GetUserReply(w http.ResponseWriter,r *http.Request){
	//是否登录
	judge,sess := dao.IsLogin(r)


	//若未登录
	if judge == false {
		_,_ = w.Write([]byte("请先登录！"))
	} else {
		content := r.PostFormValue("content")
		postID := r.PostFormValue("postID")
		//获取用户
		user := dao.FindUserByUserID(sess.UserID)

		PostID,_ := strconv.ParseInt(postID,10,64)
		//将内容存进数据库中并获取id
		replyID := dao.AddUserReply(PostID,user.UserID)

		//将内容放进文件夹中
		replyPath := "views/static/userReply/" +  "id" + strconv.FormatInt(replyID,10) + ".txt"
		file,_ := os.OpenFile(replyPath ,os.O_CREATE | os.O_APPEND,0666)
		defer file.Close()
		file.WriteString(content)

		//修改数据库中该帖子回复的回复的内容路径
		dao.SetReplyContentByReplyID(replyID,replyPath)
		_, _ = w.Write([]byte("回复成功！"))
	}

}

//PostIsThumb 点赞功能
func PostIsThumb(w http.ResponseWriter,r *http.Request){
	//是否登录
	judge,sess := dao.IsLogin(r)

	//若未登录
	if judge == false {
		_,_ = w.Write([]byte("请先登录！"))
	} else {
		//获取吧名帖子名当前用户与是否点赞
		barName := r.PostFormValue("barName")
		title := r.PostFormValue("title")
		isThumb := r.PostFormValue("judge")
		user := dao.FindUserByUserID(sess.UserID)
		if isThumb == "1" {
			//要点赞
			//增加一个点赞的项
			dao.ThumbPost(barName,user.UserID,title)

			_,_ = w.Write([]byte("成功点赞！"))

		} else {
			//取消点赞
			//将原先点赞的项删除
			dao.DisThumbPost(barName,user.UserID,title)

			_,_ = w.Write([]byte("已取消点赞！"))

		}
	}
}

//PostIsLiked 收藏功能
func PostIsLiked(w http.ResponseWriter,r *http.Request){
	//是否登录
	judge,sess := dao.IsLogin(r)
	//若未登录
	if judge == false {
		_,_ = w.Write([]byte("请先登录！"))
	} else {
		//获取吧名帖子名用户和是否收藏
		barName := r.PostFormValue("barName")
		title := r.PostFormValue("title")
		user := dao.FindUserByUserID(sess.UserID)
		kind := r.PostFormValue("kind")

		if kind == "liked" {
			//收藏
			//添加一个收藏的项
			dao.LikedPost(barName,user.UserID,title)
			_,_ = w.Write([]byte("成功收藏！"))

		} else {
			//取消收藏
			//将原先收藏的项删除
			dao.DisLikedPost(barName,user.UserID,title)
			_,_ = w.Write([]byte("已取消收藏！"))
		}
	}
}

//DeletePostHistory 删除访问帖子记录
func DeletePostHistory(w http.ResponseWriter,r *http.Request){
	//获取当前用户
	_,sess := dao.IsLogin(r)
	barName := r.FormValue("barName")
	title := r.FormValue("title")
	time := r.FormValue("time")
	dao.DeletePostHistory(barName,title,sess.UserID,time)
	LookHistory(w,r)
}

//ApplicationToRecoverPost 申请恢复帖子
func ApplicationToRecoverPost(w http.ResponseWriter,r *http.Request){
	//获取吧名和帖子标题
	barName := r.PostFormValue("barName")
	title := r.PostFormValue("title")
	//获取当前用户
	_,sess := dao.IsLogin(r)

	//查看是否已经发送过
	id := dao.FindPostApplicationToRecoverMessageByBarNameAndTitle(barName,title)

	if id <= 0 {
		//若未发送过
		//获取该吧吧主id
		receiverID := dao.FindBarOwnerIDByBarNameAndKind(barName)

		//向吧主发送申请恢复
		dao.AddPostApplicationToRecoverMessage(barName,title,sess.UserID,receiverID)

	}
	_,_ = w.Write([]byte("已成功发送申请！"))

}

//DisAgreeToRecoverPost 拒绝申请恢复帖子
func DisAgreeToRecoverPost(w http.ResponseWriter,r *http.Request){
	//获取吧名和帖子标题
	barName := r.PostFormValue("barName")
	title := r.PostFormValue("title")
	//将该信息删除
	dao.DeletePostApplicationToRecoverMessage(barName,title)
	_,_ = w.Write([]byte("已拒绝申请恢复帖子！"))
}

//AgreeToRecoverPost 同意申请恢复帖子
func AgreeToRecoverPost(w http.ResponseWriter,r *http.Request){
	//获取吧名和帖子标题
	barName := r.PostFormValue("barName")
	title := r.PostFormValue("title")
	//查看该贴吧是否被封禁
	status := dao.FindBarStatusByBarNameAndKind(barName)
	if status == true {
		//该贴吧未被封禁
		//设置该帖子信息
		dao.DisDeletePostByBarNameAndTitle(barName,title)
		//将该信息删除
		dao.DeletePostApplicationToRecoverMessage(barName,title)
		_,_ = w.Write([]byte("成功恢复帖子！"))
	} else {
		//将该信息删除
		dao.DeletePostApplicationToRecoverMessage(barName,title)
		_,_ = w.Write([]byte("该帖子的贴吧已被封禁，恢复失败！"))
	}
}

//ReportPost 举报帖子
func ReportPost(w http.ResponseWriter,r *http.Request){
	//获取吧名和标题
	barName := r.PostFormValue("barName")
	title := r.PostFormValue("title")
	//查看该帖子是否被举报过
	id := dao.FindReportPostByBarNameAndTitle(barName,title)
	if id <= 0 {
		//还未被举报过
		dao.AddReportPostMessage(barName,title)
	}
	_,_ = w.Write([]byte("举报成功！"))
}

//BannedPost 封禁帖子
func BannedPost(w http.ResponseWriter,r *http.Request){
	//获取帖子和吧名
	title := r.FormValue("title")
	barName := r.FormValue("barName")
	//调用数据库将该帖子状态设置为0
	dao.DeletePostFake(barName,title,0)
	GetReportMessage(w,r)
}

//PostsOrderBy 在吧内对帖子进行排序
func PostsOrderBy(w http.ResponseWriter,r *http.Request){
	barName := r.FormValue("barName")
	kind := r.FormValue("kind")
	var posts []*model.Post
	 if kind == "time" {
		 posts = dao.FindPostsByBarNameOrderByTime(barName)
	 } else {
	 	posts = dao.FindPostsByBarNameOrderByPostReply(barName)
	 }
	//查看是否登录
	judge,sess := dao.IsLogin(r)
	//调用数据库根据吧名查看吧的信息
	bar := dao.FindBarByBarNameAndKind(barName)
	//获取吧主信息
	barOwner := dao.FindUserByUserID(bar.BarHostID)
	bar.BarHostName = barOwner.UserName
	//获取该吧的公告
	bar.MyNotice = dao.FindNoticesByBarName(bar.BarName)
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
