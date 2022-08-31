package httpserver

import (
	"context"
	"fmt"
	"github.com/dnsdao/nft.pass/config"
	"github.com/dnsdao/nft.pass/server/httpserver/api"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"regexp"
	"time"
)

const (
	CTokenId = "/api/nft/c/{token_id}"
	MTokenId = "/api/nft/m/{token_id}"
)

type WebProxyServer struct {
	listenAddr string
	quit       chan struct{}
	server     *http.Server
}

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHander struct {
	routes []*route
}

func (rh *RegexpHander) Handle(pattern string, handler http.Handler) {
	rh.routes = append(rh.routes, &route{pattern: regexp.MustCompilePOSIX(pattern), handler: handler})
}

func (rh *RegexpHander) HandleFunc(pattern string, handleFunc func(http.ResponseWriter, *http.Request)) {
	rh.routes = append(rh.routes, &route{pattern: regexp.MustCompilePOSIX(pattern), handler: http.HandlerFunc(handleFunc)})
}

func (rh *RegexpHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range rh.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

func NewWebServer(networkAddr string) *WebProxyServer {
	ws := WebProxyServer{
		listenAddr: networkAddr,
		quit:       make(chan struct{}, 8),
	}

	return ws.init()

}

func (ws *WebProxyServer) init() *WebProxyServer {
	//rh := &RegexpHander{
	//	routes: make([]*route, 0),
	//}
	wapi := api.NewApi()
	//rh.HandleFunc(TokenId, wapi.TokenId)
	r := mux.NewRouter()
	r.HandleFunc(CTokenId, wapi.CTokenId)
	r.HandleFunc(MTokenId, wapi.MTokenId)
	server := &http.Server{
		Handler: r,
	}

	ws.server = server

	return ws
}

func (ws *WebProxyServer) Start() error {
	if l, err := net.Listen("tcp4", ws.listenAddr); err != nil {
		panic("start wss server failed")
	} else {
		return ws.server.Serve(l)
	}

}

func (ws *WebProxyServer) Shutdown() error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ws.server.Shutdown(ctx)
}

var webServer *WebProxyServer

func StartWebDaemon() {

	c := config.GetRConf().WConf

	webServer = NewWebServer(c.ListenServer)

	fmt.Println("start proxy at ", webServer.listenAddr, "  ...")

	webServer.Start()
}

func StopWebDaemon() {
	webServer.Shutdown()

	fmt.Println("stop proxy ...")
}
