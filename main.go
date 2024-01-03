package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/gologview/go/glogweb"
	"github.com/xxl6097/gologview/go/util"
	"net"
	"net/http"
	"time"
)

func init() {
	//glogweb.GetLogApi().RunAndSetUserPass(8086, "admin", "admin")
}

func TestFileServer() {
	router := mux.NewRouter()
	glogweb.GetLogApi().HandlerLogView(router, "admin", "admin")

	address := fmt.Sprintf(":%d", 8080)
	server := &http.Server{
		Addr:    address,
		Handler: router,
	}
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	fmt.Printf("please open http://localhost%s\n", server.Addr)
	fmt.Printf("please open http://localhost%s/fserver/\n", server.Addr)
	_ = server.Serve(ln)
}

func start(port int) {
	router := mux.NewRouter()

	glogweb.GetLogApi().HandlerLogView(router, "admin", "admin")

	address := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:    address,
		Handler: router,
	}
	ln, err := net.Listen("tcp", address)
	if err != nil {
		glog.Error(err)
		return
	}

	ip := util.GetHostIp()
	glog.Errorf("logv webside http://%s%s/logview/\n", ip, server.Addr)
	fmt.Printf("please open http://localhost%s/fserver/\n", server.Addr)
	_ = server.Serve(ln)
}

func main() {
	start(8080)
	for {
		fmt.Println("aaaaaa这个不会在web上显示哦")
		glog.Info("info...")
		time.Sleep(time.Second * 5)
	}

	//glogweb.TestFileServer()
}
