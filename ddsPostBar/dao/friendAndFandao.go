package dao

import (
	"MyPostBar/ddsPostBar/model"
	"MyPostBar/ddsPostBar/utils"
)

//FindIsFriend 查看是否为好友
func FindIsFriend (userID1 int64,userID2 int64) bool{
	sqlStr := "select id from friend_fan where sender_id = ? and receiver_id = ? and kind = ? and status = 1;"
	row := utils.Db.QueryRow(sqlStr,userID1,userID2,"friend")
	var id int64
	_ = row.Scan(&id)
	if id > 0 {
		return true
	}
	row = utils.Db.QueryRow(sqlStr,userID2,userID1,"friend")
	_ = row.Scan(&id)
	if id > 0 {
		return true
	}
	return false
}

//FindIsLiked 查看是否关注了
func FindIsLiked(fanID int64,starID int64)bool{
	sqlStr := "select id from friend_fan where sender_id = ? and receiver_id = ? and kind = ?;"
	row := utils.Db.QueryRow(sqlStr,fanID,starID,"liked")
	var id int64
	_ = row.Scan(&id)
	if id > 0 {
		return true
	}
	return false
}

//IsAskToMakeFriend 是否已经发送过申请
func IsAskToMakeFriend(senderID int64,receiverID int64)bool{
	sqlStr := "select id from friend_fan where sender_id = ? and receiver_id = ? and kind = ?;"
	row := utils.Db.QueryRow(sqlStr,senderID,receiverID,"friend")
	var id int64
	_ = row.Scan(&id)
	if id > 0 {
		return true
	} else {
		return false
	}
}

//AddMakeFriend 增加好友申请
func AddMakeFriend(senderID int64,receiverID int64){
	sqlStr := "insert into friend_fan (sender_id,receiver_id,kind) values (?,?,?);"
	_,_ =utils.Db.Exec(sqlStr,senderID,receiverID,"friend")
}

//FindMakeFriendUserIDByReceiverID 根据接收者id获取其收到的申请成为好友的用户id
func FindMakeFriendUserIDByReceiverID(id int64)[]*model.Users{
	sqlStr := "select sender_id from friend_fan where receiver_id = ? and kind = ? and status = 0;"
	rows,_ := utils.Db.Query(sqlStr,id,"friend")
	var users []*model.Users
	for rows.Next(){
		user := &model.Users{}
		_ = rows.Scan(&user.UserID)
		users = append (users,user)
	}
	return users
}

//AgreeToMakeFriend 同意成为好友
func AgreeToMakeFriend(senderID int64,receiverID int64){
	sqlStr := "update friend_fan set status = ? where sender_id = ? and receiver_id = ? and kind = ?;"
	_,_ = utils.Db.Exec(sqlStr,1,senderID,receiverID,"friend")
}

//DeleteMakeFriend 删除添加好友
func DeleteMakeFriend(senderID int64,receiverID int64){
	sqlStr := "delete from friend_fan where sender_id = ? and receiver_id = ? and kind = ?;"
	_,_ = utils.Db.Exec(sqlStr,senderID,receiverID,"friend")
}

//GetFriendByUserID 根据用户id查找好友
func GetFriendByUserID(id int64)[]*model.Users{
	sqlStr := "select sender_id from friend_fan where receiver_id = ? and kind = ?;"
	rows,_ := utils.Db.Query(sqlStr,id,"friend")
	var users []*model.Users
	for rows.Next(){
		user := &model.Users{}
		_ = rows.Scan(&user.UserID)
		users = append(users,user)
	}
	sqlStr = "select receiver_id from friend_fan where sender_id = ? and kind = ?;"
	rows,_ = utils.Db.Query(sqlStr,id,"friend")
	for rows.Next(){
		user := &model.Users{}
		_ = rows.Scan(&user.UserID)
		users = append(users,user)
	}
	return users
}

//DeleteFriend 删除好友
func DeleteFriend(senderID int64,receiverID int64){
	sqlStr := "delete from friend_fan where sender_id = ? and receiver_id = ? and kind = ?;"
	_,_ = utils.Db.Exec(sqlStr,senderID,receiverID,"friend")
	_,_ = utils.Db.Exec(sqlStr,receiverID,senderID,"friend")
}

//AddLikedUser 关注用户
func AddLikedUser(receiverID int64,senderID int64){
	sqlStr := "insert into friend_fan (sender_id,receiver_id,kind) values (?,?,?);"
	_,_ = utils.Db.Exec(sqlStr,senderID,receiverID,"liked")
}

//DeleteLikedUser 取消关注用户
func DeleteLikedUser(receiverID int64,senderID int64){
	sqlStr := "delete from friend_fan where receiver_id = ? and sender_id = ? and kind = ?;"
	_,_ = utils.Db.Exec(sqlStr,receiverID,senderID,"liked")
}

//FindLikedUserBySenderID 根据用户id查看其关注的用户
func FindLikedUserBySenderID(id int64)[]*model.Users{
	sqlStr := "select receiver_id from friend_fan where sender_id = ? and kind = ?;"
	rows,_ := utils.Db.Query(sqlStr,id,"liked")
	var users []*model.Users
	for rows.Next(){
		user := &model.Users{}
		_ = rows.Scan(&user.UserID)
		users = append(users,user)
	}
	return users
}

//FindMyFan 查看我的粉丝
func FindMyFan(id int64)[]*model.Users{
	sqlStr := "select sender_id from friend_fan where receiver_id = ? and kind = ?;"
	rows,_ := utils.Db.Query(sqlStr,id,"liked")
	var users []*model.Users
	for rows.Next(){
		user := &model.Users{}
		_ = rows.Scan(&user.UserID)
		users = append(users,user)
	}
	return users
}