package main

import (
	"dino/communicationlayer/dinoproto2"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"github.com/golang/protobuf/proto"
)

/*
	1- will serialize some data via dinoproto2
	2- will sen this data via TCP to a different service
	3- will deserialize the dta via proto2, and print out the extracted values

	A- A TCP client needs to be written to send the data
	B- A TCP Server to receive the data
*/

func main() {
	op := flag.String("op", "s", "s for server, and c for client") // proto2test -op s => will run as a server. proto2test -op c => will run as a client.
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		RunProto2Server()
	case "c":
		RunProto2Client()
	}
}

func RunProto2Client() {
	a := &dinoproto2.Animal{
		Id:         proto.Int(1),
		AnimalType: proto.String("Raptor"),
		Nickname:   proto.String("rapto"),
		Zone:       proto.Int(3),
		Age:        proto.Int(20),
	}
	data, err := proto.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	SendData(data)
}
func SendData(data []byte) {
	c, err := net.Dial("tcp", "127.0.0.1:8282")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c.Write(data)
}

func RunProto2Server() {
	l, err := net.Listen("tcp", ":8282")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// c Connection Handler
		c, err := l.Accept() // wait till it receives an incoming TCP request from a TCP client
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()
		go func(c net.Conn) {
			defer c.Close()
			data, err := ioutil.ReadAll(c)
			if err != nil {
				return
			}
			a := &dinoproto2.Animal{}
			proto.Unmarshal(data, a)
			fmt.Println(a)
		}(c)
	}
}
