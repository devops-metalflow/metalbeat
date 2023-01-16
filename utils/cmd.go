package utils

import (
	"fmt"
	"os/exec"
	"path"
	"strings"
)

// ExcShell  execute the cmd
func ExcShell(input string) (ret string, err error) {
	if !IsSafetyCmd(input) {
		err = fmt.Errorf("命令[%s]包含危险操作命令，不能执行", input)
		return
	}
	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")

	// confirm the cmd is in path
	_, err = exec.LookPath(args[0])
	if err != nil {
		err = fmt.Errorf("命令%s不是可执行命令", args[0])
		return
	}

	cmd := exec.Command("/bin/sh", "-c", input)
	data, err := cmd.CombinedOutput()
	ret = string(data)
	return
}

// IsSafetyCmd confirm the cmd is safe
func IsSafetyCmd(cmd string) bool {
	// 避免rm * 或 rm /*等命令直接出现, 删除命令指定全路径
	c := path.Clean(strings.ToLower(cmd))
	if strings.Contains(c, "rm") {
		if len(strings.Split(c, "/")) <= 1 {
			return false
		}
	}
	return true
}
