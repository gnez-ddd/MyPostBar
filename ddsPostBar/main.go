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

	//是否同意创建贴吧
	http.HandleFunc("/isAgreeToCreateBar",controller.IsAgreeToCreateBar)

	//是否关注贴吧
	http.HandleFunc("/isLikedBar",controller.IsLikedBar)


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
	http.HandleFunc("/getUserReply",controller.GetUserReply)

	//点赞帖子
	http.HandleFunc("/PostIsThumb",controller.PostIsThumb)

	//收藏帖子
	http.HandleFunc("/PostIsLiked",controller.PostIsLiked)

	//获取历史记录
	http.HandleFunc("/lookHistory",controller.LookHistory)

	//删除贴吧或帖子访问记录
	http.HandleFunc("/deleteBarOrPostHistory",controller.DeleteBarOrPostHistory)


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

	//举报贴吧或帖子
	http.HandleFunc("/reportBarOrPost",controller.ReportBarOrPost)



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

	//查看用户
	http.HandleFunc("/look_user",controller.LookUser)

	//申请添加好友
	http.HandleFunc("/aksToMakeFriend",controller.AskToMakeFriend)

	//获取申请成为好友消息
	http.HandleFunc("/getMakeFriendMessage",controller.GetMakeFriendMessage)

	//是否同意成为好友
	http.HandleFunc("/isAgreeToMakeFriend",controller.IsAgreeToMakeFriend)

	//查看我的好友
	http.HandleFunc("/getMyFriend",controller.GetMyFriend)

	//删除好友
	http.HandleFunc("/deleteFriend",controller.DeleteFriend)

	//关注用户
	http.HandleFunc("/likedUser",controller.LikedUser)

	//取消关注
	http.HandleFunc("/disLikedUser",controller.DisLikedUser)

	//查看我关注的用户
	http.HandleFunc("/getMyLikedUser",controller.GetMyLikedUser)

	//查看我的粉丝
	http.HandleFunc("/getMyFan",controller.GetMyFan)

	//通过Ajax判断吧名是否可用
	http.HandleFunc("/checkBarName",controller.CheckBarName)

	//通过Ajax判断用户名或是否可用
	http.HandleFunc("/checkUserNameOrEmail",controller.CheckUserNameOrEmail)

	//通过Ajax判断验证码是否正确
	http.HandleFunc("/checkCode",controller.CheckCode)

	//通过发送验证码
	http.HandleFunc("/sendEmail",controller.SendEmail)


	_ = http.ListenAndServe(":8080", nil)

}
