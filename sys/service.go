package sys

import (
	"context"
	"fmt"

	"gitlab.plaso.cn/module-go/demo/basic"
)

// @goservlet servlet
type Service struct {
}
type Request struct {
	Name string
}
type Response struct {
	Hello string
}

// @goservlet url="/hello"; title="HelloWorld";
func (service *Service) HelloWorld(ctx context.Context, param *Request) (res *Response, err basic.Error) {
	fmt.Printf("HelloWorld: %v\n", param)
	res = &Response{
		Hello: "Hello " + param.Name,
	}
	return
}
