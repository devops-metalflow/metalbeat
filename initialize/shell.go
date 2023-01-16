package initialize

import (
	"fmt"
	"metalbeat/global"
	"metalbeat/utils"
	"sync"
)

// Shell 用于执行初始化任务
func Shell() {
	if !global.Conf.Shell.InitShell {
		global.Log.Info("未设置初始化任务, 无需执行初始化命令")
		return
	}
	// 假定各个任务都是独立的，因此使用并发执各个任务
	errChan := make(chan error, 1)
	okChan := make(chan struct{})
	var wg sync.WaitGroup
	for _, command := range global.Conf.Shell.InitCommands {
		wg.Add(1)
		go func(cmd global.MetalTaskConfiguration) {
			defer wg.Done()
			global.Log.Info(fmt.Sprintf("开始执行任务%s的命令", cmd.Name))
			for _, c := range cmd.Commands {
				_, err := utils.ExcShell(c)
				if err != nil {
					global.Log.Errorf("执行命令%s失败", c)
					errChan <- err
				}
			}
		}(command)
	}

	go func() {
		wg.Wait()
		close(okChan)
	}()

	select {
	case <-okChan:
		global.Log.Info("所有初始化任务执行完成")
	case err := <-errChan:
		global.Log.Error("执行初始化任务失败：", err)
	}
}
