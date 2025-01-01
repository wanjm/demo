package main

import (
	"github.com/wanjm/demo/gen"
)

func main() {
	wg := gen.Run(gen.Config{
		Cors:       true,
		Addr:       ":8080",
		ServerName: "servlet",
	})
	wg.Wait()
}
