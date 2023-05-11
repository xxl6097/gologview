package logview

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xxl6097/gologview/assets"
	_ "github.com/xxl6097/gologview/assets/tcptest"
	"github.com/xxl6097/gologview/go/util"
	frpNet "github.com/xxl6097/gologview/go/util/net"
	"net"
	"net/http"
	"time"
)

var (
	httpServerReadTimeout  = 60 * time.Second
	httpServerWriteTimeout = 60 * time.Second
)

type LogApi struct {
	wsapi     util.WebSocketUtil
	serv      *GeneralResponse
	router    *mux.Router
	subRouter *mux.Router
}

func New() *LogApi {
	assets.Load("")
	api := &LogApi{
		wsapi: util.NewWebSocket(),
	}
	api.router = mux.NewRouter()
	api.subRouter = api.router.NewRoute().Subrouter()
	api.serv = NewService()
	return api
}

func (this *LogApi) Start(port int) {
	user, passwd := "admin", "het002402"
	this.subRouter.Use(frpNet.NewHTTPAuthMiddleware(user, passwd).Middleware)
	// api, see admin_api.go
	this.subRouter.HandleFunc("/api/status", this.serv.ApiStatus).Methods("GET")
	this.subRouter.HandleFunc("/echo", this.wsapi.Echo).Methods("GET")

	//this.subRouter.Handle("/file", http.FileServer(http.Dir("."))).Methods("GET")
	// view
	this.subRouter.Handle("/favicon.ico", http.FileServer(assets.FileSystem)).Methods("GET")
	this.subRouter.PathPrefix("/logview").Handler(frpNet.MakeHTTPGzipHandler(http.StripPrefix("/logview", http.FileServer(assets.FileSystem)))).Methods("GET")
	this.subRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/logview/", http.StatusMovedPermanently)
	})

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

	fmt.Printf("please open http://localhost%s", server.Addr)
	_ = server.Serve(ln)
}

func (this *LogApi) Send(message []byte) {
	this.wsapi.SendAll(message)
}
