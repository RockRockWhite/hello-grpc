package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"log"
	"protobuf_go/src/first"
)

func main() {
	var msg proto.Message
	msg = NewPersonMessage()

	dataBytes, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dataBytes)

	var pb first.PersonMessage
	proto.Unmarshal(dataBytes, &pb)
	fmt.Println(&pb)

	marshaler := jsonpb.Marshaler{Indent: "    "}
	json, _ := marshaler.MarshalToString(&pb)
	fmt.Println(json)

	var jpb first.PersonMessage
	jsonpb.UnmarshalString(json,&jpb)
	fmt.Println(&jpb)
}

func NewPersonMessage() *first.PersonMessage {
	pm := first.PersonMessage{
		Id:          1234,
		IsAdult:     true,
		Name:        "Hello",
		LuckyNumbes: []int32{1, 2, 3, 4, 5, 6},
	}

	return &pm
}
