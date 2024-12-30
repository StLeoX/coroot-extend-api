package main

import (
	"fmt"
	eventv1 "github.com/StLeoX/coroot-extend-api/api/proto/coroot/event/v1"
	"google.golang.org/grpc"
)

func main() {
	span := &eventv1.ServerSpan{}
	fmt.Println(span)

	fmt.Println(grpc.SupportPackageIsVersion9)
}
