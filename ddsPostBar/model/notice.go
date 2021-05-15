package model

type Notice struct {
	//属于的吧名
	BarName string
	//内容
	Content string
	//发布日期
	Time string
	//当前用户是否为吧主
	IsBarOwner bool
}
