package models

import "fmt"

type AdminMenu struct {
	Id         int    `json:"id"`
	ParentId   int    `json:"parent_id"`
	Orders     int    `json:"orders"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	IsCategory int    `json:"is_category"`
	Uri        string `json:"uri"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type AdminMenuList struct {
	AdminMenu
	ParentName   string `json:"parent_name"`
	CategoryName string `json:"category_name"`
}

type AdminMenuTree struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parent_id"`
	Title    string `json:"title"`
	Count    int    `json:"count"`
	Color    string `json:"color"`
	Checked  bool   `json:"checked"`
	Children []AdminMenuTree
}

func GetMenus() []AdminMenuTree {
	u := []AdminMenu{}
	result := PDO.Where(" is_category =? ", 1).Find(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	Tree := GetMenusTree(u, 0)
	return Tree
}

func GetAllMenus() []AdminMenuTree {
	u := []AdminMenu{}
	result := PDO.Find(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	Tree := GetMenusTree(u, 0)

	TreeCount := make([]AdminMenuTree, len(Tree))

	for k, v := range Tree {
		TreeCount[k].Count = len(v.Children)
		TreeCount[k].Id = v.Id
		TreeCount[k].Children = v.Children
		TreeCount[k].ParentId = v.ParentId
		TreeCount[k].Title = v.Title
		if k%2 == 1 {
			TreeCount[k].Color = "#f0f0f0"
		} else {
			TreeCount[k].Color = "#e7e4e4"
		}
	}
	return TreeCount
}

/**
获取 菜单树
*/
func GetMenusTree(menuList []AdminMenu, pid int) []AdminMenuTree {
	treeList := []AdminMenuTree{}
	for _, v := range menuList {
		if v.ParentId == pid {
			child := GetMenusTree(menuList, v.Id)
			node := AdminMenuTree{
				Id:       v.Id,
				Title:    v.Title,
				ParentId: v.ParentId,
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}
	return treeList
}

/**
获取栏目
*/
func GetMenuByCategorByStatus(is_category int) (Adm []AdminMenu, err error) {
	u := []AdminMenu{}
	result := PDO.Where(" is_category =? ", is_category).Find(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

/**
获取栏目
*/
func GetMenuInfoById(id string) (Adm AdminMenu, err error) {
	u := AdminMenu{}
	result := PDO.Where(" id =? ", id).First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

/**
栏目列表
*/
func GetMenuList(title string, page, limit int) (Adm []AdminMenu, count int64, err error) {
	u := []AdminMenu{}
	offset := page * limit
	Ps := PDO.Order(" id desc ")
	if title != "" {
		fmt.Println(title)
		Ps = Ps.Where(" title like ?", "%"+title+"%")
	}
	Ps.Model(&AdminMenu{}).Count(&count)
	result := Ps.Limit(limit).Offset(offset).Find(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, count, result.Error
}

/**
添加
*/
func MenuAdd(u AdminMenu) (b bool, err error) {
	result := PDO.Create(&u)
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

/**
更新
*/
func MenuEdit(id string, u AdminMenu) (b bool, err error) {
	result := PDO.Model(&AdminMenu{}).Select("is_category", "updated_at", "uri", "icon", "parent_id", "orders", "title").Where(" id=? ", id).Updates(u)
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

/**
删除
*/
func MenuDel(id string) (b bool, err error) {
	u := AdminMenu{}
	result := PDO.Where(" id = ? ", id).Delete(&u)
	if result.Error != nil {
		return false, result.Error
	}
	return true, result.Error
}

/**
in 查询  只差栏目的就行
*/
func GetMenusInMenuId(menuid []int) (Adm []AdminMenu, err error) {
	u := []AdminMenu{}
	result := PDO.Where(" is_category =? ", 1).Where(" id in ?", menuid).Find(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return u, result.Error
}

/**
获得 栏目
*/
func GetAllCategoryMenus(menuid []int, types int) (Adm []AdminMenu, err error) {
	u := []AdminMenu{}
	PS := PDO.Where(" is_category =? ", 1)
	if types == 1 {
		PS = PS.Where(" parent_id =?", 0)
	} else {
		PS = PS.Where(" parent_id !=?", 0)
	}
	result := PS.Where(" id in ? ", menuid).Order(" orders desc ").Find(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

func GetAllCategoryMenusRoles(menuid []int) (Adm []AdminMenu, err error) {
	u := []AdminMenu{}
	result := PDO.Where("parent_id !=?", 0).Where(" id in ? ", menuid).Find(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

/**
查询个数
*/
func GetMenusCount(Uri string, id int) (Adm AdminMenu, err error) {
	u:=AdminMenu{}
	PS := PDO.Where(" uri=? ", Uri).Where("parent_id!=? ", 0)
	if id != 0 {
		PS = PS.Where(" id !=?", id)
	}

	result := PS.First(&u)
	if result.Error != nil {
		return u, result.Error
	}
	return u, result.Error
}
