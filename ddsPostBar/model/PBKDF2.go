package model

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
)

//GetSalt 生成盐值
func GetSalt()(salt string){
	bytes := make([]byte,32)
	_,_ = rand.Read(bytes)
	//将bytes转为16进制字符串
	salt = hex.EncodeToString(bytes)
	return
}

//GetPBKDF2 获取密文
func GetPBKDF2(pwd string,salt string)(DK string){
	//将字符串转为byte数组
	bytes,_ := hex.DecodeString(salt)
	dk := pbkdf2.Key([]byte(pwd),bytes,1,32,sha256.New)
	DK = hex.EncodeToString(dk)
	return
}

