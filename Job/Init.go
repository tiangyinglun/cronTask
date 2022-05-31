package Job

import (
	"Gin/models"
	"Gin/until"
	"fmt"
)

func InitJobs() {
	IpLong := until.GetCmdIpLong()
	fmt.Println(IpLong)
	list, _ := models.GetAllTaskList(IpLong, 10000)
	for _, task := range *list {
		job, err := NewJobFromTask(&task)
		if err != nil {
			continue
		}
		fmt.Println(job)
		AddJob(task.CronSpec, job)
	}
}
