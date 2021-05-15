package dao

import (
	"MyPostBar/ddsPostBar/model"
	"MyPostBar/ddsPostBar/utils"
)

/*
create table messages
(
	id          int                  not null
        primary key,
    receiver_id int                  not null,
    sender_id   int                  not null,
    content     varchar(500)         not null,
    status      tinyint(1) default 0 not null,
    constraint massages_users__id_fk
        foreign key (receiver_id) references users (id),
    constraint massages_users_id_fk
        foreign key (sender_id) references users (id)
);
 */

//AddCreateBarMessage 增加申请创建贴吧信息
func AddCreateBarMessage(receiverID int64,senderID int64,barName string){
	//写sql语句
	sqlStr := "insert into createBarMessages (receiver_id,sender_id,bar_name) values (?,?,?);"
	//执行
	_,_ = utils.Db.Exec(sqlStr,receiverID,senderID,barName)
}

//GetCreateBarMessageByReceiverID 根据接收者id查看申请创建贴吧信息
func GetCreateBarMessageByReceiverID(id int64)([]*model.Message){
	//写sql语句
	sqlStr := "select * from createBarMessages where receiver_id = ?;"
	//执行
	rows,_ := utils.Db.Query(sqlStr,id)
	var createBarMessages []*model.Message
	for rows.Next(){
		//设置一个遍历接收每一条message
		createBarMessage := &model.Message{}
		_ = rows.Scan(&createBarMessage.ReceiverID,&createBarMessage.SenderID,&createBarMessage.BarName,&createBarMessage.ID)
		//获取信息发送者
		sender := FindUserByUserID(createBarMessage.SenderID)
		createBarMessage.SenderName = sender.UserName

		//将消息添加到切片中
		createBarMessages = append(createBarMessages,createBarMessage)
	}
	return createBarMessages
}


//DeleteCreateBarMessageByBarName 根据吧名删除申请创建贴吧消息
func DeleteCreateBarMessageByBarName(barName string){
	//写sql语句
	sqlStr := "delete from createBarMessages where bar_name = ?;"
	//执行
	_,_ = utils.Db.Exec(sqlStr,barName)
}

//AddPostApplicationToRecoverMessage 增加申请恢复帖子信息
func AddPostApplicationToRecoverMessage(barName string,title string,senderID int64,receiverID int64){
	sqlStr := "insert into post_recover_message (sender_id,receiver_id,barName,title) value (?,?,?,?);"
	_,_ = utils.Db.Exec(sqlStr,senderID,receiverID,barName,title)
}

//FindPostApplicationToRecoverMessage 查找用户所拥有的申请恢复帖子消息
func FindPostApplicationToRecoverMessage(receiverID int64)[]*model.Message{
	sqlStr := "select sender_id,receiver_id,barName,title from post_recover_message where receiver_id = ?;"
	rows,_ := utils.Db.Query(sqlStr,receiverID)
	var messages []*model.Message
	for rows.Next(){
		message := &model.Message{}
		_ = rows.Scan(&message.SenderID,&message.ReceiverID,&message.BarName,&message.PostTitle)
		messages = append(messages,message)
	}
	return messages
}

//FindPostApplicationToRecoverMessageByBarNameAndTitle 查看具体某个申请
func FindPostApplicationToRecoverMessageByBarNameAndTitle(barName string,title string)(id int64){
	sqlStr := "select id from post_recover_message where barName = ? and title = ? ;"
	row := utils.Db.QueryRow(sqlStr,barName,title)
	_ = row.Scan(&id)
	return
}

//DeletePostApplicationToRecoverMessage 删除申请帖子恢复信息
func DeletePostApplicationToRecoverMessage(barName string,title string){
	sqlStr := "delete from post_recover_message where barName = ? and title = ? ;"
	_,_ = utils.Db.Exec(sqlStr,barName,title)
}

//FindReportBarByBarName 根据吧名查看该吧是否被举报过
func FindReportBarByBarName(barName string)(id int64){
	sqlStr := "select id from report_message where barName = ? and post_title = ?;"
	row := utils.Db.QueryRow(sqlStr,barName,"nil")
	_ = row.Scan(&id)
	return
}

//FindReportPostByBarNameAndTitle 查看该帖子是否被举报过
func FindReportPostByBarNameAndTitle(barName string,title string)(id int64){
	sqlStr := "select id from report_message where barName = ? and post_title = ?;"
	row := utils.Db.QueryRow(sqlStr,barName,title)
	_ = row.Scan(&id)
	return
}

//AddReportBarMessage 增加举报贴吧信息
func AddReportBarMessage(barName string){
	sqlStr := "insert into report_message (barName,post_title) values (?,?);"
	_,_ = utils.Db.Exec(sqlStr,barName,"nil")
}

//AddReportPostMessage 增加帖子举报消息
func AddReportPostMessage(barName string,title string){
	sqlStr := "insert into report_message (barName,post_title) values (?,?);"
	_,_ = utils.Db.Exec(sqlStr,barName,title)
}

//FindAllBarReportMessage 获取贴吧举报消息
func FindAllBarReportMessage()(messages []*model.Message){
	sqlStr := "select barName from report_message where post_title = ?;"
	rows,_ := utils.Db.Query(sqlStr,"nil")
	for rows.Next(){
		message := &model.Message{}
		_ = rows.Scan(&message.BarName)
		messages = append(messages,message)
	}
	return
}

//FindAllPostReportMessage 获取帖子举报消息
func FindAllPostReportMessage()(messages []*model.Message){
	sqlStr := "select barName,post_title from report_message where post_title != ? ;"
	rows,_ := utils.Db.Query(sqlStr,"nil")
	for rows.Next(){
		message := &model.Message{}
		_ = rows.Scan(&message.BarName,&message.PostTitle)
		messages = append(messages,message)
	}
	return
}

//DeleteBarReport 删除贴吧举报消息
func DeleteBarReport(barName string){
	sqlStr := "delete from report_message where barName = ? and post_title = ?;"
	_,_ = utils.Db.Exec(sqlStr,barName,"nil")
}

//DeletePostReport 删除帖子举报消息
func DeletePostReport(barName string,title string){
	sqlStr := "delete from report_message where barName = ? and post_title = ?;"
	_,_ = utils.Db.Exec(sqlStr,barName,title)
}