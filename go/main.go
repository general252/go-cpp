package main

import (
	"github.com/general252/go-cpp/go/cpp"
	"github.com/general252/go-cpp/go/cpp/go4c_proto"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	cpp.InitLibrary(func(data []byte) {
		log.Printf("%v\n", string(data))
	})

	var d = go4c_proto.CmdLogin{
		Host:      "192.168.6.80",
		Port:      5566,
		Username:  "admin",
		Password:  "147258",
		ClientId:  "cli_23",
		UserAgent: "agent_45",
	}
	if data, err := proto.Marshal(&d); err != nil {
		log.Printf("xx %v", err)
	} else {
		_ = cpp.Command(data)
	}
}
