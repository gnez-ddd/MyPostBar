package dao

import (
	"MyPostBar/ddsPostBar/model"
	"MyPostBar/ddsPostBar/utils"
)

//AddPostHistory 增加帖子访问记录
func AddPostHistory(barName string,title string,id int64){
	sqlStr := "insert into lookHistory (bar_name,post_title,user_id,time) values (?,?,?,now());"
	_,_ = utils.Db.Exec(sqlStr,barName,title,id)
}

//AddBarHistory 增加贴吧访问记录
func AddBarHistory(barName string,id int64){
	sqlStr := "insert into lookHistory (bar_name,user_id,time) values (?,?,now());"
	_,_ = utils.Db.Exec(sqlStr,barName,id)
}

//DeletePostHistory 删除帖子访问记录
func DeletePostHistory(barName string,title string,id int64,time string){
	sqlStr := "delete from lookHistory where bar_name = ? and post_title = ? and user_id = ? and time = ? ;"
	_,_ = utils.Db.Exec(sqlStr,barName,title,id,time)
}

//DeleteBarHistory 删除贴吧访问记录
func DeleteBarHistory(barName string,id int64,time string){
	sqlStr := "delete from lookHistory where bar_name = ? and post_title = ? and user_id = ? and time = ? ;"
	_,_ = utils.Db.Exec(sqlStr,barName,"nil",id,time)
}

//FindHistoryBarsByUserID 根据用户id查找其访问过的贴吧按时间降序
func FindHistoryBarsByUserID(id int64)[]*model.Bar{
	sqlStr := "select bar_name,time from lookHistory where user_id = ? and post_title = ? order by time desc ;"
	rows,_ := utils.Db.Query(sqlStr,id,"nil")
	var bars []*model.Bar
	for rows.Next(){
		bar := &model.Bar{}
		_ = rows.Scan(&bar.BarName,&bar.Time)
		bars = append(bars,bar)
	}
	return bars
}

//FindHistoryPostsByUserID 根据用户id查找其访问过的帖子
func FindHistoryPostsByUserID(id int64)[]*model.Post{
	sqlStr := "select bar_name,post_title,time from lookHistory where user_id = ? and post_title != ? order by time desc ;"
	rows,_ := utils.Db.Query(sqlStr,id,"nil")
	var posts []*model.Post
	for rows.Next(){
		post := &model.Post{}
		_ = rows.Scan(&post.BarName,&post.PostTitle,&post.PostDate)
		posts = append(posts,post)
	}
	return posts
}