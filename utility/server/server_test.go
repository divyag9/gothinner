package main

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/divyag9/gothinner/utility"
)

var cases = []struct {
	request          *pb.Request
	expectedResponse *pb.Response
	expectedError    error
}{
	{
		request:          &pb.Request{DelayMilliSeconds: 1},
		expectedResponse: &pb.Response{PingResponse: "Utility.Ping - Waited 1 ms"},
		expectedError:    nil,
	},
	{
		request:          &pb.Request{},
		expectedResponse: &pb.Response{PingResponse: "Utility.Ping - Waited 0 ms"},
		expectedError:    nil,
	},
}

func TestCallServiceBus(t *testing.T) {
	server := &Server{}
	for _, c := range cases {
		response, err := server.Ping(context.Background(), c.request)
		if !reflect.DeepEqual(err, c.expectedError) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedError, err)
		}

		if !reflect.DeepEqual(c.expectedResponse, response) {
			t.Errorf("Expected %q but got %q", c.expectedResponse, response)
		}
	}
}
