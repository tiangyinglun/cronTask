package controllers

import (
	"Gin/models"
	"Gin/until"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GroupList(c *gin.Context) {
	if c.Request.Method == "POST" {
		pagestr := c.DefaultPostForm("page", "1")
		limitstr := c.DefaultPostForm("limit", "15")
		page := 1
		limit := 0
		limit, _ = strconv.Atoi(limitstr)
		page, _ = strconv.Atoi(pagestr)
		page = page - 1
		userid := c.GetInt("user_id")

		result, count, err := models.GetTaskGroupList(userid, page, limit, "1")
		if err != nil {
			fmt.Println(err)
		}
		Tresult := make([]models.AdminTaskGroup, len(result))
		for k, v := range result {
			Tresult[k].Id = v.Id
			Tresult[k].UserId = v.UserId
			Tresult[k].GroupName = v.GroupName
			Tresult[k].CreatedAt = until.TimeLocateToFormat(v.CreatedAt)
		}
		data := until.ReturnCount(until.Success, count, Tresult, "")
		c.JSON(http.StatusOK, data)
		return
	}

	data := gin.H{}
	c.HTML(http.StatusOK, "grouplist.html", data)
}

/**
添加
*/
func GroupAdd(c *gin.Context) {
	if c.Request.Method == "POST" {
		name := c.PostForm("name")
		result := ""
		u := models.AdminTaskGroup{}
		if name == "" {
			data := until.Return(until.Success, result, "名称可以为空")
			c.JSON(http.StatusOK, data)
			return
		}
		user_id := c.GetInt("user_id")

		user, err := models.GetTaskGroupByName(user_id, name)
		if err == nil || user.Id > 0 {
			data := until.Return(until.Success, result, "名称不可重复")
			c.JSON(http.StatusOK, data)
			return
		}
		u.GroupName = name
		u.UserId=c.GetInt("user_id")
		u.CreatedAt = until.GetNowTime()

		b, _ := models.TaskGroupAdd(u)
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
	c.HTML(http.StatusOK, "groupadd.html", data)
}

/**
编辑
*/
func GroupEdit(c *gin.Context) {
	if c.Request.Method == "POST" {
		name := c.PostForm("name")
		id := c.PostForm("id")
		if id == "" {
			data := until.Return(until.Success, "", "id 不可以为空")
			c.JSON(http.StatusOK, data)
			return
		}
		result := ""
		u := models.AdminTaskGroup{}
		u.GroupName = name

		b, _ := models.TaskGroupEdit(id, u)
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

	group, err := models.GetTaskGroupInfoById(id)
	if err != nil {
		until.LogObj().WithFields(logrus.Fields{
			"func": "RolesEdit",
			"id":   id,
		}).Error("获取参数：", err)
	}

	data := gin.H{
		"group": group,
	}
	c.HTML(http.StatusOK, "groupedit.html", data)
}

/**
删除
*/
func GroupDel(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		data := until.Return(until.Success, "", "id 不可以为空")
		c.JSON(http.StatusOK, data)
		return
	}
	bool, err := models.TaskGroupDel(id)
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
