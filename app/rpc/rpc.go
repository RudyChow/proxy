package rpc

import (
	"github.com/RudyChow/proxy/app/rpc/entities"
	"github.com/RudyChow/proxy/config"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func StartRpcServer()  {
	rpc.Register(new(entities.ProxyServer))

	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", config.Conf.Rpc.Addr)
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}

	log.Println("rpc start connection")

	http.Serve(lis, nil)
}