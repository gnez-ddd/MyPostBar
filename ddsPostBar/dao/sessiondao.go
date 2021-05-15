package dao

import (
	"MyPostBar/ddsPostBar/model"
	"MyPostBar/ddsPostBar/utils"
	"net/http"
)

/*
 create table sessions(
    -> session_id varchar(100) primary key,
    -> username varchar(100) not null,
    -> user_id int not null,
    -> foreign key(user_id) references users(id)
    -> );
 */

// AddSession 添加sess
func AddSession(sess *model.Session){
	//写sql语句
	sqlStr := "insert into sessions(session_id,user_id) values (?,?) ;"
	//执行sql语句
	_,_ = utils.Db.Exec(sqlStr,sess.SessionID,sess.UserID)
}

// DeleteSession 删除sess
func DeleteSession(sessionID string){
	//写sql语句
	sqlStr := "delete from sessions where session_id = ?"
	//执行删除语句
	_, _ = utils.Db.Exec(sqlStr, sessionID)
}

//GetSession 根据session的Id值从数据库中查询Session
func GetSession(sessID string) (*model.Session, error) {
	//写sql语句
	sqlStr := "select session_id,user_id from sessions where session_id = ?"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	//执行
	row := inStmt.QueryRow(sessID)
	//创建Session
	sess := &model.Session{}
	//扫描数据库中的字段值为Session的字段赋值
	_ = row.Scan(&sess.SessionID, &sess.UserID)
	return sess, nil
}

//IsLogin 判断用户是否已经登录 false 没有登录 true 已经登录
func IsLogin(r *http.Request) (bool, *model.Session) {
	//根据Cookie的name获取Cookie
	cookie, _ := r.Cookie("userhandler")
	if cookie != nil {
		//获取Cookie的value
		cookieValue := cookie.Value
		//根据cookieValue去数据库中查询与之对应的Session
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			//已经登录
			return true, session
		}
	}
	//没有登录
	return false, nil
}

