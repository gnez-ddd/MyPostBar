package dao

import "MyPostBar/ddsPostBar/utils"

//FindCodeIdByEmail 根据邮箱查找id
func FindCodeIdByEmail(email string)(id int64){
	sqlStr := "select id from code where email = ?;"
	row := utils.Db.QueryRow(sqlStr,email)
	_ = row.Scan(&id)
	return
}

//SetCodeByID 根据id修改code
func SetCodeByID(id int64,code string){
	sqlStr := "update code set code_num = ? where id = ?;"
	_,_ = utils.Db.Exec(sqlStr,code,id)
}

//AddCode 新增邮箱
func AddCode(email string,code string){
	sqlStr := "insert into code (email,code_num) values (?,?);"
	_,_ = utils.Db.Exec(sqlStr,email,code)
}

//FindCodeNumByEmail 根据邮箱查找验证码
func FindCodeNumByEmail(email string)(code string){
	sqlStr := "select code_num from code where email = ?;"
	row := utils.Db.QueryRow(sqlStr,email)
	_ = row.Scan(&code)
	return
}

//SetCodeNumNil 将验证码设置为空
func SetCodeNumNil(email string){
	sqlStr := "update code set code_num = ? where email = ?;"
	_,_ = utils.Db.Exec(sqlStr,"",email)
}