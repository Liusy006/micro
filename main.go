package main

import (
	"context"
	"fmt"

	proto "go-micro/proto"

	micro "github.com/micro/go-micro"
)

type User struct{}

func (u *User) Hello(ctx context.Context, req *proto.Request, reps *proto.Response) error {
	reps.Msg = "hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("user"),
	)

	service.Init()

	proto.RegisterUserHandler(service.Server(), new(User))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
