package model

type Page struct {
	//是否登录
	IsLogin bool
	//用户
	User *Users
	//帖子
	Posts []*Post
	//贴吧
	Bar []*Bar
	//当前贴吧
	BarNow string
	//搜索的信息
	Find string
	//是不是为系统管理员
	IsAdmin bool
}
