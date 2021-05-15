package model

type Message struct {
	ID int64
	//接收者
	ReceiverID int64
	//接收者
	ReceiverName string
	//发送者的id
	SenderID int64
	//发送者
	SenderName string
	BarName string
	PostTitle string
	Status bool
	UserStatus bool
	BarStatus bool
	PostStatus bool
}

//ReportMessages 举报的所有消息
type ReportMessages struct {
	BarMessages []*Message
	PostMessages []*Message
}
