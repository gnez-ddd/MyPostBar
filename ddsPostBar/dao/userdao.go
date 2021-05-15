// Package dao 用户数据库
//用户表
/*
create table users(
    -> id int primary key auto_increment,
    -> username varchar(100) not null unique,
    -> password varchar(100) not null,
    -> email varchar(100),
	-> head_path varchar(100) not null
	-> );
 */

package dao

import (
	"MyPostBar/ddsPostBar/model"
	"MyPostBar/ddsPostBar/utils"
)

//FindUserByUserID 根据用户id查找用户
func FindUserByUserID(id int64)*model.Users{
	//写sql语句
	sqlStr := "select status,id,username,password,email,head_path,salt,experience,signInTime from users where id = ? ;"
	//执行查找语句
	row := utils.Db.QueryRow(sqlStr,id)
	//创建一个结构体用于存储查找出来的用户
	user := &model.Users{}
	//将查找出来的信息存储起来
	_ = row.Scan(&user.Status,&user.UserID,&user.UserName,&user.PassWord,&user.Email,&user.HeadPath,&user.Salt,&user.Experience,&user.LastSign)
	return user
}

// FindUserByUserName 根据用户名查找用户1
func FindUserByUserName(username string)*model.Users{
	//写sql语句
	sqlStr := "select id,username,password,email,head_path,status from users where username = ? ;"
	//执行查找语句
	row := utils.Db.QueryRow(sqlStr,username)
	//创建一个结构体用于存储查找出来的用户
	user := &model.Users{}
	//将查找出来的信息存储起来
	_ = row.Scan(&user.UserID,&user.UserName,&user.PassWord,&user.Email,&user.HeadPath,&user.Status)
	return user
}

//FindUserByEmail 根据邮箱查找用户1
func FindUserByEmail(email string)*model.Users{
	//写sql语句
	sqlStr := "select id,username,password,email,head_path,status from users where email = ? ;"
	//执行查找语句
	row := utils.Db.QueryRow(sqlStr,email)
	//创建一个结构体用户存储查找出来的用户
	user := &model.Users{}
	//将查找出来的信息存储起来
	_ = row.Scan(&user.UserID,&user.UserName,&user.PassWord,&user.Email,&user.HeadPath,&user.Status)
	return user
}

// CheckUserNameAndPassWord 判断用户名与密码是否正确
func CheckUserNameAndPassWord(username string,password string)*model.Users{
	//写sql语句
	sqlStr := "select id,username,password,email,head_path from users where username = ? and password = ? ;"
	//执行sql语句
	row := utils.Db.QueryRow(sqlStr,username,password)
	//创建一个结构体用于保存该用户信息
	user := &model.Users{}
	//将用户信息保存起来
	_  = row.Scan(&user.UserID,&user.UserName,&user.PassWord,&user.Email,&user.HeadPath)
	return user
}

//AddUser 添加用户
func AddUser(username string,password string,email string,salt string){
	//写sql语句
	sqlStr := "insert into users(username,password,email,head_path,salt) values(?,?,?,?,?) ;"
	//执行插入语句
	_,_ = utils.Db.Exec(sqlStr,username,password,email,"../../static/img/head/head.JPG",salt)
}

//SetPassword 修改密码
func SetPassword(password string,email string,salt string){
	//写sql语句
	sqlStr := "update users set password = ?,salt = ? where email = ? ;"
	//执行修改语句
	_,_ = utils.Db.Exec(sqlStr,password,salt,email)
}

//SetUserInformation 修改用户信息
func SetUserInformation(username string,email string,headPath string,userID int64,salt string,psw string){
	//写sql语句
	sqlStr := "update users set username = ?,email = ?,head_path = ?,salt = ?,password = ? where id = ?;"
	//执行修改语句
	_,_ = utils.Db.Exec(sqlStr,username,email,headPath,salt,psw,userID)
}



//FindSaltByUserID 根据用户id查看用户的盐值
func FindSaltByUserID(id int64)(salt string){
	sqlStr := "select salt from users where id = ? ;"
	row := utils.Db.QueryRow(sqlStr,id)
	_ = row.Scan(&salt)
	return
}

//UserSignIn 签到
func UserSignIn(userName string,now string,experience int64){
	sqlStr := "update users set experience = ?,signInTime = ? where username = ? ; "
	_,_ = utils.Db.Exec(sqlStr,experience,now,userName)
}

//SetUserStatusByID 修改用户状态
func SetUserStatusByID(status int64,id int64){
	sqlStr := "update users set status = ? where id = ?;"
	_,_ = utils.Db.Exec(sqlStr,status,id)
}