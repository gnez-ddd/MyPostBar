package dao

import (
	"MyPostBar/ddsPostBar/model"
	"MyPostBar/ddsPostBar/utils"
)

//AddNotice 添加公告
func AddNotice(barName string,content string){
	//写sql语句
	sqlStr := "insert into notices (bar_name,content,time) values (?,?,now()) ;"
	//执行
	_,_ = utils.Db.Exec(sqlStr,barName,content)
}

//FindNoticesByBarName 根据吧名查看公告按时间排序
func FindNoticesByBarName(barName string)[]*model.Notice{
	//写sql语句
	sqlStr := "select * from notices where bar_name = ? order by time desc ;"
	//执行查找语句
	rows,_ := utils.Db.Query(sqlStr,barName)
	//创建一个切片用于存储
	var notices []*model.Notice
	for rows.Next(){
		//设置一个变量接收每一天公告
		notice := &model.Notice{}
		_ = rows.Scan(&notice.BarName,&notice.Content,&notice.Time)
		//将公告添加到切片中
		notices = append(notices,notice)
	}
	return notices
}

//DeleteNoticeByBarNameAndTime 根据吧名和时间删除公告
func DeleteNoticeByBarNameAndTime(barName string,time string){
	//写sql语句
	sqlStr := "delete from notices where bar_name = ? and time = ?;"
	//执行
	_,_ = utils.Db.Exec(sqlStr,barName,time)
}