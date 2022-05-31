package controllers

import (
	"Gin/models"
	"Gin/until"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Index(c *gin.Context) {
	user, bool := until.GetStrInterfaceToString(c)
	if bool {
		fmt.Println(bool)
	}

	Hdata := gin.H{
		"title": "我是测试",
		"ce":    "123456",
		"user":  user,
	}
	c.HTML(http.StatusOK, "index.html", Hdata)
}

func MenuInit(c *gin.Context) {
	//	initjson := `{"homeInfo":{"title":"首页","href":"page/welcome"},
	//"logoInfo":{"title":"LAYUI MINI","image":"/assets/images/logo.png","href":""},
	//"menuInfo":[{"title":"常规管理","icon":"fa fa-address-book","href":"","target":"_self",
	//"child":[
	//{"title":"主页模板","href":"","icon":"fa fa-home","target":"_self",
	//"child":[{"title":"主页一","href":"page/welcome","icon":"fa fa-tachometer","target":"_self"},
	//{"title":"用户列表","href":"user/list","icon":"fa fa-tachometer","target":"_self"},
	//{"title":"菜单管理","href":"/menu/list","icon":"fa fa-tachometer","target":"_self"}]},
	//{"title":"菜单管理","href":"/menu/list","icon":"fa fa-window-maximize","target":"_self"},
	//{"title":"系统设置","href":"page/setting.html","icon":"fa fa-gears","target":"_self"},
	//{"title":"表格示例","href":"page/table.html","icon":"fa fa-file-text","target":"_self"},
	//{"title":"表单示例","href":"","icon":"fa fa-calendar","target":"_self",
	//"child":[{"title":"普通表单","href":"page/form.html","icon":"fa fa-list-alt","target":"_self"},
	//{"title":"分步表单","href":"page/form-step.html","icon":"fa fa-navicon","target":"_self"}]},
	//{"title":"登录模板","href":"","icon":"fa fa-flag-o","target":"_self",
	//"child":[{"title":"登录-1","href":"page/login-1.html","icon":"fa fa-stumbleupon-circle","target":"_blank"},
	//{"title":"登录-2","href":"page/login-2.html","icon":"fa fa-viacoin","target":"_blank"},
	//{"title":"登录-3","href":"page/login-3.html","icon":"fa fa-tags","target":"_blank"}]},
	//{"title":"异常页面","href":"","icon":"fa fa-home","target":"_self",
	//"child":[{"title":"404页面","href":"page/404.html","icon":"fa fa-hourglass-end","target":"_self"}]},
	//{"title":"其它界面","href":"","icon":"fa fa-snowflake-o","target":"",
	//"child":[{"title":"按钮示例","href":"page/button.html","icon":"fa fa-snowflake-o","target":"_self"},
	//{"title":"弹出层","href":"page/layer.html","icon":"fa fa-shield","target":"_self"}]}]},{"title":"组件管理","icon":"fa fa-lemon-o","href":"","target":"_self","child":[{"title":"图标列表","href":"page/icon.html","icon":"fa fa-dot-circle-o","target":"_self"},{"title":"图标选择","href":"page/icon-picker.html","icon":"fa fa-adn","target":"_self"},{"title":"颜色选择","href":"page/color-select.html","icon":"fa fa-dashboard","target":"_self"},{"title":"下拉选择","href":"page/table-select.html","icon":"fa fa-angle-double-down","target":"_self"},{"title":"文件上传","href":"page/upload.html","icon":"fa fa-arrow-up","target":"_self"},{"title":"富文本编辑器","href":"page/editor.html","icon":"fa fa-edit","target":"_self"},{"title":"省市县区选择器","href":"page/area.html","icon":"fa fa-rocket","target":"_self"}]},{"title":"其它管理","icon":"fa fa-slideshare","href":"","target":"_self","child":[{"title":"多级菜单","href":"","icon":"fa fa-meetup","target":"","child":[{"title":"按钮1","href":"page/button.html?v=1","icon":"fa fa-calendar","target":"_self","child":[{"title":"按钮2","href":"page/button.html?v=2","icon":"fa fa-snowflake-o","target":"_self","child":[{"title":"按钮3","href":"page/button.html?v=3","icon":"fa fa-snowflake-o","target":"_self"},{"title":"表单4","href":"page/form.html?v=1","icon":"fa fa-calendar","target":"_self"}]}]}]},
	//{"title":"失效菜单","href":"page/error.html","icon":"fa fa-superpowers","target":"_self"}]}]}`

	user, bool := until.GetStrInterfaceToString(c)
	if bool {
		fmt.Println(bool)
	}

	RoleId := user.RoleId

	menuInterface := make(map[string]interface{}, 0)

	Title := make(map[string]string, 0)
	Title["title"] = "首页"
	Title["href"] = "page/welcome"

	LogoInfo := make(map[string]string, 0)
	LogoInfo["title"] = "CRONTASK"
	LogoInfo["image"] = "/assets/images/logo.png"
	LogoInfo["href"] = ""

	menuinfo := make([]map[string]interface{}, 0)

	menuinfoslice := make(map[string]interface{}, 0)

	//常规
	menuinfoCategory := GetUserMenuCategory(RoleId)

	menuinfoslice["title"] = "常规管理"
	menuinfoslice["icon"] = "fa fa-address-book"
	menuinfoslice["href"] = ""
	menuinfoslice["target"] = "_self"

	menuinfoslice["child"] = menuinfoCategory

	menuinfo = append(menuinfo, menuinfoslice)

	menuInterface["homeInfo"] = Title
	menuInterface["logoInfo"] = LogoInfo
	menuInterface["menuInfo"] = menuinfo

	jsbyte, er := json.Marshal(&menuInterface)
	if er != nil {
		fmt.Println(er)
	}
	//fmt.Println(string(jsbyte))

	c.String(http.StatusOK, string(jsbyte))
}

/**
获取该权限下的 栏目
*/
func GetUserMenuCategory(roleId int) []map[string]interface{} {
	roleMenu, err := models.GetAllRolesMenu(roleId)
	if err != nil {
		fmt.Println(err)
	}

	listmenu := make([]int, 0)
	for _, v := range roleMenu {
		listmenu = append(listmenu, v.MenuId)
	}
	//一级栏目
	catemenus, _ := models.GetAllCategoryMenus(listmenu, 1)
	menuinfoCategory := make([]map[string]interface{}, 0)
	//一级栏目
	twomenus, _ := models.GetAllCategoryMenus(listmenu, 2)
	//fmt.Println(twomenus)
	for _, v := range catemenus {
		menuinfoCategorySlice := make(map[string]interface{}, 0)
		menuinfoCategorySlice["title"] = v.Title
		menuinfoCategorySlice["href"] = ""
		menuinfoCategorySlice["icon"] = "fa " + v.Icon
		menuinfoCategorySlice["target"] = "_self"

		menuSlice := make([]map[string]interface{}, 0)

		for _, mv := range twomenus {
			if v.Id == mv.ParentId {
				mc := make(map[string]interface{}, 0)
				mc["title"] = mv.Title
				mc["href"] = mv.Uri
				mc["icon"] = "fa " + mv.Icon
				mc["target"] = "_self"
				menuSlice = append(menuSlice, mc)
			}

		}
		menuinfoCategorySlice["child"] = menuSlice
		menuinfoCategory = append(menuinfoCategory, menuinfoCategorySlice)
	}

	return menuinfoCategory

}

func Welcome(c *gin.Context) {
	userId := c.GetInt("user_id")
	//任务数量
	TaskCount := models.GetTaskCountByUserId(userId)

	taskLogCount := models.GetTaskRunCount(userId, 0)
	successLogCount := models.GetTaskRunCount(userId, 1)
	errorLogCount := taskLogCount - successLogCount
	//最近7天发送量

	linedata := GetSevenData()

	Hdata := gin.H{
		"user":            userId,
		"taskcount":       TaskCount,
		"taskLogCount":    taskLogCount,
		"successLogCount": successLogCount,
		"errorLogCount":   errorLogCount,
		"linedata":        linedata,
	}

	c.HTML(http.StatusOK, "welcome.html", Hdata)
}

func NoRole(c *gin.Context) {
	Hdata := gin.H{"title": "我是测试", "ce": "123456",
	}
	c.HTML(http.StatusOK, "norole.html", Hdata)
}

func GetSevenData() string {
	t := time.Now()
	DateTime := t.Format("2006-01-02 15:04:05")
	DateSice := DateTime[0:10] + " 00:00:00"
	p, _ := time.Parse("2006-01-02 15:04:05", DateSice)
	v := p.Unix() - 7*86400
	ts := time.Unix(v, 0)
	date := ts.Format("2006-01-02 15:04:05")
	dateData := models.GetStatistic(date, 0)
	formatDateData := formatData(dateData)
	dateDataSuccess := models.GetStatistic(date, 1)
	dateDataErr := models.GetStatistic(date, 2)
	formatdateDataSuccess := formatData(dateDataSuccess)
	formatdateDataErr := formatData(dateDataErr)

	var total []int
	var totalSuccess []int
	var totalError []int
	dateslice := GetDataSlice()
	for _, v := range dateslice {
		if _, ok := formatDateData[v]; ok {
			total = append(total, formatDateData[v])
		} else {
			total = append(total, 0)
		}
		if _, ok := formatdateDataSuccess[v]; ok {
			totalSuccess = append(totalSuccess, formatdateDataSuccess[v])
		} else {
			totalSuccess = append(totalSuccess, 0)
		}
		if _, ok := formatdateDataErr[v]; ok {
			totalError = append(totalError, formatdateDataErr[v])
		} else {
			totalError = append(totalError, 0)
		}

	}
	lines := make(map[string]interface{})
	lines["date"] = dateslice
	lines["all"] = total
	lines["succ"] = totalSuccess
	lines["err"] = totalError
	jb, err := json.Marshal(lines)
	if err != nil {
		fmt.Println(err)
	}
	return string(jb)
}

func formatData(data []models.StatisticCount) map[string]int {
	fdata := make(map[string]int, 0)
	for _, v := range data {
		datetime := v.Datetime[0:10]
		fdata[datetime] = v.Count
	}
	return fdata
}
/**

 */
func GetDataSlice() []string {
	t := time.Now()
	DateTime := t.Format("2006-01-02 15:04:05")
	DateSice := DateTime[0:10] + " 00:00:00"
	p, _ := time.Parse("2006-01-02 15:04:05", DateSice)
	v := p.Unix()
	dateslice := make([]string, 0)
	for i := 7; i >= 0; i-- {
		n := v - int64(i*86400)
		ts := time.Unix(n, 0)
		date := ts.Format("2006-01-02")
		dateslice = append(dateslice, date)
	}
	return dateslice
}
