package model

type Bar struct{
	//贴吧id
	BarID int64
	//贴吧名
	BarName string
	//当前用户id
	BarHostID int64
	//当前用户姓名
	BarHostName string
	//帖子
	MyBarPosts []*Post
	//粉丝
	MyFans []*Users
	//公告
	MyNotice []*Notice
	//帖子数量
	TotalPost int64
	//粉丝数量
	TotalFan int64
	//是否存在状态
	Status bool
	//对于当前用户该吧的种类
	Kind string
	//是否关注
	IsLiked bool
	Time string
	//当前用户是否为吧主
	IsBarOwner bool

}

type MyBar struct {
	//我创建的贴吧
	BarCreate []*Bar
	//我关注的贴吧
	BarLiked []*Bar
}

//GetTotalPost 获取贴吧中帖子总数量
func (bar *Bar)GetTotalPost()int64{
	var totalPost int64
	//遍历帖子切片
	for _,_ = range bar.MyBarPosts {
		totalPost++
	}
	return totalPost
}

//GetTotalFans 获取贴吧中粉丝总数量
func (bar *Bar)GetTotalFans()int64{
	var totalFan int64
	for _,_ = range bar.MyFans {
		totalFan++
	}
	return totalFan
}


