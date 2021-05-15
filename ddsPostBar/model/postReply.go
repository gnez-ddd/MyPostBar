package model

type PostReply struct {
	//回复ID
	ID int64
	//所在吧名
	BarName string
	//所在帖子标题
	Title string
	//回复内容
	Content string
	//用户名
	UserName string
	//用户id
	UserID int64
	//回复的排序
	Rank1 string
	Rank2 string
	//发表时间
	Time string
	//他人的回复
	UserReply []*UserReply
	//他人回复的总数
	TotalUserReply int64
	TotalUserReplyRank string
}
