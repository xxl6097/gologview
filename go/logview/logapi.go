package logview

import (
	"fmt"
	"github.com/gorilla/mux"
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
)

type LogApi struct {
	wsapi              util.WebSocketUtil
	serv               *GeneralResponse
	router             *mux.Router
	subRouter          *mux.Router
	Username, Password string
}

var lock = &sync.Mutex{}
var instance *LogApi

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
	api := &LogApi{
		wsapi:    util.NewWebSocket(),
		Username: "admin",
		Password: "het002402",
	}
	api.router = mux.NewRouter()
	api.subRouter = api.router.NewRoute().Subrouter()

	api.router.Headers("Access-Control-Allow-Origin", "*") //允许访问所有域
	// 必须，设置服务器支持的所有跨域请求的方法
	api.router.Headers("Access-Control-Allow-Methods", "*")
	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	api.router.Headers("Access-Control-Allow-Headers", "*")
	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
	api.router.Headers("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
	api.router.Headers("Access-Control-Allow-Credentials", "true")

	api.subRouter.Headers("Content-Type", "application/json")
	api.subRouter.Headers("Access-Control-Allow-Origin", "*") //允许访问所有域
	// 必须，设置服务器支持的所有跨域请求的方法
	api.subRouter.Headers("Access-Control-Allow-Methods", "*")
	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	api.subRouter.Headers("Access-Control-Allow-Headers", "*")
	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
	api.subRouter.Headers("Access-Control-Expose-Headeapirs", "Access-Control-Allow-Headers, Token")
	// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
	api.subRouter.Headers("Access-Control-Allow-Credentials", "true")

	api.router.Use(mux.CORSMethodMiddleware(api.router))
	api.serv = NewService()
	return api
}
func fileServer(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/files/", http.FileServer(http.Dir("/"))).ServeHTTP(w, r)
}

func (this *LogApi) SetUser(username, password string) *LogApi {
	this.Username = username
	this.Password = password
	return this
}

func (this *LogApi) Start(port int) {
	//user, passwd := "admin", "het002402"
	//this.subRouter.Use(util.NewHTTPAuthMiddleware(this.Username, this.Password).Middleware)
	// api, see admin_api.go
	this.subRouter.HandleFunc("/api/status", this.serv.ApiStatus).Methods("GET")
	this.subRouter.HandleFunc("/api/files", this.serv.ApiFiles).Methods("GET")
	//this.router.HandleFunc("/api/dirs/", this.serv.ApiDir)
	this.subRouter.HandleFunc("/echo", this.wsapi.Echo).Methods("GET")

	this.subRouter.PathPrefix("/fserver/").Handler(http.StripPrefix("/fserver/", http.FileServer(http.Dir("/"))))

	// view
	this.subRouter.Handle("/favicon.ico", http.FileServer(assets.FileSystem)).Methods("GET")
	this.subRouter.PathPrefix("/").Handler(util.MakeHTTPGzipHandler(http.StripPrefix("/", http.FileServer(assets.FileSystem)))).Methods("GET")
	//this.subRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	//})
	address := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:         address,
		Handler:      this.router,
		ReadTimeout:  httpServerReadTimeout,
		WriteTimeout: httpServerWriteTimeout,
	}
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return
	}

	//go func() {
	//	for {
	//		this.wsapi.SendAll([]byte("---------"))
	//		time.Sleep(time.Second * 5)
	//	}
	//}()

	fmt.Printf("logview webside http://localhost%s\n", server.Addr)
	_ = server.Serve(ln)
}

func (this *LogApi) Send(message []byte) {
	this.wsapi.SendAll(message)
}
