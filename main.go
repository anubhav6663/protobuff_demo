package main

import (
	"fmt"
	"io/ioutil"
	"log"

	simplepb "github.com/anubhavsingh6663/protobuff_demo/src/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()
	readAndWritDemo(sm)
	jsonDemo(sm)

}

func jsonDemo(sm proto.Message) {

	smAsString := toJSON(sm)
	fmt.Println(" JSON formatted proto message: \n ", smAsString)
	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Print(" Successfully created proto struct\n", sm2)
}
func readAndWritDemo(sm proto.Message) {

	witetofile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(" Read the content \n", sm2)
}

func toJSON(pb proto.Message) string {
	marshaller := jsonpb.Marshaler{}
	out, err := marshaller.MarshalToString(pb)
	if err != nil {
		fmt.Println(" Can't convert to JSON ")
		return ""
	}
	return out
}

func fromJSON(msg string, pb proto.Message) {

	err := jsonpb.UnmarshalString(msg, pb)
	if err != nil {
		log.Fatalln(" Can't unmarshal the msg to pb struct", err)
	}

}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln(" SOmething went wrong while reading the file", err)
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln(" Couldn't put the bytes into the protocol buffers", err2)
		return err2
	}
	return nil
}

func witetofile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln(" Can't serialize to byte", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln(" Can't write to file", err)
		return err
	}

	fmt.Println(" Data written to file ")
	return nil
}
func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         123,
		IsSimple:   true,
		Name:       " Simpl message just for fun ",
		SampleList: []int32{1, 2, 4, 5, 7},
	}
	return &sm
}
