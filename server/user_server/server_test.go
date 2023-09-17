package user_server_test

import (
	"context"
	"github.com/rn-consider/grpcservice/dao"
	"github.com/rn-consider/grpcservice/models"
	"github.com/rn-consider/grpcservice/protos/user"
	"github.com/rn-consider/grpcservice/server/user_server"
	"log"
	"os"
	"testing"
)

// 创建一个测试服务器，可以在测试函数中使用
func createTestServer() *user_server.UserServiceServer {
	return &user_server.UserServiceServer{}
}

func TestCreateUser(t *testing.T) {
	// 创建测试服务器
	server := createTestServer()

	// 构建 CreateUser 请求
	request := &user.CreateUserRequest{
		Username: "testuser",
		Email:    "testuser@example.com",
	}

	// 调用 CreateUser 方法进行测试
	response, err := server.CreateUser(context.Background(), request)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	// 检查返回的响应是否符合预期
	if response == nil {
		t.Fatal("CreateUser response is nil")
	}
	if response.UserId == 0 {
		t.Fatal("User ID is empty")
	}
	if response.Username != request.Username {
		t.Fatalf("Username does not match. Expected: %s, Got: %s", request.Username, response.Username)
	}
	if response.Email != request.Email {
		t.Fatalf("Email does not match. Expected: %s, Got: %s", request.Email, response.Email)
	}
}

func TestGetUser(t *testing.T) {
	// 创建测试服务器
	server := createTestServer()

	// 构建 GetUser 请求
	request := &user.GetUserRequest{
		UserId: 1, // 替换为你需要的用户 ID
	}

	// 调用 GetUser 方法进行测试
	response, _ := server.GetUser(context.Background(), request)
	// 检查返回的响应是否符合预期
	if response == nil {
		t.Fatal("GetUser response is nil")
	}
	//if response.UserId != request.UserId {
	//	t.Fatalf("User ID does not match. Expected: %v, Got: %v", request.UserId, response.UserId)
	//}
	// 检查其他字段是否符合预期
}
func TestUpdateUser(t *testing.T) {
	// 创建测试服务器
	server := createTestServer()

	// 构建 GetUser 请求
	request := &user.UpdateUserRequest{
		UserId:      1, // 替换为你需要的用户 ID
		NewUsername: "changed",
		NewEmail:    "weqweqwe@test.com",
	}

	// 调用 GetUser 方法进行测试
	response, err := server.UpdateUser(context.Background(), request)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	// 检查返回的响应是否符合预期
	if response == nil {
		t.Fatal("GetUser response is nil")
	}
	if response.UserId != request.UserId {
		t.Fatalf("User ID does not match. Expected: %v, Got: %v", request.UserId, response.UserId)
	}
	// 检查其他字段是否符合预期
}
func TestDeleteUser(t *testing.T) {
	// 创建测试服务器
	server := createTestServer()

	// 构建 GetUser 请求
	request := &user.DeleteUserRequest{
		UserId: 1, // 替换为你需要的用户 ID
	}

	// 调用 DeleteUser 函数
	response, err := server.DeleteUser(context.Background(), request)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	// 检查返回的响应是否符合预期
	if response == nil {
		t.Fatal("GetUser response is nil")
	}
	if response.UserId != request.UserId {
		t.Fatalf("User ID does not match. Expected: %v, Got: %v", request.UserId, response.UserId)
	}
	print(response.Email)
	// 检查其他字段是否符合预期
}

// 其他测试函数...

func TestMain(m *testing.M) {
	err := dao.InitMySQL()
	if err != nil {
		log.Fatalf("error happend when try to InitMysql %s", err)
	}
	// 自动迁移数据库表（根据您的数据模型定义）
	if err := dao.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("error happend when try to migrate tables %s", err)
	}

	if err != nil {
		return
	}
	// 在测试之前可以进行一些初始化操作，如数据库连接等
	log.Println("Starting user_server tests...")

	// 执行测试并获取结果
	result := m.Run()

	// 在测试完成后可以进行一些清理操作，如关闭数据库连接等
	log.Println("Finishing user_server tests...")

	// 返回测试结果
	os.Exit(result)
}
