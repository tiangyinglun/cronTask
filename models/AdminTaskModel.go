package models

import (
	"fmt"
)

type AdminTask struct {
	Id           int    `json:"id"`
	UserId       int    `json:"user_id"`
	GroupId      int    `json:"group_id"`
	TaskName     string `json:"task_name"`
	TaskType     string `json:"task_type"`
	HttpUrl      string `json:"http_url"`
	Iplong       int    `json:"iplong"`
	Host         string `json:"host"`
	Description  string `json:"description"`
	CronSpec     string `json:"cron_spec"`
	Concurrent   int    `json:"concurrent"`
	Command      string `json:"command"`
	Status       int    `json:"status"`
	Notify       int    `json:"notify"`
	NotifyEmail  string `json:"notify_email"`
	Timeout      int    `json:"timeout"`
	ExecuteCount int    `json:"execute_count"`
	PrevTime     int64  `json:"prev_time"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

const (
	TASK_SUCCESS = 0  // 任务执行成功
	TASK_ERROR   = -1 // 任务执行出错
	TASK_TIMEOUT = -2 // 任务执行超时
)

type AdminTaskList struct {
	AdminTask
	PrevTimeDate string `json:"prev_time_date"`
	GroupName    string `json:"group_name"`
	TaskStatus   string `json:"task_status"`
}

/**
查询用户
*/
func GetTaskInfo(username string) (user AdminTask, err error) {
	u := AdminTask{}
	result := PDO.Where("task_name=?", username).First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

func GetTaskByUserName(taskname string) (user AdminTask, err error) {
	u := AdminTask{}
	result := PDO.Where("task_name=?", taskname).First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

/**
根据id 查询用户
*/
func GetTaskInfoById(id string) (user AdminTask, err error) {
	u := AdminTask{}
	result := PDO.Where(" id =? ", id).First(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, result.Error
}

/**
获取用户列表
*/
func GetTaskList(name string, userid, group_id, page, size int, sort string) (user []AdminTask, count int64, err error) {
	u := []AdminTask{}
	offset := page * size
	Ps := PDO.Where("user_id=?", userid).Order("id desc")
	if name != "" {
		Ps = Ps.Where("task_name like ?", "%"+name+"%")
	}
	if group_id != 0 {
		Ps = Ps.Where("group_id=?", group_id)
	}

	retcount := Ps.Model(&AdminTask{}).Count(&count)
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
必须使用ip标识
*/
func GetAllTaskList(iplong, limit int) (Adm *[]AdminTask, err error) {
	u := []AdminTask{}
	result := PDO.Where("status=?", 1).Where("iplong=?", iplong).Limit(limit).Find(&u)
	return &u, result.Error
}

/**
添加用户
*/
func TaskAdd(u AdminTask) (b bool, err error) {
	result := PDO.Create(&u)
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

/**
用户更新
*/
func TaskEdit(id interface{}, u map[string]interface{}) (b bool, err error) {
	result := PDO.Model(&AdminTask{}).Where(" id=? ", id).Updates(u)
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

/**
用户更新
*/
func TaskEditStatus(id interface{}, u AdminTask) (b bool, err error) {
	result := PDO.Model(&AdminTask{}).Where(" id=? ", id).Select("status", "updated_at").Updates(u)
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

/**
删除任务
*/
func TaskDel(id string) (b bool, err error) {
	result := PDO.Where(" id = ? ", id).Delete(&AdminTask{})
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

func GetAllTask(userid int) (Adm []AdminTask, err error) {
	u := []AdminTask{}
	ps:=PDO.Where("1=1")
	if userid>0{
		ps=ps.Where("user_id=?",userid)
	}
	result := ps.Find(&u)

	return u, result.Error
}


func GetTaskCountByUserId(userId int) (count int64) {
	PDO.Model(&AdminTask{}).Where("user_id=?", userId).Count(&count)
	return
}

func TaskIpLongUpdate(iplong int)(b bool, err error)   {
	update:=make(map[string]interface{})
	update["iplong"]=iplong
	update["updated_at"]= GetNowTime()
	result := PDO.Model(&AdminTask{}).Where("id>1").Updates(update)
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}