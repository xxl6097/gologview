package main

import (
	"github.com/xxl6097/gologview/go/logview"
	"net/http"
)

//func loadhtml() {
//	engine := gin.Default()
//	engine.LoadHTMLGlob("./assets/tcptest/logview/index.html")
//	engine.StaticFS("/static", http.Dir("./assets/tcptest/logview"))
//	engine.Run(":7777")
//}

func StaticServer(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./assets/tcptest/logview/"+r.URL.Path)
	//http.FileServer(assets.FileSystem)
}
func main() {
	//assets.Load("")
	//router := mux.NewRouter()
	////当前端访问：http://127.0.0.1:8000/ 地址时返回当前目录下的 index.html
	//router.HandleFunc("/", StaticServer)
	////当前端访问路径前缀为 /static/js/ 时返回访问地址指向的js文件内容
	////当网页解析到 <script defer="defer" src="/static/js/main.c2c1cea9.js"></script> 时会执行
	//router.PathPrefix("/js/").HandlerFunc(StaticServer)
	////当前端访问路径前缀为 /static/css/ 时返回访问地址指向的css文件内容
	//router.PathPrefix("/css/").HandlerFunc(StaticServer)
	//router.Handle("/staticfile/", http.StripPrefix("/staticfile/", http.FileServer(http.Dir("/Users/uuxia/Desktop/work/code/go/gologview"))))
	//
	//srv := &http.Server{
	//	Handler:      router,
	//	Addr:         ":9999",
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}
	//
	//log.Fatal(srv.ListenAndServe())

	//https://cyjy-iot.chengyang.gov.cn/clink/gtbx/log/
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

	logview.New().Start(9999)
}
