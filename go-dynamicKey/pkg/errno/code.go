package errno

import "fmt"

var (
	codes = map[int]struct{}{}
)

func New(e int) int {
	if e <= 0 {
		panic("code must greater than zero")
	}
	return add(e)
}

func add(e int) int {
	if _, ok := codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	codes[e] = struct{}{}
	return e
}

var (
	OK                  	= &Errno{Code: New(1), 	 Message: "OK"}
	ErrInternalServer	 	= &Errno{Code: New(1001), Message: "error Internal server"}
	ErrBind             	= &Errno{Code: New(1002), Message: "error bind struct"}
	ErrHeartBeat			= &Errno{Code: New(1003), Message: "error heartbeat"}
	ErrReqDecode			= &Errno{Code: New(1004), Message: "error heartbeat"}
	ErrParams				= &Errno{Code: New(1005), Message: "error params"}
)
