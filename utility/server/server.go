package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	pb "github.com/divyag9/gothinner/utility"
)

// Server servicebus
type Server struct {
}

// Ping a service
func (s *Server) Ping(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	time.Sleep(time.Duration(request.DelayMilliSeconds) * time.Millisecond)
	response := &pb.Response{}
	response.PingResponse = fmt.Sprintf("Utility.Ping - Waited %v ms", request.DelayMilliSeconds)
	return response, nil
}

func main() {
	tls := flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile := flag.String("cert_file", "testdata/server1.pem", "The TLS cert file")
	keyFile := flag.String("key_file", "testdata/server1.key", "The TLS key file")
	port := flag.Int("port", 10000, "The server port")

	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUtilityServer(grpcServer, &Server{})
	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println("Failed to serve: ", err)
	}
}
