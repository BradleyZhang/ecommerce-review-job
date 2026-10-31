package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/BradleyZhang/ecommerce-review-job/common"
	"github.com/BradleyZhang/ecommerce-review-job/internal/job"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	initResources()

	cl, err := job.NewKgoCl(job.NewDefaultKgoCfg())
	if err != nil {
		common.FatalLog("failed to new kgo client,err:", err)
	}
	esCli, err := job.NewESClient(job.NewESDefaultCfg())
	if err != nil {
		common.FatalLog("failed to new es client, err: ", err)
	}
	jobWorker := job.NewJobWorker(cl, esCli)
	go func() {
		err := jobWorker.Start(ctx)
		if err != nil {
			common.FatalLog("job worker 错误: ", err)
		}
	}()
	common.SysLog("review-job start...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	common.SysLog("received signal: ", sig, " shutting down...")

	jobWorker.Stop()
}
func initResources() {
	err := godotenv.Load(".env")
	if err != nil {
		common.SysLog("No .env file found,useing default environment variables.If needed, please create a .env file and set the relevent variables.")
	}

	common.InitEnv()
}
