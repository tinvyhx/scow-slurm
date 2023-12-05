package main

import (
	"context"
	pb "scow-slurm-adapter/gen/go"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestAddUserToAccount(t *testing.T) {

	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:8972", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	// Call the Add RPC with test data
	req := &pb.AddUserToAccountRequest{
		UserId:      "test06",
		AccountName: "x_admin",
	}
	_, err = client.AddUserToAccount(context.Background(), req)
	if err != nil {
		t.Fatalf("AddUserToAccount failed: %v", err)
	}

	// Check the result, 通过判断错误为nil 来决定是否执行成功
	assert.Empty(t, err)
	// log.Println(res)
}
