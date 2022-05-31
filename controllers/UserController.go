package controllers

import (
	"Gin/models"
	"Gin/until"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UserSearch struct {
	Username string
	Email    string
	Mobile   string
	Name     string
}

func UserList(c *gin.Context) {
     if c.Request.Method=="POST"{
		 pagestr := c.DefaultPostForm("page", "1")
		 limitstr := c.DefaultPostForm("limit", "15")
		 searchParams := c.PostForm("searchParams")
		 username := ""
		 mobile := ""
		 name := ""
		 email := ""
		 page := 1
		 limit := 0
		 limit, _ = strconv.Atoi(limitstr)
		 page, _ = strconv.Atoi(pagestr)
		 page = page - 1
		 if searchParams != "" {
			 params := UserSearch{}
			 err := json.Unmarshal([]byte(searchParams), &params)
			 if err != nil {
				 until.LogObj().WithFields(logrus.Fields{
					 "func": "UserListData",
				 }).Error("获取参数：", err)
			 }
			 username = params.Username
			 mobile = params.Mobile
			 email = params.Email
			 name = params.Name
		 }

		 result, count, err := models.GetUserList(username, mobile, email, name, page, limit, "1")
		 if err != nil {
			 fmt.Println(err)
		 }
		 Roles, _ := models.GetAllRoles()
		 RolesMap := make(map[int]string, len(Roles))
		 for _, v := range Roles {
			 RolesMap[v.Id] = v.Name
		 }
		 Adulist := make([]models.AdminUsersList, len(result))
		 for k, v := range result {
			 Adulist[k].Name = v.Name
			 RoleName := ""
			 if v.Status == 1 {
				 RoleName = `<button type="button" class="layui-btn layui-btn-xs">正常</button>`
			 } else {
				 RoleName = `<button type="button" class="layui-btn layui-btn-xs layui-btn-danger">禁用</button>`
			 }
			 Adulist[k].StatusName = RoleName
			 Adulist[k].Username = v.Username
			 Adulist[k].UpdatedAt = until.TimeLocateToFormat(v.UpdatedAt)
			 Adulist[k].Mobile = v.Mobile
			 Adulist[k].Email = v.Email
			 Adulist[k].Id = v.Id
			 Adulist[k].RolesName = `<button type="button" class="layui-btn layui-btn-xs layui-btn-normal">` + RolesMap[v.RoleId] + `</button>`
		 }
		 data := until.ReturnCount(until.Success, count, Adulist, "")
		 c.JSON(http.StatusOK, data)
		 return
	 }
	data := gin.H{}
	c.HTML(http.StatusOK, "userlist.html", data)
}



func UserAdd(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		status := c.PostForm("status")
		mobile := c.PostForm("mobile")
		email := c.PostForm("email")
		name := c.PostForm("name")
		remarks := c.PostForm("remarks")
		password := c.PostForm("password")
		role_id := c.PostForm("role_id")
		result := ""
		u := models.AdminUsers{}
		if username == "" {
			data := until.Return(until.Success, result, "用户名不可以为空")
			c.JSON(http.StatusOK, data)
			return
		}

		user, err := models.GetUserByUserName(username)
		if err == nil || user.Id > 0 {
			data := until.Return(until.Success, result, "用户名不可重复")
			c.JSON(http.StatusOK, data)
			return
		}

		if password != "" {
			u.Password = until.Md5(password)
		} else {
			data := until.Return(until.Success, result, "密码不可为空")
			c.JSON(http.StatusOK, data)
			return
		}

		rolesInt, _ := strconv.Atoi(role_id)

		u.RoleId = rolesInt
		u.Username = username
		u.Name = name
		u.Email = email
		statusInt, _ := strconv.Atoi(status)

		u.Status = statusInt
		u.Mobile = mobile
		u.Remarks = remarks

		u.CreatedAt = until.GetNowTime()
		u.UpdatedAt = until.GetNowTime()
		b, _ := models.UserAdd(u)
		if !b {
			data := until.Return(until.Success, result, "添加失败")
			c.JSON(http.StatusOK, data)
			return
		}

		data := until.Return(until.Success, result, "")
		c.JSON(http.StatusOK, data)
		return
	}

	roles, err := models.GetAllRoles()
	if err != nil {
		fmt.Println(err)
	}
	data := gin.H{
		"roles": roles,
	}

	c.HTML(http.StatusOK, "useradd.html", data)
}

/**
用户编辑
*/
func UserEdit(c *gin.Context) {
	if c.Request.Method == "POST" {
		status := c.PostForm("status")
		mobile := c.PostForm("mobile")
		email := c.PostForm("email")
		name := c.PostForm("name")
		remarks := c.PostForm("remarks")
		password := c.PostForm("password")
		id := c.PostForm("id")
		role_id := c.PostForm("role_id")
		result := ""
		u := models.AdminUsers{}

		if password != "" {
			u.Password = until.Md5(password)
		}
		u.Name = name
		u.Email = email
		statusInt, _ := strconv.Atoi(status)
		rolesInt, _ := strconv.Atoi(role_id)
		u.RoleId = rolesInt
		u.Status = statusInt
		u.Mobile = mobile
		u.Remarks = remarks
		u.UpdatedAt = until.GetNowTime()

		b, _ := models.UserEdit(id, u)
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
	}
	user, err := models.GetUserInfoById(id)
	if err != nil {
		until.LogObj().WithFields(logrus.Fields{
			"func": "UserEdit",
			"id":   id,
		}).Error("获取参数：", err)
	}

	roles, err := models.GetAllRoles()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)

	data := gin.H{
		"user":  user,
		"roles": roles,
	}
	c.HTML(http.StatusOK, "useredit.html", data)
}

/**
del
*/
func UserDel(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		data := until.Return(until.Success, "", "id 不可以为空")
		c.JSON(http.StatusOK, data)
	}
	bool, err := models.UserDel(id)
	if err != nil {
		until.LogObj().WithFields(logrus.Fields{
			"func": "UserDel",
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
