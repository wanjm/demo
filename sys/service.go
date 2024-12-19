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
	Name string `json:"name"` //名字
}
type Response struct {
	Hello string `json:"hello"` //Hello提示语
}

// @goservlet url="/hello"; title="HelloWorld";
func (service *Service) HelloWorld(ctx context.Context, param *Request) (res *Response, err basic.Error) {
	fmt.Printf("HelloWorld: %v\n", param)
	res = &Response{
		Hello: "Hello " + param.Name,
	}
	return
}
