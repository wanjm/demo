
# 环境搭建
需要使用私有化module，
1. 配置环境变量: GOPRIVATE=gitlab.plaso.cn
2. 配置路径映射: 运行： git config --global url.git@gitlab.plaso.cn:.insteadOf https://gitlab.plaso.cn/

# go_servlet的程序编辑
1. 在任意目录下载go_servlet代码，git@gitlab.plaso.cn:biz_show/go_servlet.git
2. 编译代码go build -o go_servlet main.go, 生成go_servlet文件 （windows下请使用go_servlet.exe）
3. 将go_servlet文件放置到go_servlet的path环境变量中；

# 添加servlet接口
1. 参照本提交的sys/service.go文件
2. 分别是Service，Request，Response类；
3. 实现HelloWorld方法
4. Service添加注释// @goservlet servlet表示为servlet请求；
5. HelloWorld添加//@goservlet url="/hello"; title="HelloWorld"; 配置url地址和描述；
6. 运行go_servlet，生成代码；（window下也可以运行go_servlet，无需exe后缀，系统会自动识别）
7. `go mod tidy`
8. `go run main.go`
9. 发送http请求；
```
	curl 'http://localhost:8080/hello' \
	  -H 'content-type: application/json; charset=utf-8' \
	  --data-raw '{"name":"world"}'
```

