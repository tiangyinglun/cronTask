package route

import (
	"Gin/controllers"
	"Gin/models"
	"Gin/until"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func Route() *gin.Engine {
	//usercache = make(map[string]interface{}, 0)
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("views/**/*")
	//加载配置文件

	r.Use(MiddleWare())
	{
		r.GET("/", controllers.Index)
		r.GET("/norole", controllers.NoRole)
		r.GET("/api/init.json", controllers.MenuInit)

		//用户
		r.Any("/user/list", controllers.UserList)
		r.Any("/user/add", controllers.UserAdd)
		r.Any("/user/edit", controllers.UserEdit)
		r.POST("/user/del", controllers.UserDel)

		//菜单
		r.Any("/menu/list", controllers.MenuList)
		r.Any("/menu/add", controllers.MenuAdd)
		r.Any("/menu/edit", controllers.MenuEdit)
		r.POST("/menu/del", controllers.MenuDel)

		//权限管理
		r.Any("/roles/list", controllers.RolesList)
		r.Any("/roles/add", controllers.RolesAdd)
		r.Any("/roles/edit", controllers.RolesEdit)
		r.POST("/roles/del", controllers.RolesDel)
		r.Any("/roles/roles", controllers.RolesRoles)

		//分类管理
		r.Any("/group/list", controllers.GroupList)
		r.Any("/group/add", controllers.GroupAdd)
		r.Any("/group/edit", controllers.GroupEdit)
		r.POST("/group/del", controllers.GroupDel)

		//任务管理
		r.Any("/task/list", controllers.TaskList)
		r.Any("/task/add", controllers.TaskAdd)
		r.Any("/task/edit", controllers.TaskEdit)
		r.POST("/task/del", controllers.TaskDel)
		r.POST("/task/start", controllers.TaskStart)

		r.Any("/log/list", controllers.LogList)
		r.GET("/log/detail", controllers.LogDetail)

		//侧边栏
		r.GET("/page/welcome", controllers.Welcome)
		r.GET("/json", controllers.Sjson)

	}

	//help nolog
	r.GET("/help/index", controllers.Help)
	r.GET("/log/del", controllers.DelLogByTime)
	//重置task表中的 iplong
	r.GET("/help/resetiplong", controllers.RetsetIpLong)

	r.Any("/login", controllers.Login)
	r.GET("/loginout", controllers.LoginOut)
	r.GET("/captcha/image", controllers.Captchimage)

	return r
}

/**
不需要登录就可执行 配置  比较懒就卸载这里了
*/
func NoNeedLogUrl() map[string]bool {
	NoneedMap := make(map[string]bool, 0)

	NoneedMap["help/index"] = true
	NoneedMap["log/del"] = true
	NoneedMap["loginout"] = true
	NoneedMap["login"] = true
	NoneedMap["captcha/image"] = true
	NoneedMap["help/resetiplong"] = true

	return NoneedMap

}

/**
检测是否为不需要登录地址
*/
func CheckNoNeedLogUrl(uri string) bool {
	trimUrl := strings.Trim(uri, "/")
	NoneedMap := NoNeedLogUrl()
	_, ok := NoneedMap[trimUrl]
	if ok {
		return true
	}
	return false
}

/**
验证全局中间件  后期会验证权限的  会单独提出来的
*/
func MiddleWare() gin.HandlerFunc {
	f := func(c *gin.Context) {
		cookie, err := c.Cookie("auth")
		if err == nil {
			userinfo := strings.Split(cookie, "|")
			c.Set("username", userinfo[0])
			c.Set("user_id", userinfo[1])

			//有条件应该缓存下来
			user, err := models.GetUserInfoById(userinfo[1])
			if err != nil {
				c.Redirect(http.StatusFound, "/login")
			}
			Btye, errs := json.Marshal(&user)
			if errs != nil {
				fmt.Println(errs)
			}
			c.Set("userinfo", string(Btye))
			c.Set("user_id", user.Id)
			status := CheckRole(user.RoleId, c.Request.RequestURI)

			if !status {
				if c.Request.Method == "POST" {
					msg := until.Return(until.NoRole, "", "")
					c.JSON(http.StatusUnauthorized, msg)
				} else {
					msg := gin.H{
						"error": "Unauthorized",
					}
					c.HTML(http.StatusUnauthorized, "norole.html", msg)
				}
				c.Abort()
				return
			}

		} else {
			url := ""
			URL := c.Request.RequestURI
			mt := strings.Index(URL, "?")
			if mt > 0 {
				url = URL[0:mt]
			} else {
				url = URL
			}
			// 检测是否需要登录
			b := CheckNoNeedLogUrl(url)
			if b {
				return
			}
			if url != "/login" && url != "/captcha/image" && url != "/user/login" && url != "/loginout" {
				t := time.Now().Unix()
				times := strconv.FormatInt(t, 10)
				//跳转不同地址防止301被浏览器缓存
				c.Redirect(http.StatusFound, "/login?time="+times)
			}
		}
		return
	}
	return f
}

func CheckRole(RoleId int, URL string) bool {
	if URL == "/" {
		return true
	}
	urls := ""
	mt := strings.Index(URL, "?")
	if mt > 0 {
		urls = URL[0:mt]
	} else {
		urls = URL
	}
	//不需要验证
	b := CheckNoNeedLogUrl(urls)
	if b {
		return true
	}
	currentUrl := strings.Trim(urls, "/")
	//必须登录当时不需要权限地址
	if currentUrl == "page/welcome" || currentUrl == "api/init.json" || currentUrl == "norole" {
		return true
	}
	roleMenu, err := models.GetAllRolesMenu(RoleId)
	if err != nil {
		fmt.Println(err)
	}
	if len(roleMenu) == 0 {
		return false
	}
	listmenu := make([]int, 0)
	for _, v := range roleMenu {
		listmenu = append(listmenu, v.MenuId)
	}
	menus, err := models.GetAllCategoryMenusRoles(listmenu)
	if err != nil {
		fmt.Println(err)
	}
	if len(menus) == 0 {
		return false
	}
	UrlMap := make(map[string]bool, 0)
	for _, v := range menus {
		if v.Uri != "" {
			vur := strings.Trim(v.Uri, "/")
			UrlMap[vur] = true
		}
	}
	if _, ok := UrlMap[currentUrl]; ok {
		return true
	}
	return false
}
