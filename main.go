package main

import (
	"fmt"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/gologview/go/logview"
	"log"
	"net/http"
	"time"
)

//func loadhtml() {
//	engine := gin.Default()
//	engine.LoadHTMLGlob("./assets/tcptest/logview/index.html")
//	engine.StaticFS("/static", http.Dir("./assets/tcptest/logview"))
//	engine.Run(":7777")
//}

func StaticServer(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, ".")
	pp := r.FormValue("path")
	http.ServeFile(w, r, pp)
	//http.FileServer(assets.FileSystem)
	//http.StripPrefix("/file/", http.FileServer(http.Dir("/Users/uuxia/Desktop/work/code/go/gologview"))).ServeHTTP(w, r)
}

func serveFile1() {
	router := http.NewServeMux()
	//router := mux.NewRouter()
	//router.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("/Users/uuxia/Desktop/work/code/go/gologview"))))
	router.HandleFunc("/files/", StaticServer)
	//http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("/Users/uuxia/Desktop/work/code/go/gologview"))))
	//router.PathPrefix("/js/").HandlerFunc(StaticServer)
	////当前端访问路径前缀为 /static/css/ 时返回访问地址指向的css文件内容
	//router.PathPrefix("/css/").HandlerFunc(StaticServer)
	//router.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("/Users/uuxia/Desktop/work/code/go/gologview"))))

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func serveFile() {
	http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("/Users/uuxia/Desktop/work/code/go/gologview"))))
	http.ListenAndServe(":8080", nil)
}

func onLogHook(log []byte) {
	logview.GetLogApi().Send(log)
}

func init() {
	glog.Hook(onLogHook)
	logview.GetLogApi().SetUser("admin", "admin")
	go func() {
		logview.GetLogApi().Start(8080)
	}()
}

func main() {
	//serveFile1()
	//r := mux.NewRouter()
	//r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("/Users/uuxia/Desktop/work/code/go/gologview"))))
	////log.Fatal(http.ListenAndServe(":8888", r))
	//address := fmt.Sprintf(":%d", 8888)
	//server := &http.Server{
	//	Addr:    address,
	//	Handler: r,
	//}
	//ln, err := net.Listen("tcp", address)
	//if err != nil {
	//	return
	//}
	//fmt.Printf("please open http://localhost%s", server.Addr)
	//_ = server.Serve(ln)

	//serveFile()
	//http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("/Users/uuxia/Desktop/work/code/go/gologview"))))

	//https://cyjy-iot.chengyang.gov.cn/clink/gtbx/log/
	//wss://cyjy-iot.chengyang.gov.cn/clink/gtbx/echo
	//aa := http.Dir("../")//20230511111437
	//fmt.Println(aa)
	//
	//dirs, err := os.ReadDir("./assets")
	//if err != nil {
	//	return
	//}
	//for _, dir := range dirs {
	//	if !dir.IsDir() {
	//		arr, _ := os.ReadFile(filepath.Join("./assets", dir.Name()))
	//		fmt.Println(string(arr))
	//	}
	//}

	//static file handler.
	//http.Handle("/staticfile/", http.StripPrefix("/staticfile/", http.FileServer(http.Dir("/Users/uuxia/Desktop/work/code/go/gologview/"))))
	////Listen on port 8080
	//http.ListenAndServe(":9999", nil)

	//this.subRouter.Handle("/file", http.FileServer()).Methods("GET")

	//router := mux.NewRouter()
	//address := fmt.Sprintf(":%d", 9999)
	//server := &http.Server{
	//	Addr:    address,
	//	Handler: router,
	//}
	//router.Handle("/staticfile/", http.StripPrefix("/staticfile/", http.FileServer(http.Dir("/Users/uuxia/Desktop/work/code/go/gologview/"))))
	//ln, err := net.Listen("tcp", address)
	//if err != nil {
	//	return
	//}
	//fmt.Printf("please open http://localhost%s", server.Addr)
	//_ = server.Serve(ln)
	//aa := "/User/uuxia/a.txt/"
	//bb := strings.HasSuffix(aa, "/")
	//fmt.Println(bb)
	for {
		fmt.Println("aaaaaa")
		glog.Info("info...")
		time.Sleep(time.Second)
	}
}
