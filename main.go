package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gologview/assets"
	_ "gologview/assets/tcptest"
	"gologview/go/logview"
	"gologview/go/util"
	frpNet "gologview/go/util/net"
	"net"
	"net/http"
	"time"
)

// 设置websocket
// CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func loadwebsocket(w http.ResponseWriter, r *http.Request) {

	fmt.Println("loadwebsocket...")

	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close() //返回前关闭
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("recv", message)
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}

}

func loadfile() {
	var (
		httpServerReadTimeout  = 60 * time.Second
		httpServerWriteTimeout = 60 * time.Second
	)
	assets.Load("")

	router := mux.NewRouter()

	subRouter := router.NewRoute().Subrouter()

	//user, passwd := "admin", "admin"
	//subRouter.Use(frpNet.NewHTTPAuthMiddleware(user, passwd).Middleware)

	serv := logview.NewService()

	ws := util.NewWebSocket()

	// api, see admin_api.go
	subRouter.HandleFunc("/api/status", serv.ApiStatus).Methods("GET")
	subRouter.HandleFunc("/echo", ws.Echo).Methods("GET")

	// view
	subRouter.Handle("/favicon.ico", http.FileServer(assets.FileSystem)).Methods("GET")
	subRouter.PathPrefix("/static/").Handler(frpNet.MakeHTTPGzipHandler(http.StripPrefix("/static/", http.FileServer(assets.FileSystem)))).Methods("GET")
	subRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/", http.StatusMovedPermanently)
	})

	address := ":8081"
	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  httpServerReadTimeout,
		WriteTimeout: httpServerWriteTimeout,
	}
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return
	}

	go func() {
		for {
			ws.SendAll([]byte("hello------>>>>"))
			time.Sleep(time.Second)
		}
	}()
	_ = server.Serve(ln)
}

func loadhtml() {
	//engine := gin.Default()
	//engine.LoadHTMLGlob("./static/index.html")
	//engine.StaticFS("/static", http.Dir("./static"))
	//engine.Run(":8080")
}

func main() {
	//loadfile()
	logview.New().Start()
}
