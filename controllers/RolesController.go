package controllers

import (
	"Gin/models"
	"Gin/until"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

func RolesList(c *gin.Context) {
	if c.Request.Method == "POST" {
		pagestr := c.DefaultPostForm("page", "1")
		limitstr := c.DefaultPostForm("limit", "15")
		name := ""
		page := 1
		limit := 0
		limit, _ = strconv.Atoi(limitstr)
		page, _ = strconv.Atoi(pagestr)
		page = page - 1

		result, count, err := models.GetRolesList(name, page, limit, "1")
		if err != nil {
			fmt.Println(err)
		}
		Tresult := make([]models.AdminRoles, len(result))
		for k, v := range result {
			Tresult[k].Id = v.Id
			Tresult[k].Name = v.Name
			Tresult[k].Slug = v.Slug
			Tresult[k].CreatedAt = until.TimeLocateToFormat(v.CreatedAt)
			Tresult[k].UpdatedAt = until.TimeLocateToFormat(v.UpdatedAt)
		}

		data := until.ReturnCount(until.Success, count, Tresult, "")
		c.JSON(http.StatusOK, data)
		return
	}
	data := gin.H{}
	c.HTML(http.StatusOK, "roleslist.html", data)
}

/**
添加
*/
func RolesAdd(c *gin.Context) {
	if c.Request.Method == "POST" {
		name := c.PostForm("name")
		result := ""
		u := models.AdminRoles{}
		if name == "" {
			data := until.Return(until.Success, result, "名称可以为空")
			c.JSON(http.StatusOK, data)
			return
		}

		user, err := models.GetRolesByName(name)
		if err == nil || user.Id > 0 {
			data := until.Return(until.Success, result, "名称不可重复")
			c.JSON(http.StatusOK, data)
			return
		}
		u.Name = name
		u.Slug = name
		u.CreatedAt = until.GetNowTime()
		u.UpdatedAt = until.GetNowTime()
		b, _ := models.RolesAdd(u)
		if !b {
			data := until.Return(until.Success, result, "添加失败")
			c.JSON(http.StatusOK, data)
			return
		}

		data := until.Return(until.Success, result, "")
		c.JSON(http.StatusOK, data)
		return
	}
	data := gin.H{}
	c.HTML(http.StatusOK, "rolesadd.html", data)
}

/**
编辑
*/
func RolesEdit(c *gin.Context) {
	if c.Request.Method == "POST" {
		name := c.PostForm("name")
		id := c.PostForm("id")
		if id == "" {
			data := until.Return(until.Success, "", "id 不可以为空")
			c.JSON(http.StatusOK, data)
			return
		}
		result := ""
		u := models.AdminRoles{}
		u.Name = name
		u.Slug = name
		u.UpdatedAt = until.GetNowTime()
		b, _ := models.RolesEdit(id, u)
		if !b {
			data := until.Return(until.Success, result, "添加失败")
			c.JSON(http.StatusOK, data)
			return
		}

		data := until.Return(until.Success, result, "")
		c.JSON(http.StatusOK, data)
		return
	}
	id := c.Query("id")
	if id == "" {
		data := until.Return(until.Error, "", "id 不可以为空")
		c.JSON(http.StatusOK, data)
		return
	}

	roles, err := models.GetRolesInfoById(id)
	if err != nil {
		until.LogObj().WithFields(logrus.Fields{
			"func": "RolesEdit",
			"id":   id,
		}).Error("获取参数：", err)
	}

	data := gin.H{
		"roles": roles,
	}
	c.HTML(http.StatusOK, "rolesedit.html", data)
}

/**
删除
*/
func RolesDel(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		data := until.Return(until.Success, "", "id 不可以为空")
		c.JSON(http.StatusOK, data)
		return
	}
	bool, err := models.RolesDel(id)
	if err != nil {
		until.LogObj().WithFields(logrus.Fields{
			"func": "RolesDel",
			"id":   id,
		}).Error("获取参数：", err)

		data := until.Return(until.Error, "", "删除失败")
		c.JSON(http.StatusOK, data)
		return
	}
	if bool {
		data := until.Return(until.Success, "", "")
		c.JSON(http.StatusOK, data)
		return
	}
	data := until.Return(until.Error, "", "")
	c.JSON(http.StatusOK, data)
}

/**
权限
*/
func RolesRoles(c *gin.Context) {
	if c.Request.Method == "POST" {
		roleid := c.PostForm("roleid")
		menuIdstr := c.PostForm("menuid")
		menuslice := strings.Split(menuIdstr, ",")
		err := models.InsertRoleMenu(roleid, menuslice)
		if err != nil {
			data := until.Return(until.Error, "", "编辑失败")
			c.JSON(http.StatusOK, data)
			return
		}

		data := until.Return(until.Success, "", "编辑成功")
		c.JSON(http.StatusOK, data)
		return
	}
	rolesId := c.Query("id")
	rolesinfo, err := models.GetRolesInfoById(rolesId)
	if err != nil {
		until.LogObj().WithFields(logrus.Fields{
			"func": "RolesRoles",
			"id":   rolesId,
		}).Error("获取参数：", err)
	}

	RoleIdInt, _ := strconv.Atoi(rolesId)
	menus := GetCheckMenu(RoleIdInt)
	fmt.Println(menus)
	data := gin.H{
		"rolesinfo": rolesinfo,
		"menus":     menus,
	}
	c.HTML(http.StatusOK, "rolesroles.html", data)
}

/**
只有三级就简单的搞定了
*/
func GetCheckMenu(roleId int) []models.AdminMenuTree {
	menus := models.GetAllMenus()
	Roles := GetMapRoles(roleId)

	rolesCheck := make([]models.AdminMenuTree, len(menus))
	for k, v := range menus {
		rolesCheck[k].Color = v.Color
		rolesCheck[k].Id = v.Id
		rolesCheck[k].ParentId = v.ParentId
		rolesCheck[k].Title = v.Title
		rolesCheck[k].Count = v.Count
		_, ok := Roles[v.Id]
		if ok {
			rolesCheck[k].Checked = true
		} else {
			rolesCheck[k].Checked = false
		}
		if len(v.Children) > 0 {
			listOne := make([]models.AdminMenuTree, len(v.Children))
			for kv, tf := range v.Children {
				listOne[kv].Id = tf.Id
				listOne[kv].Color = tf.Color
				listOne[kv].ParentId = tf.ParentId
				listOne[kv].Title = tf.Title
				listOne[kv].Count = tf.Count
				_, fok := Roles[tf.Id]
				if fok {
					listOne[kv].Checked = true
				} else {
					listOne[kv].Checked = false
				}
				if len(tf.Children) > 0 {
					listTwo := make([]models.AdminMenuTree, len(tf.Children))
					for ktv, rtf := range tf.Children {
						listTwo[ktv].Id = rtf.Id
						listTwo[ktv].Color = rtf.Color
						listTwo[ktv].ParentId = rtf.ParentId
						listTwo[ktv].Title = rtf.Title
						listTwo[ktv].Count = rtf.Count
						_, ffok := Roles[rtf.Id]
						if ffok {
							listTwo[ktv].Checked = true
						} else {
							listTwo[ktv].Checked = false
						}
					}
					listOne[kv].Children = listTwo
				}
			}
			rolesCheck[k].Children = listOne
		} else {
			rolesCheck[k].Children = v.Children
		}
	}
	return rolesCheck
}

/**
获取权限 map
*/
func GetMapRoles(roleId int) map[int]int {
	roleMenu, errs := models.GetAllRolesMenu(roleId)
	if errs != nil {
		until.LogObj().WithFields(logrus.Fields{
			"func": "RolesRoles",
			"id":   roleId,
		}).Error("获取参数：", errs)
	}
	roleMap := make(map[int]int, 0)
	for _, v := range roleMenu {
		roleMap[v.MenuId] = v.MenuId
	}
	return roleMap
}
