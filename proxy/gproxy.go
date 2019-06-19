package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type h struct {
	host string
	port string
}

func (p *h) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("http://" + p.host + ":" + p.port)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func startServer() {
	//被代理的服务器host和port
	h := &h{host: "127.0.0.1", port: "80"}
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}
