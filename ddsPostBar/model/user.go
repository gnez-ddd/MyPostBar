package model

// Users 创建用户结构体
type Users struct{
	UserID int64
	UserName string
	PassWord string
	Email string
	HeadPath string
	Salt string
	//是否登录
	IsLogin bool
	//上一次签到时间
	LastSign string
	//经验值
	Experience int64
	//今日是否签到过
	IsSignIn bool
	Status int64
}