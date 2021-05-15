package dao

import (
	"MyPostBar/ddsPostBar/model"
	"MyPostBar/ddsPostBar/utils"
	"strconv"
)

//AddUserReply 增加回复他人回复内容并获取id
func AddUserReply(postReplyID int64,userID int64)int64{
	//写sql语句
	sqlStr := "insert into userReply (post_reply_id,user_id,time) values (?,?,now());"
	ret,_ := utils.Db.Exec(sqlStr,postReplyID,userID)
	id,_ := ret.LastInsertId()
	return id
}

func SetReplyContentByReplyID(id int64,content string){
	//写sql语句
	sqlStr := "update userReply set content = ? where id = ? ;"
	_,_ = utils.Db.Exec(sqlStr,content,id)
}

//FindUserReplyByPostReplyID 根据帖子回复id获取用户回复按时间降序
func FindUserReplyByPostReplyID(postID int64)[]*model.UserReply{
	//写sql语句
	sqlStr := "select id,post_reply_id,user_id,content,time from userReply where post_reply_id = ? order by time desc ; "
	rows,_ := utils.Db.Query(sqlStr,postID)
	var userReplys []*model.UserReply
	var i int64 = 0
	for rows.Next(){
		userReply := &model.UserReply{}
		_ = rows.Scan(&userReply.UserReplyID,&userReply.PostReplyID,&userReply.UserID,&userReply.Content,&userReply.Time)

		userReply.UserRank1 = "userReplyID" + strconv.FormatInt(i,10)
		userReply.UserRank2 = "userReply" + strconv.FormatInt(i,10)
		i++

		userReplys = append(userReplys,userReply)
	}
	return userReplys
}

//DeleteUserReplyByID 根据id将用户回复删除
func DeleteUserReplyByID(id int64){
	sqlStr := "delete from userReply where id = ? ;"
	_,_ = utils.Db.Exec(sqlStr,id)
}
