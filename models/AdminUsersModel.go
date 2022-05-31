package models

import (
	"fmt"
)

type AdminUsers struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Mobile    string `json:"mobile"`
	RoleId    int    `json:"role_id"`
	Status    int    `json:"status"`
	Email     string `json:"email"`
	Remarks   string `json:"remarks"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AdminUsersList struct {
	AdminUsers
	RolesName  string `json:"roles_name"`
	StatusName string `json:"status_name"`
}

/**
查询用户
*/
func GetUserInfo(username, pwd string) (user AdminUsers, err error) {
	u := AdminUsers{}
	result := PDO.Where("username=?", username).Where("password=?", pwd).First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

func GetUserByUserName(username string) (user AdminUsers, err error) {
	u := AdminUsers{}
	result := PDO.Where("username=?", username).First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

/**
根据id 查询用户
*/
func GetUserInfoById(id string) (user AdminUsers, err error) {
	u := AdminUsers{}
	result := PDO.Where(" id =? ", id).First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

/**
获取用户列表
*/
func GetUserList(username, mobile, email, name string, page, size int, sort string) (user []AdminUsers, count int64, err error) {
	u := []AdminUsers{}
	offset := page * size
	Ps := PDO.Order("id desc")
	if username != "" {
		Ps = Ps.Where("username like ?", "%"+username+"%")
	}
	if mobile != "" {
		Ps = Ps.Where("mobile like ?", "%"+mobile+"%")
	}
	if email != "" {
		Ps = Ps.Where("email like ?", "%"+email+"%")
	}
	if name != "" {
		Ps = Ps.Where("name like ?", "%"+name+"%")
	}

	retcount := Ps.Model(&AdminUsers{}).Count(&count)
	if retcount.Error != nil {
		fmt.Println(retcount)
	}
	if sort == "0" {
		Ps = Ps.Order("id desc")
	} else {
		Ps = Ps.Order("id asc")
	}

	result := Ps.Limit(size).Offset(offset).Find(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, count, result.Error
}

/**
添加用户
*/
func UserAdd(u AdminUsers) (b bool, err error) {
	result := PDO.Create(&u)
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

/**
用户更新
*/
func UserEdit(id string, u AdminUsers) (b bool, err error) {
	result := PDO.Model(&AdminUsers{}).Where(" id=? ", id).Updates(u)
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

func UserDel(id string) (b bool, err error) {
	result := PDO.Where(" id = ? ", id).Delete(&AdminUsers{})
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}
