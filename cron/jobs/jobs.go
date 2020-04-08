package jobs

import "wxGameWebSocket/utils/logger"

var log = logger.New("jobs")

type Job interface {
	Run()
	Cron() string
	Name() string
}

var jobs []Job

func addJobs(runner Job) {
	jobs = append(jobs, runner)
}

// GetJobs 得到所有定时任务
func GetJobs() []Job {
	return jobs
}