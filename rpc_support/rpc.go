package rpc_support

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//新起一个rpc服务，并注册service
func ServeRpc(host string, service ...interface{}) error {
	var err error
	for _, v := range service {
		err = rpc.Register(v)
		if err != nil {
			log.Println("register service fail, %s", err)
		}
	}

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	log.Println("Listen", host, "...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error %v", err)
		}
		go jsonrpc.ServeConn(conn)
	}
}

//新起一个rpc的client
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	//这里切忌不能conn.Close()！！！  因为关闭之后后续就没法连接了
	return jsonrpc.NewClient(conn), nil
}
