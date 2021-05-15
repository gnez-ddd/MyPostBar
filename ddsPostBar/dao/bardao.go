package dao

import (
	"MyPostBar/ddsPostBar/model"
	"MyPostBar/ddsPostBar/utils"
)

/*
create table bars
(
    id         int auto_increment
        primary key,
    total_post int        default 0 not null,
    total_fan  int        default 0 not null,
    host_id    int                  not null,
    bar_name   varchar(100)         not null,
    status     tinyint(1) default 0 not null,
    constraint bars_ibfk_1
        foreign key (host_id) references users (id)
);
 */

// AddBar 增加贴吧
func AddBar(userID int64,barName string,kind string){
	//写sql语句
	sqlStr := "insert bars(host_id,bar_name,kind) values (?,?,?) ;"
	//执行
	_,_ = utils.Db.Exec(sqlStr,userID,barName,kind)
}

//FindBarByBarNameAndKind 根据吧名和创建类型查找贴吧
func FindBarByBarNameAndKind(barName string)*model.Bar{
	//写sql语句
	sqlStr := "select id,total_post,total_fan,host_id,bar_name from bars where bar_name = ? and kind = ? ;"
	//执行
	row := utils.Db.QueryRow(sqlStr,barName,"create")
	//创建一个吧结构体用于存储
	bar := &model.Bar{}
	_ = row.Scan(&bar.BarID,&bar.TotalPost,&bar.TotalFan,&bar.BarHostID,&bar.BarName)
	return bar
}

//DeleteBarByBarName 根据吧名删除贴吧
func DeleteBarByBarName(barName string){
	//写sql语句
	sqlStr := "delete from bars where bar_name = ? ;"
	//执行
	_,_ = utils.Db.Exec(sqlStr,barName)
}

//SetBarStatusByBarName 根据吧名修改贴吧状态
func SetBarStatusByBarName(barName string,status int){
	//写sql语句
	sqlStr := "update bars set status = ? where bar_name = ? ;"
	//执行
	_,_ = utils.Db.Exec(sqlStr,status,barName)
}

//GetBarsByUserID 根据用户ID查看创建和关注的贴吧
func GetBarsByUserID(userID int64)*model.MyBar{
	//写sql语句查看用户的吧
	sqlStr := "select id,total_post,total_fan,host_id,bar_name from bars where host_id = ? and kind = ? and status = ?;"

	//执行查看用户创建的吧
	rows1,_ := utils.Db.Query(sqlStr,userID,"create",1)
	//创建一个吧切用于存储创建的贴吧
	var barCreate []*model.Bar
	for rows1.Next(){
		//设置一个变量接收每一个bar
		bar := &model.Bar{}
		_ = rows1.Scan(&bar.BarID,&bar.TotalPost,&bar.TotalFan,&bar.BarHostID,&bar.BarName)
		//将bar添加到切片中
		barCreate = append(barCreate,bar)
	}

	//查看用户关注的贴吧
	rows2,_ := utils.Db.Query(sqlStr,userID,"liked",1)
	//创建一个吧切片用于存储关注的贴吧
	var barLiked []*model.Bar
	for rows2.Next(){
		//设置一个变量接收每一个bar
		bar := &model.Bar{}
		_ = rows2.Scan(&bar.BarID,&bar.TotalPost,&bar.TotalFan,&bar.BarHostID,&bar.BarName)
		//将bar添加到切片中
		barLiked = append(barLiked,bar)
	}

	mybar := &model.MyBar{
		BarCreate: barCreate,
		BarLiked: barLiked,
	}
	return mybar
}

//GetAllBars 获取所有创建的贴吧
func GetAllBars()[]*model.Bar{
	//写sql语句
	sqlStr := "select id,total_post,total_fan,host_id,bar_name from bars where kind = ? and status = ?;"
	//执行查询语句
	rows,_ := utils.Db.Query(sqlStr,"create",1)
	//创建一个切片用于存储
	var bars []*model.Bar
	for rows.Next(){
		//设置一个变量接收每一个bar
		bar := &model.Bar{}
		_ = rows.Scan(&bar.BarID,&bar.TotalPost,&bar.TotalFan,&bar.BarHostID,&bar.BarName)
		//将bar添加到切片中
		bars = append(bars,bar)
	}
	return bars
}

//FindBarOwnerIDByBarNameAndKind 根据吧名和类型查找吧主id
func FindBarOwnerIDByBarNameAndKind(barName string)(id int64){
	//写sql语句
	sqlStr := "select host_id from bars where bar_name = ? and kind = ? ;"
	//执行
	row := utils.Db.QueryRow(sqlStr,barName,"create")
	_ = row.Scan(&id)
	return
}

//IsLikedByBarNameAndUserID 根据吧名和用户id查看是否已经关注
func IsLikedByBarNameAndUserID(barName string,id int64)bool{
	//写sql语句
	sqlStr := "select id from bars where bar_name = ? and host_id = ? and kind = ? and status = 1;"
	//执行
	row := utils.Db.QueryRow(sqlStr,barName,id,"liked")
	var barID int64
	_ = row.Scan(&barID)
	if barID > 0 {
		return true
	} else {
		return false
	}
}

//LikedBar 关注贴吧
func LikedBar(bar *model.Bar){
	//写sql语句
	sqlStr := "insert into bars (total_post,total_fan,host_id,bar_name,status,kind) values (?,?,?,?,?,?);"
	//执行
	_,_ = utils.Db.Exec(sqlStr,bar.TotalPost,bar.TotalFan,bar.BarHostID,bar.BarName,1,bar.Kind)

}

//SetTotalFanByBarName 修改贴吧的粉丝信息
func SetTotalFanByBarName(totalFan int64,barName string){
	sqlStr := "update bars set total_fan = ? where bar_name = ? ;"
	_,_ = utils.Db.Exec(sqlStr,totalFan,barName)
}


//DisLikedBar 取消关注贴吧
func DisLikedBar(barName string,hostID int64){
	//写sql语句
	sqlStr := "delete from bars where bar_name = ? and host_id = ? and kind = ? ;"
	_,_ = utils.Db.Exec(sqlStr,barName,hostID,"liked")
}

//FindBars 模糊搜索贴吧
func FindBars(find string)[]*model.Bar{
	//写sql语句
	sqlStr := "select bar_name from bars where kind = ? and status = 1 and bar_name like ? ;"
	rows,_ := utils.Db.Query(sqlStr,"create",find)
	var bars []*model.Bar
	for rows.Next(){
		bar := &model.Bar{}
		_ = rows.Scan(&bar.BarName)
		bars = append(bars,bar)
	}
	return bars
}

//FindBarStatusByBarNameAndKind 根据吧名查找吧的状态
func FindBarStatusByBarNameAndKind(barName string)(status bool){
	sqlStr := "select status from bars where bar_name = ? and kind = ?;"
	row := utils.Db.QueryRow(sqlStr,barName,"create")
	var judge int64
	_ = row.Scan(&judge)
	if judge == 1 {
		status = true
	} else {
		status = false
	}
	return
}

