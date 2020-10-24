package main

import (
	"context"
	"fmt"

	proto "go-micro/proto"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
)

type User struct{}

func (u *User) Hello(ctx context.Context, req *proto.Request, reps *proto.Response) error {
	reps.Msg = "hello " + req.Name
	return nil
}

func main() {
	reg := etcdv3.NewRegistry()
	service := micro.NewService(
		micro.Name("user"),
		micro.Registry(reg),
	)

	service.Init()

	proto.RegisterUserHandler(service.Server(), new(User))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
