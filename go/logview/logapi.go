package logview

import (
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

func (this *LogApi) Start() {
	user, passwd := "admin", "het002402"
	this.subRouter.Use(frpNet.NewHTTPAuthMiddleware(user, passwd).Middleware)
	// api, see admin_api.go
	this.subRouter.HandleFunc("/api/status", this.serv.ApiStatus).Methods("GET")
	this.subRouter.HandleFunc("/echo", this.wsapi.Echo).Methods("GET")

	// view
	this.subRouter.Handle("/favicon.ico", http.FileServer(assets.FileSystem)).Methods("GET")
	this.subRouter.PathPrefix("/static/").Handler(frpNet.MakeHTTPGzipHandler(http.StripPrefix("/static/", http.FileServer(assets.FileSystem)))).Methods("GET")
	this.subRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/", http.StatusMovedPermanently)
	})

	address := ":8080"
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

	_ = server.Serve(ln)
}

func (this *LogApi) Send(message []byte) {
	this.wsapi.SendAll(message)
}
