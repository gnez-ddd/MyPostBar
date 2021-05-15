package model

type UserReply struct {
	//回复ID
	UserReplyID int64
	//回复的帖子回复ID
	PostReplyID int64
	//回复内容
	Content string
	//用户名
	UserID int64
	UserName string
	//回复的排序
	UserRank1 string
	UserRank2 string
	//发表时间
	Time string
}
