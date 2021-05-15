package dao

import (
	"MyPostBar/ddsPostBar/model"
	"MyPostBar/ddsPostBar/utils"
	"strconv"
)

//AddPostReply 增加帖子回复信息并返回回复id
func AddPostReply (barName string,title string,userID int64)int64{
	//写sql语句
	sqlStr := "insert into postReply (bar_name,title,user_id,time,content) values (?,?,?,now(),'unknown');"
	//执行插入语句
	ret,_ := utils.Db.Exec(sqlStr,barName,title,userID)
	//获取数据的id
	id,_ := ret.LastInsertId()
	return id
}

//SetContentByID 根据id修改帖子回复内容路径
func SetContentByID(id int64,content string){
	//写sql语句
	sqlStr := "update postReply set content = ? where id = ? ;"
	//执行
	_,_ = utils.Db.Exec(sqlStr,content,id)
}

//FindReplyByBarNameAndTitleOrderByTime 根据吧名和帖子标题获取帖子的回复内容按时间降序排序
func FindReplyByBarNameAndTitleOrderByTime(barName string,title string)[]*model.PostReply{
	//写sql语句
	sqlStr := "select id,bar_name,title,user_id,time,content from postReply where bar_name = ? and title = ? order by time desc ;"
	rows,_ := utils.Db.Query(sqlStr,barName,title)
	var replys []*model.PostReply
	var i int64 = 1
	for rows.Next(){
		reply := &model.PostReply{}
		_ = rows.Scan(&reply.ID,&reply.BarName,&reply.Title,&reply.UserID,&reply.Time,&reply.Content)

		reply.Rank1 = "postReplyID" + strconv.FormatInt(i,10)
		reply.Rank2 = "postReply" + strconv.FormatInt(i,10)
		i++

		replys = append(replys,reply)
	}
	return replys
}
