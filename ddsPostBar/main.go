package main

import (
	"MyPostBar/ddsPostBar/controller"
	"net/http"
)


func main(){
	//设置梳理静态资源
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	//去首页
	http.HandleFunc("/index",controller.ToIndexPage)

	//去注册
	http.HandleFunc("/regist",controller.UserRegist)

	//去登录
	http.HandleFunc("/login",controller.UserLogin)

	//退出登录
	http.HandleFunc("/logout",controller.UserLogout)

	//去找回密码
	http.HandleFunc("/findPassword",controller.FindPassWord)

	//去用户身份验证页面
	http.HandleFunc("/toProveUserPage",controller.ToProveUserPage)

	//用户身份验证
	http.HandleFunc("/proveUser",controller.ProveUser)

	//用户信息修改
	http.HandleFunc("/setUserInformation",controller.SetUserInformation)

	//查看我的贴吧
	http.HandleFunc("/getMyBar",controller.GetMyBar)

	//进入贴吧
	http.HandleFunc("/goToBar",controller.GoToBar)

	//去发布公告页面
	http.HandleFunc("/toSendNoticePage",controller.ToSendNoticePage)

	//通过Ajax发布公共
	http.HandleFunc("/sendNotice",controller.SendNotice)

	//删除公告
	http.HandleFunc("/deleteNotice",controller.DeleteNotice)

	//创建贴吧
	http.HandleFunc("/barCreate",controller.BarCreate)

	//查看申请创建贴吧信息
	http.HandleFunc("/getCreateBarMessages",controller.GetCreateBarMessages)

	//同意创建贴吧
	http.HandleFunc("/agreeToCreateBar",controller.AgreeToCreateBar)

	//不同意创建贴吧
	http.HandleFunc("/disagreeToCreateBar",controller.DisagreeToCreateBar)

	//关注贴吧
	http.HandleFunc("/likedBar",controller.LikedBar)

	//取消关注贴吧
	http.HandleFunc("/disLikedBar",controller.DisLikedBar)

	//创建帖子
	http.HandleFunc("/postCreate",controller.PostCreate)

	//查看我的帖子
	http.HandleFunc("/getMyPosts",controller.GetMyPosts)

	//查看帖子
	http.HandleFunc("/goToPost",controller.GoToPost)

	//删除帖子
	http.HandleFunc("/deletePost",controller.DeletePostFake)

	//恢复帖子
	http.HandleFunc("/disDeletePost",controller.DisDeletePost)

	//申请恢复帖子
	http.HandleFunc("/applicationToRecoverPost",controller.ApplicationToRecoverPost)

	//模糊搜索贴吧和帖子
	http.HandleFunc("/findBarsAndPosts",controller.FindBarsAndPosts)

	//模糊搜索帖子
	http.HandleFunc("/findPosts",controller.FindPosts)

	//在吧内将帖子按进行排序
	http.HandleFunc("/postsOrderBy",controller.PostsOrderBy)

	//在搜索帖子内按时间由近到远将帖子排序
	http.HandleFunc("/postsOrderByTimeInFindPosts",controller.PostsOrderByTimeInFindPosts)

	//写帖子回复内容
	http.HandleFunc("/getPostReply",controller.GetPostReply)

	//去写帖子内回复他人回复内容页面
	http.HandleFunc("/toReplyUserInPostPage",controller.ToReplyUserInPostPage)

	//获取回复他人回复内容
	http.HandleFunc("/getUerReply",controller.GetUserReply)

	//点赞帖子
	http.HandleFunc("/PostIsThumb",controller.PostIsThumb)

	//收藏帖子
	http.HandleFunc("/PostIsLiked",controller.PostIsLiked)

	//获取历史记录
	http.HandleFunc("/lookHistory",controller.LookHistory)

	//删除贴吧访问记录
	http.HandleFunc("/deleteBarHistory",controller.DeleteBarHistory)

	//删除访问帖子记录
	http.HandleFunc("/deletePostHistory",controller.DeletePostHistory)

	//去查看信息页面
	http.HandleFunc("/toGetMessagesPage",controller.ToGetMessagesPage)

	//得到申请帖子恢复消息
	http.HandleFunc("/getPostApplicationToRecover",controller.GetPostApplicationToRecover)

	//不同意申请帖子恢复
	http.HandleFunc("/disagreeToRecoverPost",controller.DisAgreeToRecoverPost)

	//同意申请帖子恢复
	http.HandleFunc("/agreeToRecoverPost",controller.AgreeToRecoverPost)

	//去签到
	http.HandleFunc("/signIn",controller.SignIn)

	//举报贴吧
	http.HandleFunc("/reportBar",controller.ReportBar)

	//举报帖子
	http.HandleFunc("/reportPost",controller.ReportPost)

	//获取举报消息
	http.HandleFunc("/getReportMessage",controller.GetReportMessage)

	//封禁贴吧
	http.HandleFunc("/bannedBar",controller.BannedBar)

	//封禁帖子
	http.HandleFunc("/bannedPost",controller.BannedPost)

	//封禁用户
	http.HandleFunc("/bannedUser",controller.BannedUser)

	//删除举报消息
	http.HandleFunc("/deleteReport",controller.DeleteReport)


	//通过Ajax判断吧名是否可用
	http.HandleFunc("/checkBarName",controller.CheckBarName)

	//通过Ajax判断用户名是否可用
	http.HandleFunc("/checkUserName",controller.CheckUserName)

	//通过Ajax判断邮箱是否可用
	http.HandleFunc("/checkEmail",controller.CheckEmail)

	//通过Ajax判断验证码是否正确
	http.HandleFunc("/checkCode",controller.CheckCode)

	//根据用户名验证验证码是否正确
	http.HandleFunc("/checkCodeByUserName",controller.CheckCodeByUserName)

	//通过发送验证码
	http.HandleFunc("/sendEmail",controller.SendEmail)

	//根据用户名发送验证码
	http.HandleFunc("/sendEmailByUserName",controller.SendEmailByUserName)

	_ = http.ListenAndServe(":8080", nil)

}
