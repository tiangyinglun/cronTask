package models

import (
	"strconv"
	"time"
)

type AdminRoleMenu struct {
	RoleId    int    `json:"role_id"`
	MenuId    int    `json:"menu_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

/**
获得所有权限
*/
func GetAllRolesMenu(roleId int) (AdM []AdminRoleMenu, err error) {
	roleMenu := []AdminRoleMenu{}
	result := PDO.Where(" role_id=? ", roleId).Find(&roleMenu)
	return roleMenu, result.Error
}



/**
h获取权限  事务提交
*/
func InsertRoleMenu(roleId string, menu []string) (err error) {
	tx := PDO.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	roleMenu := AdminRoleMenu{}
	ret := tx.Where("role_id=?", roleId).Delete(&roleMenu)
	if ret.Error != nil {
		tx.Rollback()
		return ret.Error
	}
	if len(menu) == 0 {
		return tx.Commit().Error
	}
	AdMBox := []AdminRoleMenu{}
	roleIdint, _ := strconv.Atoi(roleId)
	for _, v := range menu {
		slic := AdminRoleMenu{}
		slic.RoleId = roleIdint
		vint, _ := strconv.Atoi(v)
		slic.MenuId = vint
		slic.CreatedAt = GetNowTime()
		slic.UpdatedAt = GetNowTime()
		AdMBox = append(AdMBox, slic)
	}
	result := tx.Create(&AdMBox)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

func GetNowTime() string {
	t := time.Now()
	nowTime := t.Format("2006-01-02 15:04:05")
	return nowTime
}
