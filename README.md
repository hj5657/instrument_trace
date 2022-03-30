一个用go实现的简单的函数调用链查看工具

编译生成instrument工具可执行文件
````
go build -o output/instrument github.com/hj5657/instrument_trace/cmd/instrument
````

使用instrument工具向源文件注入函数跟踪设施
````
output/instrument -w examples/demo/demo.go
````

根据build tags 来判断是否要输出trace
````
调试模式 go build -tags debug -o output/demo examples/demo/demo.go
生产    go build -o output/demo examples/demo/demo.go
````