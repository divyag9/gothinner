package main

import (
	"context"
	"flag"
	"log"
	"time"

	"fmt"

	pb "github.com/divyag9/gothinner/utility"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	tls := flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile := flag.String("ca_file", "testdata/ca.pem", "The file containning the CA root cert file")
	serverAddr := flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride := flag.String("server_host_override", "", "The server name use to verify the hostname returned by TLS handshake")
	delayMilliSeconds := flag.Int("delayMilliSeconds", 0, "Number of seconds to sleep")

	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		var sn string
		if *serverHostOverride != "" {
			sn = *serverHostOverride
		}
		var creds credentials.TransportCredentials
		if *caFile != "" {
			var err error
			creds, err = credentials.NewClientTLSFromFile(*caFile, sn)
			if err != nil {
				grpclog.Fatalf("Failed to create TLS credentials %v", err)
			}
		} else {
			creds = credentials.NewClientTLSFromCert(nil, sn)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewUtilityClient(conn)

	count := 0
	var elapsed time.Duration
	for count < 20 {
		start := time.Now()
		response, err := client.Ping(context.Background(), &pb.Request{DelayMilliSeconds: int32(*delayMilliSeconds)})
		if err != nil {
			log.Fatalf("Error Pinging: %v", err)
		}
		fmt.Println("Response: ", response.PingResponse)

		elapsed = elapsed + time.Since(start)
		count++
	}
	average := elapsed / time.Duration(count)
	fmt.Println("Average elapsed: ", average)
}
