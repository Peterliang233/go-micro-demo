package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"
	example "github.com/peterliang/demo/user/proto"
)

type Example struct{}


func (e *Example) Handle(ctx context.Context, msg *example.LoginRequest) error {
	log.Fatalf("handler had received the message", msg)
	return nil
}

func (e *Example) Handler (ctx context.Context, msg *example.LoginRequest) error {
	log.Fatalf("function had received the message", msg)
	return nil
}