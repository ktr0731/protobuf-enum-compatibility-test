package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/k0kubun/pp"
	apiv1 "github.com/ktr0731/protobuf-enum-compatibility-test/v1"
	apiv2 "github.com/ktr0731/protobuf-enum-compatibility-test/v2"
)

func main() {
	v2msg := &apiv2.Message{ErrorCode: apiv2.ErrorCode_UPDATED_ERROR1}
	pp.Printf("--- v2 ---\n%s\n%s\n\n", v2msg, v2msg.GetErrorCode().String())
	v2bytes, err := proto.Marshal(v2msg)
	if err != nil {
		log.Fatal(err)
	}

	var v1msg apiv1.Message
	if err := proto.Unmarshal(v2bytes, &v1msg); err != nil {
		log.Fatal(err)
	}
	pp.Printf("--- v1 ---\n%s\n%s\n", v1msg, v1msg.GetErrorCode().String(), apiv2.ErrorCode(apiv1.ErrorCode_ERROR1) == apiv2.ErrorCode_UPDATED_ERROR1)
}
