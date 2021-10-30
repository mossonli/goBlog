package model

import (
	"encoding/base64"
	"fmt"
	"github.com/jinzhu/gorm"
	"goBlog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(32);not null" json:"username" validate:"required,min=4,max=18" label:"用户名"`
	Password string `gorm:"type:varchar(100); not null" json:"password" validate:"required,min=6,max=32" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色"`
}
// 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPwd(data.Password)
	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUserList(pageSize, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, 0
	}
	return users, total
}
/*
使用map更新多个值，只会更新其中发生变化的属性
使用struct更新多个值，只会更新其中有变化且非0的值
*/

// 编辑用户
func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	fmt.Println("edit data", data)
	err := db.Model(&User{}).Where("id= ?", id).Updates(maps).Error // Updates 多列 Update 单列
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DelUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}


// 密码加密
func ScryptPwd(password string)string  {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{22,35,64,65,44,32,31,12} // 里面的数随便定义
	HashPwd, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, KeyLen)
	if err != nil{
		log.Fatal(err)
	}
	fpwd := base64.StdEncoding.EncodeToString(HashPwd)
	// 加密之后的密码入库
	// 方式1 gorm的钩子函数 beforesave

	// 方式2 CreateUser 的时候把密码进行加密

	return fpwd
}

// 登录验证
func CheckLogin(username, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0{
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPwd(password) != user.Password{
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
