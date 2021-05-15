package dao

import (
	"MyPostBar/ddsPostBar/model"
	"MyPostBar/ddsPostBar/utils"
)

//AddPost 发帖
func AddPost(barName string,hostID int64,title string,content string){
	//写sql语句
	sqlStr := "insert into posts (bar_name,host_id,title,content,kind,time) values (?,?,?,?,?,now()) ;"
	//执行
	_,_ = utils.Db.Exec(sqlStr,barName,hostID,title,content,"create")
}

//FindCreatePostByBarNameAndTitle 通过吧名和标题查找帖子
func FindCreatePostByBarNameAndTitle(barName string,title string)*model.Post{
	//写sql语句
	sqlStr := "select id,host_id,bar_name,kind,title,time,content,totalpostReply from posts where bar_name = ? and title = ? and kind = ? and status = 1 ;"
	//执行查找语句
	row := utils.Db.QueryRow(sqlStr,barName,title,"create")
	post := &model.Post{}
	_ = row.Scan(&post.PostID,&post.PostHostID,&post.BarName,&post.Kind,&post.PostTitle,&post.PostDate,&post.PostContent,&post.TotalReply)
	return post
}

//FindOwnerIDByBarNameAndTitle 根据帖子吧名标题查找发布者
func FindOwnerIDByBarNameAndTitle(barName string,title string)(id int64){
	//写sql语句
	sqlStr := "select host_id from posts where bar_name = ? and title = ? and kind = ? ;"
	//执行
	row := utils.Db.QueryRow(sqlStr,barName,title,"create")
	_ = row.Scan(&id)
	return
}

//FindCreatePostsByBarName 根据吧名获取所有创建的帖子按用户活跃度排序
func FindCreatePostsByBarName(barName string)[]*model.Post{
	//写sql语句
	sqlStr := "select bar_name,title from posts,users where posts.bar_name = ? and posts.kind = ? and posts.status = 1 and posts.host_id = users.id order by users.experience desc ;"
	//执行查找语句
	rows,_ := utils.Db.Query(sqlStr,barName,"create")
	//建立切片
	var posts []*model.Post
	for rows.Next(){
		//创建一个变量用于存储
		post := &model.Post{}
		_ = rows.Scan(&post.BarName,&post.PostTitle)
		//添加到切片中
		posts = append(posts,post)
	}
	return posts
}

//FindPostsByUserIDAndKind 根据用户ID和种类查找帖子
func FindPostsByUserIDAndKind(id int64)*model.MyPost{
	//写sql语句
	sqlStr := "select id,host_id,bar_name,kind,title,time,content from posts where host_id = ? and status = 1;"
	//执行查找语句
	rows,_ := utils.Db.Query(sqlStr,id)
	var postsCreate []*model.Post
	var postsLiked []*model.Post
	for rows.Next(){
		//创建变量用于存储
		post := &model.Post{}
		_ = rows.Scan(&post.PostID,&post.PostHostID,&post.BarName,&post.Kind,&post.PostTitle,&post.PostDate,&post.PostContent)

		if post.Kind == "create"{
			postsCreate = append(postsCreate,post)
		} else {
			postsLiked = append(postsLiked,post)
		}
	}

	//查找当前用户删除的帖子
	sqlStr1 := "select id,host_id,bar_name,kind,title,time,content from posts where kind= ? and performer_id = ? and status = 0;"
	rows1,_ := utils.Db.Query(sqlStr1,"create",id)
	var deletePosts []*model.Post
	for rows1.Next(){
		post := &model.Post{}
		_ = rows1.Scan(&post.PostID,&post.PostHostID,&post.BarName,&post.Kind,&post.PostTitle,&post.PostDate,&post.PostContent)
		deletePosts = append(deletePosts,post)
	}

	//查找当前用户被删除的帖子
	sqlStr2 := "select bar_name,title from posts where kind = ? and host_id = ? and status = 0 ;"
	rows2,_ := utils.Db.Query(sqlStr2,"create",id)
	var postsDeleted []*model.Post
	for rows2.Next(){
		post := &model.Post{}
		_ = rows2.Scan(&post.BarName,&post.PostTitle)
		postsDeleted = append(postsDeleted,post)
	}


	myPosts := &model.MyPost{
		PostCreate: postsCreate,
		PostLiked: postsLiked,
		PostDelete: deletePosts,
		PostDeleted: postsDeleted,
	}
	return myPosts
}

//FindAllDeletePosts 查找所有被删除的帖子
func FindAllDeletePosts()[]*model.Post{
	sqlStr := "select bar_name,title from posts where kind = ?  and status = 0;"
	rows,_ := utils.Db.Query(sqlStr,"create")
	var deletePosts []*model.Post
	for rows.Next(){
		post := &model.Post{}
		_ = rows.Scan(&post.BarName,&post.PostTitle)
		deletePosts = append(deletePosts,post)
	}
	return deletePosts
}

//DisDeletePostByBarNameAndTitle 恢复帖子
func DisDeletePostByBarNameAndTitle(barName string,title string){
	sqlStr := "update posts set status = 1 where bar_name = ? and title = ? ;"
	_,_ = utils.Db.Exec(sqlStr,barName,title)
}

//DeletePostsByBarNameAndPostTitle 根据吧名和标题删除帖子
func DeletePostsByBarNameAndPostTitle(barName string,postTitle string){
	//写sql语句
	sqlStr := "delete from posts where bar_name = ? and title = ?;"
	_,_ = utils.Db.Exec(sqlStr,barName,postTitle)
}

//GetAllCreatePostOrderByExperience 获取所有创建的帖子
func GetAllCreatePostOrderByExperience()[]*model.Post{
	//写sql语句
	sqlStr := "select bar_name,title from posts,users where posts.kind = ? and posts.status = 1 and posts.host_id = users.id order by users.experience desc ;"
	rows,_ := utils.Db.Query(sqlStr,"create")
	var posts []*model.Post
	for rows.Next(){
		post := &model.Post{}
		_ = rows.Scan(&post.BarName,&post.PostTitle)
		posts = append(posts,post)
	}
	return posts
}

//FindPostsByFind 模糊搜索帖子
func FindPostsByFind(find string)[]*model.Post{
	//写sql语句
	sqlStr := "select bar_name,title from posts,users where posts.kind = ? and posts.title like ? and posts.status = 1 and posts.host_id = users.id order by  users.experience desc ;"
	rows,_ := utils.Db.Query(sqlStr,"create",find)
	var posts []*model.Post
	for rows.Next(){
		post := &model.Post{}
		_ = rows.Scan(&post.BarName,&post.PostTitle)
		posts = append(posts,post)
	}
	return posts
}

//FindPostsByBarNameAndFind 模糊搜索贴吧内帖子
func FindPostsByBarNameAndFind(barName string,find string)[]*model.Post{
	//写sql语句
	sqlStr := "select bar_name,title from posts,users where posts.kind = ? and posts.bar_name = ? and posts.title like ? and posts.status = 1 and posts.host_id = users.id order by users.experience desc ;"
	rows,_ := utils.Db.Query(sqlStr,"create",barName,find)
	var posts []*model.Post
	for rows.Next(){
		post := &model.Post{}
		_ = rows.Scan(&post.BarName,&post.PostTitle)
		posts = append(posts,post)
	}
	return posts
}

//FindPostsByBarNameAndFindOrderByTime 根据时间与搜索信息按照时间排序查找
func FindPostsByBarNameAndFindOrderByTime(barName string,find string)[]*model.Post{
	//写sql语句
	sqlStr := "select Bar_name,title from posts where bar_name = ? and kind= ? and status = 1 and title like ? order by time desc; "
	rows,_ := utils.Db.Query(sqlStr,barName,"create",find)
	var posts []*model.Post
	for rows.Next(){
		post := &model.Post{}
		_ = rows.Scan(&post.BarName,&post.PostTitle)
		posts = append(posts,post)
	}
	return posts
}

//IsPostThumbByBarNameAndUserIDAndKind 是否点赞
func IsPostThumbByBarNameAndUserIDAndKind(barName string,hostID int64,title string)bool{
	sqlStr := "select id from posts where bar_name = ? and host_id = ? and title = ? and kind = ? and is_thumb = ? ;"
	row := utils.Db.QueryRow(sqlStr,barName,hostID,title,"nil",1)
	var id int
	_ = row.Scan(&id)
	if id > 0 {
		return true
	} else {
		return false
	}
}

//IsPostLikedByBarNameAndUserID 是否收藏
func IsPostLikedByBarNameAndUserID(barName string,hostID int64,title string)bool{
	sqlStr := "select id from posts where bar_name = ? and host_id = ? and title = ? and kind = ? ;"
	row := utils.Db.QueryRow(sqlStr,barName,hostID,title,"liked")
	var id int
	_ = row.Scan(&id)
	if id > 0 {
		return true
	} else {
		return false
	}
}

//ThumbPost 点赞
func ThumbPost (barName string,hostID int64,title string){
	sqlStr := "insert into posts (bar_name,host_id,title,is_thumb) values (?,?,?,?);"
	_,_ = utils.Db.Exec(sqlStr,hostID,barName,title,1)
}

//DisThumbPost 取消点赞
func DisThumbPost(barName string,hostID int64,title string){
	sqlStr := "delete from posts where bar_name = ? and host_id = ? and title = ? and is_thumb = ? ;"
	_,_ = utils.Db.Exec(sqlStr,barName,hostID,title,1)
}

//LikedPost 收藏帖子
func LikedPost (barName string,hostID int64,title string){
	sqlStr := "insert into posts (host_id,bar_name,title,kind) values (?,?,?,?);"
	_,_ = utils.Db.Exec(sqlStr,hostID,barName,title,"liked")
}

//DisLikedPost 取消收藏帖子
func DisLikedPost(barName string,hostID int64,title string){
	sqlStr := "delete from posts where bar_name = ? and host_id = ? and title = ? and kind = ? ;"
	_,_ = utils.Db.Exec(sqlStr,barName,hostID,title,"liked")
}

//DeletePostFake 虚假删除帖子
func DeletePostFake(barName string,title string,id int64){
	sqlStr := "update posts set status = 0,performer_id = ? where bar_name = ? and title = ? ;"
	_,_ = utils.Db.Exec(sqlStr,id,barName,title)
}

//FindStatusByBarNameAndTitleAndKind 查找帖子状态
func FindStatusByBarNameAndTitleAndKind(barName string,title string)(status bool){
	sqlStr := "select status from posts where bar_name = ? and title = ? and kind = ?;"
	row := utils.Db.QueryRow(sqlStr,barName,title,"create")
	var judge int64
	_ = row.Scan(&judge)
	if judge == 0 {
		status = false
	} else {
		status = true
	}
	return
}

//SetTotalPostReply 修改帖子回复量
func SetTotalPostReply(barName string,title string,total int64){
	sqlStr := "update posts set totalpostReply = ? where bar_name = ? and title = ? ;"
	_,_ = utils.Db.Exec(sqlStr,total,barName,title)
}

//FindPostsByBarNameOrderByTime 在吧内将帖子按时间排序
func FindPostsByBarNameOrderByTime(barName string)(posts []*model.Post){
	sqlStr := "select bar_name,title from posts where bar_name = ? and kind = ? and status = 1 order by time desc ;"
	rows,_ := utils.Db.Query(sqlStr,barName,"create")
	for rows.Next(){
		post := &model.Post{}
		_ = rows.Scan(&post.BarName,&post.PostTitle)
		posts = append(posts,post)
	}
	return
}

//FindPostsByBarNameOrderByPostReply 在吧内将帖子按回复量排序
func FindPostsByBarNameOrderByPostReply(barName string)(posts []*model.Post){
	sqlStr := "select bar_name,title from posts where bar_name = ? and kind = ? and status = 1 order by totalpostReply desc ;"
	rows,_ := utils.Db.Query(sqlStr,barName,"create")
	for rows.Next(){
		post := &model.Post{}
		_ = rows.Scan(&post.BarName,&post.PostTitle)
		posts = append(posts,post)
	}
	return
}

