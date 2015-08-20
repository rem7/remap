package main

import (
	"fmt"
	"os"
	"time"
)

func LogPrintf(format string, a ...interface{}) (int, error) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	var args []interface{}
	args = append(args, timestamp)
	args = append(args, a...)
	return fmt.Fprintf(os.Stdout, "[%v] "+format+"\n", args...)
}
