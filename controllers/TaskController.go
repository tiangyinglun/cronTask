package controllers

import (
	"Gin/Job"
	"Gin/models"
	"Gin/until"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jakecoffman/cron"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type TaskSearch struct {
	TaskName string `json:"task_name"`
	GroupId  string `json:"group_id"`
}

func TaskList(c *gin.Context) {
	if c.Request.Method == "POST" {
		pagestr := c.DefaultPostForm("page", "1")
		limitstr := c.DefaultPostForm("limit", "15")
		searchParams := c.PostForm("searchParams")
		taskname := ""
		group_id := ""
		var group_id_int int
		page := 1
		limit := 0
		limit, _ = strconv.Atoi(limitstr)
		page, _ = strconv.Atoi(pagestr)
		page = page - 1
		if searchParams != "" {
			params := TaskSearch{}
			fmt.Println(searchParams)
			err := json.Unmarshal([]byte(searchParams), &params)
			if err != nil {
				until.LogObj().WithFields(logrus.Fields{
					"func": "TaskList",
				}).Error("获取参数：", err)
			}
			taskname = params.TaskName
			group_id = params.GroupId
			group_id_int, _ = strconv.Atoi(group_id)

		}
		user_id := c.GetInt("user_id")

		result, count, err := models.GetTaskList(taskname, user_id, group_id_int, page, limit, "1")
		if err != nil {
			fmt.Println(err)
		}
		group, _ := models.GetAllTaskGroup(user_id)
		groupMap := make(map[int]string, len(group))
		for _, v := range group {
			groupMap[v.Id] = v.GroupName
		}
		Adulist := make([]models.AdminTaskList, len(result))
		for k, v := range result {
			Adulist[k].TaskName = v.TaskName
			Adulist[k].Status = v.Status
			if v.Status == 0 {
				Adulist[k].TaskStatus = `<i class="fa fa-pause reds"></i>暂停`
			} else {
				Adulist[k].TaskStatus = `<i class="fa fa-play greens"></i>开启`
			}
			Adulist[k].CronSpec = v.CronSpec
			Adulist[k].UpdatedAt = until.TimeLocateToFormat(v.UpdatedAt)
			Adulist[k].PrevTime = v.PrevTime
			Adulist[k].Id = v.Id
			Adulist[k].TaskType = v.TaskType
			Adulist[k].HttpUrl = v.HttpUrl
			Adulist[k].UserId = v.UserId
			Adulist[k].PrevTimeDate = until.TimestampString(v.PrevTime)
			Adulist[k].GroupName = groupMap[v.GroupId]
		}
		data := until.ReturnCount(until.Success, count, Adulist, "")
		c.JSON(http.StatusOK, data)
		return
	}
	user_id := c.GetInt("user_id")
	group, _ := models.GetAllTaskGroup(user_id)
	data := gin.H{
		"group": group,
	}
	c.HTML(http.StatusOK, "tasklist.html", data)
}

func TaskAdd(c *gin.Context) {
	if c.Request.Method == "POST" {

		defer func() {
			if err := recover(); err != nil {
				until.LogObj().WithFields(logrus.Fields{
					"func": "Run",
				}).Error("获取参数：", err)
				data := until.Return(until.TimeParse, "", "")
				c.JSON(http.StatusOK, data)
				return
			}
		}()

		task_name := c.PostForm("task_name")
		task_type := c.PostForm("task_type")
		http_url := c.PostForm("http_url")
		cron_spec := c.PostForm("cron_spec")
		host := c.PostForm("host")
		command := c.PostForm("command")
		group_id := c.PostForm("group_id")
		timeout := c.PostForm("timeout")
		notify := c.PostForm("notify")
		notify_email := c.PostForm("notify_email")

		result := ""
		u := models.AdminTask{}
		if task_name == "" {
			data := until.Return(until.Success, result, "任务名不可以为空")
			c.JSON(http.StatusOK, data)
			return
		}

		cron.Parse(cron_spec)

		timeoutInt, _ := strconv.Atoi(timeout)
		groupidInt, _ := strconv.Atoi(group_id)

		u.UserId = c.GetInt("user_id")
		u.TaskName = task_name
		u.TaskType = task_type
		u.HttpUrl = http_url
		notifyInt, _ := strconv.Atoi(notify)
		u.Command = command
		u.NotifyEmail = notify_email
		u.Iplong = until.GetCmdIpLong()

		u.Notify = notifyInt
		u.Timeout = timeoutInt
		u.GroupId = groupidInt
		u.Host = host
		u.CronSpec = cron_spec

		u.CreatedAt = until.GetNowTime()
		u.UpdatedAt = until.GetNowTime()
		b, _ := models.TaskAdd(u)
		if !b {
			data := until.Return(until.Success, result, "添加失败")
			c.JSON(http.StatusOK, data)
			return
		}

		data := until.Return(until.Success, result, "")
		c.JSON(http.StatusOK, data)
		return
	}
	user_id := c.GetInt("user_id")
	group, err := models.GetAllTaskGroup(user_id)
	if err != nil {
		fmt.Println(err)
	}

	data := gin.H{
		"group": group,
	}
	c.HTML(http.StatusOK, "taskadd.html", data)
}

/**
用户编辑
*/
func TaskEdit(c *gin.Context) {
	if c.Request.Method == "POST" {

		defer func() {
			if err := recover(); err != nil {
				until.LogObj().WithFields(logrus.Fields{
					"func": "Run",
				}).Error("获取参数：", err)
				data := until.Return(until.TimeParse, "", "")
				c.JSON(http.StatusOK, data)
				return
			}
		}()

		task_name := c.PostForm("task_name")
		task_type := c.PostForm("task_type")
		http_url := c.PostForm("http_url")
		cron_spec := c.PostForm("cron_spec")
		host := c.PostForm("host")
		command := c.PostForm("command")
		group_id := c.PostForm("group_id")
		id := c.PostForm("id")

		timeout := c.PostForm("timeout")
		notify := c.PostForm("notify")
		notify_email := c.PostForm("notify_email")

		result := ""
		u := models.AdminTask{}
		if task_name == "" {
			data := until.Return(until.Success, result, "任务名不可以为空")
			c.JSON(http.StatusOK, data)
			return
		}
		//验证 时间表达式
		cron.Parse(cron_spec)
		timeoutInt, _ := strconv.Atoi(timeout)
		groupidInt, _ := strconv.Atoi(group_id)

		update := make(map[string]interface{})
		notifyInt, _ := strconv.Atoi(notify)

		u.UserId = c.GetInt("user_id")
		update["user_id"] = c.GetInt("user_id")
		update["task_name"] = task_name
		update["task_type"] = task_type
		update["http_url"] = http_url
		update["notify"] = notifyInt
		update["command"] = command
		update["iplong"] = until.GetCmdIpLong()
		update["status"] = 0
		update["timeout"] = timeoutInt
		update["group_id"] = groupidInt
		update["host"] = host
		update["cron_spec"] = cron_spec
		update["updated_at"] = until.GetNowTime()
		update["notify_email"] = notify_email

		b, _ := models.TaskEdit(id, update)
		if !b {
			data := until.Return(until.Success, result, "编辑失败")
			c.JSON(http.StatusOK, data)
			return
		}
		//删除 任务
		Job.RemoveJob(id)

		data := until.Return(until.Success, result, "")
		c.JSON(http.StatusOK, data)
		return
	}
	id := c.Query("id")
	user_id := c.GetInt("user_id")
	group, err := models.GetAllTaskGroup(user_id)
	if err != nil {
		fmt.Println(err)
	}
	task, err := models.GetTaskInfoById(id)
	if err != nil {
		fmt.Println(err)
	}
	data := gin.H{
		"group": group,
		"task":  task,
	}
	fmt.Println(task)
	c.HTML(http.StatusOK, "taskedit.html", data)
}

/**
del
*/
func TaskDel(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		data := until.Return(until.Success, "", "id 不可以为空")
		c.JSON(http.StatusOK, data)
	}
	bool, err := models.TaskDel(id)
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
		//删除任务 就停止 任务
		Job.RemoveJob(id)
		data := until.Return(until.Success, "", "")
		c.JSON(http.StatusOK, data)
		return
	}
	data := until.Return(until.Error, "", "")
	c.JSON(http.StatusOK, data)
}

func TaskStart(c *gin.Context) {

	id := c.PostForm("id")
	if id == "" {
		data := until.Return(until.Success, "", "id 不可以为空")
		c.JSON(http.StatusOK, data)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			until.LogObj().WithFields(logrus.Fields{
				"func": "Run",
			}).Error("获取参数：", err)
			data := until.Return(until.TimeParse, "", "")
			c.JSON(http.StatusOK, data)
			return
		}
	}()

	adm, err := models.GetTaskInfoById(id)
	if err != nil {
		data := until.Return(until.Error, "", "id 不存在")
		c.JSON(http.StatusOK, data)
		return
	}
	if adm.Id == 0 {
		data := until.Return(until.Error, "", "id 不存在")
		c.JSON(http.StatusOK, data)
		return
	}
	status := adm.Status
	updateStatus := 0
	attrs := ""
	if status == 1 {
		attrs = `<i class="fa fa-pause reds"></i>暂停`
		updateStatus = 0
	} else {
		attrs = `<i class="fa fa-play greens"></i>开启`
		updateStatus = 1
	}

	//如果格式化失败 就停止了 如果报错了 就异常
	cron.Parse(adm.CronSpec)

	TaskUpdate := models.AdminTask{}
	TaskUpdate.Status = updateStatus
	TaskUpdate.UpdatedAt = until.GetNowTime()
	bool, _ := models.TaskEditStatus(id, TaskUpdate)
	fmt.Println(bool)
	if bool {
		if updateStatus == 0 {
			//关闭定时
			Job.RemoveJob(id)
		} else {
			//检测是否存在
			ljob := Job.GetEntryById(id)
			if ljob != nil {
				Job.RemoveJob(id)
			}
			TaskJob, _ := models.GetTaskInfoById(id)
			job, err := Job.NewJobFromTask(&TaskJob)
			if err != nil {
				until.LogObj().WithFields(logrus.Fields{
					"func": "start",
					"id":   id,
				}).Error("获取参数：", err)
			}
			Job.AddJob(TaskJob.CronSpec, job)

		}
		data := until.Return(until.Success, "", attrs)
		c.JSON(http.StatusOK, data)
		return
	}
	data := until.Return(until.Success, "", "id 不存在")
	c.JSON(http.StatusOK, data)
	return
}



