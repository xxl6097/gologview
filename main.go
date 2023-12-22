package main

import (
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/gologview/go/glogweb"
	"time"
)

func init() {
	glogweb.GetLogApi().RunAndSetUserPass(8086, "admin", "admin")
}

func main() {
	for {
		fmt.Println("aaaaaa这个不会在web上显示哦")
		glog.Info("info...")
		time.Sleep(time.Second * 5)
	}
}
