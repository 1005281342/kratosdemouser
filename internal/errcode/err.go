package errcode

import "fmt"

var (
	ErrCheckParam        = fmt.Errorf("check param failed")
	ErrCheckRowsAffected = fmt.Errorf("check rows affected failed")
)
