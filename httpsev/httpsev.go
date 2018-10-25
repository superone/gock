package httpsev

import(
	"fmt"
	_"net/http"
)
type Request struct {
	Method string "method" 
	Params string "params" 
} 
type Response struct { 
	Code string "code" 
	Body string "body" 
} 

type Server interface{
	name string "gobx"
	Handle(method, params string) *Response
}

type IpcServer struct { 
	Server 
} 

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}
