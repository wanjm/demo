package main

import (
	"context"

	"gitlab.plaso.cn/module-go/common"
)

func main() {
	common.InitLogger()
	common.Info(context.Background(), "Hello, World!")
}
