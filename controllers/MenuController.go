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
	"strings"
)

type MenuSearch struct {
	Title string
}
/**
菜单列表
 */
func MenuList(c *gin.Context) {
	if c.Request.Method == "POST" {
		pagestr := c.DefaultPostForm("page", "1")
		limitstr := c.DefaultPostForm("limit", "15")
		searchParams := c.PostForm("searchParams")
		name := ""
		page := 1
		limit := 0
		limit, _ = strconv.Atoi(limitstr)
		page, _ = strconv.Atoi(pagestr)
		page = page - 1
		if searchParams != "" {
			params := MenuSearch{}
			err := json.Unmarshal([]byte(searchParams), &params)
			if err != nil {
				fmt.Println(err)
			}
			name = params.Title
		}
		result, count, err := models.GetMenuList(name, page, limit)
		if err != nil {
			fmt.Println(err)
		}
		retlist := make([]models.AdminMenuList, len(result))
		catMap := GetCategoryMap()
		for k, v := range result {
			retlist[k].UpdatedAt = until.TimeLocateToFormat(v.UpdatedAt)
			retlist[k].Id = v.Id
			retlist[k].Title = v.Title
			retlist[k].Orders = v.Orders
			retlist[k].Icon = v.Icon
			retlist[k].Uri = v.Uri
			Cname := ""
			if v.IsCategory == 1 {
				Cname = `<button type="button" class="layui-btn layui-btn-xs">是</button>`
			} else {
				Cname = `<button type="button" class="layui-btn layui-btn-xs layui-btn-danger">否</button>`
			}
			retlist[k].CategoryName = Cname
			_, ok := catMap[v.Id]
			if ok && v.ParentId != 0 {
				retlist[k].ParentName = catMap[v.ParentId]
			} else {
				if v.ParentId != 0 {
					retlist[k].ParentName = catMap[v.ParentId]
				} else {
					retlist[k].ParentName = "root"
				}
			}
		}
		data := until.ReturnCount(until.Success, count, retlist, "")
		c.JSON(http.StatusOK, data)
		return
	}
	data := gin.H{}
	c.HTML(http.StatusOK, "menulist.html", data)
}

/**
栏目对象
*/
func GetCategoryMap() (cap map[int]string) {
	categoryName, err := models.GetMenuByCategorByStatus(1)
	if err != nil {
		fmt.Println(err)
	}
	//栏目 map
	CateMap := make(map[int]string, 0)
	for _, v := range categoryName {
		CateMap[v.Id] = v.Title
	}
	return CateMap
}

/**
添加 菜单
*/
func MenuAdd(c *gin.Context) {
	if c.Request.Method == "POST" {
		Title := c.PostForm("title")
		is_category := c.PostForm("is_category")
		parent_idstr := c.PostForm("parend_id")
		Uri := c.PostForm("uri")
		icon := c.PostForm("icon")
		orderstr := c.PostForm("order")
		result := ""
		u := models.AdminMenu{}
		if Title == "" {
			data := until.Return(until.Error, result, "菜单名不可以为空")
			c.JSON(http.StatusOK, data)
			return
		}
		url := strings.Trim(Uri, "/")
		u.Title = Title
		u.Uri = url
		adm, _ := models.GetMenusCount(url, 0)
		if adm.Id > 0 {
			data := until.Return(until.Error, result, "路径地址不可以重复")
			c.JSON(http.StatusOK, data)
			return
		}
		is_categoryInt, _ := strconv.Atoi(is_category)
		orderInt, _ := strconv.Atoi(orderstr)
		parentInt, _ := strconv.Atoi(parent_idstr)
		u.IsCategory = is_categoryInt
		u.Orders = orderInt
		u.ParentId = parentInt
		u.Icon = icon
		if parentInt == 0 {
			u.Uri = ""
		}
		u.CreatedAt = until.GetNowTime()
		u.UpdatedAt = until.GetNowTime()
		b, _ := models.MenuAdd(u)
		if !b {
			data := until.Return(until.Success, result, "添加失败")
			c.JSON(http.StatusOK, data)
			return
		}
		data := until.Return(until.Success, result, "")
		c.JSON(http.StatusOK, data)
		return
	}
	cate := models.GetMenus()
	fmt.Println(cate)
	data := gin.H{
		"cate": cate,
	}
	c.HTML(http.StatusOK, "menuadd.html", data)
}

/**
添加 菜单
*/
func MenuEdit(c *gin.Context) {
	if c.Request.Method == "POST" {
		Title := c.PostForm("title")
		is_category := c.PostForm("is_category")
		parent_idstr := c.PostForm("parend_id")
		Uri := c.PostForm("uri")
		icon := c.PostForm("icon")
		orderstr := c.PostForm("order")
		id := c.PostForm("id")

		result := ""
		u := models.AdminMenu{}
		if Title == "" {
			data := until.Return(until.Error, result, "菜单名不可以为空")
			c.JSON(http.StatusOK, data)
			return
		}

		url := strings.Trim(Uri, "/")
		idint, _ := strconv.Atoi(id)

		adm, _ := models.GetMenusCount(url, idint)

		fmt.Println(url)

		fmt.Println(adm)

		if adm.Id > 0 {
			data := until.Return(until.Error, result, "路径地址不可以重复")
			c.JSON(http.StatusOK, data)
			return
		}

		u.Title = Title
		u.Uri = Uri
		is_categoryInt, _ := strconv.Atoi(is_category)

		orderInt, _ := strconv.Atoi(orderstr)
		parentInt, _ := strconv.Atoi(parent_idstr)

		u.IsCategory = is_categoryInt
		u.Orders = orderInt
		u.ParentId = parentInt
		u.Icon = icon

		u.UpdatedAt = until.GetNowTime()

		b, _ := models.MenuEdit(id, u)
		if !b {
			data := until.Return(until.Success, result, "编辑失败")
			c.JSON(http.StatusOK, data)
			return
		}

		data := until.Return(until.Success, result, "")
		c.JSON(http.StatusOK, data)
		return
	}
	cate := models.GetMenus()
	id := c.Query("id")
	menus, err := models.GetMenuInfoById(id)
	if err != nil {
		until.LogObj().WithFields(logrus.Fields{
			"func": "MenuEdit",
			"id":   id,
		}).Error("获取参数：", err)
	}
	data := gin.H{
		"cate": cate,
		"menu": menus,
	}
	c.HTML(http.StatusOK, "menuedit.html", data)
}

/**
删除
*/
func MenuDel(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		data := until.Return(until.Success, "", "id 不可以为空")
		c.JSON(http.StatusOK, data)
	}
	bool, err := models.MenuDel(id)
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
