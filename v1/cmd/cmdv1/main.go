package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/k0kubun/pp"
	apiv1 "github.com/ktr0731/protobuf-enum-compatibility-test/v1"
)

func main() {
	v2bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var v1msg apiv1.Message
	if err := proto.Unmarshal(v2bytes, &v1msg); err != nil {
		log.Fatal(err)
	}
	pp.Fprintf(os.Stderr, "\n--- v1 ---\n%s\n%s\n\n", v1msg, v1msg.GetErrorCode().String())
}
