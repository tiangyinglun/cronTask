package Job

import (
	"github.com/jakecoffman/cron"
	"strconv"
	"sync"
)

var (
	mainCron *cron.Cron
	workPool chan bool
	lock     sync.Mutex
)

/**
初始化
*/
func init() {
	//初始换只给500 个
	workPool = make(chan bool, 500)
	mainCron = cron.New()
	mainCron.Start()
}

/**
添加任务
*/
func AddJob(spec string, job *Job) bool {
	lock.Lock()
	defer lock.Unlock()

	name := strconv.Itoa(job.id)
	//防止重复
	if GetEntryById(name) != nil {
		return false
	}
	//添加任务
	mainCron.AddJob(spec, job, name)
	return true
}

/**
删除任务
*/
func RemoveJob(name string) {
	mainCron.RemoveJob(name)
}

/**
获取任务id
*/
func GetEntryById(name string) *cron.Entry {
	entries := mainCron.Entries()
	for _, e := range entries {
		if name == e.Name {
			return e
		}
	}
	return nil
}

/**
获取数量
*/
func GetEntries(size int) []*cron.Entry {
	ret := mainCron.Entries()
	if len(ret) > size {
		return ret[:size]
	}
	return ret
}
