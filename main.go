package main
import ("gitlab.plaso.cn/module-go/demo/gen")

func main() {
	wg:=gen.Run(gen.Config{
		Cors: true,
		Addr: ":8080",
		ServerName: "servlet",
	})
	wg.Wait()
}
	