package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-test/protocol"
	"io"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if nil != err {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	c := protocol.NewUserServiceClient(conn)

	user := &protocol.User{
		Username: "user1",
		UserId:   10,
	}

	response, err2 := c.GetUser(context.Background(), user)
	if nil != err2 {
		fmt.Printf("get user error: %s \n", err2)
		return
	}

	fmt.Printf("response msg: %s \n", response)

	resStream, err3 := c.GetUsers(context.Background(), user)
	if nil != err3 {
		fmt.Println(err3.Error())
		return
	}

	for {
		res, err4 := resStream.Recv()
		if err4 == io.EOF {
			break
		}
		if nil != err4 {
			fmt.Println(err4.Error())
			break
		}
		fmt.Println("get stream response ", res)
	}

	resStream2, err5 := c.SaveUsers(context.Background())
	if nil != err5 {
		fmt.Println(err5.Error())
		return
	}

	resStream2.Send(user)
	resStream2.Send(user)
	reply, err6 := resStream2.CloseAndRecv()
	if nil != err6 {
		fmt.Println(err6.Error())
		return
	}
	fmt.Println("saveUsers reply ", reply)
}
