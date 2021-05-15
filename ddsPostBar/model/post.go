package model

// Post 创建帖子结构体
type Post struct{
	PostID int64
	//帖子标题
	PostTitle string
	//当前用户
	PostHostName string
	//当前用户id
	PostHostID int64
	//帖子所属的贴吧
	BarName string
	//帖子内容
	PostContent string
	//帖子发布时间
	PostDate string
	//帖子对于当前用户的种类
	Kind string
	//帖子对于当前用户的状态
	status bool
	//帖子的回复
	Reply []*PostReply
	//帖子回复的总数
	TotalReply int64
	//是否收藏
	IsLiked bool
	//是否点赞
	IsThumb bool
	//当前用户是否为吧主
	IsBarOwner bool
}

type MyPost struct {
	//我创建的帖子
	PostCreate []*Post
	//我关注的帖子
	PostLiked []*Post
	//删除的帖子
	PostDelete []*Post
	//被删除的帖子
	PostDeleted []*Post
}