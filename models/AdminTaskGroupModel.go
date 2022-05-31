package models

import "fmt"

type AdminTaskGroup struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	GroupName   string `json:"group_name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

/**
获取所有角色
*/
func GetAllTaskGroup(Uid int) (user []AdminTaskGroup, err error) {
	u := []AdminTaskGroup{}
	result := PDO.Where("user_id=?", Uid).Find(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

func GetTaskGroupByName(uid int, name string) (group AdminTaskGroup, err error) {
	u := AdminTaskGroup{}
	result := PDO.Where("user_id=?", uid).Where(" group_name=? ", name).First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

/**
根据id 查询用户
*/
func GetTaskGroupInfoById(id string) (group AdminTaskGroup, err error) {
	u := AdminTaskGroup{}
	result := PDO.Where(" id =? ", id).First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

/**
获取用户列表
*/
func GetTaskGroupList(user_id, page, size int, sort string) (group []AdminTaskGroup, count int64, err error) {
	u := []AdminTaskGroup{}
	offset := page * size
	Ps := PDO.Where("user_id=?", user_id).Order("id desc")
	retcount := Ps.Model(&AdminTaskGroup{}).Count(&count)
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
func TaskGroupAdd(u AdminTaskGroup) (b bool, err error) {
	result := PDO.Create(&u)
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

/**
用户更新
*/
func TaskGroupEdit(id string, u AdminTaskGroup) (b bool, err error) {
	result := PDO.Model(&AdminTaskGroup{}).Where(" id=? ", id).Updates(u)
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

/**
删除分类
*/
func TaskGroupDel(id string) (b bool, err error) {
	result := PDO.Where(" id = ? ", id).Delete(&AdminTaskGroup{})
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}
