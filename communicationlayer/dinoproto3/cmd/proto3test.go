package main

import (
	"dino/communicationlayer/dinoproto3"
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
	op := flag.String("op", "s", "s for server, c for client")
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		RunProto3Server()
	case "c":
		RunProto3Client()
	}
}

func RunProto3Server() {
	l, err := net.Listen("tcp", ":8282")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
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
			a := &dinoproto3.Animal{} // from protodata.pb.go Animal
			err = proto.Unmarshal(data, a)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(a)
		}(c)
	}
}

func RunProto3Client() {
	a := &dinoproto3.Animal{
		Id:         1,
		AnimalType: "Raptor",
		Nickname:   "rapto",
		Zone:       3,
		Age:        20,
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
