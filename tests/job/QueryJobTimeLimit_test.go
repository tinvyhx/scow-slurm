package main

import (
	"context"
	pb "scow-slurm-adapter/gen/go"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestQueryJobTimeLimit(t *testing.T) {

	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:8972", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewJobServiceClient(conn)

	// Call the Add RPC with test data
	req := &pb.QueryJobTimeLimitRequest{
		JobId: 1266,
	}
	res, err := client.QueryJobTimeLimit(context.Background(), req)
	if err != nil {
		t.Fatalf("QueryJobTimeLimit failed: %v", err)
	}

	// Check the result, 通过判断错误为nil 来决定是否执行成功
	// assert.Empty(t, err)
	assert.IsType(t, uint64(1), res.TimeLimitMinutes)
}
