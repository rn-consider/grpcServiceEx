package main

import (
	"github.com/rn-consider/grpcservice/dao"
	"github.com/rn-consider/grpcservice/protos/helloworld"
	"github.com/rn-consider/grpcservice/protos/user"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type userService struct {
	user.UnimplementedUserServer
}
type helloworldService struct {
	helloworld.UnimplementedGreeterServer
}

func registerServices(s *grpc.Server) {
	user.RegisterUserServer(s, &userService{})
	helloworld.RegisterGreeterServer(s, &helloworldService{})
	reflection.Register(s)
}
func startServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
}
func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Reading config fatal!")
	}
	Shost := viper.GetString("service.host")
	Sport := viper.GetString("service.port")
	err = dao.InitMySQL()
	if err != nil {
		return
	}
	listenAddr := Shost + ":" + Sport
	if len(listenAddr) == 0 {
		listenAddr = ":50051"
	}
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	registerServices(s)

	// 输出服务器启动信息
	log.Printf("Server is starting on %s...", listenAddr)

	log.Fatal(startServer(s, lis))
}
