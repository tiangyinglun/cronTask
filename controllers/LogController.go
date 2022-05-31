package controllers

import (
	"Gin/models"
	"Gin/until"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

type LogSearch struct {
	Httpcode string `json:"httpcode"`
	TaskId   string `json:"task_id"`
}

func LogList(c *gin.Context) {
	if c.Request.Method == "POST" {
		pagestr := c.DefaultPostForm("page", "1")
		limitstr := c.DefaultPostForm("limit", "15")
		taskIs := c.DefaultPostForm("task_id", "0")
		searchParams := c.PostForm("searchParams")
		task_id := ""
		httpcode := ""
		var task_id_int int
		var httcode_int int
		page := 1
		limit := 0
		limit, _ = strconv.Atoi(limitstr)
		page, _ = strconv.Atoi(pagestr)
		page = page - 1
		if searchParams != "" {
			params := LogSearch{}
			fmt.Println(searchParams)
			err := json.Unmarshal([]byte(searchParams), &params)
			if err != nil {
				until.LogObj().WithFields(logrus.Fields{
					"func": "TaskList",
				}).Error("获取参数：", err)
			}
			httpcode = params.Httpcode
			task_id = params.TaskId
			task_id_int, _ = strconv.Atoi(task_id)
			httcode_int, _ = strconv.Atoi(httpcode)
		}
		if taskIs != "0" {
			task_id_int, _ = strconv.Atoi(taskIs)
		}
		userid := c.GetInt("user_id")
		result, count, err := models.GetLogList(userid,httcode_int, task_id_int, page, limit)
		if err != nil {
			fmt.Println(err)
		}

		tasks, _ := models.GetAllTask(userid)
		taskMap := make(map[int]string, 0)
		for _, v := range tasks {
			taskMap[v.Id] = v.TaskName
		}
		Adulist := make([]models.AdminTaskLogList, len(result))
		for k, v := range result {
			tname := ""
			_, ok := taskMap[v.TaskId]
			if ok {
				tname = taskMap[v.TaskId]
			} else {
				tname = "已删除"
			}
			Adulist[k].TaskId = v.TaskId
			Adulist[k].Output = v.Output
			Adulist[k].TaskName = tname
			Adulist[k].Status = v.Status
			Adulist[k].Httpcode = v.Httpcode
			Adulist[k].UpdatedAt = until.TimeLocateToFormat(v.UpdatedAt)
			Adulist[k].Id = v.Id
		}
		data := until.ReturnCount(until.Success, count, Adulist, "")
		c.JSON(http.StatusOK, data)
		return
	}
	task_id := c.DefaultQuery("task_id", "0")
	userid := c.GetInt("user_id")
	tasks, _ := models.GetAllTask(userid)
	data := gin.H{
		"task":    tasks,
		"task_id": task_id,
	}
	c.HTML(http.StatusOK, "loglist.html", data)
}

/**
获取详情
*/
func LogDetail(c *gin.Context) {
	id := c.Query("id")
	log, _ := models.GetLogInfoById(id)

	task_id := strconv.Itoa(log.TaskId)
	Task, _ := models.GetTaskInfoById(task_id)
	datetime := until.TimeLocateToFormat(log.CreatedAt)
	data := gin.H{
		"id":       id,
		"log":      log,
		"task":     Task,
		"datetime": datetime,
	}
	c.HTML(http.StatusOK, "logdetail.html", data)
}

//清理日志数据
func DelLogByTime(c *gin.Context) {
	savetime := viper.GetString("log.savetime")
	savetimeint, err := strconv.ParseInt(savetime, 10, 0)
	if err != nil {
		data := until.Return(until.ParamsError, "", "删除失败")
		c.JSON(http.StatusOK, data)
		return
	}
	if savetimeint < 2 {
		data := until.Return(until.ParamsError, "", "执行日志必须保存2天以上")
		c.JSON(http.StatusOK, data)
		return
	}
	b, err := models.DelLogByDateTime(savetimeint)
	if err != nil || !b {
		data := until.Return(until.ParamsError, "", "删除失败")
		c.JSON(http.StatusOK, data)
		return
	}
	data := until.Return(until.Success, "", "")
	c.JSON(http.StatusOK, data)
}
