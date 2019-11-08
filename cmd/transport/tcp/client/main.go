package main

import (
	"github.com/brewlin/net-protocol/pkg/logging"
	"github.com/brewlin/net-protocol/protocol/transport/tcp/client"
	_ "github.com/brewlin/net-protocol/stack/stackinit"
	"fmt"
)

func init() {
	logging.Setup()
}
func main() {
	con := client.NewClient("127.0.0.1", 8080)
	con.Connect()
	con.Write([]byte("send msg"))
	res := con.Read()
	fmt.Println(res)
}
