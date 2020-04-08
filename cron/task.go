package task

import (
	"github.com/robfig/cron"
	"os"
	"os/signal"
	"syscall"
	"wxGameWebSocket/cron/jobs"
	"wxGameWebSocket/utils/logger"
)

var log = logger.New("cron")

// cron 定时任务处理
func TaskInit() {
	log.Info("Starting...")
	// 根据本地时间创建一个新（空白）的 Cron job runner
	c := cron.New()
	for _, job := range jobs.GetJobs() {
		log.Info("job启动", "job name", job.Name())
		c.AddFunc(job.Cron(), func() {
			defer func() {
				if err := recover(); err != nil {
					log.Error("job 运行出错", "job name", job.Name(), "error", err)
				}
			}()
			// 执行任务
			job.Run()
		})
	}
	c.Start()
	defer c.Stop()

	// 阻塞main协程退出，直到手动退出程序
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	killSignal := <-interrupt

	log.Info("退出定时任务", "signal", killSignal)
}
