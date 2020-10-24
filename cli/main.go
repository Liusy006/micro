package main

import (
	"context"
	"fmt"
	proto "go-micro/proto"

	micro "github.com/micro/go-micro"
)

func main() {
	services := micro.NewService(
		micro.Name("user.client"),
	)
	services.Init()

	user := proto.NewUserService("user", services.Client())
	res, err := user.Hello(context.Background(), &proto.Request{Name: "hello"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Msg)
}
