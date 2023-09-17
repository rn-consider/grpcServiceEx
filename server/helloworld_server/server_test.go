package helloworld_server

import (
	"context"
	"testing"

	"github.com/rn-consider/grpcservice/protos/helloworld"
)

func TestServer_SayHello(t *testing.T) {
	// 创建 Server 实例
	server := &Server{}

	// 创建一个 HelloRequest 请求
	request := &helloworld.HelloRequest{
		Name: "John",
	}

	// 创建一个上下文
	ctx := context.Background()

	// 调用 SayHello 方法
	response, err := server.SayHello(ctx, request)

	// 检查错误
	if err != nil {
		t.Fatalf("Error calling SayHello: %v", err)
	}

	// 检查返回的消息
	expectedMessage := "Hello John"
	if response.Message != expectedMessage {
		t.Errorf("Expected message: %s, but got: %s", expectedMessage, response.Message)
	}
}

func TestServer_SayHelloAgain(t *testing.T) {
	// 创建 Server 实例
	server := &Server{}

	// 创建一个 HelloRequest 请求
	request := &helloworld.HelloRequest{
		Name: "Alice",
	}

	// 创建一个上下文
	ctx := context.Background()

	// 调用 SayHelloAgain 方法
	response, err := server.SayHelloAgain(ctx, request)

	// 检查错误
	if err != nil {
		t.Fatalf("Error calling SayHelloAgain: %v", err)
	}

	// 检查返回的消息
	expectedMessage := "Hello again Alice"
	if response.Message != expectedMessage {
		t.Errorf("Expected message: %s, but got: %s", expectedMessage, response.Message)
	}
}
