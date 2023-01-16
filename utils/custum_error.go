package utils

const (
	Ok    = 201
	NotOk = 405
)

const (
	OkMsg    = "执行成功"
	NotOkMsg = "执行失败"
)

var CustomError = map[int]string{
	Ok:    OkMsg,
	NotOk: NotOkMsg,
}
