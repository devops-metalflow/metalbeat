package register

import (
	"context"
	"encoding/json"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"metalbeat/global"
	"metalbeat/utils"
	"net/http"
	"runtime"
	"strings"
	"time"
)

// Service 服务注册
func Service() {
	// 获取本机ip
	localIp, err := utils.LocalIp()
	if err != nil {
		global.Log.Fatal("获取本机ip失败：", err)
	}
	global.Log.Info("获取本机ip为： ", localIp.String())
	// 拼接服务健康检查的网址
	healthCheckUrl := fmt.Sprintf("http://%s:%d%s",
		localIp.String(),
		global.Conf.Service.Port,
		global.Conf.Service.CheckSuffix)

	registration := new(consulapi.AgentServiceRegistration)
	// service的id以service的name和服务本机的ip组成
	registration.ID = fmt.Sprintf("%s:%s", global.Conf.Service.Name, localIp.String())
	registration.Name = localIp.String()
	registration.Tags = strings.Split(global.Conf.Service.Tags, ",")
	registration.Port = global.Conf.Service.Port
	registration.Address = localIp.String()
	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           healthCheckUrl,
		Timeout:                        fmt.Sprintf("%ds", global.Conf.Service.CheckTimeout),
		Interval:                       fmt.Sprintf("%ds", global.Conf.Service.CheckInterval),
		DeregisterCriticalServiceAfter: fmt.Sprintf("%ds", global.Conf.Service.DeregisterAfter),
	}
	// 将运行机器的os信息添加到consul的kv配置中心
	// get os info.
	os := runtime.GOOS
	// Get a handle to the KV API.
	kv := global.Consul.KV()
	// put os info KV pair.
	p := &consulapi.KVPair{
		Key:   localIp.String(),
		Value: []byte(os),
	}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}
	// 注册
	if err := global.Consul.Agent().ServiceRegister(registration); err != nil {
		global.Log.Fatal("注册服务失败：", err)
	}
	global.Log.Info("注册服务成功：", global.Conf.Service.Name)
}

// HealthCheckHandler 创建一个健康检查的服务
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	global.Log.Info("请求健康检查：", r.URL.Path)
	if _, err := fmt.Fprintf(w, "service health check success"); err != nil {
		global.Log.Error(err)
	}
}

// ShellExecHandler add cmd execute handler
func ShellExecHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		resp := Resp{
			Code: utils.NotOk,
			Msg:  "非法请求方式",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	// tell client the response is json
	w.Header().Set("Content-Type", "application/json")
	// receive request and handle
	var req Request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		resp := Resp{
			Code: utils.NotOk,
			Msg:  "接受数据不合法",
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	global.Log.Info("开始执行命令：", req.Cmd)
	execute := false

	timeoutChan := make(chan struct{})
	retChan := make(chan Resp)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(global.Conf.Shell.Timeout)*time.Second)
	defer cancel()
	go func() {
		for { //nolint:gosimple
			select {
			case <-ctx.Done():
				if !execute {
					// broadcast
					close(timeoutChan)
				}
				// 此处需return避免协程空跑
				return
			}
		}
	}()

	// execute the command
	go func() {
		code := utils.Ok
		msg := utils.CustomError[code]
		ret, err := utils.ExcShell(req.Cmd)
		if err != nil {
			code = utils.NotOk
			msg = fmt.Sprintf("%s:%s", utils.CustomError[code], err)
		}
		global.Log.Info("命令执行结束：", req.Cmd)
		resp := Resp{
			Code: code,
			Data: ret,
			Msg:  msg,
		}
		retChan <- resp
		execute = true
	}()

	select {
	case <-timeoutChan:
		resp := Resp{
			Code: utils.NotOk,
			Msg:  "命令执行时间过长，已强制结束",
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)
	case ret := <-retChan:
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(ret)
	}
}

type Request struct {
	Cmd string `json:"cmd"`
}

// Resp http response data struct
type Resp struct {
	Code int         `json:"code"` // 错误代码
	Data interface{} `json:"data"` // 数据内容
	Msg  string      `json:"msg"`  // 消息提示
}
