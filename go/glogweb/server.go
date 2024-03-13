package glogweb

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/gologview/assets"
	_ "github.com/xxl6097/gologview/assets/html"
	"github.com/xxl6097/gologview/go/util"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	httpServerReadTimeout  = 60 * time.Second
	httpServerWriteTimeout = 60 * time.Second
	pifix                  = "/logview"
)
var lock = &sync.Mutex{}
var instance *LogApi

type LogApi struct {
	wsapi              util.WebSocketUtil
	serv               *GeneralResponse
	router             *mux.Router
	subRouter          *mux.Router
	Username, Password string
}

func GetLogApi() *LogApi {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = new()
		} else {
		}
	} else {
	}
	return instance
}

func new() *LogApi {
	assets.Load("")
	this := &LogApi{
		wsapi:    util.NewWebSocket(),
		Username: "admin",
		Password: "admin",
	}
	glog.Hook(this.hook)
	return this
}
func (this *LogApi) Send(message []byte) {
	this.wsapi.SendAll(message)
}
func (this *LogApi) hook(log []byte) {
	this.Send(log)
}
func (this *LogApi) setUserPass(username, password string) *LogApi {
	this.Username = username
	this.Password = password
	return this
}
func (this *LogApi) RunAndSetUserPass(port int, username, password string, fun func(*mux.Router)) {
	this.setUserPass(username, password)
	go this.start(port, fun)
}

func (this *LogApi) HandlerLogView(router *mux.Router, username, password string) {
	this.Username = username
	this.Password = password
	this.handlelog(router, nil)
}

func (this *LogApi) handlelog(router *mux.Router, fun func(*mux.Router)) {
	this.serv = NewService()

	router.Use(util.NewHTTPAuthMiddleware(this.Username, this.Password).Middleware)
	if fun != nil {
		fun(router)
	}
	router.HandleFunc(pifix+"/api/status", this.serv.ApiStatus).Methods("GET")
	router.HandleFunc(pifix+"/api/files", this.serv.ApiFiles).Methods("GET")
	router.HandleFunc(pifix+"/echo", this.wsapi.Echo).Methods("GET")
	//router.PathPrefix(pifix + "/fserver/").Handler(http.StripPrefix("/fserver/", http.FileServer(http.Dir("/"))))
	router.PathPrefix("/fserver/").Handler(http.StripPrefix("/fserver/", http.FileServer(http.Dir("/"))))
	// view
	router.Handle(pifix+"/favicon.ico", http.FileServer(assets.FileSystem)).Methods("GET")
	router.PathPrefix(pifix + "/").Handler(util.MakeHTTPGzipHandler(http.StripPrefix("/", http.FileServer(assets.FileSystem)))).Methods("GET")
}

func (this *LogApi) start(port int, fun func(*mux.Router)) {
	this.router = mux.NewRouter()
	this.subRouter = this.router.NewRoute().Subrouter()

	this.handlelog(this.subRouter, fun)

	address := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:         address,
		Handler:      this.router,
		ReadTimeout:  httpServerReadTimeout,
		WriteTimeout: httpServerWriteTimeout,
	}
	ln, err := net.Listen("tcp", address)
	if err != nil {
		glog.Error(err)
		return
	}

	ip := util.GetHostIp()
	glog.Errorf("logv webside http://%s%s/logview/\n", ip, server.Addr)
	_ = server.Serve(ln)
}

func TestFileServer() {
	router := mux.NewRouter()
	router.NewRoute().Subrouter().PathPrefix("/fserver/").Handler(http.StripPrefix("/fserver/", http.FileServer(http.Dir("/"))))
	//router.PathPrefix("/").Handler(util.MakeHTTPGzipHandler(http.StripPrefix("/", http.FileServer(assets.FileSystem)))).Methods("GET")

	address := fmt.Sprintf(":%d", 8080)
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
	fmt.Printf("please open http://localhost%s\n", server.Addr)
	fmt.Printf("please open http://localhost%s/fserver/\n", server.Addr)
	_ = server.Serve(ln)
}
