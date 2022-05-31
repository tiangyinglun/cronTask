package models

import (
	"fmt"
	"time"
)

type AdminTaskLog struct {
	Id          int64  `json:"id"`
	UserId      int    `json:"user_id"`
	TaskId      int    `json:"task_id"`
	Output      string `json:"output"`
	Httpcode    int    `json:"httpcode"`
	Error       string `json:"error"`
	Status      int    `json:"status"`
	ProcessTime int    `json:"process_time"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
type AdminTaskLogList struct {
	AdminTaskLog
	TaskName string `json:"task_name"`
}

type StatisticCount struct {
	Datetime string `json:"datetime"`
	Count    int    `json:"count"`
}

/**
添加用户
*/
func LogAdd(u AdminTaskLog) (id int64, err error) {
	result := PDO.Create(&u)
	if result.Error != nil {
		return 0, err
	}

	return u.Id, result.Error
}

/**
获取日志
*/
func GetLogList(user_id, httpcode int, task_id, page, size int) (adtlog []AdminTaskLog, count int64, err error) {
	u := []AdminTaskLog{}
	offset := page * size
	Ps := PDO.Order("id desc")

	if task_id != 0 {
		Ps = Ps.Where("task_id=?", task_id)
	}
	if user_id != 0 {
		Ps = Ps.Where("user_id=?", user_id)
	}

	if httpcode != 0 {
		Ps = Ps.Where("httpcode=?", httpcode)
	}

	retcount := Ps.Model(&AdminTaskLog{}).Count(&count)
	if retcount.Error != nil {
		fmt.Println(retcount)
	}

	result := Ps.Limit(size).Offset(offset).Find(&u)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	return u, count, result.Error
}

/**
获取详情
*/
func GetLogInfoById(id string) (Adlog AdminTaskLog, err error) {
	u := AdminTaskLog{}
	result := PDO.Where("id=?", id).Find(&u)
	if result.Error != nil {
		return u, result.Error
	}
	return u, result.Error
}

/**
删除
*/
func DelLogByDateTime(day int64) (b bool, err error) {
	now := time.Now().Unix()

	Deletetime := now - day*86400

	tm := time.Unix(Deletetime, 0)
	DateTime := tm.Format("2006-01-02 15:04:05")
	result := PDO.Where("created_at<?", DateTime).Delete(&AdminTaskLog{})
	if result.Error != nil {
		return false, err
	}
	return true, result.Error
}

/**

 */
func GetTaskRunCount(userId int, types int) (count int64) {
	PS := PDO.Model(&AdminTaskLog{}).Where("user_id=?", userId)
	if types == 1 {
		PS = PS.Where("httpcode=?", 200)
	}
	PS.Count(&count)
	return
}

func GetStatistic(start string, Status int) []StatisticCount {
	var static []StatisticCount
	Sql := "SELECT count(*) as count,DATE(created_at) as datetime  FROM admin_task_log where created_at>='" + start + "'"
	if Status == 1 {
		Sql += " and httpcode=200 "
	} else if Status == 2 {
		Sql += " and httpcode !=200 "
	}
	Sql += " group by DATE(created_at) "
	PDO.Raw(Sql).Scan(&static)
 
	return static
}
