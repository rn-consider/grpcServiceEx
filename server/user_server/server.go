package user_server

import (
	"context"
	"github.com/rn-consider/grpcservice/models"
	"github.com/rn-consider/grpcservice/protos/user"
	"log"
)

// UserServiceServer 是用户服务的 gRPC 服务器
type UserServiceServer struct{}

// CreateUser 实现 CreateUser gRPC 方法
func (s *UserServiceServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error) {
	// 创建 User 模型
	userModel := &models.User{
		Name:  req.GetUsername(),
		Email: req.GetEmail(),
	}

	// 调用模型中的 CreateUser 方法
	if err := models.CreateUser(userModel); err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, err
	}

	// 构建响应
	userResponse := &user.UserResponse{
		UserId:   uint64(userModel.ID), // 使用新创建的用户的ID
		Username: userModel.Name,
		Email:    userModel.Email,
	}

	return userResponse, nil
}

// GetUser 实现 GetUser gRPC 方法
func (s *UserServiceServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.UserResponse, error) {
	// 调用模型中的 GetAUser 方法
	userModel, err := models.GetAUser(int(req.GetUserId()))
	if err != nil {
		return &user.UserResponse{}, nil // 返回一个空的用户响应对象
	}

	// 构建响应
	userResponse := &user.UserResponse{
		UserId:   uint64(userModel.ID),
		Username: userModel.Name,
		Email:    userModel.Email,
	}

	return userResponse, nil
}

// UpdateUser 实现 UpdateUser gRPC 方法
func (s *UserServiceServer) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UserResponse, error) {
	// 先获取用户信息
	userModel, err := models.GetAUser(int(req.GetUserId()))
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		return nil, err
	}

	// 更新用户信息
	userModel.Name = req.GetNewUsername()
	userModel.Email = req.GetNewEmail()

	// 调用模型中的 UpdatedAtUser 方法
	if err := models.UpdatedAtUser(userModel); err != nil {
		log.Printf("Failed to update user: %v", err)
		return nil, err
	}

	// 构建响应
	userResponse := &user.UserResponse{
		UserId:   uint64(userModel.ID),
		Username: userModel.Name,
		Email:    userModel.Email,
	}

	return userResponse, nil
}

// DeleteUser 实现 DeleteUser gRPC 方法
func (s *UserServiceServer) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.UserResponse, error) {
	// 调用模型中的 DeleteUser 方法
	if err := models.DeleteUser(int(req.GetUserId())); err != nil {
		log.Printf("Failed to delete user: %v", err)
		return nil, err
	}

	// 构建响应
	userResponse := &user.UserResponse{
		UserId: req.GetUserId(),
	}

	return userResponse, nil
}

// 在这里你可以添加其他辅助函数和逻辑
