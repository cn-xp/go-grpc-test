package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-test/protocol"
	"io"
	"net"
)

type UserServer struct {
}

func (u *UserServer) GetUser(ctx context.Context, user *protocol.User) (*protocol.User, error) {
	fmt.Println("get user from server	", user)
	return &protocol.User{
		Username: user.Username,
		Name:     "the new name",
	}, nil
}

func (u *UserServer) GetUsers(user *protocol.User, stream protocol.UserService_GetUsersServer) error {
	user1 := &protocol.User{
		Username: user.Username,
		Name:     "stream new name",
	}
	if err := stream.Send(user1); nil != err {
		fmt.Println("send user1 error", err.Error())
		return err
	}

	user2 := &protocol.User{
		Username: user.Username,
		Name:     "stream new name2",
	}
	if err := stream.Send(user2); err != nil {
		fmt.Println("send user2 error", err.Error())
		return err
	}
	return nil
}

func (u *UserServer) SaveUsers(stream protocol.UserService_SaveUsersServer) error {
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&protocol.User{
				Username: "back username",
				Name:     "name",
			})
		}

		if nil != err {
			fmt.Println("saveUser error", err.Error())
			return err
		}
		fmt.Println("saveUsers get user", user)
	}
}

func NewServer() *UserServer {
	return &UserServer{}
}

func StartServer() {
	fmt.Println("start...")
	lis, err := net.Listen("tcp", "127.0.0.1:9091")
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	protocol.RegisterUserServiceServer(grpcServer, NewServer())

	grpcServer.Serve(lis)
	fmt.Println("start success...")
}
