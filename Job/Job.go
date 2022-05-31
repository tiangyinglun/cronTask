package Job

import (
	"Gin/mail"
	"Gin/models"
	"Gin/until"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Job struct {
	id         int                                      // 任务ID
	logId      int64                                    // 日志记录ID
	name       string                                   // 任务名称
	task       *models.AdminTask                        // 任务对象
	runFunc    func(time.Duration) (string, int, error) // 执行函数
	status     int                                      // 任务状态，大于0表示正在执行中
	Concurrent bool                                     // 同一个任务是否允许并行执行

}

var mailTpl *template.Template

func init() {
	mailTpl, _ = template.New("mail_tpl").Parse(`
	你好 {{.username}}，<br/>

<p>以下是任务执行结果：</p>

<p>
任务ID：{{.task_id}}<br/>
任务名称：{{.task_name}}<br/>       
执行时间：{{.start_time}}<br />
执行耗时：{{.process_time}}秒<br />
执行状态：{{.status}}
</p>
<p>-------------以下是任务执行输出-------------</p>
<p>{{.output}}</p>
<p>
--------------------------------------------<br />
本邮件由系统自动发出，请勿回复<br />
如果要取消邮件通知，请登录到系统进行设置<br />
</p>
`)

}

func NewJobFromTask(task *models.AdminTask) (*Job, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if task.Id < 1 {
		return nil, fmt.Errorf("ToJob: 缺少id")
	}
	if task.TaskType == "POST" { //POST
		job := NewPostTJop(task)
		job.task = task
		job.Concurrent = task.Concurrent == 1
		return job, nil
	}
	//GET
	job := NewGetJop(task)
	job.task = task
	job.Concurrent = task.Concurrent == 1
	return job, nil
}

/**
get 请求
*/
func NewGetJop(task *models.AdminTask) *Job {
	job := &Job{
		id:   task.Id,
		name: task.TaskName,
	}

	job.runFunc = func(timeout time.Duration) (string, int, error) {
		client := &http.Client{}
		newurl := task.HttpUrl
		Urlhost := ""
		//替换host
		if task.Host != "" {
			Urlhost = UrlHost(task.HttpUrl)
			newurl = ReplaceHost(task.HttpUrl, Urlhost, task.Host)
		}
		//拼接参数
		if task.Command != "" {
			if strings.Contains(newurl, "?") {
				newurl += "&" + task.Command
			} else {
				newurl += "?" + task.Command
			}
		}
		req, err := http.NewRequest("GET", newurl, nil)

		if err != nil {
			fmt.Println(err.Error())
		}
		if Urlhost != "" {
			req.Host = Urlhost
		}

		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.88 Safari/537.36")
		resp, err := client.Do(req)
		if err != nil {
			errstring := fmt.Sprintf("%s", err)
			return errstring, 0, err
		}

		httpCode := resp.StatusCode
		defer resp.Body.Close()
		bufOutByte, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		bufOutString := string(bufOutByte)
		errstring := "无"
		if err != nil {
			errstring = fmt.Sprintf("%s", err)
		}
		if httpCode == http.StatusOK {
			return bufOutString, httpCode, err
		}
		return errstring, httpCode, err
	}
	return job
}

/**
post
*/
func NewPostTJop(task *models.AdminTask) *Job {
	job := &Job{
		id:   task.Id,
		name: task.TaskName,
	}
	job.runFunc = func(timeout time.Duration) (string, int, error) {
		client := &http.Client{}
		newurl := task.HttpUrl
		Urlhost := ""
		if task.Host != "" {
			Urlhost = UrlHost(task.HttpUrl)
			newurl = ReplaceHost(task.HttpUrl, Urlhost, task.Host)
		}
		payload := strings.NewReader(task.Command)
		req, err := http.NewRequest("POST", newurl, payload)
		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.88 Safari/537.36")
		if Urlhost != "" {
			req.Host = Urlhost
		}
		if err != nil {
			errstring := fmt.Sprintf("%s", err)
			return errstring, 0, err
		}
		resp, err := client.Do(req)
		defer resp.Body.Close()

		bufOutByte, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println(err2)
		}
		httpCode := resp.StatusCode

		bufOutString := string(bufOutByte)
		errstring := ""
		if err != nil {
			errstring = fmt.Sprintf("%s", err)
		}
		if httpCode == http.StatusOK {
			return bufOutString, httpCode, err
		}
		return errstring, httpCode, err
	}
	return job
}

/**
替换host
*/
func UrlHost(httpst string) (urls string) {
	u, err := url.Parse(httpst)
	if err != nil {
		panic(err)
	}
	return u.Host
}

func ReplaceHost(tarusrl, urls, host string) string {
	newUrl := strings.Replace(tarusrl, urls, host, 1)
	return newUrl
}

func (j *Job) Status() int {
	return j.status
}

func (j *Job) GetName() string {
	return j.name
}

func (j *Job) GetId() int {
	return j.id
}

func (j *Job) GetLogId() int64 {
	return j.logId
}
func (j *Job) Run() {
	if !j.Concurrent && j.status > 0 {
		fmt.Sprintf("任务[%d]上一次执行尚未结束，本次被忽略。", j.id)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			until.LogObj().WithFields(logrus.Fields{
				"func": "Run",
			}).Error("获取参数：", err)
		}
	}()

	if workPool != nil {
		workPool <- true
		defer func() {
			<-workPool
		}()
	}
	j.status++
	defer func() {
		j.status--
	}()
	t := time.Now()
	//执行超时时间 如果 0 超时时间 就24小时
	timeout := time.Duration(j.task.Timeout)
	if j.task.Timeout > 0 {
		timeout = time.Second * time.Duration(j.task.Timeout)
	} else {
		timeout = time.Second * time.Duration(24*3600)
	}
	cmdOut, httpcode, err := j.runFunc(timeout)

	ut := time.Now().Sub(t) / time.Millisecond
	log := models.AdminTaskLog{}
	log.TaskId = j.id
	log.Output = cmdOut
	log.UserId = j.task.UserId

	log.ProcessTime = int(ut)
	log.CreatedAt = until.GetNowTime()
	log.UpdatedAt = until.GetNowTime()
	log.Httpcode = httpcode
	log.Error = "无"
	if httpcode != http.StatusOK {
		log.Status = models.TASK_TIMEOUT
		log.Error = fmt.Sprintf("%s", err)
		log.Error = fmt.Sprintf("任务执行超过 %d 秒\n----------------------\n%s\n", int(timeout/time.Second), fmt.Sprintf("%s", err))
	}
	if err != nil {
		log.Status = models.TASK_ERROR
		log.Error = fmt.Sprintf("%s", err)
		log.Error = err.Error() + ":" + fmt.Sprintf("%s", err)
	}

	j.logId, _ = models.LogAdd(log)

	TaskUpdate := make(map[string]interface{})
	TaskUpdate["prev_time"] = t.Unix()
	TaskUpdate["execute_count"] = j.task.ExecuteCount + 1
	TaskUpdate["updated_at"] = until.GetNowTime()
	models.TaskEdit(j.id, TaskUpdate)
	// 发送邮件通知
	if (j.task.Notify == 1 && err != nil) || j.task.Notify == 2 {
		tk := *j.task
		userStrin := strconv.Itoa(tk.UserId)
		user, uerr := models.GetUserInfoById(userStrin)
		if uerr != nil {
			return
		}
		var title string
		data := make(map[string]interface{})
		data["task_id"] = j.task.Id
		data["username"] = user.Name
		data["task_name"] = j.task.TaskName
		data["start_time"] = until.GetNowTime()
		data["process_time"] = float64(ut) / 1000
		data["output"] = cmdOut

		if httpcode != http.StatusOK {
			title = fmt.Sprintf("任务执行结果通知 #%d: %s", j.task.Id, "失败")
			data["status"] = "失败（" + err.Error() + "）"
		} else {
			title = fmt.Sprintf("任务执行结果通知 #%d: %s", j.task.Id, "成功")
			data["status"] = "成功"
		}

		content := new(bytes.Buffer)
		mailTpl.Execute(content, data)
		ccList := make([]string, 0)
		if j.task.NotifyEmail != "" {
			ccList = strings.Split(j.task.NotifyEmail, "\n")
		}
		contentStr := content.String()
		ret := mail.SendEmail(title, contentStr, []string{user.Email}, ccList)
		until.LogObj().WithFields(logrus.Fields{
			"func": "Run",
		}).Error("发送邮件：", ret)
	}

}
