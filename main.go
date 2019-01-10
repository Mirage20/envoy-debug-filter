package main

import (
	"context"
	"fmt"
	ext_authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2alpha"
	google_rpc "github.com/gogo/googleapis/google/rpc"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

type server struct {
	mode string
}


func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go listen(":8081", &server{mode: "GATEWAY"})
	go listen(":8082", &server{mode: "SIDECAR_INBOUND"})
	go listen(":8083", &server{mode: "SIDECAR_OUTBOUND"})

	<-c
}

func listen(address string, serverType *server) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ext_authz.RegisterAuthorizationServer(s, serverType)
	reflection.Register(s)
	fmt.Printf("Starting %q reciver on %q\n", serverType.mode, address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Check(ctx context.Context, req *ext_authz.CheckRequest) (*ext_authz.CheckResponse, error) {

	fmt.Printf("======================================== %-24s ========================================\n", fmt.Sprintf("%s Start", s.mode))
	defer fmt.Printf("======================================== %-24s ========================================\n\n", fmt.Sprintf("%s End", s.mode))

	m := jsonpb.Marshaler{Indent: "  "}
	js, err := m.MarshalToString(req)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(js)
	}

	//fmt.Printf("%+v\n", req.Attributes.Source.Address)
	//fmt.Printf("%+v\n", req.Attributes.Destination.Address)
	//
	//var keys []string
	//for k := range req.Attributes.Request.Http.Headers {
	//	keys = append(keys, k)
	//}
	//sort.Strings(keys)
	//
	//for _, k := range keys {
	//	fmt.Printf("%+v:%+v\n", k, req.Attributes.Request.Http.Headers[k])
	//}

	resp := &ext_authz.CheckResponse{
		Status: &google_rpc.Status{Code: int32(google_rpc.OK)},
	}
	return resp, nil
}
