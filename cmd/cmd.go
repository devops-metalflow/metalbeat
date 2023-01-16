package cmd

import (
	"context"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"metalbeat/global"
	"metalbeat/initialize"
	"metalbeat/register"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

const version = "1.3.0"

var (
	app        = kingpin.New("metalbeat", "MetalBeat").Version(version)
	configFile = app.Flag("config-file", "Configfile (.yml)").String()
)

func Run() {
	// 捕获异常,并写入日志
	defer func() {
		if err := recover(); err != nil {
			global.Log.Error(fmt.Sprintf("启动metalbeat失败：%v\n堆栈信息：%v", err, string(debug.Stack())))
		}
	}()

	kingpin.MustParse(app.Parse(os.Args[1:]))
	// 初始化配置文件
	initialize.Config(*configFile)
	// 初始化日志
	initialize.Logger()
	// 初始化consul
	initialize.Consul()
	// 服务注册
	register.Service()
	// 初始化相关shell任务
	initialize.Shell()

	// 打开健康检查服务并监听
	host := "0.0.0.0"
	port := global.Conf.Service.Port
	mux := http.NewServeMux()
	mux.HandleFunc(global.Conf.Service.CheckSuffix, register.HealthCheckHandler)
	// add cmd execute service
	mux.HandleFunc(global.Conf.Shell.ShellSuffix, register.ShellExecHandler)
	// 服务器启动以及优雅的关闭
	// 参考地址https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", host, port),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Error("listen error: ", err)
		}
	}()

	global.Log.Info(fmt.Sprintf("Health-Check server is running at %s:%d%s", host, port, global.Conf.Service.CheckSuffix))
	global.Log.Info(fmt.Sprintf("Shell-cmd server is running at %s:%d%s", host, port, global.Conf.Shell.ShellSuffix))

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Log.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //nolint:gomnd
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Log.Error("Server forced to shutdown: ", err)
	}

	global.Log.Info("Server exiting")
}
