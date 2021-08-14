package http

import (
	"fmt"
	nativeHttp "net/http"

	"github.com/gorilla/mux"
)

type HttpMethod string
type HttpHandler func(w nativeHttp.ResponseWriter, r *nativeHttp.Request)

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

type HttpRouter struct {
	router *mux.Router
}

func NewHttpRouter() *HttpRouter {
	r := mux.NewRouter()
	return &HttpRouter{
		router: r,
	}
}

func (h *HttpRouter) Register(method HttpMethod, route string, handler HttpHandler) {
	h.router.HandleFunc(route, handler).Methods(string(method))
}

func (h *HttpRouter) ListenOnPort(port int) {
	nativeHttp.ListenAndServe(":"+fmt.Sprint(port), h.router)
}
