package main

import (
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/k0kubun/pp"
	apiv2 "github.com/ktr0731/protobuf-enum-compatibility-test/v2"
)

func main() {
	v2msg := &apiv2.Message{ErrorCode: apiv2.ErrorCode_UPDATED_ERROR1}
	pp.Fprintf(os.Stderr, "--- v2 ---\n%s\n%s\n\n", v2msg, v2msg.GetErrorCode().String())
	v2bytes, err := proto.Marshal(v2msg)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(v2bytes)
}
