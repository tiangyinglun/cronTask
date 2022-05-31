package models

import (
	"fmt"
)

type AdminRoles struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

/**
获取所有角色
 */
func GetAllRoles() (user []AdminRoles, err error) {
	u := []AdminRoles{}
	result := PDO.Find(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}




func GetRolesByName(username string) (user AdminRoles, err error) {
	u := AdminRoles{}
	result := PDO.Where(" name=? ", username).First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

/**
根据id 查询用户
*/
func GetRolesInfoById(id string) (user AdminRoles, err error) {
	u := AdminRoles{}
	result := PDO.Where(" id =? ", id).First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

/**
获取用户列表
*/
func GetRolesList(name string, page, size int, sort string) (user []AdminRoles, count int64, err error) {
	u := []AdminRoles{}
	offset := page * size
	Ps := PDO.Order("id desc")
	if name != "" {
		Ps = Ps.Where("name like ?", "%"+name+"%")
	}

	retcount := Ps.Model(&AdminRoles{}).Count(&count)
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
func RolesAdd(u AdminRoles) (b bool, err error) {
	result := PDO.Create(&u)
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

/**
用户更新
*/
func RolesEdit(id string, u AdminRoles) (b bool, err error) {
	result := PDO.Model(&AdminRoles{}).Where(" id=? ", id).Updates(u)
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

func RolesDel(id string) (b bool, err error) {
	result := PDO.Where(" id = ? ", id).Delete(&AdminRoles{})
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}


