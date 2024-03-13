package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/gologview/go/glogweb"
	"net"
	"net/http"
	"os"
	"time"
)

func init() {
	glogweb.GetLogApi().RunAndSetUserPass(8086, "admin", "admin", func(router *mux.Router) {
		router.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
			_, _ = writer.Write([]byte("hello glogweb"))
		})
	})
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

func main() {
	fmt.Println("=======>" + os.Getenv("ENV_TYPE"))
	for {
		fmt.Println("aaaaaa这个不会在web上显示哦")
		glog.Info("info...")
		time.Sleep(time.Second * 5)
	}
}
